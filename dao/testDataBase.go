package dao

import (
	"fmt"
	"goProject/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int64     `gorm:"column:id"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:createdate"`
}

func (u User) TableName() string {
	return "users"
}
func ConnAndInsertAndSelect() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("ConnandSelect err,ere is:%s", p)
		}
	}()
	username := "lqf"
	password := "Wangfei222@"
	host := "192.168.1.103"
	port := 3306
	dbName := "blogs"
	// Mysql dsn格式： {username}:{password}@tcp({host}:{port})/
	//{Dbname}?charset=utf8&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print(err)
		return
	}
	u := User{
		Id:       2,
		Username: "lqf2",
		Password: "wangfei222",
	}
	if err := db.Create(&u).Error; err != nil {
		fmt.Print(err)
		return
	} else {
		fmt.Println("success")
	}
	var nums int64
	dbConn.Model(&models.User{}).Select("id").Where("username=", "lqf").Count(&nums)
	if nums != 0 {
		nums = 1
	}

}
