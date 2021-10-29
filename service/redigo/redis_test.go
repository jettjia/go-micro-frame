package redigo

import (
	"fmt"
	"testing"

	"github.com/gomodule/redigo/redis"
)

func Test_Alone(t *testing.T) {
	addr := []string{
		"10.4.7.71:6379",
	}
	conn := NewRedisClient(addr, "", 1)
	reply, err := redis.String(conn.Do("GET", "some-key"))
	if err != nil {
		t.Errorf("GET failed: %v", err)
	}
	defer conn.Close()
	fmt.Println(reply)
}

func Test_Cluster(t *testing.T) {
	addr := []string{
		"10.4.7.71:6379",
	}
	conn := NewRedisClient(addr, "", 3)

	// call commands on it
	//_, err := redis.String(conn.Do("SET", "some-key", "2222"))
	//if err != nil {
	//	t.Errorf("SET failed: %v", err)
	//}

	s, err := redis.String(conn.Do("GET", "some-key"))
	if err != nil {
		t.Errorf("GET failed: %v", err)
	}

	defer conn.Close()
	fmt.Println(s)
}

func Test_Sentinel(t *testing.T) {
	addr := []string{
		"10.4.7.71:6379",
	}
	conn := NewRedisClient(addr, "", 2)
	s, err := redis.String(conn.Do("GET", "some-key"))
	if err != nil {
		t.Errorf("GET failed: %v", err)
	}

	defer conn.Close()
	fmt.Println(s)
}
