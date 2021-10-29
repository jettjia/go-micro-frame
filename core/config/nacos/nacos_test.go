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

// 读取配置
func Test_GetClientContent(t *testing.T) {
	configFileName := fmt.Sprintf("nacos-config.yaml")

	// 读取文件配置内容
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 把内容设置到变量 NacosConfig
	var nConfig NacosConfig
	if err := v.Unmarshal(&nConfig); err != nil {
		panic(err)
	}

	// 读取配置
	c := NewConfigClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)
	content, err := c.GetConfig(nConfig.DataId, nConfig.Group)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

func Test_GetConfigClient(t *testing.T) {
	configFileName := fmt.Sprintf("nacos-config.yaml")

	// 读取文件配置内容
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 把内容设置到变量 NacosConfig
	var nConfig NacosConfig
	if err := v.Unmarshal(&nConfig); err != nil {
		panic(err)
	}

	// 读取配置
	c := NewConfigClient(nConfig.Host, nConfig.Port, nConfig.Namespace, nConfig.User, nConfig.Password)
	configClient, err := c.GetConfigClient()
	if err != nil {
		panic(err)
	}

	fmt.Println(configClient)
}
