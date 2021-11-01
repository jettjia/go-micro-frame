package redigo

import (
	"fmt"
	"testing"
)

// 分布式锁
func TestRedisRedsync_Lock(t *testing.T) {
	addr := []string{
		"10.4.7.71:6379",
	}

	// 1 alone, 2 sentinel, 3 cluster
	typ := 1
	client := NewRedisRedsync(addr, "", typ)

	key := []string{"redisLock"}
	uid := 100
	expire := 30

	bo := client.Lock(key, uid, expire)
	fmt.Println("结果：", bo)
}

// 释放锁
func TestRedisRedsync_UnLock(t *testing.T) {
	addr := []string{
		"10.4.7.71:6379",
	}

	// 1 alone, 2 sentinel, 3 cluster
	typ := 1
	client := NewRedisRedsync(addr, "", typ)

	key := []string{"redisLock"}
	uid := 100

	bo := client.UnLock(key, uid)
	fmt.Println("结果：", bo)
}
