package main

import (
	"fmt"
	"sync"

	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/configures"
	"github.com/jugglechat/im-server/commons/dbs"
	"github.com/jugglechat/im-server/commons/imstarters"
	"github.com/jugglechat/im-server/commons/logs"
	"github.com/jugglechat/im-server/services/apigateway"
	"github.com/jugglechat/im-server/services/chatroom"
	"github.com/jugglechat/im-server/services/connectmanager"
	"github.com/jugglechat/im-server/services/message"
	"github.com/jugglechat/im-server/services/usermanager"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	//init configures
	if err := configures.InitConfigures(); err != nil {
		//logs.Error("Init Configures failed.", err)
		fmt.Println("Init Configures failed.", err)
		return
	}
	//init logs
	logs.InitLogs()
	//init mysql
	if err := dbs.InitMysql(); err != nil {
		logs.Error("Init Mysql failed.", err)
		return
	}
	//init cluster
	if err := clusters.InitCluster(); err != nil {
		logs.Error("Init Cluster failed.", err)
		return
	}

	imstarters.Loaded(&connectmanager.ConnectManager{})
	imstarters.Loaded(&apigateway.ApiGateway{})
	imstarters.Loaded(&message.MessageManager{})
	imstarters.Loaded(&chatroom.ChatRoomManager{})
	imstarters.Loaded(&usermanager.UserManager{})

	imstarters.Startup()

	waitgroup.Wait()
}
