package alipay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"

	"github.com/jettjia/go-micro-frame/util/pay/common"
	utilLocal "github.com/jettjia/go-micro-frame/util/pay/util"
)

var defaultAliAppClient *AliAppClient

type AliAppClient struct {
	*AliClient
}

func InitAliAppClient(c *AliAppClient) {
	defaultAliAppClient = c
}

// DefaultAliAppClient 得到默认支付宝app客户端
func DefaultAliAppClient() *AliAppClient {
	return defaultAliAppClient
}

// alipay.trade.app.pay(app支付接口2.0)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
func (a AliAppClient) Pay(charge *common.Charge) (map[string]string, error) {
	client, err := alipay.NewClient(a.AppID, a.PrivateKey, false)
	if err != nil {
		return nil, err
	}

	client.SetReturnUrl(charge.CallbackURL).SetNotifyUrl(charge.CallbackURL)

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", utilLocal.TruncatedText(charge.Describe, 256)).
		Set("out_trade_no", charge.TradeNo).
		Set("total_amount", utilLocal.AliyunMoneyFeeToString(charge.MoneyFee))

	// 手机APP支付参数请求
	payParam, err := client.TradeAppPay(context.TODO(), bm)
	if err != nil {
		return nil, err
	}
	xlog.Debug("payParam:", payParam)

	// todo
	// 支付结果

	return nil, nil
}
