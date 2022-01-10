package external_ip

import (
	"fmt"
	"testing"
)

func TestExternalIP(t *testing.T) {
	ip, err := ExternalIP()
	if err != nil {
		t.Error()
	}

	fmt.Println(ip.String())
}
