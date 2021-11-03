package wechat

import (
	"context"
	"errors"
	"time"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"

	"github.com/jettjia/go-micro-frame/util/pay/common"
	"github.com/jettjia/go-micro-frame/util/pay/constant"
	utilLocal "github.com/jettjia/go-micro-frame/util/pay/util"
)

var defaultWechatMiniProgramClient *WechatMiniProgramClient

func InitWxMiniProgramClient(c *WechatMiniProgramClient) {
	defaultWechatMiniProgramClient = c
}

func DefaultWechatMiniProgramClient() *WechatMiniProgramClient {
	return defaultWechatMiniProgramClient
}

// WechatMiniProgramClient 微信小程序
type WechatMiniProgramClient struct {
	*WechatClient
}

// JSAPI/小程序下单API
//	Code = 0 is success
//	商户JSAPI文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_1.shtml
//	商户小程序文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_5_1.shtml
func (w WechatMiniProgramClient) Pay(charge *common.Charge) (map[string]string, error) {
	client, err := wechat.NewClientV3(w.MchID, w.SerialNo, w.Key, w.PrivateKey)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}
	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	//初始化参数Map
	bm := make(gopay.BodyMap)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm.Set("appid", w.AppID).
		Set("description", utilLocal.TruncatedText(charge.Describe, 32)).
		Set("out_trade_no", charge.TradeNo).
		Set("time_expire", expire).
		Set("notify_url", charge.CallbackURL).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", charge.MoneyFee).
				Set("currency", "CNY")
		}).
		SetBodyMap("payer", func(bm gopay.BodyMap) {
			bm.Set("openid", charge.OpenID)
		})

	wxRsp, err := client.V3TransactionJsapi(context.TODO(), bm)
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
