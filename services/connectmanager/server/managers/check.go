package managers

import (
	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/commons/caches"
	"github.com/jugglechat/im-server/commons/tokens"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

func CheckLogin(ctx netty.InboundContext, msg *codec.ConnectMsgBody) int32 {
	appkey := msg.Appkey
	tokenStr := msg.Token

	tokenWrap, err := tokens.ParseTokenString(tokenStr)
	if err == nil {
		appinfo := caches.GetAppInfo(appkey)
		if appinfo != nil {
			token, err := tokens.ParseToken(tokenWrap, []byte(appinfo.AppSecureKey))
			if err == nil {
				utils.SetContextAttr(ctx, utils.StateKey_UserID, token.UserId)
				return utils.ConnectAckState_Access
			}
		}
	}
	return utils.ConnectAckState_AuthorizeFailed
}
