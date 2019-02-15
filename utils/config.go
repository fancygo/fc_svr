package utils

import (
	"encoding/json"
	"github.com/fancygo/fc_svr/iface"
	"io/ioutil"
)

type Config struct {
	TcpServer iface.IServer
	Host      string
	TcpPort   int
	Name      string
	Version   string

	MaxPacketSize int
	MaxConn       int
}

var GlobalConfig *Config

func init() {
	GlobalConfig = &Config{
		Host:    "0.0.0.0",
		TcpPort: 7002,
		Name:    "fcsvr",
		Version: "v0.1",

		MaxPacketSize: 4096,
		MaxConn:       10000,
	}

	GlobalConfig.Reload()
}

func (this *Config) Reload() {
	data, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, GlobalConfig)
	if err != nil {
		panic(err)
	}
}
