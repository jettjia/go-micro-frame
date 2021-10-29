package nacos

import (
	"fmt"
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

// 注册服务
func Test_Register(t *testing.T) {
	// 获取client
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

	client := NewRegistryClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)

	err := client.Register("127.0.0.1", 8898, "go-micro-frame-srv", "dev", 10)
	if err != nil {
		panic(err)
	}
}

// 删除服务
func Test_DelRegister(t *testing.T) {
	// 获取client
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

	client := NewRegistryClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)

	err := client.DelRegister("127.0.0.1", 8898, "go-micro-frame-srv", "dev")
	if err != nil {
		panic(err)
	}
}

// 获取服务
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

	client := NewRegistryClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)

	_, err := client.Discovery("go-micro-frame-srv", "dev")
	if err != nil {
		panic(err)
	}
}