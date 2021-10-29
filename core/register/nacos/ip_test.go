package nacos

import (
	"fmt"
	"testing"
)

func Test_Ip(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ip.String())
}
