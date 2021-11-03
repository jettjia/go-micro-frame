package common

// PayClient 支付客户端接口
type PayClient interface {
	// 用户下单付款
	Pay(charge *Charge) (map[string]string, error)
}
