package actorsystem

import (
	"sync"

	"github.com/Jeffail/tunny"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem/rpc"
)

type IExecutor interface {
	Execute(req *rpc.RpcMessageRequest, msgSender *MsgSender)
}

type ActorExecutor struct {
	wraperChan     chan wraper
	executePool    *tunny.Pool
	actorPool      sync.Pool
	actorCreateFun func() IUntypedActor
}

func NewActorExecutor(concurrentCount int, actorCreateFun func() IUntypedActor) *ActorExecutor {
	pool := sync.Pool{
		New: func() interface{} {
			return actorCreateFun()
		},
	}
	executor := &ActorExecutor{
		wraperChan:     make(chan wraper, buffersize),
		executePool:    tunny.NewCallback(concurrentCount),
		actorPool:      pool,
		actorCreateFun: actorCreateFun,
	}
	go actorExecute(executor)
	return executor
}

func (executor *ActorExecutor) Execute(req *rpc.RpcMessageRequest, msgSender *MsgSender) {
	executor.wraperChan <- commonExecute(req, msgSender, executor.actorCreateFun())
}

func actorExecute(executor *ActorExecutor) {
	for {
		wraper := <-executor.wraperChan
		go executor.executePool.Process(func() {
			actorObj := executor.actorPool.Get()

			senderHandler, ok := actorObj.(ISenderHandler)
			if ok {
				senderHandler.SetSender(wraper.sender)
			}

			receiveHandler, ok := actorObj.(IReceiveHandler)
			if ok {
				receiveHandler.OnReceive(wraper.msg)
			}
			executor.actorPool.Put(actorObj)
		})
	}
}
