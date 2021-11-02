package client

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"

	"github.com/jettjia/go-micro-frame/util/pay/common"
	utilLocal "github.com/jettjia/go-micro-frame/util/pay/util"
)

var defaultAliWapClient *AliWapClient

type AliWapClient struct {
	AppID string // 应用ID

	PrivateKey string
}

func InitAliWapClient(c *AliWapClient) {
	defaultAliWapClient = c
}

// DefaultAliWapClient 得到默认支付宝Wap客户端
func DefaultAliWapClient() *AliWapClient {
	return defaultAliWapClient
}

// alipay.trade.wap.pay(手机网站支付接口2.0)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay
func (a AliWapClient) Pay(charge *common.Charge) (map[string]string, error) {
	// 初始化支付宝客户端
	//    WapId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(a.AppID, a.PrivateKey, false)
	if err != nil {
		return nil, err
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOn

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", utilLocal.TruncatedText(charge.Describe, 256)).
		Set("out_trade_no", charge.TradeNo).
		Set("quit_url", charge.QuitUrl).
		Set("total_amount", utilLocal.AliyunMoneyFeeToString(charge.MoneyFee)).
		Set("product_code", charge.ProductCode)

	// 手机网站支付请求
	payUrl, err := client.TradeWapPay(context.TODO(), bm)
	if err != nil {
		return nil, err
	}
	xlog.Debug("payUrl:", payUrl)

	// todo
	// 支付结果

	return nil, nil
}
