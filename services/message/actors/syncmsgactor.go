package actors

import (
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/services/message/services"
	"google.golang.org/protobuf/proto"
)

type SyncMsgActor struct {
	clusters.BaseActor
}

func (actor *SyncMsgActor) OnReceive(input proto.Message) {
	if syncMsg, ok := input.(*pbobjs.SyncMsgReq); ok {
		code, messages := services.SyncMessages(actor.Context, syncMsg)
		rpcMsg := clusters.CreateQueryAckWraper(code, messages, actor.Context)
		actor.Sender.Tell(rpcMsg, actorsystem.NoSender)
	}
}

func (actor *SyncMsgActor) CreateInputObj() proto.Message {
	return &pbobjs.SyncMsgReq{}
}
