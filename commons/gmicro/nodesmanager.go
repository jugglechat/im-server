package gmicro

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	//"github.com/yuwnloyblog/commonutils"
	"github.com/jugglechat/im-server/commons/gmicro/logs"
	"github.com/jugglechat/im-server/commons/gmicro/utils"
	"github.com/yuwnloyblog/commonutils"
)

type Node struct {
	Name      string          `json:"name"`
	Ip        string          `json:"ip"`
	Port      int             `json:"port"`
	Methods   []string        `json:"methods"`
	methodMap map[string]bool `json:"-"`
}

func NewNode(name, ip string, port int) *Node {
	node := &Node{
		Name:      name,
		Ip:        ip,
		Port:      port,
		Methods:   []string{},
		methodMap: make(map[string]bool),
	}
	return node
}

func (node *Node) AddMethod(method string) {
	node.methodMap[method] = true
	methodArr := make([]string, 0, len(node.methodMap))
	for method := range node.methodMap {
		methodArr = append(methodArr, method)
	}
	node.Methods = methodArr
}

type NodesManager struct {
	basePath          string
	zkConn            *zk.Conn
	nodeMap           map[string]*Node
	method2RingMapStr map[string]string
	ringMap           map[string]*commonutils.ConsistentHash
}

func NewNodesManager(basePath string, zkAddress []string) (*NodesManager, error) {
	conn, _, err := zk.Connect(zkAddress, time.Second*5)

	if err != nil {
		return nil, err
	}
	manager := &NodesManager{
		basePath: basePath,
		zkConn:   conn,
	}
	createPathIfNotExist(conn, basePath+"/nodes")
	go manager.WatchChildrensChange(basePath+"/nodes", NodesChangeListener)
	return manager, nil
}

func createPathIfNotExist(conn *zk.Conn, path string) {
	pathStr := strings.Split(path, "/")
	fmt.Println(pathStr)
	nodePath := strings.Builder{}
	for _, node := range pathStr {
		if node == "" {
			continue
		}
		nodePath.WriteString("/")
		nodePath.WriteString(node)
		tmpPath := nodePath.String()
		exist, _, _ := conn.Exists(tmpPath)
		if !exist {
			rst, err := conn.Create(tmpPath, []byte{}, 0, zk.WorldACL(zk.PermAll))
			if err == nil {
				logs.Error("Failed to create the path[", rst, "]", err)
			} else {
				break
			}
		}
	}
}

func NodesChangeListener(children []string, manager *NodesManager) {
	for _, child := range children {
		fmt.Println("Node:", child)
	}
	//reInit
	manager.initHashRingByZk()
}

func (manager *NodesManager) WatchChildrensChange(path string, listener func(children []string, manager *NodesManager)) {
	for {
		children, _, child_ch, err := manager.zkConn.ChildrenW(path)
		if err != nil {
			fmt.Println("Can not get the list of path [" + path + "]")
			time.Sleep(30 * time.Second)
		} else {
			listener(children, manager)
			<-child_ch
		}
	}
}

func (manager *NodesManager) initHashRingByNodes(nodes []*Node) {
	tmpNodeMap := make(map[string]*Node)
	tmpMethod2Nodes := make(map[string]map[string]bool)
	tmpMethod2RingMapStr := make(map[string]string)
	tmpRingMap := make(map[string]*commonutils.ConsistentHash)

	for _, node := range nodes {
		tmpNodeMap[node.Name] = node

		for _, method := range node.Methods {
			nodeNameMap := tmpMethod2Nodes[method]
			if nodeNameMap == nil {
				nodeNameMap = make(map[string]bool)
				tmpMethod2Nodes[method] = nodeNameMap
			}
			isExist := nodeNameMap[node.Name]
			if !isExist {
				nodeNameMap[node.Name] = true
			}
		}

	}

	for method, nodeNameSet := range tmpMethod2Nodes {
		nodeNameSlice := make([]string, 0, len(nodeNameSet))
		for name := range nodeNameSet {
			nodeNameSlice = append(nodeNameSlice, name)
		}
		sort.Strings(nodeNameSlice)
		namesStr := strings.Join(nodeNameSlice, ":")
		tmpMethod2RingMapStr[method] = namesStr

		hashRing := generateConsistentHashRing(tmpNodeMap, nodeNameSlice)
		tmpRingMap[namesStr] = hashRing
	}
	manager.nodeMap = tmpNodeMap
	manager.method2RingMapStr = tmpMethod2RingMapStr
	manager.ringMap = tmpRingMap
}

func (manager *NodesManager) initHashRingByZk() {
	nodesPath := manager.basePath + "/nodes"
	children, _, err := manager.zkConn.Children(nodesPath)
	if err != nil {
		log.Fatal(err)
		return
	}
	var nodes []*Node
	for _, childName := range children {
		node, err := manager.getNodeWithZk(childName)
		if err == nil {
			nodes = append(nodes, node)
		}
	}
	manager.initHashRingByNodes(nodes)
}

func generateConsistentHashRing(nodeMap map[string]*Node, names []string) *commonutils.ConsistentHash {
	hashRing := commonutils.NewConsistentHash(true)

	for _, name := range names {
		nod := nodeMap[name]
		if nod != nil {
			hashRing.Add(nod.Name, nod.Name, 1)
		}
	}
	return hashRing
}

func (manager *NodesManager) getNodeWithZk(name string) (*Node, error) {
	nodePath := manager.basePath + "/nodes/" + name
	bytes, _, err := manager.zkConn.Get(nodePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	node := &Node{}
	err = utils.JsonUnMarshal(bytes, node)
	return node, err
}

func (manager *NodesManager) GetTargetNode(method, targetId string) *Node {
	nodeNameStr := manager.method2RingMapStr[method]
	hashRing := manager.ringMap[nodeNameStr]
	if hashRing != nil {
		nod := hashRing.Get(targetId)
		if nod != nil {
			return manager.nodeMap[nod.Name]
		}
	}
	return nil
}

/**
* The path of node data:   /gmicro/clusters/{clusterName}/nodes/{node.Name}
*
**/
func (manager *NodesManager) RegisterSelf2ZK(node Node) {
	nodePath := manager.basePath + "/nodes/" + node.Name
	data, _ := utils.JsonMarshal(node)
	isExist, state, _ := manager.zkConn.Exists(nodePath)
	var version int32 = 0
	var err error
	if isExist {
		version = state.Version
		_, err = manager.zkConn.Set(nodePath, data, version)
	} else {
		// flags有4种取值：
		// 0:永久，除非手动删除
		// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
		// zk.FlagSequence  = 2:会自动在节点后面添加序号
		// 3:Ephemeral和Sequence，即，短暂且自动添加序号
		_, err = manager.zkConn.Create(nodePath, data, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
	}
	if err != nil {
		log.Fatal(err)
	}
}

/**
* destroy
**/
func (manager *NodesManager) Destroy() {
	manager.zkConn.Close()
}
