package ali

import (
	"fmt"
	"github.com/jettjia/go-micro-frame/util/pay/common"
	"testing"

	"github.com/jettjia/go-micro-frame/util/pay/client/alipay"
)

// 支付订单查询
func Test_Ali_QueryOrder(t *testing.T) {
	c := alipay.NewAliClient(appID, privateKey, alipayPublicContentRSA2, appPublicContent,
		alipayRootContent, debugSwitch)

	charge := new(common.QueryOrder)
	charge.TradeNo = "3ff232" // 订单号，本系统生成的单号

	fdata, err := c.QueryOrder(charge)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}

// 退款
func Test_Ali_Refund(t *testing.T) {
	c := alipay.NewAliClient(appID, privateKey, alipayPublicContentRSA2, appPublicContent,
		alipayRootContent, debugSwitch)

	charge := new(common.AliRefundReq)
	charge.TradeNo = "3ff232" // 订单号，本系统生成的单号
	charge.RefundNo = "111"   // 退款单号
	charge.RefundFee = 100    // 退款金额

	err := c.Refund(charge)
	if err != nil {
		t.Error(err)
	}
}

// 退款查询
func Test_Ali_RefundQuery(t *testing.T) {
	c := alipay.NewAliClient(appID, privateKey, alipayPublicContentRSA2, appPublicContent,
		alipayRootContent, debugSwitch)

	charge := new(common.AliRefundQueryReq)
	charge.TradeNo = "111"  // 订单号，本系统生成
	charge.RefundNo = "zzz" // 退款单号，本系统生成的退款记录单号，由退款的时候生成

	fdata, err := c.RefundQuery(charge)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}
