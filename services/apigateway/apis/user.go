package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/services/apigateway/models"
	"google.golang.org/protobuf/proto"
)

func Register(ctx *gin.Context) {
	var userInfo models.UserRegReq
	if err := ctx.BindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr(err.Error()))
		return
	}

	resp, err := SyncApiCall(ctx, "regUser", userInfo.UserId, &pbobjs.UserRegReq{
		UserId:       userInfo.UserId,
		Nickname:     userInfo.Nickname,
		UserPortrait: userInfo.UserPortrait,
	}, func() proto.Message {
		return &pbobjs.UserReqResp{}
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.DefaultErr(err.Error()))
		return
	}

	rpcResp, ok := resp.(*pbobjs.UserReqResp)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, models.DefaultErr("pb not match."))
		return
	}

	ctx.JSON(http.StatusOK, models.Success(models.UserReqResp{
		UserId: rpcResp.UserId,
		Token:  rpcResp.Token,
	}))
}
