// wechat 微信支付测试
package wechat

import (
	"fmt"
	"testing"

	"github.com/jettjia/go-micro-frame/util/pay"
	"github.com/jettjia/go-micro-frame/util/pay/client"
	"github.com/jettjia/go-micro-frame/util/pay/client/wechat"
	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

var (
	appID             = "11"
	mchID             = "11"
	key               = "11"
	isProd            = true
	serialNo          = "11"
	privateKeyContent = `-----BEGIN PRIVATE KEY-----

-----END PRIVATE KEY-----
`
)

// 小程序支付
func Test_Wechat_Mini(t *testing.T) {
	initClient()

	charge := new(common.Charge)
	charge.PayMethod = int64(constant.Pay_Wechat_Mini)
	charge.MoneyFee = 100
	charge.Describe = "test-pay"
	charge.TradeNo = "111111111223"
	charge.OpenID = "oUpF8uMuAJO_M2pxb1Q9zNjWeS6o"
	charge.CallbackURL = "http://127.0.0.1/callback/wechatcallback"

	fdata, err := pay.Pay(charge)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}

func initClient() {
	client.InitWxMiniProgramClient(&client.WechatMiniProgramClient{
		AppID:      appID,
		MchID:      mchID,
		Key:        key,
		IsProd:     isProd,
		SerialNo:   serialNo,
		PrivateKey: privateKeyContent,
	})
}

// 小程序查询订单
func Test_Wechat_QueryOrder(t *testing.T) {
	c := wechat.NewWechat(appID, mchID, key, serialNo, privateKeyContent, isProd)

	tradeNo := "3ff232" // 本系统单号

	fdata, err := c.QueryPay(tradeNo)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}

// 退款
func Test_Wechat_Refund(t *testing.T) {
	c := wechat.NewWechat(appID, mchID, key, serialNo, privateKeyContent, isProd)

	charge := new(common.RefundReq)
	charge.TradeNo = "wwwwwwwwwww"                                // 订单号，本系统生成的单号
	charge.RefundNo = "zzz"                                       // 退款单号，本系统生成的退款记录单号
	charge.MoneyFee = 100                                         //订单金额
	charge.RefundFee = 100                                        // 退款金额

	fdata, err := c.Refund(charge)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}
