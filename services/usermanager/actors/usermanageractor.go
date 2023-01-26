package actors

import (
	"fmt"
	"time"

	"github.com/jugglechat/im-server/commons/caches"
	"github.com/jugglechat/im-server/commons/clusters"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tokens"
	"github.com/jugglechat/im-server/services/usermanager/services"
	"google.golang.org/protobuf/proto"
)

type UserManagerActor struct {
	clusters.BaseActor
}

func (actor *UserManagerActor) OnReceive(input proto.Message) {
	if req, ok := input.(*pbobjs.UserRegReq); ok {
		token := tokens.ImToken{
			AppKey:    actor.Context.AppKey,
			UserId:    req.UserId,
			DeviceId:  "",
			TokenTime: time.Now().UnixMilli(),
		}
		appInfo := caches.GetAppInfo(token.AppKey)
		if appInfo != nil {
			tokenStr, _ := token.ToTokenString([]byte(appInfo.AppSecureKey))
			queryAck := clusters.CreateQueryAckWraper(0, &pbobjs.UserReqResp{
				UserId: req.UserId,
				Token:  tokenStr,
			}, actor.Context)
			actor.Sender.Tell(queryAck, actorsystem.NoSender)
		}
		services.AddUser(actor.Context, req.UserId, req.Nickname, req.UserPortrait)
	} else {
		fmt.Println("err:", input)
	}
}

func (actor *UserManagerActor) CreateInputObj() proto.Message {
	return &pbobjs.UserRegReq{}
}
