package ali

import (
	"fmt"
	"testing"

	"github.com/jettjia/go-micro-frame/util/pay/client/alipay"
)

// 支付订单查询
func Test_Ali_QueryOrder(t *testing.T) {
	c := alipay.NewAliClient(appID, privateKey, alipayPublicContentRSA2, appPublicContent,
		alipayRootContent, debugSwitch)

	tradeNo := "3ff232" // 本系统单号

	fdata, err := c.QueryOrder(tradeNo)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", fdata)
}
