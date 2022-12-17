package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty-transport/websocket"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

type ImWebsocketServer struct {
	MessageListener ImListener
	bootstrap       netty.Bootstrap
}

func (server *ImWebsocketServer) SyncStart(port int) {
	// setup websocket params.
	options := &websocket.Options{
		Timeout:  time.Second * 5,
		ServeMux: http.NewServeMux(),
	}
	//check page.
	options.ServeMux.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
	})
	//initializer
	setupCodec := func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.PacketCodec(65536 * 4)).
			AddLast(codec.NewReadTimeoutHandler(300 * time.Second)).
			AddLast(codec.ImWebsocketCodecHandler{}).
			AddLast(IMWebsocketMsgHandler{server.MessageListener})
		utils.InitCtxAttrByChannel(channel)
	}
	//setup bootstrap & startup server
	server.bootstrap = netty.NewBootstrap(netty.WithChildInitializer(setupCodec), netty.WithTransport(websocket.New()))

	server.bootstrap.Listen(fmt.Sprintf("0.0.0.0:%d/im", port), websocket.WithOptions(options)).Sync()
}

func (server *ImWebsocketServer) Stop() {
	if server.bootstrap != nil {
		server.bootstrap.Shutdown()
	}
}
