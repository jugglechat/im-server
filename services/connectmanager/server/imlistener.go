package server

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/logs"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/jugglechat/im-server/services/connectmanager/server/managers"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

type ImListener interface {
	Create(ctx netty.ActiveContext)
	Close(ctx netty.InactiveContext)
	ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception)

	Connected(msg *codec.ConnectMsgBody, ctx netty.InboundContext)
	Diconnected(msg *codec.DisconnectMsgBody, ctx netty.InboundContext)
	PublishArrived(msg *codec.PublishMsgBody, qos int, ctx netty.InboundContext)
	PubAckArrived(msg *codec.PublishAckMsgBody, ctx netty.InboundContext)
	QueryArrived(msg *codec.QueryMsgBody, ctx netty.InboundContext)
	QueryConfirmArrived(msg *codec.QueryConfirmMsgBody, ctx netty.InboundContext)
	PingArrived(ctx netty.InboundContext)
}

type ImListenerImpl struct{}

func (*ImListenerImpl) Create(ctx netty.ActiveContext) {
	// utils.SetContextAttr(ctx, utils.StateKey_ConnectSession, tools.GenerateUUIDShortString())
	// utils.SetContextAttr(ctx, utils.StateKey_ConnectCreateTime, time.Now().UnixMilli())
	// utils.SetContextAttr(ctx, utils.StateKey_CtxLocker, &sync.Mutex{})
}
func (*ImListenerImpl) Close(ctx netty.InactiveContext) {
}
func (*ImListenerImpl) ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception) {}
func (*ImListenerImpl) Connected(msg *codec.ConnectMsgBody, ctx netty.InboundContext) {
	clientIp := msg.ClientIp
	if clientIp == "" {
		clientIp = ctx.Channel().RemoteAddr()
	}
	//check something
	if code := managers.CheckLogin(ctx, msg); code > 0 {
		msgAck := codec.NewConnectAckMessage(&codec.ConnectAckMsgBody{
			Code:      code,
			Session:   utils.GetConnSession(ctx),
			Timestamp: time.Now().UnixMilli(),
		})
		ctx.Channel().Write(msgAck)
		ctx.Channel().Close(errors.New("Failed to Login."))
		return
	}
	userId := utils.GetContextAttrString(ctx, utils.StateKey_UserID)
	//success
	logs.Info(utils.GetConnSession(ctx), utils.Action_Connect, msg.Appkey, userId, msg.SdkVersion, msg.DeviceId, msg.Platform, msg.DeviceCompany, msg.DeviceModel, msg.DeviceOsVersion, msg.NetworkId, msg.IspNum, clientIp)

	utils.SetContextAttr(ctx, utils.StateKey_Appkey, msg.Appkey)
	utils.SetContextAttr(ctx, utils.StateKey_Platform, msg.Platform)
	utils.SetContextAttr(ctx, utils.StateKey_Version, msg.SdkVersion)
	utils.SetContextAttr(ctx, utils.StateKey_ClientIp, clientIp)
	managers.PutInContextCache(ctx)

	msgAck := codec.NewConnectAckMessage(&codec.ConnectAckMsgBody{
		Code:      utils.ConnectAckState_Access,
		UserId:    msg.Token,
		Session:   utils.GetConnSession(ctx),
		Timestamp: time.Now().UnixMilli(),
	})
	ctx.Channel().Write(msgAck)
}
func (*ImListenerImpl) Diconnected(msg *codec.DisconnectMsgBody, ctx netty.InboundContext) {
	logs.Info(utils.GetConnSession(ctx), utils.Action_Disconnect, msg.Code)
	ctx.Close(fmt.Errorf("dissconnect"))
	managers.RemoveFromContextCache(ctx)
}
func (*ImListenerImpl) PublishArrived(msg *codec.PublishMsgBody, qos int, ctx netty.InboundContext) {
	logs.Info(utils.GetConnSession(ctx), utils.Action_UserPub, msg.Index, msg.Topic, msg.TargetId, len(msg.Data))
	clusters.UnicastRoute(&pbobjs.RpcMessageWraper{
		RpcMsgType:   pbobjs.RpcMsgType_UserPub,
		AppKey:       utils.GetContextAttrString(ctx, utils.StateKey_Appkey),
		Session:      utils.GetConnSession(ctx),
		Method:       msg.Topic,
		RequesterId:  utils.GetContextAttrString(ctx, utils.StateKey_UserID),
		ReqIndex:     msg.Index,
		Qos:          int32(qos),
		AppDataBytes: msg.Data,
		TargetId:     msg.TargetId,
	}, "connect")
}
func (*ImListenerImpl) PubAckArrived(msg *codec.PublishAckMsgBody, ctx netty.InboundContext) {
	logs.Info(utils.GetConnSession(ctx), utils.Action_ServerPubAck, msg.Index)
}
func (*ImListenerImpl) QueryArrived(msg *codec.QueryMsgBody, ctx netty.InboundContext) {
	logs.Info(utils.GetConnSession(ctx), utils.Action_Query, msg.Index, msg.Topic, msg.TargetId, len(msg.Data))

	qos := codec.QoS_NoAck
	r := rand.Intn(2)
	if r == 1 {
		qos = codec.QoS_NeedAck
	}
	//Debug
	retData := []byte{1, 2, 3}
	ctx.Write(codec.NewQueryAckMessage(&codec.QueryAckMsgBody{
		Index:     msg.Index,
		Code:      0,
		Timestamp: time.Now().UnixMilli(),
		Data:      retData,
	}, qos))
	logs.Info(utils.GetConnSession(ctx), utils.Action_QueryAck, msg.Index, 0, len(retData))
}
func (*ImListenerImpl) QueryConfirmArrived(msg *codec.QueryConfirmMsgBody, ctx netty.InboundContext) {
	logs.Info(utils.GetConnSession(ctx), utils.Action_QueryConfirm, msg.Index)
}
func (*ImListenerImpl) PingArrived(ctx netty.InboundContext) {
	ctx.Write(codec.NewPonMessage())
}
