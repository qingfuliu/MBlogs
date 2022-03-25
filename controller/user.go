package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goProject/dao"
	"goProject/logic"
	"goProject/models"
)

func SingUpHandle(c *gin.Context) {
	newUser := &models.UserRegister{}
	if err := c.ShouldBindJSON(newUser); err != nil || newUser.PassWord != newUser.ConfirmPassword {
		if validationError, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithData(c, CodeInvaildParams,
				removeStructHeader(validationError.Translate(translator)))
			return
		}
		ResponseError(c, CodeInvaildParams)
		return
	}
	if err := logic.Register(newUser); err != nil {
		if errors.Is(err, dao.ErrorUserExisted) {
			ResponseErrorWithData(c, CodeInvaildParams, err)
			return
		} else if errors.Is(err, dao.ErrorInsertFailed) {
			ResponseErrorWithData(c, CodeSeverBase, err)
			return
		}
	}
	ResponseSuccess(c, nil)
}

func LoginHandle(c *gin.Context) {
	userLogin := &models.UserLoginForm{}
	if err := c.ShouldBindJSON(userLogin); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			ResponseErrorWithData(c, CodeInvaildParams, removeStructHeader(errs.Translate(translator)))
			return
		}
		ResponseError(c, CodeInvaildParams)
		return
	}
	if err := logic.Login(userLogin); err != nil {
		ResponseError(c, CodeCertifiedFailed)
		return
	}
	ResponseSuccess(c, nil)
}
