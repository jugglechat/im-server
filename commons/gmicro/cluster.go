package gmicro

import (
	"time"

	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"google.golang.org/protobuf/proto"
)

type Cluster struct {
	Name         string
	currentNode  *Node
	actorSystem  *actorsystem.ActorSystem
	nodesManager *NodesManager
	isSingle     bool
}

type IActorRegister interface {
	RegisterActor(method string, actorCreateFun func() actorsystem.IUntypedActor, concurrentCount int)
}

func NewSingleCluster(clustername, nodename string) *Cluster {
	actorSystem := actorsystem.NewActorSystemNoRpc(nodename)
	//current Node
	curNode := NewNode(nodename, actorSystem.Host, actorSystem.Prot)
	cluster := &Cluster{
		Name:        clustername,
		currentNode: curNode,
		actorSystem: actorSystem,
		isSingle:    true,
	}
	return cluster
}

func NewCluster(clustername, nodename, host string, port int, zkAddress []string) *Cluster {
	actorSystem := actorsystem.NewActorSystem(nodename, host, port)

	//current Node
	curNode := NewNode(nodename, host, port)
	//start nodesmanager
	nodesManager, _ := NewNodesManager("/gmicro/clusters/"+clustername, zkAddress)
	cluster := &Cluster{
		Name:         clustername,
		currentNode:  curNode,
		actorSystem:  actorSystem,
		nodesManager: nodesManager,
	}
	return cluster
}

func (cluster *Cluster) RegisterActor(method string, actorCreateFun func() actorsystem.IUntypedActor, concurrentCount int) {
	cluster.actorSystem.RegisterActor(method, actorCreateFun, concurrentCount)
	cluster.currentNode.AddMethod(method)
}

func (cluster *Cluster) Startup() {
	if !cluster.isSingle {
		cluster.nodesManager.RegisterSelf2ZK(*cluster.currentNode)
	}
}

func (cluster *Cluster) Shutdown() {
	if !cluster.isSingle {
		cluster.nodesManager.Destroy()
	}
}

func (cluster *Cluster) getTargetNode(method, targetId string) *Node {
	if cluster.isSingle {
		return cluster.currentNode
	} else {
		return cluster.nodesManager.GetTargetNode(method, targetId)
	}
}

func (cluster *Cluster) getNodeList(method string) []*Node {
	if cluster.isSingle {
		return []*Node{
			cluster.currentNode,
		}
	} else {
		return []*Node{}
	}
}

func (cluster *Cluster) LocalActorOf(method string) actorsystem.ActorRef {
	return cluster.actorSystem.LocalActorOf(method)
}

func (cluster *Cluster) ActorOf(host string, port int, method string) actorsystem.ActorRef {
	return cluster.actorSystem.ActerOf(host, port, method)
}

func (cluster *Cluster) CallbackActorOf(ttl time.Duration, actor actorsystem.ICallbackUntypedActor) actorsystem.ActorRef {
	return cluster.actorSystem.CallbackActerOf(ttl, actor)
}

func (cluster *Cluster) UnicastRouteWithNoSender(method, targetId string, obj proto.Message) {
	cluster.UnicastRoute(method, targetId, obj, actorsystem.NoSender)
}

func (cluster *Cluster) UnicastRoute(method, targetId string, obj proto.Message, sender actorsystem.ActorRef) {
	nod := cluster.getTargetNode(method, targetId)
	if nod != nil {
		cluster.baseRoute(method, nod.Ip, nod.Port, obj, sender)
	}
}

func (cluster *Cluster) baseRoute(method string, host string, port int, obj proto.Message, sender actorsystem.ActorRef) {
	actor := cluster.actorSystem.ActerOf(host, port, method)
	actor.Tell(obj, sender)
}

func (cluster *Cluster) BroadcastWithNoSender(method string, obj proto.Message) {
	cluster.BroadcastRoute(method, obj, actorsystem.NoSender)
}

func (cluster *Cluster) BroadcastRoute(method string, obj proto.Message, sender actorsystem.ActorRef) {
	nodes := cluster.getNodeList(method)
	for _, node := range nodes {
		cluster.baseRoute(method, node.Ip, node.Port, obj, sender)
	}
}
