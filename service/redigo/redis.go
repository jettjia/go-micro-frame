package redigo

import (
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/gomodule/redigo/redis"
)

// redis 有三种模式
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

// 分布式锁
func NewRedsyncLockPool(host string, port int, user string, password string) redsyncredis.Pool {
	return GetRedsyncLockPool(host, port, user, password)
}
