package actors

import (
	"fmt"
	"time"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"google.golang.org/protobuf/proto"
)

type ChatMsgActor struct {
	clusters.BaseActor
}

func (actor ChatMsgActor) OnReceive(input proto.Message) {
	if upMsg, ok := input.(*pbobjs.UpMsg); ok {
		fmt.Println(upMsg)
		userPubAck := clusters.CreateUserPubAckWraper(0, "msg-id", time.Now().UnixMilli(), actor.Context)
		actor.Sender.Tell(userPubAck, actorsystem.NoSender)
	}
}

func (actor ChatMsgActor) CreateInputObj() proto.Message {
	return &pbobjs.UpMsg{}
}
