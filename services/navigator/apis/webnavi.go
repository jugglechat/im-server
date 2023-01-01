package apis

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/caches"
	"github.com/jugglechat/im-server/commons/configures"
	"github.com/jugglechat/im-server/commons/tokens"
	"github.com/jugglechat/im-server/services/navigator/models"
)

func WebNavi(ctx *gin.Context) {
	var navReq models.NaviReq
	if err := ctx.BindJSON(&navReq); err != nil {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr(err.Error()))
		return
	}
	appKey := navReq.AppKey
	tokenStr := navReq.Token

	appInfo := caches.GetAppInfo(appKey)
	if appInfo != nil {
		tokenWrap, err := tokens.ParseTokenString(tokenStr)
		if err == nil {
			appinfo := caches.GetAppInfo(appKey)
			if appinfo != nil {
				token, err := tokens.ParseToken(tokenWrap, []byte(appinfo.AppSecureKey))
				if err == nil {
					ctx.JSON(http.StatusOK, models.DefaultResp{
						Code: 0,
						Data: models.NaviResp{
							AppKey:  appKey,
							UserId:  token.UserId,
							Servers: []string{fmt.Sprintf("%s:%d", configures.Config.ConnectManager.WsHost, configures.Config.ConnectManager.WsPort)},
						},
					})
				}
			}
		}
	}
}
