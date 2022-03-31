package dao

import (
	"fmt"
	"go.uber.org/zap"
	"goProject/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reflect"
	"strings"
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
		logger.Mlogger.Info("db 连接失败，err is ", zap.Error(err))
		return
	} else {
		logger.Mlogger.Info("db 成功")
	}
}

func IfIsExisted(arg interface{}) (bool, error) {
	var nums int64
	rfv := arg.(reflect.Value)
	rft := reflect.TypeOf(rfv.Interface())
	var db *gorm.DB
	if _, ok := rft.FieldByName("Ktable_name"); ok {
		db = dbConn.Table(rfv.FieldByName("Ktable_name").String())
	} else {
		return false, ErrorQueryFailed
	}
	for index := 0; index < rft.NumField(); index++ {
		fieldVal := rfv.Field(index).String()
		fieldTag := rft.Field(index).Name[1:] + "=?"
		if !strings.Contains(fieldTag, "table_name") {
			db = db.Where(fieldTag, fieldVal)
		}
	}
	err := db.Count(&nums).Error
	if err != nil {
		return false, ErrorQueryFailed
	}
	if nums != 0 {
		return true, ErrorUserExisted
	}
	return false, ErrorUserNotExisted
}

func makeStruct(m map[string]interface{}) reflect.Value {
	var structFields []reflect.StructField
	for key, Val := range m {
		typeOf := reflect.TypeOf(Val)
		structElem := reflect.StructField{
			Name: "K" + key,
			Type: typeOf,
		}
		structFields = append(structFields, structElem)
	}
	st := reflect.StructOf(structFields)
	so := reflect.New(st).Elem()
	ret := reflect.TypeOf(so.Interface())
	for index := 0; index < so.NumField(); index++ {
		so.Field(index).SetString(m[ret.Field(index).Name[1:]].(string))
	}
	return so
}
