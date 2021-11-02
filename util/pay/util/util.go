package util

import (
	"net"
	"strings"

	"github.com/shopspring/decimal"
)

// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

// 过滤文本内容的特殊字符
func TruncatedText(data string, length int) string {
	data = FilterTheSpecialSymbol(data)
	if len([]rune(data)) > length {
		return string([]rune(data)[:length-1])
	}
	return data
}

//过滤特殊符号
func FilterTheSpecialSymbol(data string) string {
	// 定义转换规则
	specialSymbol := func(r rune) rune {
		if r == '`' || r == '[' || r == '~' || r == '!' || r == '@' || r == '#' || r == '$' ||
			r == '^' || r == '&' || r == '*' || r == '~' || r == '(' || r == ')' || r == '=' ||
			r == '~' || r == '|' || r == '{' || r == '}' || r == '~' || r == ':' || r == ';' ||
			r == '\'' || r == ',' || r == '\\' || r == '[' || r == ']' || r == '.' || r == '<' ||
			r == '>' || r == '/' || r == '?' || r == '~' || r == '！' || r == '@' || r == '#' ||
			r == '￥' || r == '…' || r == '&' || r == '*' || r == '（' || r == '）' || r == '—' ||
			r == '|' || r == '{' || r == '}' || r == '【' || r == '】' || r == '‘' || r == '；' ||
			r == '：' || r == '”' || r == '“' || r == '\'' || r == '"' || r == '。' || r == '，' ||
			r == '、' || r == '？' || r == '%' || r == '+' || r == '_' || r == ']' || r == '"' || r == '&' {
			return ' '
		}
		return r
	}
	data = strings.Map(specialSymbol, data)
	return strings.Replace(data, "\n", " ", -1)
}


// 微信金额浮点转字符串
func WechatMoneyFeeToString(moneyFee float64) string {
	aDecimal := decimal.NewFromFloat(moneyFee)
	bDecimal := decimal.NewFromFloat(100)
	return aDecimal.Mul(bDecimal).Truncate(0).String()
}

// 支付宝金额转字符串
func AliyunMoneyFeeToString(moneyFee float64) string {
	return decimal.NewFromFloat(moneyFee).Truncate(2).String()
}