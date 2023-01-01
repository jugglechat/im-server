package clusters

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jugglechat/im-server/commons/configures"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"google.golang.org/protobuf/proto"
)

const (
	Mod_Signal  = "signal"
	Mod_Cluster = "cluster"
)

var cluster *gmicro.Cluster

func InitCluster() error {
	clusterMod := configures.Config.ClusterMod
	if clusterMod == Mod_Cluster {
		zkAddress := strings.Split(configures.Config.Zookeeper.Address, ",")
		cluster = gmicro.NewCluster(configures.Config.ClusterName, configures.Config.NodeName, configures.Config.RpcHost, configures.Config.RpcPort, zkAddress)
	} else {
		cluster = gmicro.NewSingleCluster(configures.Config.ClusterName, configures.Config.NodeName)
	}
	return nil
}

type IRoute interface {
	GetMethod() string
	GetTargetId() string
}

func GetCluster() *gmicro.Cluster {
	return cluster
}

type ApiCallbackActor struct {
	actorsystem.UntypedActor
	respChan chan *ApiRespWraper
}
type ApiRespWraper struct {
	Msg *pbobjs.RpcMessageWraper
	Err error
}

func (actor *ApiCallbackActor) OnReceive(input proto.Message) {
	if rpcMsg, ok := input.(*pbobjs.RpcMessageWraper); ok {
		actor.respChan <- &ApiRespWraper{
			Msg: rpcMsg,
			Err: nil,
		}
		fmt.Println("recev:", rpcMsg)
	} else {
		fmt.Println("need log.")
	}
}
func (actor *ApiCallbackActor) CreateInputObj() proto.Message {
	return &pbobjs.RpcMessageWraper{}
}
func (actor *ApiCallbackActor) OnTimeout() {
	actor.respChan <- &ApiRespWraper{
		Msg: nil,
		Err: errors.New("time out1"),
	}
}

func SyncUnicastRoute(msg IRoute, ttl time.Duration) (*pbobjs.RpcMessageWraper, error) {
	respChan := make(chan *ApiRespWraper, 1)
	sender := cluster.CallbackActorOf(ttl, &ApiCallbackActor{
		respChan: respChan,
	})
	fmt.Println("method:", msg.GetMethod(), msg.GetTargetId())
	cluster.UnicastRoute(msg.GetMethod(), msg.GetTargetId(), msg.(*pbobjs.RpcMessageWraper), sender)

	select {
	case resp := <-respChan:
		return resp.Msg, resp.Err
	case <-time.After(ttl + time.Millisecond*1000):
		return nil, errors.New("time out2")
	}
}

func UnicastRouteWithCallback(msg IRoute, callbackActor actorsystem.ICallbackUntypedActor, ttl time.Duration) {
	sender := cluster.CallbackActorOf(ttl, callbackActor)
	cluster.UnicastRoute(msg.GetMethod(), msg.GetTargetId(), msg.(*pbobjs.RpcMessageWraper), sender)
}

func UnicastRoute(msg IRoute, sendMethod string) {
	sender := cluster.LocalActorOf(sendMethod)
	cluster.UnicastRoute(msg.GetMethod(), msg.GetTargetId(), msg.(*pbobjs.RpcMessageWraper), sender)
}

func UnicastRouteWithNoSender(msg IRoute) {
	cluster.UnicastRouteWithNoSender(msg.GetMethod(), msg.GetTargetId(), msg.(*pbobjs.RpcMessageWraper))
}

func UnicastRouteWithSenderActor(msg IRoute, sender actorsystem.ActorRef) {
	cluster.UnicastRoute(msg.GetMethod(), msg.GetTargetId(), msg.(*pbobjs.RpcMessageWraper), sender)
}

func Startup() {
	if cluster != nil {
		cluster.Startup()
	}
}

func Shutdown() {

}
