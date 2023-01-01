package clusters

import (
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"

	"google.golang.org/protobuf/proto"
)

func BaseProcessActor(actor actorsystem.IUntypedActor) actorsystem.IUntypedActor {
	return &baseProcessActor{exeActor: actor}
}

type IContextHandler interface {
	SetContext(ctx BaseContext)
}

type baseProcessActor struct {
	exeActor actorsystem.IUntypedActor
}

type BaseActor struct {
	actorsystem.UntypedActor
	Context BaseContext
}

type BaseContext struct {
	SeqIndex     int
	AppKey       string
	Qos          int
	Session      string
	Method       string
	SourceMethod string
	RequesterId  string
	TargetId     string
	PublishType  int
	IsFromApi    bool

	ExtParams map[string]string
}

const (
	CtxKey_SdkVersion    string = "CtxKey_SdkVersion"
	CtxKey_DeviceId      string = "CtxKey_DeviceId"
	CtxKey_ClientOs      string = "CtxKey_ClientOs"
	CtxKey_ClientAddress string = "CtxKey_ClientAddress"
	CtxKey_PackageName   string = "CtxKey_PackageName"
	CtxKey_TerminalCount string = "CtxKey_TerminalCount"
)

func (actor *BaseActor) SetContext(ctx BaseContext) {
	actor.Context = ctx
}

func (actor *baseProcessActor) CreateInputObj() proto.Message {
	return &pbobjs.RpcMessageWraper{}
}

func (actor *baseProcessActor) OnReceive(input proto.Message) {
	var err error
	if input != nil {
		ssRequest, ok := input.(*pbobjs.RpcMessageWraper)
		if ok {
			ctx := BaseContext{
				SeqIndex:     int(ssRequest.ReqIndex),
				AppKey:       ssRequest.AppKey,
				Qos:          int(ssRequest.Qos),
				Session:      ssRequest.Session,
				Method:       ssRequest.Method,
				SourceMethod: ssRequest.SourceMethod,
				RequesterId:  ssRequest.RequesterId,
				TargetId:     ssRequest.TargetId,
				PublishType:  int(ssRequest.PublishType),
				IsFromApi:    ssRequest.IsFromApi,
				ExtParams:    ssRequest.ExtParams,
			}

			ctxHandler, ok := actor.exeActor.(IContextHandler)
			if ok {
				ctxHandler.SetContext(ctx)
			}

			msgBytes := ssRequest.AppDataBytes
			createInputHandler, ok := actor.exeActor.(actorsystem.ICreateInputHandler)
			if ok {
				msg := createInputHandler.CreateInputObj()
				err = tools.PbUnMarshal(msgBytes, msg)
				if err == nil {
					receiveHandler, ok := actor.exeActor.(actorsystem.IReceiveHandler)
					if ok {
						receiveHandler.OnReceive(msg)
					}
				}
			}
		}
	}
}

func (actor *baseProcessActor) SetSender(sender actorsystem.ActorRef) {
	senderHandler, ok := actor.exeActor.(actorsystem.ISenderHandler)
	if ok {
		senderHandler.SetSender(sender)
	}
}
func (actor *baseProcessActor) SetSelf(self actorsystem.ActorRef) {
	selfHandler, ok := actor.exeActor.(actorsystem.ISelfHandler)
	if ok {
		selfHandler.SetSelf(self)
	}
}

func (actor *baseProcessActor) OnTimeout() {
	timeoutHandler, ok := actor.exeActor.(actorsystem.ITimeoutHandler)
	if ok {
		timeoutHandler.OnTimeout()
	}
}
