package apis

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/caches"
	"github.com/jugglechat/im-server/commons/errs"
	"github.com/jugglechat/im-server/services/apigateway/services"
)

const (
	Header_AppKey    string = "appkey"
	Header_Nonce     string = "nonce"
	Header_Timestamp string = "timestamp"
	Header_Signature string = "signature"
)

func Signature(ctx *gin.Context) {
	appKey := ctx.Request.Header.Get(Header_AppKey)
	nonce := ctx.Request.Header.Get(Header_Nonce)
	tsStr := ctx.Request.Header.Get(Header_Timestamp)
	signature := ctx.Request.Header.Get(Header_Signature)
	if appKey == "" {
		ctx.JSON(http.StatusBadRequest, errs.ERR_API_APPKEY_REQUIRED)
		ctx.Abort()
		return
	}
	if nonce == "" {
		ctx.JSON(http.StatusBadRequest, errs.ERR_API_NONCE_REQUIRED)
		ctx.Abort()
		return
	}
	if tsStr == "" {
		ctx.JSON(http.StatusBadRequest, errs.ERR_API_TIMESTAMP_REQUIRED)
		ctx.Abort()
		return
	}
	if signature == "" {
		ctx.JSON(http.StatusBadRequest, errs.ERR_API_SIGNATURE_REQUIRED)
		ctx.Abort()
		return
	}
	appInfo := caches.GetAppInfo(appKey)
	if appInfo != nil {
		str := fmt.Sprintf("%s%s%s", appInfo.AppSecret, nonce, tsStr)
		sig := SHA1(str)
		if sig == signature {
			ctx.Set(services.CtxKey_AppKey, appKey)
		} else {
			ctx.JSON(http.StatusForbidden, errs.ERR_API_SIGNATURE_FAIL)
			ctx.Abort()
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, errs.ERR_API_APP_NOT_EXISTED)
		ctx.Abort()
		return
	}
}

func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

func TestSha1() {
	str := fmt.Sprintf("%s%s%s", "appsecret", "nonce", "1672568121910")
	sig := SHA1(str)
	fmt.Println(sig)
}
