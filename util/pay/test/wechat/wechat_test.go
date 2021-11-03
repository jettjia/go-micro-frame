package wechat

import (
	"fmt"
	"testing"

	"github.com/jettjia/go-micro-frame/util/pay/client/wechat"
	"github.com/jettjia/go-micro-frame/util/pay/common"
)

// 支付订单查询
func Test_Wechat_QueryOrder(t *testing.T) {
	c := wechat.NewWechatClient(appID, mchID, key, serialNo, privateKeyContent, isProd)

	tradeNo := "3ff232" // 本系统单号

	fdata, err := c.QueryOrder(tradeNo)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}

// 退款
func Test_Wechat_Refund(t *testing.T) {
	c := wechat.NewWechatClient(appID, mchID, key, serialNo, privateKeyContent, isProd)

	charge := new(common.WxRefundReq)
	charge.TradeNo = "wwwwwwwwwww" // 订单号，本系统生成的单号
	charge.RefundNo = "zzz"        // 退款单号，本系统生成的退款记录单号
	charge.MoneyFee = 100          //订单金额
	charge.RefundFee = 100         // 退款金额

	fdata, err := c.Refund(charge)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}