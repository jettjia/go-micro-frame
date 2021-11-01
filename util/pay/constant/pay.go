package constant

type PayType int32

const (
	Pay_Wechat_Mini  PayType = 11 // 小程序支付
	Pay_Wechat_App   PayType = 12 // app支付
	Pay_Wechat_JsApi PayType = 13 // jsapi 支付

	Pay_Ali_Wap  PayType = 21 // 手机网站支付
	Pay_Ali_App  PayType = 22 // app支付
	Pay_Ali_Page PayType = 22 // 电脑网站支付
)
