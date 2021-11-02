package wechat

import (
	"context"
	"errors"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

// 关闭订单API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_3.shtml
func (w *Wechat) CloseOrder(outTradeNo string) error {
	client, err := wechat.NewClientV3(w.MchID, w.SerialNo, w.Key, w.PrivateKey)
	wxRsp, err := client.V3TransactionCloseOrder(context.TODO(), outTradeNo)
	if err != nil {
		return err
	}
	if wxRsp.Code != constant.Success {
		return errors.New("关闭订单，未知错误")
	}

	return nil
}
