package redigo

import (
	"fmt"

	goredislib "github.com/go-redis/redis/v8"
	redsyncredis "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

// 获取分布式redis锁的 pool
func GetRedsyncLockPool(host string, port int, user string, password string) redsyncredis.Pool {
	client := goredislib.NewClient(&goredislib.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		PoolSize:     5,
		MinIdleConns: 10,
		Username:     user,
		Password:     password,
	})

	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	return pool
}
