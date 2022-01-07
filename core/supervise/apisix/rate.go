package apisix

import (
	"github.com/guonaihong/gout"
	"github.com/jettjia/go-micro-frame/util/methodset"
	"strconv"
)

func (a *Apisix) RateRouter() (err error) {
	id := methodset.GenUniqueStringToInt(a.Host + strconv.Itoa(a.Port))
	url := a.ApisixUrlPrefix() + UrlCreateRouter + "/" + strconv.Itoa(id)
	rsp := RspBody{}
	header := RspHeader{}

	err = gout.
		PUT(url).
		Debug(true).
		SetHeader(gout.H{"X-API-KEY": a.ApisixToken}).
		SetJSON(
			gout.H{
				"uri": "/" + a.ServiceName + "/*",
				"plugins":
				gout.H{"limit-req":
				gout.H{
					"rate":          a.Rate,
					"burst":         0,
					"rejected_code": 503,
					"key_type":      "var",
					"key":           "remote_addr",
				},
				},
			}).
		BindJSON(&rsp).
		BindHeader(&header).
		Do()

	return
}
