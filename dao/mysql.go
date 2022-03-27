package dao

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func init() {
	userName := "lqf"
	passWord := "Wangfei222@"
	ip := "192.168.1.103"
	port := 3306
	dbName := "blogs"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", userName, passWord, ip, port, dbName)
	var err error
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("连接失败，错误: %s", err)
		return
	} else {
		fmt.Println("连接成功")
	}
	return
}

func IfIsExisted(colsName string, colsValue string, tableName string) (bool, error) {
	var nums int64
	err := dbConn.Table(tableName).Where(colsName+"=?", colsValue).Count(&nums).Error
	if err != nil {
		return false, errors.New("query failed")
	}
	if nums != 0 {
		return true, errors.New("has been existed")
	}
	return false, errors.New("not existed")
}
