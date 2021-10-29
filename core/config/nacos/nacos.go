package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type Config struct {
	Host      string
	Port      uint64
	Namespace string
	User      string
	Password  string
}

type ConfigClient interface {
	GetConfig(dataId string, groupName string) (confContent string, err error)
	GetConfigClient() (config_client.IConfigClient, error)
}

func NewConfigClient(host string, port uint64, namespace string, user string, password string) ConfigClient {
	return &Config{
		Host:      host,
		Port:      port,
		Namespace: namespace,
		User:      user,
		Password:  password,
	}
}

// 直接读取 nacos 的配置信息
func (t *Config) GetConfig(dataId string, groupName string) (confContent string, err error) {
	configClient, err := t.GetConfigClient()
	if err != nil {
		return "", err
	}

	confContent, err = configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  groupName,
	})

	if err != nil {
		return "", err
	}
	return confContent, nil
}

// 获取nacos的 IConfigClient
func (t *Config) GetConfigClient() (config_client.IConfigClient, error) {
	sc := []constant.ServerConfig{
		{
			IpAddr: t.Host,
			Port:   t.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         t.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		Username:            t.User,
		Password:            t.Password,
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		return nil, err
	}

	return configClient, nil
}
