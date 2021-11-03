package wechat

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

// Refund 退款
func (w *WechatClient) Refund(charge *common.WxRefundReq) (map[string]string, error) {
	client, err := w.InitNewWechatClient()
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", charge.TradeNo).
		Set("out_refund_no", charge.RefundNo).
		Set("notify_url", charge.CallbackURL).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", charge.MoneyFee).
				Set("refund", charge.RefundFee).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3Refund(context.TODO(), bm)
	if err != nil {
		return nil, err
	}

	// todo
	if wxRsp.Code == constant.Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return nil, err
	}

	return nil, nil
}

// RefundQuery 退款结果查询
// 查询单笔退款API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_10.shtml
//	服务商文档：https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_10.shtml
func (w *WechatClient) RefundQuery(charge *common.WxRefundQueryReq) (map[string]string, error) {
	client, err := w.InitNewWechatClient()
	if err != nil {
		return nil, err
	}

	wxRsp, err := client.V3RefundQuery(context.TODO(), charge.RefundNo)
	if err != nil {
		return nil, err
	}

	if wxRsp.Code == constant.Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return nil, err
	}

	// todo, 这里要处理返回结果
	return nil, nil
}
