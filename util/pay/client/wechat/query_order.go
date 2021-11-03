package wechat

import (
	"context"

	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"

	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

// 查询订单API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml
func (w *WechatClient) QueryOrder(outTradeNo string) (map[string]string, error) {
	client, err := w.InitNewWechatClient()
	if err != nil {
		return nil, err
	}

	wxRsp, err := client.V3TransactionQueryOrder(context.TODO(), wechat.OrderNoType(constant.OutTradeNo), outTradeNo)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	if wxRsp.Code == constant.Success {
		xlog.Debugf("wxRsp: %#v", wxRsp.Response)
		return nil, err
	}
	xlog.Errorf("wxRsp:%s", wxRsp.Error)

	// todo 查询结果返回

	return nil, nil
}
