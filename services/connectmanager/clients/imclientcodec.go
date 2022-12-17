package clients

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	connUtils "github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

type ImClientCodecHandler struct{}

func (ImClientCodecHandler) CodecName() string {
	return "im-client-codec"
}

func (ImClientCodecHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	var imMsg codec.IMessage
	tmpBs := make([]byte, 1)
	reader.Read(tmpBs)
	version := tmpBs[0]
	if version == codec.Version_1 {
		msgHeader := &codec.MsgHeader{Version: codec.Version_1}
		msgHeader.DecodeHeader(reader)
		var msgBodyBytes []byte
		if msgHeader.MsgBodySize > 0 {
			msgBodyBytes = make([]byte, msgHeader.MsgBodySize)
			reader.Read(msgBodyBytes)
			obfuscationCode := getObfuscationCodeFromCtx(ctx)
			codec.DoObfuscation(obfuscationCode, msgBodyBytes)
		} else {
			msgBodyBytes = []byte{}
		}
		//validate checksum
		ok := msgHeader.ValidateChecksum(msgBodyBytes)
		if !ok {
			ctx.Close(fmt.Errorf("checksum failed"))
		}
		switch msgHeader.GetCmd() {
		case codec.Cmd_ConnectAck:
			imMsg = codec.NewConnectAckMessageWithHeader(msgHeader)
		case codec.Cmd_Disconnect:
			imMsg = codec.NewDisconnectMessageWithHeader(msgHeader)
		case codec.Cmd_Pong:
			imMsg = codec.NewPongMessageWithHeader(msgHeader)
		case codec.Cmd_Publish:
			imMsg = codec.NewServerPublishMessageWithHeader(msgHeader)
		case codec.Cmd_PublishAck:
			imMsg = codec.NewUserPublishAckMessageWithHeader(msgHeader)
		case codec.Cmd_QueryAck:
			imMsg = codec.NewQueryAckMessageWithHeader(msgHeader)
		default:
			return
		}
		err := imMsg.DecodeBody(msgBodyBytes)
		if err != nil {
			ctx.Close(err)
		}
	}
	ctx.HandleRead(imMsg)
}
func (ImClientCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case codec.IMessage:
		msgBody, err := s.EncodeBody()
		if err == nil {
			buf := bytes.NewBuffer([]byte{})
			s.EncodeHeader(buf, msgBody) //encode header
			if len(msgBody) > 0 {
				obfuscationCode := getObfuscationCodeFromCtx(ctx)
				codec.DoObfuscation(obfuscationCode, msgBody)
				buf.Write(msgBody) //write msg body
			}
			ctx.HandleWrite(buf)
		} else {
			ctx.HandleWrite(message)
		}
	default:
		ctx.HandleWrite(message)
	}
}

func getObfuscationCodeFromCtx(ctx netty.HandlerContext) [8]byte {
	obfuscationCodeObj := connUtils.GetContextAttr(ctx, connUtils.StateKey_ObfuscationCode)
	obfuscationCode := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	if obfuscationCodeObj == nil {
		obfuscationCode = randomObfuscationCode()
		connUtils.SetContextAttr(ctx, connUtils.StateKey_ObfuscationCode, obfuscationCode)
	} else {
		obfuscationCode = obfuscationCodeObj.([8]byte)
	}
	return obfuscationCode
}

func randomObfuscationCode() [8]byte {
	code := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 8; i++ {
		code[i] = byte(rand.Intn(256))
	}
	return code
}
