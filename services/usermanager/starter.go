package usermanager

import (
	"fmt"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/services/usermanager/actors"
)

type UserManager struct{}

func (manager *UserManager) RegisterActors(register gmicro.IActorRegister) {
	register.RegisterActor("regUser", func() actorsystem.IUntypedActor {
		return clusters.BaseProcessActor(&actors.UserManagerActor{})
	}, 64)
}

func (manager *UserManager) Startup(args map[string]interface{}) {
	fmt.Println("Startup usermanager.")
}
func (manager *UserManager) Shutdown() {
	fmt.Println("Shutdown usermanager.")
}
