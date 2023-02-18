package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error
	//配置MySQL连接参数
	username := "root"
	password := "password"
	host := "127.0.0.1"
	port := 3306
	Dbname := "douyin5th"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		username, password, host, port, Dbname)
	//dsn := "root:54321@tcp(127.0.0.1:3306)/bytedance_qingxunying_douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
