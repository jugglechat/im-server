package managers

import (
	"hash/crc32"
	"strings"
	"sync"

	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

var OnlineUserConnectMap sync.Map    //map[useridentifier]map[session]netty.HandlerContext
var OnlineSessionConnectMap sync.Map // map[session]netty.HandlerContext
var lockArray [512]*sync.Mutex

func init() {
	for i := 0; i < 512; i++ {
		lockArray[i] = &sync.Mutex{}
	}
}
func GetConnectCtxBySession(session string) netty.HandlerContext {
	if obj, ok := OnlineSessionConnectMap.Load(session); ok {
		ctx := obj.(netty.HandlerContext)
		return ctx
	}
	return nil
}
func GetConnectCtxByUser(appkey, userid string) map[string]netty.HandlerContext {
	identifier := getUserIdentifier(appkey, userid)
	if ctxMapObj, ok := OnlineUserConnectMap.Load(identifier); ok {
		ctxMap := ctxMapObj.(map[string]netty.HandlerContext)
		return ctxMap
	}
	return map[string]netty.HandlerContext{}
}
func PutInContextCache(ctx netty.HandlerContext) {
	session := utils.GetConnSession(ctx)
	if session != "" {
		OnlineSessionConnectMap.Store(session, ctx)

		appkey := utils.GetContextAttrString(ctx, utils.StateKey_Appkey)
		userid := utils.GetContextAttrString(ctx, utils.StateKey_UserID)
		identifier := getUserIdentifier(appkey, userid)

		lock := GetLock(identifier)
		lock.Lock()
		defer lock.Unlock()
		var userSessionMap map[string]netty.HandlerContext
		if tmpUserSessionMap, ok := OnlineUserConnectMap.Load(identifier); ok {
			userSessionMap = tmpUserSessionMap.(map[string]netty.HandlerContext)
		} else {
			userSessionMap = map[string]netty.HandlerContext{}
			OnlineUserConnectMap.Store(identifier, userSessionMap)
		}
		userSessionMap[session] = ctx
	}
}
func RemoveFromContextCache(ctx netty.HandlerContext) {
	session := utils.GetContextAttrString(ctx, utils.StateKey_ConnectSession)
	if session != "" {
		OnlineSessionConnectMap.Delete(session)
	}
	appkey := utils.GetContextAttrString(ctx, utils.StateKey_Appkey)
	userid := utils.GetContextAttrString(ctx, utils.StateKey_UserID)
	identifier := getUserIdentifier(appkey, userid)

	lock := GetLock(identifier)
	lock.Lock()
	defer lock.Unlock()
	if userSessionMapObj, ok := OnlineUserConnectMap.Load(identifier); ok {
		userSessionMap := userSessionMapObj.(map[string]netty.HandlerContext)
		delete(userSessionMap, session)
		if len(userSessionMap) <= 0 {
			OnlineUserConnectMap.Delete(identifier)
		}
	}
}
func GetLock(identifier string) *sync.Mutex {
	v := int(crc32.ChecksumIEEE([]byte(identifier)))
	if v < 0 {
		v = -v
	}
	return lockArray[v%512]
}
func getUserIdentifier(appkey, userid string) string {
	return strings.Join([]string{appkey, userid}, "_")
}
