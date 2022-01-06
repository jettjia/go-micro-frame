package register

import (
	"fmt"
	"github.com/jettjia/go-micro-frame/core/register/consul"
	"github.com/jettjia/go-micro-frame/core/register/nacos"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

type Reg struct {
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

type RegClient interface {
	Register(conf Reg) error
	DelRegister(conf Reg) error
	Discovery(conf Reg) (*grpc.ClientConn, error)
}

func NewRegClient(conf Reg) RegClient {
	return &Reg{
		Typ:         conf.Typ,
		Host:        conf.Host,
		Port:        conf.Port,
		Namespace:   conf.Namespace,
		User:        conf.User,
		Password:    conf.Password,
		ServiceHost: conf.ServiceHost,
		ServicePort: conf.ServicePort,
		ServiceName: conf.ServiceName,
		GroupName:   conf.GroupName,
		Weight:      conf.Weight,
		Tags:        conf.Tags,
	}
}

// 注册服务
func (r *Reg) Register(conf Reg) (err error) {
	if r.Typ == "nacos" {
		client := nacos.NewRegistryClient(conf.Host, conf.Port, conf.Namespace, conf.User, conf.Password)
		err = client.Register(conf.ServiceHost, conf.ServicePort, conf.ServiceName, conf.GroupName, conf.Weight)
	}

	if r.Typ == "consul" {
		client := consul.NewRegistryClient(conf.Host, int(conf.Port))
		err = client.Register(conf.ServiceHost, int(conf.ServicePort), conf.ServiceName, conf.Tags, fmt.Sprintf("%s", uuid.NewV4()))
	}

	return err
}

func (r *Reg) DelRegister(conf Reg) (err error) {
	if r.Typ == "nacos" {
		client := nacos.NewRegistryClient(conf.Host, conf.Port, conf.Namespace, conf.User, conf.Password)
		err = client.DelRegister(conf.ServiceHost, conf.ServicePort, conf.ServiceName, conf.GroupName)
	}

	if r.Typ == "consul" {
		client := consul.NewRegistryClient(conf.Host, int(conf.Port))
		err = client.DelRegister(conf.ServiceId)
	}

	return
}

func (r *Reg) Discovery(conf Reg) (grpcClient *grpc.ClientConn, err error) {
	if r.Typ == "nacos" {
		client := nacos.NewRegistryClient(conf.Host, conf.Port, conf.Namespace, conf.User, conf.Password)
		grpcClient, err = client.Discovery(conf.ServiceName, conf.GroupName)
	}

	if r.Typ == "consul" {
		client := consul.NewRegistryClient(conf.Host, int(conf.Port))
		client.Discovery(conf.ServiceHost, int(conf.ServicePort), conf.ServiceName)
	}

	return grpcClient, err
}
