package utils

import (
	"sync"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/commons/tools"
)

func InitCtxAttrByChannel(channel netty.Channel) {
	SetContextAttrByChannel(channel, StateKey_ConnectSession, tools.GenerateUUIDShort11())
	SetContextAttrByChannel(channel, StateKey_ConnectCreateTime, time.Now().UnixMilli())
	SetContextAttrByChannel(channel, StateKey_CtxLocker, &sync.Mutex{})
}

func SetContextAttr(ctx netty.HandlerContext, key string, value interface{}) {
	if ctx.Attachment() == nil {
		attMap := make(map[string]interface{})
		ctx.SetAttachment(attMap)
	}
	attMap := ctx.Attachment().(map[string]interface{})
	attMap[key] = value
	ctx.SetAttachment(attMap)
}
func SetContextAttrByChannel(channel netty.Channel, key string, value interface{}) {
	if channel.Attachment() == nil {
		attMap := make(map[string]interface{})
		channel.SetAttachment(attMap)
	}
	attMap := channel.Attachment().(map[string]interface{})
	attMap[key] = value
	channel.SetAttachment(attMap)
}
func GetContextAttr(ctx netty.HandlerContext, key string) interface{} {
	if ctx.Attachment() != nil {
		attMap := ctx.Attachment().(map[string]interface{})
		return attMap[key]
	}
	return nil
}
func GetContextAttrString(ctx netty.HandlerContext, key string) string {
	ret := GetContextAttr(ctx, key)
	if ret != nil {
		str, ok := ret.(string)
		if ok {
			return str
		}
	}
	return ""
}
func GetConnSession(ctx netty.HandlerContext) string {
	return GetContextAttrString(ctx, StateKey_ConnectSession)
}
func GetCtxLocker(ctx netty.HandlerContext) *sync.Mutex {
	obj := GetContextAttr(ctx, StateKey_CtxLocker)
	if obj == nil {
		lock := &sync.Mutex{}
		SetContextAttr(ctx, StateKey_CtxLocker, lock)
		return lock
	} else {
		return obj.(*sync.Mutex)
	}
}
func GetServerIndexAfterIncrease(ctx netty.HandlerContext) uint16 {
	lock := GetCtxLocker(ctx)
	lock.Lock()
	defer lock.Unlock()
	var index uint16 = 0
	indexObj := GetContextAttr(ctx, StateKey_ServerMsgIndex)
	if indexObj != nil {
		index = indexObj.(uint16)
	}
	index = index + 1
	SetContextAttr(ctx, StateKey_ServerMsgIndex, index)
	return index
}

func PutServerPubCallback(ctx netty.HandlerContext, index int32, callback func()) {
	lock := GetCtxLocker(ctx)
	lock.Lock()
	defer lock.Unlock()
	obj := GetContextAttr(ctx, StateKey_ServerPubCallbackMap)
	var callbackMap *sync.Map
	if obj == nil {
		callbackMap = &sync.Map{}
		SetContextAttr(ctx, StateKey_ServerPubCallbackMap, callbackMap)
	} else {
		callbackMap = obj.(*sync.Map)
	}
	callbackMap.Store(index, callback)
}

func GetAndDeleteServerPubCallback(ctx netty.HandlerContext, index int32) func() {
	obj := GetContextAttr(ctx, StateKey_ServerPubCallbackMap)
	if obj != nil {
		callbackMap := obj.(*sync.Map)
		callbackObj, ok := callbackMap.LoadAndDelete(index)
		if ok {
			callback := callbackObj.(func())
			return callback
		}
	}
	return nil
}

func RemoveServerPubCallback(ctx netty.HandlerContext, index int32) {
	obj := GetContextAttr(ctx, StateKey_ServerPubCallbackMap)
	if obj != nil {
		callbackMap := obj.(*sync.Map)
		callbackMap.Delete(index)
	}
}

func PutQueryAckCallback(ctx netty.HandlerContext, index int32, callback func()) {
	lock := GetCtxLocker(ctx)
	lock.Lock()
	defer lock.Unlock()
	obj := GetContextAttr(ctx, StateKey_QueryConfirmMap)
	var callbackMap *sync.Map
	if obj == nil {
		callbackMap = &sync.Map{}
		SetContextAttr(ctx, StateKey_QueryConfirmMap, callbackMap)
	} else {
		callbackMap = obj.(*sync.Map)
	}
	callbackMap.Store(index, callback)
}

func GetAndDeleteQueryAckCallback(ctx netty.HandlerContext, index int32) func() {
	obj := GetContextAttr(ctx, StateKey_QueryConfirmMap)
	if obj != nil {
		callbackMap := obj.(*sync.Map)
		callbackObj, ok := callbackMap.LoadAndDelete(index)
		if ok {
			callback := callbackObj.(func())
			return callback
		}
	}
	return nil
}

func RemoveQueryAckCallback(ctx netty.HandlerContext, index int32) {
	obj := GetContextAttr(ctx, StateKey_QueryConfirmMap)
	if obj != nil {
		callbackMap := obj.(*sync.Map)
		callbackMap.Delete(index)
	}
}
