package redigo

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func NewRedisAlone(addr string, password string) redis.Conn {
	conn, err := redis.Dial(
		"tcp",
		addr,
		redis.DialPassword(password),
	)
	if err != nil {
		fmt.Println("redis.Dial err", err)
	}

	return conn
}
