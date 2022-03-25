package models

import "time"

type User struct {
	ID        string    `gorm:"column:id"`
	UserName  string    `gorm:"column:username"`
	PassWord  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Gender    bool      `gorm:"column:gender"`
	CreatedAt time.Time `gorm:"column:createdate"`
	UpdatedAt time.Time `gorm:"column:updatedate"`
}

func (u User) TableName() string {
	return "users"
}
