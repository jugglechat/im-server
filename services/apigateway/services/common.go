package services

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"
	"google.golang.org/protobuf/proto"
)

const (
	CtxKey_AppKey  string = "CtxKey_AppKey"
	CtxKey_Session string = "CtxKey_Session"
)

func SyncApiCall(ctx *gin.Context, method, targetId string, req proto.Message, respFactory func() proto.Message) (proto.Message, error) {
	dataBytes, _ := tools.PbMarshal(req)
	result, err := clusters.SyncUnicastRoute(&pbobjs.RpcMessageWraper{
		RpcMsgType:   pbobjs.RpcMsgType_QueryMsg,
		AppKey:       GetCtxString(ctx, CtxKey_AppKey),
		Session:      GetCtxString(ctx, CtxKey_Session),
		Method:       method,
		RequesterId:  "",
		Qos:          1,
		AppDataBytes: dataBytes,
		TargetId:     targetId,
	}, 5*time.Second)
	if err != nil {
		return nil, err
	}
	respObj := respFactory()
	err = tools.PbUnMarshal(result.AppDataBytes, respObj)
	if err != nil {
		return nil, err
	}
	return respObj, nil
}
func GetCtxString(ctx *gin.Context, key string) string {
	val, exist := ctx.Get(key)
	if exist {
		return val.(string)
	} else {
		return ""
	}
}
