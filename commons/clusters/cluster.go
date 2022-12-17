package clusters

import (
	"strings"
	"time"

	"github.com/jugglechat/im-server/commons/configures"
	"github.com/jugglechat/im-server/commons/pbdefines/pbobjs"
	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
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
