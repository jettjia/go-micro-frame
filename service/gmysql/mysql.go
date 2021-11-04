package gmysql

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Mysql struct {
	Host         string // 服务器地址
	Port         int    // 端口
	User         string // 数据库用户名
	Password     string // 数据库密码
	Db           string // 数据名
	MaxIdleConns int    // 最大空闲连接
	MaxOpenConns int    // 最大连接数
	MaxLifetime  int    // 最大生存时间(s)
	LogMode      bool   // 是否打印日志
}

//参数含义:数据库用户名、密码、主机ip、连接的数据库、端口号
func (m *Mysql) GetDB() (*gorm.DB, error) {
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Db)
	fmt.Println(connArgs)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		return nil, err
	}
	db.SingularTable(true) //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	db.LogMode(m.LogMode)  //打印sql语句

	//开启连接池
	if m.MaxIdleConns == 0 {
		m.MaxIdleConns = 10
	}
	if m.MaxOpenConns == 0 {
		m.MaxOpenConns = 100
	}
	if m.MaxLifetime == 0 {
		m.MaxLifetime = 30
	}
	db.DB().SetMaxIdleConns(m.MaxIdleConns)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Duration(m.MaxLifetime) * time.Second)

	return db, nil
}
