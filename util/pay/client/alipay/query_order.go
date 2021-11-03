package alipay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
)

// alipay.trade.query(统一收单线下交易查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.query
func (a *AliClient) QueryOrder(outTradeNo string) (map[string]string, error) {
	client, err := a.InitAliClient()
	if err != nil {
		return nil, err
	}

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", outTradeNo)

	// 查询订单
	aliRsp, err := client.TradeQuery(context.TODO(), bm)
	if err != nil {
		return nil, err
	}

	// 同步返回验签
	ok, err := alipay.VerifySyncSignWithCert(a.AlipayPublicContentRSA2, aliRsp.SignData, aliRsp.Sign)
	if err != nil {
		return nil, err
	}

	if ok {
		// todo, 这里需要处理返回结果
	}

	return nil, nil
}
