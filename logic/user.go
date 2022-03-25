package logic

import (
	"goProject/dao"
	"goProject/generate"
	"goProject/models"
)

func Register(user *models.UserRegister) (err error) {
	if err = dao.IfIsExisted(user.UserName); err != nil {
		return
	}
	var md5Paddword string
	if md5Paddword, err = generate.Md5(user.PassWord); err != nil {
		return
	}
	newUser := models.User{
		ID:       generate.SnowFlakeUID(),
		UserName: user.UserName,
		PassWord: md5Paddword,
		Email:    user.Email,
		Gender:   user.Gender,
	}
	return dao.InsertUser(&newUser)
}

func Login(user *models.UserLoginForm) (err error) {
	if user.PassWord, err = generate.Md5(user.PassWord); err != nil {
		return
	}
	if err = dao.IfCertified(user); err != nil {
		return
	}
	return nil
}
