package wechat

import (
	"context"
	"errors"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
	utilLocal "github.com/jettjia/go-micro-frame/util/pay/util"
	"time"

	"github.com/jettjia/go-micro-frame/util/pay/common"
)

var defaultWechatAppClient *WechatAppClient

func InitWxAppClient(c *WechatAppClient) {
	defaultWechatAppClient = c
}

func DefaultWechatAppClient() *WechatAppClient {
	return defaultWechatAppClient
}

// WechatAppClient app支付
type WechatAppClient struct {
	*WechatClient
}

// APP下单API
//	Code = 0 is success
//	商户文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_2_1.shtml
func (w WechatAppClient) Pay(charge *common.Charge) (map[string]string, error) {
	client, err := w.InitNewWechatClient()
	if err != nil {
		return nil, err
	}
	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		return nil, err
	}

	//初始化参数Map
	bm := make(gopay.BodyMap)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm.Set("appid", w.AppID).
		Set("sp_appid", w.AppID).
		Set("sp_mchid", w.MchID).
		Set("sub_mchid", w.MchID).
		Set("description", utilLocal.TruncatedText(charge.Describe, 32)).
		Set("out_trade_no", charge.TradeNo).
		Set("time_expire", expire).
		Set("notify_url", charge.CallbackURL).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", charge.MoneyFee).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionApp(context.TODO(), bm)
	if err != nil {
		return nil, err
	}

	if wxRsp.Code != constant.Success {
		return nil, errors.New("未知错误")
	}
	xlog.Debugf("wxRsp: %#v", wxRsp.Response)

	// todo
	// 支付结果返回
	return nil, nil
}
