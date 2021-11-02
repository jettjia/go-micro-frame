package alipay

type AliAppClient struct {
	AppID string // 应用ID

	PrivateKey string
}

func NewAliAppClient(appID string, privateKey string) *AliAppClient {
	return &AliAppClient{
		AppID:      appID,
		PrivateKey: privateKey,
	}
}
