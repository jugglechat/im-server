package services

import (
	"fmt"
	"time"

	"github.com/jugglechat/im-server/commons/caches"
)

type UserInfo struct {
	LastSyncTime        *int64
	LastSendBoxSyncTime *int64
	LatestMsgTime       *int64 // 最新消息时间
	TerminalNum         *int
	OnlineStatus        bool //在线状态
}

var userOnlineStatusCache *caches.LruCache

func init() {
	userOnlineStatusCache = caches.NewLruCacheWithAddReadTimeout(10000, func(key, value interface{}) {}, 5*time.Minute, 10*time.Minute)
}

func GetUserInfo(appKey, targetId string) *UserInfo {
	key := getKey(appKey, targetId)
	if exist := userOnlineStatusCache.Contains(key); exist {
		val, _ := userOnlineStatusCache.Get(key)
		return val.(*UserInfo)
	} else {
		userInfo := &UserInfo{}
		//TODO init userInfo

		val, _ := userOnlineStatusCache.AddIfAbsent(key, userInfo)
		return val.(*UserInfo)
	}
}

func getKey(appKey, targetId string) string {
	return fmt.Sprintf("%s_%s", appKey, targetId)
}
