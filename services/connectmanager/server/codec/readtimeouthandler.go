package codec

import (
	"fmt"
	"time"

	"github.com/go-netty/go-netty"
)

type ReadTimeoutHandler struct {
	timeoutMillis int64
	lastTime      int64
}

func NewReadTimeoutHandler(timeout time.Duration) *ReadTimeoutHandler {
	return &ReadTimeoutHandler{
		timeoutMillis: timeout.Milliseconds(),
		lastTime:      time.Now().UnixMilli(),
	}
}

func (handler *ReadTimeoutHandler) HandleActive(ctx netty.ActiveContext) {
	time.AfterFunc(time.Duration(handler.timeoutMillis)*time.Millisecond, func() {
		handler.checkTimeout(ctx)
	})
}

func (handler *ReadTimeoutHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	handler.lastTime = time.Now().UnixMilli()
	ctx.HandleRead(message)
}

func (handler *ReadTimeoutHandler) checkTimeout(ctx netty.HandlerContext) {
	if handler.timeoutMillis > 0 {
		currentTime := time.Now().UnixMilli()
		waitTime := handler.timeoutMillis - (currentTime - handler.lastTime)
		if waitTime > 0 {
			time.AfterFunc(time.Duration(waitTime)*time.Millisecond, func() {
				handler.checkTimeout(ctx)
			})
		} else { //timeout
			ctx.Close(fmt.Errorf("read time out"))
			// time.AfterFunc(time.Duration(handler.timeoutMillis)*time.Millisecond, func() {
			// 	handler.checkTimeout(ctx)
			// })
		}
	}
}
