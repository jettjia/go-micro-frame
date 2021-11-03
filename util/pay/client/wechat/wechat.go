package wechat

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/pkg/xlog"
	"github.com/go-pay/gopay/wechat/v3"
)

type WechatClient struct {
	AppID       string // 应用ID
	MchID       string // 商户ID 或者服务商模式的 sp_mchid
	Key         string // API秘钥值, 商户平台获取
	IsProd      bool   // 是否是正式环境
	SerialNo    string // 商户API证书的证书序列号
	PrivateKey  string // 商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	PublicKey   string // 公钥文件内容
	DebugSwitch int    // 日志开启，1开0关
}

func NewWechatClient(appId, mchID, key, serialNo, privateKey string, isProd bool, debugSwitch int) *WechatClient {
	return &WechatClient{
		AppID:       appId,
		MchID:       mchID,
		Key:         key,
		IsProd:      isProd,
		SerialNo:    serialNo,
		PrivateKey:  privateKey,
		DebugSwitch: debugSwitch,
	}
}

// InitNewWechatClient 初始化微信支付客户端
func (w *WechatClient) InitNewWechatClient() (client *wechat.ClientV3, err error) {
	// NewClientV3 初始化微信客户端 V3
	//	mchid：商户ID
	// 	serialNo：商户证书的证书序列号
	//	apiV3Key：APIv3Key，商户平台获取
	//	privateKey：商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	client, err = wechat.NewClientV3(w.MchID, w.SerialNo, w.Key, w.PrivateKey)
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
	client.DebugSwitch = gopay.DebugSwitch(w.DebugSwitch)

	return
}
