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
		DbName   string `yaml:"dbName"`
	} `yaml:"mysql"`

	ConnectManager struct {
		TcpPort int `yaml:"tcpPort"`
		WsPort  int `yaml:"wsPort"`
	} `yaml:"connectManager"`
}

var Config ImConfig
var Env string

func InitConfigures() error {
	env := os.Getenv("ENV")
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
