package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func Init() (err error) {
	userName := "lqf"
	passWord := "Wangfei222@"
	ip := "192.168.1.1"
	port := 3306
	dbName := "bolgs"
	dsn := fmt.Sprintf("%s:%s:tcp(%s:%d)/%s", userName, passWord, ip, port, dbName)
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败，错误: %s", err)
		return
	} else {
		fmt.Println("连接成功")
	}
	return
}
