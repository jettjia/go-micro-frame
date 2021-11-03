package alipay

import (
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/alipay/cert"

	"github.com/jettjia/go-micro-frame/util/pay/constant"
)

type AliClient struct {
	AppID                   string // 应用ID
	PrivateKey              string // 私钥
	AlipayPublicContentRSA2 string // 支付宝公钥证书
	AppPublicContent        string // 应用公钥证书
	AlipayRootContent       string // 支付宝根证书
	DebugSwitch             int    // 日志开启，1开0关
}

func NewAliClient(appID, privateKey, alipayPublicContentRSA2, appPublicContent,
	alipayRootContent string, debugSwitch int) *AliClient {
	return &AliClient{
		AppID:                   appID,
		PrivateKey:              privateKey,
		AlipayPublicContentRSA2: alipayPublicContentRSA2,
		AppPublicContent:        appPublicContent,
		AlipayRootContent:       alipayRootContent,
		DebugSwitch:             debugSwitch,
	}
}

// InitAliClient 初始化支付宝客户端
func (a *AliClient) InitAliClient() (*alipay.Client, error) {
	//    appId：应用ID
	//    privateKey：应用私钥，支持PKCS1和PKCS8
	//    isProd：是否是正式环境
	client, err := alipay.NewClient(a.AppID, a.PrivateKey, false)
	if err != nil {
		return nil, err
	}
	// 打开Debug开关，输出日志
	client.DebugSwitch = gopay.DebugSwitch(a.DebugSwitch)

	// 配置公共参数
	client.SetCharset("utf-8").
		SetSignType(constant.RSA2)
	// SetAppAuthToken("")

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign(cert.AlipayPublicContentRSA2)

	// 传入证书内容
	err = client.SetCertSnByContent(cert.AppPublicContent, cert.AlipayRootContent, cert.AlipayPublicContentRSA2)
	if err != nil {
		return nil, err
	}

	return client, nil
}
