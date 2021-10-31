package sms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	//util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AliSms struct {
	AccessKeyId     string
	AccessKeySecret string
}

func NewAliSms(accessKeyId, accessKeySecret string) *AliSms {
	return &AliSms{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
}

func (a *AliSms) createClient() (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(a.AccessKeyId),
		AccessKeySecret: tea.String(a.AccessKeySecret),
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

// 发送国内短信
// phoneNumbers 接受手机号
// signName 短信签名名称
// templateCode 短信模板ID
// templateParam 短信模板变量对应的实际值 示例：{"name":"张三","number":"15038****76"}
// outId 外部流水扩展字段，方便去调取已经发送的短信状态
func (a *AliSms) SendSms(phoneNumbers, signName, templateCode, templateParam, outId string) (_bizId string, _err error) {
	client, _err := a.createClient()
	if _err != nil {
		return _bizId, _err
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phoneNumbers),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String(templateParam),
		OutId:         tea.String(outId),
	}
	sendRes, _err := client.SendSms(sendSmsRequest)
	if _err != nil {
		return _bizId, _err
	}

	_bizId = *sendRes.Body.BizId

	return
}

// 查询发送短信结果
// phoneNumber 接收短信的手机号码
// sendDate 短信发送日期，支持查询最近30天的记录。 格式为yyyyMMdd，例如20181225。
// pageSize 分页查看发送记录，指定每页显示的短信记录数量。 取值范围为1~50。
// currentPage 分页查看发送记录，指定发送记录的当前页码。
// bizId 发送回执ID 可以为空
func (a *AliSms) QuerySendDetails(phoneNumber string, sendDate string, pageSize int64, currentPage int64,
	bizId string) (_querySendDetailsResponse *dysmsapi20170525.QuerySendDetailsResponse, _err error) {
	client, _err := a.createClient()
	if _err != nil {
		return _querySendDetailsResponse, _err
	}

	querySendDetailsRequest := &dysmsapi20170525.QuerySendDetailsRequest{
		PhoneNumber: tea.String(phoneNumber),
		SendDate:    tea.String(sendDate),
		PageSize:    tea.Int64(pageSize),
		CurrentPage: tea.Int64(currentPage),
		BizId:       tea.String(bizId),
	}
	// 复制代码运行请自行打印 API 的返回值
	_querySendDetailsResponse, _err = client.QuerySendDetails(querySendDetailsRequest)
	if _err != nil {
		return _querySendDetailsResponse, _err
	}

	return
}

// 发送国际短信
// to
// message 参考：https://next.api.aliyun.com/document/Dysmsapi/2017-05-25/SendSms
// {"phoneNumbers" :"1111", "signName":"signName", "templateCode":"templateCode",
//	"templateParam":`{"name":"张三","number":"15038****76"}`, "outId": "outId"}
func (a *AliSms) SendGlobeSms(to, message string) (_sendMessageToGlobeResponse *dysmsapi20170525.SendMessageToGlobeResponse, _err error) {
	client, _err := a.createClient()
	if _err != nil {
		return _sendMessageToGlobeResponse, _err
	}

	sendMessageToGlobeRequest := &dysmsapi20170525.SendMessageToGlobeRequest{
		Message: tea.String(message),
		To:      tea.String(to),
	}

	_sendMessageToGlobeResponse, _err = client.SendMessageToGlobe(sendMessageToGlobeRequest)
	if _err != nil {
		return _sendMessageToGlobeResponse, _err
	}

	//console.Log(util.ToJSONString(tea.ToMap(resp)))
	return
}
