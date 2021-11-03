package wechat

import (
	"github.com/go-pay/gopay/wechat/v3"
)

type WechatClient struct {
	AppID      string // 应用ID
	MchID      string // 商户ID 或者服务商模式的 sp_mchid
	Key        string // API秘钥值, 商户平台获取
	IsProd     bool   // 是否是正式环境
	SerialNo   string // 商户API证书的证书序列号
	PrivateKey string // 商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	PublicKey  string // 公钥文件内容
}

func NewWechatClient(appId, mchID, key, serialNo, privateKey string, isProd bool) *WechatClient {
	return &WechatClient{
		AppID:      appId,
		MchID:      mchID,
		Key:        key,
		IsProd:     isProd,
		SerialNo:   serialNo,
		PrivateKey: privateKey,
	}
}

// InitNewWechatClient 初始化微信支付客户端
func (w *WechatClient) InitNewWechatClient() (client *wechat.ClientV3, err error) {
	client, err = wechat.NewClientV3(w.MchID, w.SerialNo, w.Key, w.PrivateKey)
	if err != nil {
		return nil, err
	}

	return
}
