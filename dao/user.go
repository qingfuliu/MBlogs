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

func IfUsersExisted(username string) (bool, error) {
	temp := makeStruct(map[string]interface{}{
		"username":   username,
		"table_name": models.User{}.TableName(),
	})
	return IfIsExisted(temp)
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

func ModifyUser(user *models.User) error {
	err := dbConn.Where("id=?", user.ID).Updates(user).Error
	return err
}

func QueryUser(username string) (*models.User, error) {
	user := &models.User{}
	if err := dbConn.Where("username=?", username).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
