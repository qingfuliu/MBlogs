package dao

import (
	"fmt"
	"goProject/models"
)

type UserRegister struct {
	UserName   string `json:"username,string" binding:"required"`
	PassWord   string `json:"password,string" binding:"required"`
	ConfirmPsw string `json:"config_password" binding:"required,eqfield=Password"`
	Email      string `json:"email,string"`
	Gender     bool   `json:"gender"`
}

func Register(user *UserRegister) (err error) {
	if err = ifIsExisted(user); err != nil {
		return
	}
	newUser := models.User{
		UserName: user.UserName,
		PassWord: user.PassWord,
		Email:    user.Email,
		Gender:   user.Gender,
	}

	return insertUser(&newUser)
}

func insertUser(user *models.User) (err error) {
	return dbConn.Create(user).Error
}

func ifIsExisted(user *UserRegister) (err error) {
	var nums int64 = 0
	dbConn.Model(&models.User{}).Select("id").Where("username=", user.UserName).Count(&nums)
	if nums != 0 {
		err = fmt.Errorf("userName has been existed")
	}
	return
}
