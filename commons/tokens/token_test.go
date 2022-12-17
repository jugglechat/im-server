package tokens

import (
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	token := ImToken{
		AppKey:    "appkey",
		UserId:    "userid",
		DeviceId:  "deviceid",
		TokenTime: time.Now().UnixMilli(),
	}
	secureKey := []byte("abcdefghijklmnop")
	tokenStr, err := token.ToTokenString(secureKey)

	if err == nil {
		tokenWrap, err := ParseTokenString(tokenStr)
		if err == nil {
			newToken, err := ParseToken(tokenWrap, secureKey)
			if err == nil {
				if newToken.UserId == token.UserId {
					return
				}
			}
		}
	}
	t.Error("Failed.")
}
