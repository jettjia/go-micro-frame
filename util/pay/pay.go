package pay

import (
	"github.com/jettjia/go-micro-frame/util/pay/client/alipay"
	"github.com/jettjia/go-micro-frame/util/pay/client/wechat"
	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

// 用户下单支付接口
func Pay(charge *common.Charge) (map[string]string, error) {
	ct := getPayType(charge.PayMethod)
	re, err := ct.Pay(charge)
	return re, err
}

// getPayType 得到需要支付的类型
func getPayType(payMethod int64) common.PayClient {
	//如果使用余额支付
	switch payMethod {
	case int64(constant.Pay_Ali_Wap):
		return alipay.DefaultAliWapClient()
	case int64(constant.Pay_Ali_App):
		return alipay.DefaultAliAppClient()
	case int64(constant.Pay_Ali_Pc):
		return alipay.DefaultAliPcClient()

	case int64(constant.Pay_Wechat_Web):
		return wechat.DefaultWechatWebClient()
	case int64(constant.Pay_Wechat_App):
		return wechat.DefaultWechatAppClient()
	case int64(constant.Pay_Wechat_Mini):
		return wechat.DefaultWechatMiniProgramClient()
	case int64(constant.Pay_Wechat_H5):
		return wechat.DefaultWechatMiniProgramClient()
	}
	return nil
}
