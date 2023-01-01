package imstarters

import (
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro"
)

type IServiceStarter interface{}

type IRegisterActorsHandler interface {
	RegisterActors(register gmicro.IActorRegister)
}

type IStartupHandler interface {
	Startup(args map[string]interface{})
}
type IShutdownHandler interface {
	Shutdown()
}

var serverList []IServiceStarter

func Loaded(server IServiceStarter) {
	if server != nil {
		//register actors
		if regActorsHandler, ok := server.(IRegisterActorsHandler); ok {
			regActorsHandler.RegisterActors(clusters.GetCluster())
		}
		serverList = append(serverList, server)
	}
}

func Startup() {
	for _, server := range serverList {
		//execute startup
		if startHandler, ok := server.(IStartupHandler); ok {
			startHandler.Startup(map[string]interface{}{})
		}
	}
	clusters.Startup()
}

func Shutdown() {
	//remove self from zk TODO
	for _, server := range serverList {
		//execute startup
		if shutdownHandler, ok := server.(IShutdownHandler); ok {
			shutdownHandler.Shutdown()
		}
	}
}
