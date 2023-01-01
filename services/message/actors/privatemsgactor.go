package actors

import (
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/services/message/services"
	"google.golang.org/protobuf/proto"
)

type PrivateMsgActor struct {
	clusters.BaseActor
}

func (actor *PrivateMsgActor) OnReceive(input proto.Message) {
	if upMsg, ok := input.(*pbobjs.UpMsg); ok {
		code, msgId, sendTime := services.SendMsg(actor.Context, upMsg)
		userPubAck := clusters.CreateUserPubAckWraper(code, msgId, sendTime, actor.Context)
		actor.Sender.Tell(userPubAck, actorsystem.NoSender)
	}
}

func (actor *PrivateMsgActor) CreateInputObj() proto.Message {
	return &pbobjs.UpMsg{}
}
