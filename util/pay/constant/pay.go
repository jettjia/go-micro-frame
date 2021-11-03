package constant

type PayType int64
type OrderNoType uint8

const (
	Success = 0

	Pay_Wechat_Mini PayType = 11 // 小程序支付, jsapi支付
	Pay_Wechat_App  PayType = 12 // app支付
	Pay_Wechat_Web  PayType = 13 // native 生成支付二维码
	Pay_Wechat_H5   PayType = 14 // h5支付

	Pay_Ali_Wap PayType = 21 // 手机网站支付
	Pay_Ali_App PayType = 22 // app支付
	Pay_Ali_Pc  PayType = 23 // 电脑网站支付

	// 订单号类型，1-微信订单号，2-商户订单号，3-微信侧回跳到商户前端时用于查单的单据查询id（查询支付分订单中会使用）
	TransactionId OrderNoType = 1
	OutTradeNo    OrderNoType = 2
	QueryId       OrderNoType = 3

	RSA                       = "RSA"
	RSA2                      = "RSA2"
	UTF8                      = "utf-8"
)
