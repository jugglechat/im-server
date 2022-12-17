package clients

import (
	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
)

type ImClientMessageHandler struct {
	Client *ImClient
}

func (handler ImClientMessageHandler) HandleActive(ctx netty.ActiveContext) {
	ctx.HandleActive()
}

func (handler ImClientMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	switch msg := message.(type) {
	case *codec.ConnectAckMessage:
		handler.Client.OnConnectAck(msg)
	case *codec.DisconnectMessage:
		handler.Client.OnDisconnect(msg)
	case *codec.PongMessage:
		handler.Client.OnPong(msg)
	case *codec.UserPublishAckMessage:
		handler.Client.OnPublishAck(msg)
	case *codec.ServerPublishMessage:
		handler.Client.OnPublish(msg)
	case *codec.QueryAckMessage:
		handler.Client.OnQueryAck(msg)
	default:
		break
	}

	ctx.HandleRead(message)
}

func (handler ImClientMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	handler.Client.OnInactive(ex)
	ctx.Close(ex)
	ctx.HandleInactive(ex)
}

func (handler ImClientMessageHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	handler.Client.OnException(ex)
	ctx.Close(ex)
	ctx.HandleException(ex)
}
