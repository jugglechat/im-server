package services

import (
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
)

func SyncMessages(ctx clusters.BaseContext, syncMsg *pbobjs.SyncMsgReq) (int32, *pbobjs.DownMsgSet) {
	// userId := ctx.TargetId
	// appKey := ctx.AppKey

	return 0, nil
}
