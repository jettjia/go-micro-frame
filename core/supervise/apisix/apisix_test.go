package apisix

import "testing"

// 创建路由
func Test_CreateRouter(t *testing.T) {
	conf := Apisix{
		ApisixHost:  "10.4.7.102",
		ApisixPort:  9080,
		ApisixToken: "edd1c9f034335f136f87ad84b625c8f1",
		Host:        "127.0.0.1",
		Port:        8021,
		ServiceName: "backend",
	}
	client := NewApisixClient(conf)

	err := client.CreateRouter()
	if err != nil {
		t.Error()
	}
}

// 删除路由
func Test_DeleteRouter(t *testing.T) {
	conf := Apisix{
		ApisixHost:  "10.4.7.102",
		ApisixPort:  9080,
		ApisixToken: "edd1c9f034335f136f87ad84b625c8f1",
		Host:        "127.0.0.1",
		Port:        8021,
		ServiceName: "backend",
	}
	client := NewApisixClient(conf)

	err := client.DeleteRouter()
	if err != nil {
		t.Error()
	}
}
