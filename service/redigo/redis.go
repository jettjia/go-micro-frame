package redigo

import (
	"github.com/gomodule/redigo/redis"
)

// NewRedisClient 获取redis的链接
// 1 alone, 2 sentinel, 3 cluster
func NewRedisClient(addr []string, password string, typ int) redis.Conn {
	var conn redis.Conn

	if typ == 1 {
		conn = NewRedisAlone(addr[0], password)
	}

	if typ == 2 {
		conn = NewRedisSentinel(addr, password)
	}

	if typ == 3 {
		conn = NewRedisCluster(addr, password)
	}

	return conn
}