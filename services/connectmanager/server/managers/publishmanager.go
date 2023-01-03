package managers

import (
	"time"

	"github.com/jugglechat/im-server/commons/logs"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
	"github.com/rfyiamcool/go-timewheel"
)

var callbackTimeoutTimer *timewheel.TimeWheel

func init() {
	t, err := timewheel.NewTimeWheel(1*time.Second, 360)
	if err != nil {
		logs.Error("can not init timeWheel for publish callback.")
	} else {
		callbackTimeoutTimer = t
		callbackTimeoutTimer.Start()
	}
}
func PublishServerPubMessage(appkey, userid, session string, serverPubMsg *codec.PublishMsgBody, publishType int, callback func(), notOnlineCallback func()) {
	userCtxMap := GetConnectCtxByUser(appkey, userid)
	if len(userCtxMap) > 0 { //target user is online
		isSetCallback := false
		for kSess, vCtx := range userCtxMap {
			if publishType == 1 && kSess != session { //publishType:1, 只给指定的session发送
				continue
			}
			if publishType == 2 && kSess == session { //publishType:2, 除了指定session以外，给该用户其他登录端发送
				continue
			}
			if vCtx.Channel().IsActive() {
				qos := codec.QoS_NoAck
				if callback != nil {
					qos = codec.QoS_NeedAck
				}
				index := utils.GetServerIndexAfterIncrease(vCtx)
				tmpPubMsg := codec.NewServerPublishMessage(&codec.PublishMsgBody{
					Index:     int32(index),
					Topic:     serverPubMsg.Topic,
					TargetId:  serverPubMsg.TargetId,
					Timestamp: time.Now().UnixMilli(),
					Data:      serverPubMsg.Data,
				}, qos)
				vCtx.Write(tmpPubMsg)
				logs.Info(utils.GetConnSession(vCtx), utils.Action_ServerPub, tmpPubMsg.MsgBody.Index, tmpPubMsg.MsgBody.Topic, len(tmpPubMsg.MsgBody.Data))
				if callback != nil && !isSetCallback {
					isSetCallback = true
					task := callbackTimeoutTimer.Add(20*time.Second, func() {
						//do timeout
						utils.RemoveServerPubCallback(vCtx, tmpPubMsg.MsgBody.Index)
					})
					utils.PutServerPubCallback(vCtx, tmpPubMsg.MsgBody.Index, func() {
						callbackTimeoutTimer.Remove(task) //remove from timeout timer
						callback()                        //execute
					})
				}
			}
		}
	} else { //target user is not online
		if notOnlineCallback != nil {
			notOnlineCallback()
		}
	}
}

func PublishQryAckMessage(session string, qryAckMsg *codec.QueryAckMsgBody, callback func(), notOnlineCallback func()) {
	ctx := GetConnectCtxBySession(session)
	if ctx != nil {
		qos := codec.QoS_NoAck
		if callback != nil {
			qos = codec.QoS_NeedAck
			task := callbackTimeoutTimer.Add(20*time.Second, func() {
				//do timeout
				utils.RemoveQueryAckCallback(ctx, qryAckMsg.Index)
			})
			utils.PutQueryAckCallback(ctx, qryAckMsg.Index, func() {
				callbackTimeoutTimer.Remove(task)
				callback()
			})
		}
		tmpQryAckMsg := codec.NewQueryAckMessage(qryAckMsg, qos)
		ctx.Write(tmpQryAckMsg)
		logs.Info(utils.GetConnSession(ctx), utils.Action_QueryAck, qryAckMsg.Index, qryAckMsg.Code, len(qryAckMsg.Data))
	} else {
		if notOnlineCallback != nil {
			notOnlineCallback()
		}
	}
}

func PublishUserPubAckMessage(appkey, userid, session string, pubAckMsg *codec.PublishAckMsgBody) {
	ctx := GetConnectCtxBySession(session)
	if ctx != nil {
		tmpPubAckMsg := codec.NewUserPublishAckMessage(pubAckMsg)
		ctx.Write(tmpPubAckMsg)
		logs.Info(utils.GetConnSession(ctx), utils.Action_UserPubAck, pubAckMsg.Index, pubAckMsg.Code)
	}
}
