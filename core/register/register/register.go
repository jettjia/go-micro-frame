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
	Register() error
	DelRegister() error
	Discovery() (*grpc.ClientConn, error)
}

func NewRegClient(r Reg) RegClient {
	return &Reg{
		Typ:         r.Typ,
		Host:        r.Host,
		Port:        r.Port,
		Namespace:   r.Namespace,
		User:        r.User,
		Password:    r.Password,
		ServiceHost: r.ServiceHost,
		ServicePort: r.ServicePort,
		ServiceName: r.ServiceName,
		GroupName:   r.GroupName,
		Weight:      r.Weight,
		Tags:        r.Tags,
	}
}

// 注册服务
func (r *Reg) Register() (err error) {
	if r.Typ == "nacos" {
		client := nacos.NewRegistryClient(r.Host, r.Port, r.Namespace, r.User, r.Password)
		err = client.Register(r.ServiceHost, r.ServicePort, r.ServiceName, r.GroupName, r.Weight)
	}

	if r.Typ == "consul" {
		client := consul.NewRegistryClient(r.Host, int(r.Port))
		err = client.Register(r.ServiceHost, int(r.ServicePort), r.ServiceName, r.Tags, fmt.Sprintf("%s", uuid.NewV4()))
	}

	return err
}

func (r *Reg) DelRegister() (err error) {
	if r.Typ == "nacos" {
		client := nacos.NewRegistryClient(r.Host, r.Port, r.Namespace, r.User, r.Password)
		err = client.DelRegister(r.ServiceHost, r.ServicePort, r.ServiceName, r.GroupName)
	}

	if r.Typ == "consul" {
		client := consul.NewRegistryClient(r.Host, int(r.Port))
		err = client.DelRegister(r.ServiceId)
	}

	return
}

func (r *Reg) Discovery() (grpcClient *grpc.ClientConn, err error) {
	if r.Typ == "nacos" {
		client := nacos.NewRegistryClient(r.Host, r.Port, r.Namespace, r.User, r.Password)
		grpcClient, err = client.Discovery(r.ServiceName, r.GroupName)
	}

	if r.Typ == "consul" {
		client := consul.NewRegistryClient(r.Host, int(r.Port))
		client.Discovery(r.ServiceHost, int(r.ServicePort), r.ServiceName)
	}

	return grpcClient, err
}
