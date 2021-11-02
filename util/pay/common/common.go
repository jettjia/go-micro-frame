package common

// PayClient 支付客户端接口
type PayClient interface {
	// 用户下单付款
	Pay(charge *Charge) (map[string]string, error)
}

// Charge 支付参数
type Charge struct {
	APPID       string  `json:"-"`
	TradeNo     string  `json:"tradeNo,omitempty"` // 支付单号，本系统记录，传给企业微信
	Origin      string  `json:"origin,omitempty"`
	UserID      string  `json:"userId,omitempty"`
	PayMethod   int64   `json:"payMethod,omitempty"`   // 支付方式，参考 constant
	MoneyFee    float64 `json:"MoneyFee,omitempty"`    // 支付金额，分
	CallbackURL string  `json:"callbackURL,omitempty"` // 回调地址
	ShowURL     string  `json:"showURL,omitempty"`
	Describe    string  `json:"describe,omitempty"`
	OpenID      string  `json:"openid,omitempty"`
	CheckName   bool    `json:"check_name,omitempty"`
	ReUserName  string  `json:"re_user_name,omitempty"`

	ProductCode string `json:"product_code,omitempty"` // 销售产品码，商家和支付宝签约的产品码
	QuitUrl     string `json:"quit_url,omitempty"`     // 用户付款中途退出返回商户网站的地址
	// 阿里提现
	AliAccount     string `json:"ali_account"`
	AliAccountType string `json:"ali_account_type"`
}

// 退款
type RefundReq struct {
	TradeNo     string  `json:"tradeNo,omitempty"`     // 支付单号，本系统记录，传给企业微信
	RefundNo    string  `json:"refund_no,omitempty"`   // 退款单号，本系统记录，传给微信
	MoneyFee    float64 `json:"MoneyFee,omitempty"`    // 支付金额，分
	RefundFee   float64 `json:"MoneyFee,omitempty"`    // 退款金额，分
	CallbackURL string  `json:"callbackURL,omitempty"` // 回调地址
}
