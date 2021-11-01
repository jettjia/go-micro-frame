package redigo

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/mna/redisc"
)

func NewRedisCluster(addr []string, password string) redis.Conn {
	// create the cluster
	cluster := redisc.Cluster{
		//StartupNodes: []string{
		//	"10.4.7.71:6379", "10.4.7.71:7000", "10.4.7.71:7001",
		//},
		StartupNodes: addr,
		DialOptions:  []redis.DialOption{redis.DialConnectTimeout(5 * time.Second), redis.DialPassword(password)},
		CreatePool:   createPool,
	}
	defer cluster.Close()

	// initialize its mapping
	if err := cluster.Refresh(); err != nil {
		log.Fatalf("Refresh failed: %v", err)
	}

	// grab a connection from the pool
	conn := cluster.Get()

	return conn
}

func createPool(addr string, opts ...redis.DialOption) (*redis.Pool, error) {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   10,
		IdleTimeout: time.Minute,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr, opts...)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}, nil
}
