package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/errs"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"
	"github.com/jugglechat/im-server/services/apigateway/models"
	"github.com/jugglechat/im-server/services/apigateway/services"
	"google.golang.org/protobuf/proto"
)

func Register(ctx *gin.Context) {
	var userInfo models.UserRegReq
	if err := ctx.BindJSON(&userInfo); err != nil {
		tools.ErrorHttpResp(ctx, errs.ERR_API_REQ_BODY_ILLEGAL)
		return
	}

	resp, err := services.SyncApiCall(ctx, "regUser", userInfo.UserId, &pbobjs.UserRegReq{
		UserId:       userInfo.UserId,
		Nickname:     userInfo.Nickname,
		UserPortrait: userInfo.UserPortrait,
	}, func() proto.Message {
		return &pbobjs.UserReqResp{}
	})
	if err != nil {
		tools.ErrorHttpResp(ctx, errs.ERR_API_INTERNAL_TIMEOUT)
		return
	}

	rpcResp, ok := resp.(*pbobjs.UserReqResp)
	if !ok {
		tools.ErrorHttpResp(ctx, errs.ERR_API_INTERNAL_RESP_FAIL)
		return
	}
	tools.SuccessHttpResp(ctx, models.UserReqResp{
		UserId: rpcResp.UserId,
		Token:  rpcResp.Token,
	})
}
