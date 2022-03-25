package models

import "time"

type User struct {
	ID        int64     `gorm:"column:id"`
	UserName  string    `gorm:"column:username"`
	PassWord  string    `gorm:"column:password"`
	Email     string    `gorm:"column:email"`
	Gender    bool      `gorm:"-"`
	CreatedAt time.Time `gorm:"column:createdate"`
	UpdatedAt time.Time `gorm:"column:updatedate"`
}

func (u User) TableName() string {
	return "users"
}

type UserLoginForm struct {
	UserName string `gorm:"column:username" json:"username" binding:"required"`
	PassWord string `gorm:"column:password" json:"password" binding:"required"`
}

func (u UserLoginForm) TableName() string {
	return "users"
}

type UserRegister struct {
	UserName        string `json:"username" binding:"required"`
	PassWord        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"re_password" binding:"required,eqfield=PassWord"`
	Email           string `json:"email,omitempty"`
	Gender          bool   `json:"gender,omitempty"`
}
