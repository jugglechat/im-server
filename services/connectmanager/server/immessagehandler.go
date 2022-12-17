package server

import (
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
)

type ImMessageHandler struct {
	listener ImListener
}

func (handler ImMessageHandler) HandleActive(ctx netty.ActiveContext) {
	if handler.listener != nil {
		handler.listener.Create(ctx)
	}
	fmt.Println("active")
	//ctx.HandleActive()
}

func (handler ImMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if handler.listener != nil {
		switch msg := message.(type) {
		case *codec.ConnectMessage:
			handler.listener.Connected(msg.MsgBody, ctx)
		case *codec.DisconnectMessage:
			handler.listener.Diconnected(msg.MsgBody, ctx)
		case *codec.PingMessage:
			handler.listener.PingArrived(ctx)
		case *codec.UserPublishMessage:
			handler.listener.PublishArrived(msg.MsgBody, msg.GetQoS(), ctx)
		case *codec.ServerPublishAckMessage:
			handler.listener.PubAckArrived(msg.MsgBody, ctx)
		case *codec.QueryMessage:
			handler.listener.QueryArrived(msg.MsgBody, ctx)
		case *codec.QueryConfirmMessage:
			handler.listener.QueryConfirmArrived(msg.MsgBody, ctx)
		default:
			break
		}
	}
	//ctx.HandleRead(message)
}

func (handler ImMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	if handler.listener != nil {
		handler.listener.Close(ctx)
	}
	fmt.Println("inactive", ex)
	ctx.Close(ex)
	//ctx.HandleInactive(ex)
}

func (handler ImMessageHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	if handler.listener != nil {
		handler.listener.ExceptionCaught(ctx, ex)
	}
	fmt.Println("exception")
	ctx.Close(ex)
	//ctx.HandleException(ex)
}
