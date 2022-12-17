package chatroom

import (
	"fmt"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/services/message/actors"
	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
)

type ChatRoomManager struct{}

func (manager *ChatRoomManager) RegisterActors(register gmicro.IActorRegister) {
	register.RegisterActor("cMsg", func() actorsystem.IUntypedActor {
		return clusters.BaseProcessActor(&actors.PrivateMsgActor{})
	}, 64)
}

func (manager *ChatRoomManager) Startup(args map[string]interface{}) {
	fmt.Println("Startup chatroom.")
}
func (manager *ChatRoomManager) Shutdown() {
	fmt.Println("Shutdown chatroom.")
}
