package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/jettjia/go-micro-frame/core/register/otgrpc"
)

type Registry struct {
	Host      string
	Port      uint64
	Namespace string
	User      string
	Password  string
}

type RegistryClient interface {
	Register(serviceHost string, servicePort uint64, serviceName string, groupName string, weight float64) error
	DelRegister(serviceHost string, servicePort uint64, serviceName string, groupName string) error
	Discovery(serviceName string, groupName string) (*grpc.ClientConn, error)
}

func NewRegistryClient(host string, port uint64, namespace string, user string, password string) RegistryClient {
	return &Registry{
		Host:      host,
		Port:      port,
		Namespace: namespace,
		User:      user,
		Password:  password,
	}
}

// 注册服务
func (r *Registry) Register(serviceHost string, servicePort uint64, serviceName string, group string, weight float64, ) error {
	client, err := getNamingClient(r.Host, r.Port, r.Namespace, r.User, r.Password)
	if err != nil {
		return err
	}

	param := vo.RegisterInstanceParam{
		Ip:          serviceHost,
		Port:        servicePort,
		ServiceName: serviceName,
		Weight:      weight,
		//ClusterName: cluster,
		GroupName: group,
		Enable:    true,
		Healthy:   true,
		Ephemeral: true,
	}

	_, err = client.RegisterInstance(param)
	if err != nil {
		return err
	}
	return nil
}

// 删除服务
func (r *Registry) DelRegister(serviceHost string, servicePort uint64, serviceName string, group string) error {
	client, err := getNamingClient(r.Host, r.Port, r.Namespace, r.User, r.Password)
	if err != nil {
		return err
	}

	param := vo.DeregisterInstanceParam{
		Ip:          serviceHost,
		Port:        servicePort,
		ServiceName: serviceName,
		GroupName:   group,
		Ephemeral:   true,
	}

	_, err = client.DeregisterInstance(param)
	if err != nil {
		return err
	}
	return nil
}

// 发现服务
func (r *Registry) Discovery(serviceName string, groupName string) (*grpc.ClientConn, error) {
	ins, err := r.discoveryNacos(serviceName, groupName)
	if err != nil {
		return nil, err
	}

	// nacosAddr
	nacosAddr := fmt.Sprintf("%s:%d", ins.Ip, ins.Port)
	fmt.Println(nacosAddr)
	conn, err := grpc.Dial(
		fmt.Sprintf(nacosAddr),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)
	return conn, err
}

// 获取服务, 获取一个健康的实例（加权随机轮询）
func (r *Registry) discoveryNacos(serviceName string, groupName string) (*model.Instance, error) {
	// 实例必须满足的条件：health=true,enable=true and weight>0
	namingClient, err := getNamingClient(r.Host, r.Port, r.Namespace, r.User, r.Password)
	if err != nil {
		return nil, err
	}
	instance, err := namingClient.SelectOneHealthyInstance(vo.SelectOneHealthInstanceParam{
		ServiceName: serviceName,
		GroupName:   groupName, // 默认值DEFAULT_GROUP
		//Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	})
	return instance, nil
}

// 获取注册微服务的 client
func getNamingClient(host string, port uint64, namespace string, user string, password string) (naming_client.INamingClient, error) {
	//从nacos中读取配置信息
	sc := []constant.ServerConfig{
		{
			IpAddr: host,
			Port:   port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
		Username:            user,
		Password:            password,
	}

	client, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		return nil, err
	}

	return client, nil
}
