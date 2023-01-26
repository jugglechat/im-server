package clusters

import (
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/jugglechat/im-server/commons/tools"

	"google.golang.org/protobuf/proto"
)

func CreateQueryAckWraper(code int32, respMsg proto.Message, ctx BaseContext) *pbobjs.RpcMessageWraper {
	queryAck := &pbobjs.RpcMessageWraper{
		RpcMsgType: pbobjs.RpcMsgType_QueryAck,
		ResultCode: code,
	}

	bs, _ := tools.PbMarshal(respMsg)

	queryAck.AppDataBytes = bs
	handleBaseContext(queryAck, ctx)

	return queryAck
}

func CreateServerPubWraper(requesterId, targetId, method string, msg proto.Message, ctx BaseContext) *pbobjs.RpcMessageWraper {
	serverPub := &pbobjs.RpcMessageWraper{
		RpcMsgType: pbobjs.RpcMsgType_ServerPub,
		Qos:        1,
	}
	bs, _ := tools.PbMarshal(msg)
	serverPub.AppDataBytes = bs
	handleBaseContext(serverPub, ctx)
	serverPub.RequesterId = requesterId
	serverPub.TargetId = targetId
	serverPub.Method = method
	return serverPub
}
func CreateUserPubAckWraper(code int, msgId string, msgSendTime int64, ctx BaseContext) *pbobjs.RpcMessageWraper {
	userPubAck := &pbobjs.RpcMessageWraper{
		RpcMsgType:  pbobjs.RpcMsgType_UserPubAck,
		ResultCode:  int32(code),
		MsgId:       msgId,
		MsgSendTime: msgSendTime,
	}
	handleBaseContext(userPubAck, ctx)
	return userPubAck
}

func handleBaseContext(rpcMsg *pbobjs.RpcMessageWraper, ctx BaseContext) {
	rpcMsg.ReqIndex = int32(ctx.SeqIndex)
	rpcMsg.AppKey = ctx.AppKey
	rpcMsg.Qos = int32(ctx.Qos)
	rpcMsg.Session = ctx.Session
	rpcMsg.Method = ctx.Method
	rpcMsg.SourceMethod = ctx.SourceMethod
	rpcMsg.RequesterId = ctx.RequesterId
	rpcMsg.TargetId = ctx.TargetId
	rpcMsg.PublishType = int32(ctx.PublishType)
	rpcMsg.IsFromApi = ctx.IsFromApi
	rpcMsg.ExtParams = ctx.ExtParams
}
