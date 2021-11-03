// ali-wap
package ali

import (
	"fmt"
	"testing"

	"github.com/jettjia/go-micro-frame/util/pay"
	"github.com/jettjia/go-micro-frame/util/pay/client/alipay"
	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

func Test_Ali_Wap(t *testing.T) {
	initWapClient()

	charge := new(common.Charge)
	charge.PayMethod = int64(constant.Pay_Ali_Wap)
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

func initWapClient() {
	alipay.InitAliWapClient(&alipay.AliWapClient{
		AliClient: &alipay.AliClient{
			AppID:                   appID,                   // 应用ID
			PrivateKey:              privateKey,              // 私钥
			PublicKey:               publicKey,               // 公钥
			AlipayPublicContentRSA2: alipayPublicContentRSA2, // 支付宝公钥证书
			AppPublicContent:        appPublicContent,        // 应用公钥证书
			AlipayRootContent:       alipayRootContent,       // 支付宝根证书
			DebugSwitch:             debugSwitch,             // 日志开启，1开0关
		},
	})
}
