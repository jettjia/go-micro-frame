package apisix

const  (
	UrlCreateRouter = "/apisix/admin/routes" //创建路由
)

// 用于解析 服务端 返回的http body
type RspBody struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
	Data    string `json:"data"`
}

// 用于解析 服务端 返回的http header
type RspHeader struct {
	Sid  string `header:"sid"`
	Time int    `header:"time"`
}
