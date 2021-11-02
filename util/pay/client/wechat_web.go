package client

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

var defaultWechatWebClient *WechatWebClient

func InitWxWebClient(c *WechatWebClient) {
	defaultWechatWebClient = c
}

func DefaultWechatWebClient() *WechatWebClient {
	return defaultWechatWebClient
}

// WechatWebClient web,支付二维码
type WechatWebClient struct {
	AppID      string // 应用ID
	MchID      string // 商户ID 或者服务商模式的 sp_mchid
	Key        string // API秘钥值, 商户平台获取
	IsProd     bool   // 是否是正式环境
	SerialNo   string // 商户API证书的证书序列号
	PrivateKey string // 商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	PublicKey  string // 公钥文件内容
}

// Native下单API
//	Code = 0 is success
//	文档：https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_4_1.shtml
func (w WechatWebClient) Pay(charge *common.Charge) (map[string]string, error) {
	client, err := wechat.NewClientV3(w.MchID, w.SerialNo, w.Key, w.PrivateKey)
	if err != nil {
		return nil, err
	}
	// 启用自动同步返回验签，并定时更新微信平台API证书
	err = client.AutoVerifySign()
	if err != nil {
		return nil, err
	}

	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugOff

	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)

	bm := make(gopay.BodyMap)
	bm.Set("appid", w.AppID).
		Set("description", utilLocal.TruncatedText(charge.Describe, 32)).
		Set("out_trade_no", charge.TradeNo).
		Set("time_expire", expire).
		Set("notify_url", charge.CallbackURL).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", charge.MoneyFee).
				Set("currency", "CNY")
		})

	wxRsp, err := client.V3TransactionNative(context.TODO(), bm)
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
