package consul

import (
	"fmt"
	"testing"

	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

func Test_RegistryConsul(t *testing.T) {
	//服务注册
	registerClient := NewRegistryClient(
		"10.4.7.71",
		8500,
	)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register("127.0.0.1", 5100, "consul_test", []string{"consul_test_tag"}, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}
}

func Test_DiscoveryConsul(t *testing.T) {
	// 服务发现
	registerClient := NewRegistryClient(
		"10.4.7.71",
		8500,
	)
	_, err := registerClient.Discovery("10.4.7.71", 51000, "gomicrom-srv")
	if err != nil {
		panic(err)
	}
}