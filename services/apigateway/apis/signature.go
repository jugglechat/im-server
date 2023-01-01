package apis

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jugglechat/im-server/commons/caches"
	"github.com/jugglechat/im-server/services/apigateway/models"
)

const (
	Header_AppKey    string = "AppKey"
	Header_Nonce     string = "Nonce"
	Header_Timestamp string = "Timestamp"
	Header_Signature string = "Signature"
)

func Signature(ctx *gin.Context) {
	appKey := ctx.Request.Header.Get(Header_AppKey)
	nonce := ctx.Request.Header.Get(Header_Nonce)
	tsStr := ctx.Request.Header.Get(Header_Timestamp)
	signature := ctx.Request.Header.Get(Header_Signature)
	if appKey == "" {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr("AppKey is empty."))
		ctx.Abort()
		return
	}
	if nonce == "" {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr("Nonce is empty."))
		ctx.Abort()
		return
	}
	if tsStr == "" {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr("Timestamp is empty."))
		ctx.Abort()
		return
	}
	if signature == "" {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr("Signature is empty."))
		ctx.Abort()
		return
	}
	appInfo := caches.GetAppInfo(appKey)
	if appInfo != nil {
		str := fmt.Sprintf("%s%s%s", appInfo.AppSecret, nonce, tsStr)
		sig := SHA1(str)
		fmt.Println("sig:", sig)
		if sig == signature {
			ctx.Set(CtxKey_AppKey, appKey)
		} else {
			ctx.JSON(http.StatusForbidden, models.DefaultErr("Forbidden."))
			ctx.Abort()
			return
		}
	} else {
		ctx.JSON(http.StatusBadRequest, models.DefaultErr("No AppKey."))
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
