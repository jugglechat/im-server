package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/errs"
)

func ErrorHttpResp(ctx *gin.Context, err *errs.ErrorMsg) {
	ctx.JSON(int(err.HttpCode), errs.ERR_API_REQ_BODY_ILLEGAL)
}

type SuccHttpResp struct {
	errs.ErrorMsg
	Data interface{} `json:"data"`
}

func SuccessHttpResp(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, SuccHttpResp{
		ErrorMsg: errs.ErrorMsg{
			Code: 200,
			Msg:  "",
		},
		Data: data,
	})
}
