package actors

import (
	"fmt"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"google.golang.org/protobuf/proto"
)

type UserManagerActor struct {
	clusters.BaseActor
}

func (actor *UserManagerActor) OnReceive(input proto.Message) {
	fmt.Println("user receive.")
	if upMsg, ok := input.(*pbobjs.UpMsg); ok {
		fmt.Println(upMsg)
		// queryAck := clusters.CreateQueryAckWraper(&pbobjs.UpMsg{}, actor.Context)
		// actor.Sender.Tell(queryAck, actorsystem.NoSender)
		// userPubAck := clusters.CreateUserPubAckWraper(0, "msg-id", time.Now().UnixMilli(), actor.Context)
		// actor.Sender.Tell(userPubAck, actorsystem.NoSender)
	} else {
		fmt.Println("err:", input)
	}
}

func (actor *UserManagerActor) CreateInputObj() proto.Message {
	return &pbobjs.UpMsg{}
}
