package apisix

import (
	"github.com/guonaihong/gout"
	"github.com/jettjia/go-micro-frame/util/methodset"
	"strconv"
)

// 官方文档
// https://apisix.apache.org/zh/docs/apisix/admin-api/
// https://apisix.apache.org/zh/docs/apisix/plugins/limit-req
// https://apisix.apache.org/zh/docs/apisix/plugins/api-breaker/

// CreateRouter 创建、修改路由
func (a *Apisix) CreateRouter() (err error) {
	id := methodset.GenUniqueStringToInt(a.Host + strconv.Itoa(a.Port))
	url := a.ApisixUrlPrefix() + UrlCreateRouter + "/" + strconv.Itoa(id)
	rsp := RspBody{}
	header := RspHeader{}

	err = gout.
		PUT(url).
		Debug(true).
		SetHeader(gout.H{"X-API-KEY": a.ApisixToken}).
		SetJSON(gout.H{"uri": "/" + a.ServiceName + "/*",
			"upstream": gout.H{"type": "roundrobin",
				"nodes": gout.H{a.Host + strconv.Itoa(a.Port): 1}}}).
		BindJSON(&rsp).
		BindHeader(&header).
		Do()

	return
}

// DeleteRouter 删除路由
func (a *Apisix) DeleteRouter() (err error) {
	id := methodset.GenUniqueStringToInt(a.Host + strconv.Itoa(a.Port))
	url := a.ApisixUrlPrefix() + UrlCreateRouter + "/" + strconv.Itoa(id)
	rsp := RspBody{}
	header := RspHeader{}

	err = gout.
		DELETE(url).
		Debug(true).
		SetHeader(gout.H{"X-API-KEY": a.ApisixToken}).
		BindJSON(&rsp).
		BindHeader(&header).
		Do()

	return
}
