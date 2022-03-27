package dao

import "errors"

var (
	ErrorUserExisted    = errors.New("用户已经存在")
	ErrorUserNotExisted = errors.New("用户不已经存在")
	ErrorInsertFailed   = errors.New("插入错误")
	ErrorQueryFailed    = errors.New("查询错误")
)
