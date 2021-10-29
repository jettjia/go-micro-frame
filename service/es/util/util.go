package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"net"
	"os"
	"strings"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}


//检测并补全路径右边的反斜杠
func RightAddPathPos(path string) string {
	if path[len(path)-1:len(path)] != "/" {
		path = path + "/"
	}
	return path
}


//断言
func Assertion(data interface{}) interface{} {
	switch data.(type) {
	case string:
		return data.(string)
	case int:
		return data.(int)
	case int8:
		return data.(int8)
	case int32:
		return data.(int32)
	case int64:
		return data.(int64)
	case float32:
		return data.(float32)
	case float64:
		return data.(float64)
	default:
		return data
	}
	return nil
}

//json转map数组
func JsonToMapArray(data string) []map[string]interface{} {
	var res []map[string]interface{}
	if data == "" {
		return res
	}
	err := json.Unmarshal([]byte(data), &res)
	Must(err)

	return res
}

//json转map
func JsonToMap(data string) map[string]interface{} {
	var res map[string]interface{}
	if data == "" {
		return res
	}
	err := json.Unmarshal([]byte(data), &res)
	Must(err)
	return res
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}

func ExternalIP() (string, error) {
	iFaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iFace := range iFaces {
		if iFace.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iFace.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iFace.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

//获得本机名
func HostName() string {
	hostNamePrefix := ""
	host, err := os.Hostname()
	Must(err)
	if err == nil {
		parts := strings.SplitN(host, ".", 2)
		if len(parts) > 0 {
			hostNamePrefix = parts[0]
		}
	}
	return hostNamePrefix
}