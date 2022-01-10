package apisix

import "strconv"

type Apisix struct {
	ApisixToken string
	ApisixHost  string
	ApisixPort  int
	Host        string // 服务 host
	Port        int    // 服务 port
	RouteName   string // 路由名称
	ServiceName string // 网关路由路径
	Rate        int    // 速率（以秒为单位）
}

type IApisix interface {
	// 路由
	CreateRouter() error // 创建、修改路由
	DeleteRouter() error // 删除路由
	RateRouter() error
}

func NewApisixClient(conf Apisix) IApisix {
	return &Apisix{
		ApisixToken: conf.ApisixToken,
		ApisixHost:  conf.ApisixHost,
		ApisixPort:  conf.ApisixPort,
		Host:        conf.Host,
		Port:        conf.Port,
		RouteName:   conf.RouteName,
		ServiceName: conf.ServiceName,
		Rate:        conf.Rate,
	}
}

// ApisixUrlPrefix 获取 apisix请求的路由前缀部分，比如：http://127.0.0.1:9080/apisix/admin/routes/1 部分的 http://127.0.0.1:9080
func (a *Apisix) ApisixUrlPrefix() string {
	return "http://" + a.ApisixHost + ":" + strconv.Itoa(a.ApisixPort)
}
