package register

import (
	"encoding/json"
	"fmt"
	"github.com/jettjia/go-micro-frame/core/config/nacos"
	"testing"

	"github.com/spf13/viper"
)

type NacosConfig struct {
	Host      string `mapstructure:"host" `
	Port      uint64 `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

// 解析nacos中的配置
type ServerConfig struct {
	Name         string         `mapstructure:"name" json:"name"`
	Host         string         `mapstructure:"host" json:"host"`
	Port         uint64         `mapstructure:"port" json:"port"`
	Env          string         `mapstructure:"env" json:"env"`
	Tags         []string       `mapstructure:"tags" json:"tags"`
	RegisterInfo RegisterConfig `mapstructure:"register" json:"register"`
}
type RegisterConfig struct {
	Typ       string
	Host      string // nacos,consul: Host
	Port      uint64 // nacos,consul: Port
	Namespace string // nacos
	User      string // nacos
	Password  string // nacos

	ServiceHost string   //ServiceHost
	ServicePort uint64   //ServicePort
	ServiceName string   //ServiceName
	GroupName   string   //nacos
	Weight      float64  //nacos
	Tags        []string //consul
	ServiceId   string   //consul
}

// 从nacos读取配置
func Test_nacos(t *testing.T) {
	configFileName := fmt.Sprintf("nacos-config.yaml")

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var nConfig NacosConfig
	if err := v.Unmarshal(&nConfig); err != nil {
		panic(err)
	}

	// 读取配置
	c := nacos.NewConfigClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)
	content, err := c.GetConfig(nConfig.DataId, nConfig.Group)
	if err != nil {
		panic(err)
	}
	fmt.Println("=========================")
	//fmt.Printf("%+v", content)
	fmt.Println("=========================")

	var serverConfig ServerConfig
	err = json.Unmarshal([]byte(content), &serverConfig)

	fmt.Printf("%+v", serverConfig)
}

// 注册服务
func Test_Register(t *testing.T) {
	configFileName := fmt.Sprintf("nacos-config.yaml")

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var nConfig NacosConfig
	if err := v.Unmarshal(&nConfig); err != nil {
		panic(err)
	}

	// 读取配置
	c := nacos.NewConfigClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)
	content, err := c.GetConfig(nConfig.DataId, nConfig.Group)
	if err != nil {
		panic(err)
	}

	var serverConfig ServerConfig
	err = json.Unmarshal([]byte(content), &serverConfig)

	// 注册服务
	client := NewRegClient(Reg(serverConfig.RegisterInfo))
	err = client.Register()
	if err != nil {
		fmt.Println("注册服务失败")
	}
}

// 发现服务
func Test_Discovery(t *testing.T) {
	configFileName := fmt.Sprintf("nacos-config.yaml")

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var nConfig NacosConfig
	if err := v.Unmarshal(&nConfig); err != nil {
		panic(err)
	}

	// 读取配置
	c := nacos.NewConfigClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)
	content, err := c.GetConfig(nConfig.DataId, nConfig.Group)
	if err != nil {
		panic(err)
	}

	var serverConfig ServerConfig
	err = json.Unmarshal([]byte(content), &serverConfig)

	// 注册服务
	client := NewRegClient(Reg(serverConfig.RegisterInfo))
	grpcClient, err := client.Discovery()
	if err != nil {
		fmt.Println("发现服务失败")
	}
	fmt.Println(grpcClient)
}

// 删除服务
func Test_DelRegister(t *testing.T) {
	configFileName := fmt.Sprintf("nacos-config.yaml")

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	var nConfig NacosConfig
	if err := v.Unmarshal(&nConfig); err != nil {
		panic(err)
	}

	// 读取配置
	c := nacos.NewConfigClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)
	content, err := c.GetConfig(nConfig.DataId, nConfig.Group)
	if err != nil {
		panic(err)
	}

	var serverConfig ServerConfig
	err = json.Unmarshal([]byte(content), &serverConfig)

	// 注册服务
	client := NewRegClient(Reg(serverConfig.RegisterInfo))
	err = client.DelRegister()
	if err != nil {
		fmt.Println("删除服务失败")
	}
}
