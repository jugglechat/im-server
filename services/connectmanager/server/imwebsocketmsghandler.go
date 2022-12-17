package server

import (
	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
)

type IMWebsocketMsgHandler struct {
	listener ImListener
}

func (handler IMWebsocketMsgHandler) HandleActive(ctx netty.ActiveContext) {
	if handler.listener != nil {
		handler.listener.Create(ctx)
	}
	//ctx.HandleActive()
}

func (handler IMWebsocketMsgHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if handler.listener != nil {
		wsMsg, ok := message.(*codec.ImWebsocketMsg)
		if ok {
			switch wsMsg.Cmd {
			case int32(codec.Cmd_Connect):
				handler.listener.Connected(wsMsg.GetConnectMsgBody(), ctx)
			case int32(codec.Cmd_Disconnect):
				handler.listener.Diconnected(wsMsg.GetDisconnectMsgBody(), ctx)
			case int32(codec.Cmd_Ping):
				handler.listener.PingArrived(ctx)
			case int32(codec.Cmd_Publish):
				handler.listener.PublishArrived(wsMsg.GetPublishMsgBody(), int(wsMsg.GetQos()), ctx)
			case int32(codec.Cmd_PublishAck):
				handler.listener.PubAckArrived(wsMsg.GetPubAckMsgBody(), ctx)
			case int32(codec.Cmd_Query):
				handler.listener.QueryArrived(wsMsg.GetQryMsgBody(), ctx)
			case int32(codec.Cmd_QueryConfirm):
				handler.listener.QueryConfirmArrived(wsMsg.GetQryConfirmMsgBody(), ctx)
			default:
				break
			}
		}
	}
	//ctx.HandleRead(message)
}

func (handler IMWebsocketMsgHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	if handler.listener != nil {
		handler.listener.Close(ctx)
	}
	ctx.Close(ex)
	//ctx.HandleInactive(ex)
}

func (handler IMWebsocketMsgHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	if handler.listener != nil {
		handler.listener.ExceptionCaught(ctx, ex)
	}
	ctx.Close(ex)
	//ctx.HandleException(ex)
}
