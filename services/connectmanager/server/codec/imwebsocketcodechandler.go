package codec

import (
	"bytes"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"github.com/jugglechat/im-server/commons/tools"
	"google.golang.org/protobuf/proto"
)

type ImWebsocketCodecHandler struct{}

func (ImWebsocketCodecHandler) CodecName() string {
	return "im-websocket-codec"
}

func (ImWebsocketCodecHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	bs := utils.MustToBytes(message)
	wsMsg := &ImWebsocketMsg{}
	tools.PbUnMarshal(bs, wsMsg)
	ctx.HandleRead(wsMsg)
}

func (ImWebsocketCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case proto.Message:
		buf := bytes.NewBuffer([]byte{})
		bs, err := tools.PbMarshal(s)
		if err == nil {
			buf.Write(bs)
			ctx.HandleWrite(buf)
		} else {
			ctx.HandleWrite(message)
		}
	default:
		ctx.HandleWrite(message)
	}
}
