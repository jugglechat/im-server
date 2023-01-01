package message

import (
	"fmt"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/services/message/actors"
)

type MessageManager struct{}

func (manager *MessageManager) RegisterActors(register gmicro.IActorRegister) {
	register.RegisterActor("pMsg", func() actorsystem.IUntypedActor {
		return clusters.BaseProcessActor(&actors.PrivateMsgActor{})
	}, 64)
}

func (manager *MessageManager) Startup(args map[string]interface{}) {
	fmt.Println("Startup message.")
}
func (manager *MessageManager) Shutdown() {
	fmt.Println("Shutdown message.")
}
