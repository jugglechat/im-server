package apis

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"
	"github.com/jugglechat/im-server/services/apigateway/models"
)

func Register(ctx *gin.Context) {
	var userInfo models.UserInfo
	if err := ctx.BindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr())
		return
	}

	m := &pbobjs.UpMsg{
		MsgContent: []byte{1, 2, 3},
	}
	bs, _ := tools.PbMarshal(m)

	result, err := clusters.SyncUnicastRoute(&pbobjs.RpcMessageWraper{
		RpcMsgType:   pbobjs.RpcMsgType_QueryMsg,
		AppKey:       "appkey",
		Session:      "session",
		Method:       "regUser",
		RequesterId:  "requestId",
		Qos:          1,
		AppDataBytes: bs,
		TargetId:     userInfo.UserId,
	}, 5*time.Second)

	fmt.Println(result, err)

	ctx.JSON(http.StatusOK, userInfo)
}
