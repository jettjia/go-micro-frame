package alipay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"

	"github.com/jettjia/go-micro-frame/util/pay/common"
	utilLocal "github.com/jettjia/go-micro-frame/util/pay/util"
)

var defaultAliPcClient *AliPcClient

type AliPcClient struct {
	*AliClient
}

func InitAliPcClient(c *AliPcClient) {
	defaultAliPcClient = c
}

// DefaultAliPcClient 得到默认支付宝Wap客户端
func DefaultAliPcClient() *AliPcClient {
	return defaultAliPcClient
}

// alipay.trade.page.pay(统一收单下单并支付页面接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
func (a AliPcClient) Pay(charge *common.Charge) (map[string]string, error) {
	client, err := alipay.NewClient(a.AppID, a.PrivateKey, false)
	if err != nil {
		return nil, err
	}

	client.SetReturnUrl(charge.CallbackURL).SetNotifyUrl(charge.CallbackURL)

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", utilLocal.TruncatedText(charge.Describe, 256)).
		Set("out_trade_no", charge.TradeNo).
		Set("total_amount", utilLocal.AliyunMoneyFeeToString(charge.MoneyFee)).
		Set("product_code", charge.ProductCode)

	// 电脑网站支付请求
	payUrl, err := client.TradePagePay(context.TODO(), bm)
	if err != nil {
		return nil, err
	}
	xlog.Debug("payUrl:", payUrl)

	// todo
	// 支付结果

	return nil, nil
}
