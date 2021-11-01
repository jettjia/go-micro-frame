// 分布式锁
package redigo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var (
	updateLockExpireUidScript = redis.NewScript(1, `
		local res = redis.call("SETNX", KEYS[1], ARGV[1]) 
		if res == 1 then
			return redis.call("EXPIRE", KEYS[1], ARGV[2])
		end
		return res
	`)
	deleteLockByUidScript = redis.NewScript(1, `
		local res = redis.call("GET", KEYS[1]) 
		if res == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		end
		return res 
	`)
)

type redisRedsync struct {
	Conn redis.Conn
}

func NewRedisRedsync(addr []string, password string, typ int) *redisRedsync {
	return &redisRedsync{
		Conn: NewRedisClient(addr, password, typ),
	}
}

// Lock 加锁
func (r *redisRedsync) Lock(key []string, uid int, expire int) bool {
	conn := r.Conn
	lock := false

	num := 1
	for !lock && num < 3{
		res, err := updateLockExpireUidScript.Do(conn, key, uid, expire)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if res.(int64) == 1 {
			lock = true
			//fmt.Println("获取锁成功")
			return true
		}

		num++
	}
	//fmt.Println("获取锁失败")
	return false
}

// UnLock 删除锁
func (r *redisRedsync) UnLock(key []string, uid int) bool {
	conn := r.Conn

	_, err := deleteLockByUidScript.Do(conn, key, uid)
	if err != nil {
		//fmt.Println("删除锁失败 ", err.Error())
		return false
	}
	//fmt.Println("删除锁成功")
	return true
}
