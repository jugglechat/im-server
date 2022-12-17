package codec

import (
	"bytes"
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"github.com/jugglechat/im-server/commons/tools"
	connUtils "github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

type ImCodecHandler struct{}

func (ImCodecHandler) CodecName() string {
	return "im-codec"
}

func (ImCodecHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	var imMsg IMessage
	tmpBs := make([]byte, 1)
	reader.Read(tmpBs)
	version := tmpBs[0]
	if version == Version_1 {
		msgHeader := &MsgHeader{Version: Version_1}
		msgHeader.DecodeHeader(reader)
		var msgBodyBytes []byte
		if msgHeader.MsgBodySize > 0 {
			msgBodyBytes = make([]byte, msgHeader.MsgBodySize)
			reader.Read(msgBodyBytes)

			var obfuscationCode [8]byte
			if msgHeader.GetCmd() == Cmd_Connect {
				obfuscationCode = calObfuscationCode(msgBodyBytes)
				connUtils.SetContextAttr(ctx, connUtils.StateKey_ObfuscationCode, obfuscationCode)
			} else {
				obfuscationCode = getObfuscationCodeFromCtx(ctx)
			}
			DoObfuscation(obfuscationCode, msgBodyBytes)
		} else {
			msgBodyBytes = []byte{}
		}
		//validate checksum
		ok := msgHeader.ValidateChecksum(msgBodyBytes)
		if !ok {
			ctx.Close(fmt.Errorf("checksum failed"))
		}
		switch msgHeader.GetCmd() {
		case Cmd_Connect:
			imMsg = NewConnectMessageWithHeader(msgHeader)
		case Cmd_Disconnect:
			imMsg = NewDisconnectMessageWithHeader(msgHeader)
		case Cmd_Ping:
			imMsg = NewPingMessageWithHeader(msgHeader)
		case Cmd_Publish:
			imMsg = NewUserPublishMessageWithHeader(msgHeader)
		case Cmd_PublishAck:
			imMsg = NewServerPublishAckMessageWithHeader(msgHeader)
		case Cmd_Query:
			imMsg = NewQueryMessageWithHeader(msgHeader)
		case Cmd_QueryConfirm:
			imMsg = NewQueryConfirmMessageWithHeader(msgHeader)
		default:
			return
		}
		err := imMsg.DecodeBody(msgBodyBytes)
		if err != nil {
			ctx.Close(err)
		}
	} else {
		panic(fmt.Errorf("wrong proto received"))
	}
	ctx.HandleRead(imMsg)
}
func (ImCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case IMessage:
		msgBody, err := s.EncodeBody()
		if err == nil {
			buf := bytes.NewBuffer([]byte{})
			s.EncodeHeader(buf, msgBody) //encode header
			if len(msgBody) > 0 {
				obfuscationCode := getObfuscationCodeFromCtx(ctx)
				DoObfuscation(obfuscationCode, msgBody)
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
	if obfuscationCodeObj != nil {
		obfuscationCode := obfuscationCodeObj.([8]byte)
		return obfuscationCode
	}
	return [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
}

var fixedConnMsgBytes []byte

func getFixedConnMsgBytes() []byte {
	if len(fixedConnMsgBytes) != 8 {
		connMsg := &ConnectMsgBody{
			ProtoId: ProtoId,
		}
		bs, err := tools.PbMarshal(connMsg)
		if err == nil {
			fixedConnMsgBytes = bs[:8]
		}
	}
	return fixedConnMsgBytes
}
func calObfuscationCode(connectData []byte) [8]byte {
	fixedConnBytes := getFixedConnMsgBytes()
	code := [8]byte{0, 0, 0, 0, 0, 0, 0, 0}
	if len(connectData) > 8 && len(fixedConnBytes) == 8 {
		for i := 0; i < 8; i++ {
			code[i] = connectData[i] ^ fixedConnBytes[i]
		}
	}
	return code
}

func DoObfuscation(code [8]byte, data []byte) {
	dataLen := len(data)
	for i := 0; i < dataLen; i++ {
		data[i] = data[i] ^ code[i%8]
	}
}
