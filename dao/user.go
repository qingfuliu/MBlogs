package dao

import (
	"errors"
	"goProject/models"
)

func InsertUser(user *models.User) error {
	if err := dbConn.Create(user).Error; err != nil {
		return ErrorInsertFailed
	}
	return nil
}

func IfIsExisted(username string) (err error) {
	var nums int64 = 0
	err = dbConn.Model(&models.User{}).Select("id").Where("username=?", username).Count(&nums).Error
	if err != nil {
		err = ErrorInsertFailed
		return
	}
	if nums != 0 {
		err = ErrorUserExisted
	}
	return
}

func IfCertified(user *models.UserLoginForm) (err error) {
	var md5PassWord string
	if err = dbConn.Model(user).Select("password").Where("username=?", user.UserName).Find(&md5PassWord).Error; err != nil {
		return
	}
	if md5PassWord != user.PassWord {
		err = errors.New("password not match")
	}
	return
}
