// wechat 微信支付测试
package wechat

import (
	"fmt"
	"testing"

	"github.com/jettjia/go-micro-frame/util/pay"
	"github.com/jettjia/go-micro-frame/util/pay/client/wechat"
	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

// 小程序支付
func Test_Wechat_Mini(t *testing.T) {
	initMiniClient()

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

func initMiniClient() {
	wechat.InitWxMiniProgramClient(&wechat.WechatMiniProgramClient{
		WechatClient: &wechat.WechatClient{
			AppID:      appID,
			MchID:      mchID,
			Key:        key,
			IsProd:     isProd,
			SerialNo:   serialNo,
			PrivateKey: privateKeyContent,
		},
	})
}
