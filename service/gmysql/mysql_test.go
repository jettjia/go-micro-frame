package gmysql

import (
	"fmt"
	"testing"
)

type User struct {
	Id     int    `gorm:"primary_key" json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender int    `json:"gender"` //1:男、2:女
}

func Test_GetDB(t *testing.T) {
	m := &Mysql{
		Host:         "10.4.7.71",
		Port:         3307,
		User:         "root",
		Password:     "root",
		Db:           "test",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		MaxLifetime:  20,
	}

	DB, err := m.GetDB()
	if err != nil {
		t.Error(err)
	}

	var user User
	DB.First(&user, 1)
	fmt.Println(user)
}
