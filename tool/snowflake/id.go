package snowflake

import (
	"errors"
	"log"
	"math/big"
	"net"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {

	// Create snowflake node
	n, err := snowflake.NewNode(genNodeId())
	if err != nil {
		log.Fatal(err)
	}
	// Set node
	node = n
}

// 获取id
func Id() int64 {
	return node.Generate().Int64()
}

// Parse node number
func genNodeId() int64 {
	ip, err := ExternalIP()
	if err != nil {
		log.Fatal(err)
	}
	ips := strings.Split(ip.String(), ".")
	if len(ips) != 4 {
		return 0
	}

	i1, _ := strconv.Atoi(ips[0])
	i2, _ := strconv.Atoi(ips[1])
	i3, _ := strconv.Atoi(ips[2])
	i4, _ := strconv.Atoi(ips[3])

	total := i1 + i2 + i3 + i4
	if total > 1023 {
		return int64(total - 1023)
	}

	return int64(total)
}

// 将ip地址转整形
func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

// 获取本机ip
func ExternalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
