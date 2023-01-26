package navigator

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/configures"
	"github.com/jugglechat/im-server/commons/gmicro"
	"github.com/jugglechat/im-server/services/navigator/apis"
)

type Navigator struct {
	httpServer *gin.Engine
}

func (ser *Navigator) RegisterActors(register gmicro.IActorRegister) {

}
func (ser *Navigator) Startup(args map[string]interface{}) {
	ser.httpServer = gin.Default()
	group := ser.httpServer.Group("/navigator")
	group.POST("/mobile", apis.MobileNavi)
	group.GET("/mobile", apis.MobileNaviGet)
	group.POST("/web", apis.WebNavi)
	group.GET("/web", apis.WebNaviGet)

	httpPort := configures.Config.NavGateway.HttpPort
	go ser.httpServer.Run(fmt.Sprintf(":%d", httpPort))
	fmt.Println("start Navitor with port :", httpPort)
}

func (ser *Navigator) Shutdown() {

}
