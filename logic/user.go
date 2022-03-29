package logic

import (
	"goProject/dao"
	"goProject/generate"
	"goProject/models"
)

func Register(user *models.UserRegister) error {
	if existed, err := dao.IfUsersExisted(user.UserName); existed {
		return err
	}
	var md5Paddword string
	var err error
	if md5Paddword, err = generate.Md5(user.PassWord); err != nil {
		return err
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
	var ok bool
	if ok, err = dao.IfUsersExisted(user.UserName); !ok {
		return err
	}
	if user.PassWord, err = generate.Md5(user.PassWord); err != nil {
		return
	}
	if err = dao.IfCertified(user); err != nil {
		return
	}
	return nil
}

func ModifyUser(user *models.User) error {
	oldUser, err := dao.QueryUser(user.UserName)
	if err != nil {
		return dao.ErrorQueryFailed
	}
	if oldUser != nil && oldUser.ID != user.ID {
		return dao.ErrorUserExisted
	}
	if err := dao.ModifyUser(user); err != nil {
		return dao.ErrorModifyFailed
	}
	return nil
}
