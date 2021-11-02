package wechat

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"

	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

// Refund 退款
func (w *Wechat) Refund(charge *common.RefundReq) (map[string]string, error) {
	client, err := wechat.NewClientV3(w.MchID, w.SerialNo, w.Key, w.PrivateKey)
	if err != nil {
		return nil, err
	}

	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", charge.TradeNo).
		Set("out_refund_no", charge.RefundNo).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", charge.MoneyFee).
				Set("refund", charge.RefundFee).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3Refund(context.TODO(), bm)
	if err != nil {
		return nil, err
	}

	if wxRsp.Code == constant.Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return nil, err
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	return nil, nil
}
