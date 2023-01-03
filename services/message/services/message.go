package services

import (
	"fmt"
	"time"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"
	"google.golang.org/protobuf/proto"
)

func SendMsg(ctx clusters.BaseContext, upMsg *pbobjs.UpMsg) (int, string, int64) {
	sendTime := time.Now().UnixMilli()
	msgId := tools.GenerateUUIDShort22()

	//TODO generate timestamp

	downMsg := &pbobjs.DownMsg{
		FromId:     ctx.RequesterId,
		Type:       pbobjs.ChannelType_Private,
		MsgType:    upMsg.MsgType,
		MsgId:      msgId,
		MsgContent: upMsg.MsgContent,
		MsgTime:    sendTime,
		Flags:      upMsg.Flags,
		ClientUid:  upMsg.ClientUid,
	}
	//save msg for sender
	SaveMsg(ctx.RequesterId, downMsg, MsgDirection_Send)
	//save msg for receiver
	SaveMsg(ctx.TargetId, downMsg, MsgDirection_Receive)

	rpcMsg := clusters.CreateServerPubWraper(ctx.RequesterId, ctx.TargetId, "msg", downMsg, ctx)
	clusters.UnicastRouteWithCallback(rpcMsg, &SendMsgAckActor{}, 5*time.Second)
	return 0, msgId, sendTime
}

type SendMsgAckActor struct {
	actorsystem.UntypedActor
}

func (actor *SendMsgAckActor) OnReceive(input proto.Message) {
	fmt.Println("ServerPubAck:", input)
}

func (actor *SendMsgAckActor) CreateInputObj() proto.Message {
	return &pbobjs.RpcMessageWraper{}
}
func (actor *SendMsgAckActor) OnTimeout() {

}
