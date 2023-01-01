package actors

import (
	"fmt"

	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"google.golang.org/protobuf/proto"
)

type UpstreamActor struct {
	actorsystem.UntypedActor
}

func (actor *UpstreamActor) OnReceive(input proto.Message) {
	if rpcMsg, ok := input.(*pbobjs.RpcMessageWraper); ok {
		fmt.Println(rpcMsg)
	}
}

func (actor *UpstreamActor) CreateInputObj() proto.Message {
	return &pbobjs.RpcMessageWraper{}
}
