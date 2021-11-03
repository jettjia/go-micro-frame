package free_port

import (
	"fmt"

	"testing"
)

func Test_GetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil {
		fmt.Errorf("GetFreePort error")
	}
	fmt.Println("获取到的微服务端口是：", port)
}
