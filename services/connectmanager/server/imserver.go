package server

import (
	"fmt"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

type ImServer struct {
	MessageListener ImListener
	bootstrap       netty.Bootstrap
}

func (server *ImServer) SyncStart(port int) {
	childInitializer := func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(codec.ImCodecHandler{}).
			AddLast(codec.NewReadTimeoutHandler(300 * time.Second)).
			AddLast(ImMessageHandler{server.MessageListener})
		utils.InitCtxAttrByChannel(channel)
	}

	// new bootstrap
	server.bootstrap = netty.NewBootstrap(netty.WithChildInitializer(childInitializer))

	// setup bootstrap & startup server.
	server.bootstrap.Listen(fmt.Sprintf("0.0.0.0:%d", port)).Sync()
}

func (server *ImServer) Stop() {
	if server.bootstrap != nil {
		server.bootstrap.Shutdown()
	}
}
