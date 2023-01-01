package configures

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	EnvDev  = "dev"
	EnvProd = "prod"
)

type ImConfig struct {
	ClusterName string `yaml:"clusterName"`
	NodeName    string `yaml:"nodeName"`
	ClusterMod  string `yaml:"clusterMod"`
	RpcHost     string `yaml:"rpcHost"`
	RpcPort     int    `yaml:"rpcPort"`

	Log struct {
		LogPath string `yaml:"logPath"`
		LogName string `yaml:"logName"`
	} `ymal:"log"`

	Zookeeper struct {
		Address string `yaml:"address"`
	} `ymal:"zookeeper"`

	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DbName   string `yaml:"name"`
	} `yaml:"mysql"`

	ConnectManager struct {
		TcpHost string `yaml:"tcpHost"`
		TcpPort int    `yaml:"tcpPort"`
		WsHost  string `yaml:"wsHost"`
		WsPort  int    `yaml:"wsPort"`
	} `yaml:"connectManager"`

	ApiGateway struct {
		HttpPort int `yaml:"httpPort"`
	} `yaml:"apiGateway"`

	NavGateway struct {
		HttpPort int `yaml:"httpPort"`
	} `yaml:"navGateway"`
}

var Config ImConfig
var Env string

func InitConfigures() error {
	env := "dev" //os.Getenv("ENV")
	if env == "" {
		env = EnvDev
		os.Setenv("ENV", env)
	}
	cfBytes, err := ioutil.ReadFile(fmt.Sprintf("conf/config_%s.yml", env))
	if err == nil {
		var conf ImConfig
		yaml.Unmarshal(cfBytes, &conf)
		Env = env
		Config = conf
		return nil
	} else {
		return err
	}
}
