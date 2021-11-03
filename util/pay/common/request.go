package common

// Charge 支付参数 -公共
type Charge struct {
	TradeNo     string  `json:"tradeNo,omitempty"`     // 支付单号，本系统记录，传给企业微信
	UserID      string  `json:"userId,omitempty"`      // 用户id
	PayMethod   int64   `json:"payMethod,omitempty"`   // 支付方式，参考 constant
	MoneyFee    float64 `json:"MoneyFee,omitempty"`    // 支付金额，分
	CallbackURL string  `json:"callbackURL,omitempty"` // 回调地址
	Describe    string  `json:"describe,omitempty"`    //产品描述

	// 微信
	ChargeWechat

	// 支付宝
	ChargeAli
}

type ChargeWechat struct {
	OpenID string `json:"openid,omitempty"` // 微信 openid
}

type ChargeAli struct {
	ProductCode string `json:"product_code,omitempty"` // 销售产品码，商家和支付宝签约的产品码
	QuitUrl     string `json:"quit_url,omitempty"`     // 用户付款中途退出返回商户网站的地址
}

// 查询订单 - 公共
type QueryOrder struct {
	TradeNo string `json:"tradeNo,omitempty"` // 本系统生成的订单号
}

// 退款 - 微信
type WxRefundReq struct {
	TradeNo     string  `json:"tradeNo,omitempty"`     // 支付单号，本系统记录，传给企业微信
	RefundNo    string  `json:"refund_no,omitempty"`   // 商户系统内部的退款单号
	MoneyFee    float64 `json:"MoneyFee,omitempty"`    // 支付金额，分
	RefundFee   float64 `json:"MoneyFee,omitempty"`    // 退款金额，分
	CallbackURL string  `json:"callbackURL,omitempty"` // 回调地址
}

// 退款 - 阿里
type AliRefundReq struct {
	TradeNo   string  `json:"tradeNo,omitempty"`   // 支付单号，本系统记录
	RefundNo  string  `json:"refund_no,omitempty"` // 退款单号，本系统记录
	RefundFee float64 `json:"MoneyFee,omitempty"`  // 退款金额，分
}

// 退款查询 - 微信
type WxRefundQueryReq struct {
	RefundNo string `json:"refund_no,omitempty"` // 商户系统内部的退款单号
}

// 退款查询 - 阿里
type AliRefundQueryReq struct {
	TradeNo  string `json:"tradeNo,omitempty"`   // 支付单号，本系统记录
	RefundNo string `json:"refund_no,omitempty"` // 商户系统内部的退款单号
}
