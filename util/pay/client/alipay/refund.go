package alipay

import (
	"context"
	"fmt"

	"github.com/go-pay/gopay"
	"github.com/jettjia/go-micro-frame/util/pay/common"

	utilLocal "github.com/jettjia/go-micro-frame/util/pay/util"
)

// alipay.trade.refund(统一收单交易退款接口)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.refund
func (a *AliClient) Refund(refundReq *common.AliRefundReq) error {
	client, err := a.InitAliClient()
	if err != nil {
		return err
	}

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", refundReq.TradeNo).
		Set("refund_amount", utilLocal.AliyunMoneyFeeToString(refundReq.RefundFee)).
		Set("out_request_no", refundReq.RefundNo)

	// 发起退款请求
	aliRsp, err := client.TradeRefund(context.TODO(), bm)
	if err != nil {
		return err
	}

	fmt.Println(aliRsp) // todo, 这里需要处理返回结果
	return nil
}

// RefundQuery 退款结果查询
// alipay.trade.fastpay.refund.query(统一收单交易退款查询)
//	文档地址：https://opendocs.alipay.com/apis/api_1/alipay.trade.fastpay.refund.query
func (a *AliClient) RefundQuery(req *common.AliRefundQueryReq) (map[string]string, error) {
	client, err := a.InitAliClient()
	if err != nil {
		return nil, err
	}

	// 请求参数
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", req.TradeNo).
		Set("out_request_no", req.RefundNo)

	// 发起退款查询请求
	aliRsp, err := client.TradeFastPayRefundQuery(context.TODO(), bm)
	if err != nil {
		return nil, err
	}

	fmt.Println(aliRsp)
	// todo 这里需要把请求结果返回
	return nil, nil
}
