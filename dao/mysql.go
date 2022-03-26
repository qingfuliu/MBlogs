package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	userName := "liuqingfu"
	passWord := "123456"
	ip := "101.43.252.121"
	port := 3306
	dbName := "qfdatabase"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", userName, passWord, ip, port, dbName)
	var err error
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接失败，错误: %s", err)
		return
	} else {
		fmt.Println("连接成功")
	}
	return
}
