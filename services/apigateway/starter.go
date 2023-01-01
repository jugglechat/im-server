package apigateway

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/configures"
	"github.com/jugglechat/im-server/commons/gmicro"
	"github.com/jugglechat/im-server/services/apigateway/apis"
)

type ApiGateway struct {
	httpServer *gin.Engine
}

func (ser *ApiGateway) RegisterActors(register gmicro.IActorRegister) {

}
func (ser *ApiGateway) Startup(args map[string]interface{}) {
	ser.httpServer = gin.Default()
	group := ser.httpServer.Group("/apigateway")
	group.POST("/users/register", apis.Register)

	httpPort := configures.Config.ApiGateway.HttpPort
	go ser.httpServer.Run(fmt.Sprintf(":%d", httpPort))
	fmt.Println("start ApiGateway with port :", httpPort)
}

func (ser *ApiGateway) Shutdown() {

}
