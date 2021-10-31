package sms

import (
	"fmt"
	"testing"
)

func TestAliSms_SendSms(t *testing.T) {
	aliSms := NewAliSms("xx", "yy")
	bizId, err := aliSms.SendSms("", "", "", "", "")
	if err != nil {
		t.Error("发送短信失败")
	}

	fmt.Println("SendSms response bizId:", bizId)
}
