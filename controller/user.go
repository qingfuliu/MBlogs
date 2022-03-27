package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goProject/dao"
	"goProject/generate"
	"goProject/logic"
	"goProject/models"
	"net/http"
	"strings"
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
		ResponseErrorWithData(c, CodeInvaildParams, err)
		return
	}
	aToken, rToken, err := generate.GetAssAndRefToken(userLogin.UserName)
	if err != nil {
		ResponseError(c, CodeSeverBase)
		return
	}
	ResponseSuccess(c, gin.H{
		"rToken":   rToken,
		"aToken":   aToken,
		"username": userLogin.UserName,
	})
}

func RefreshToken(c *gin.Context) {
	rtbefore := strings.SplitN(c.GetHeader("rToken"), " ", 2)
	if len(rtbefore) < 2 {
		ResponseErrorWithData(c, CodeInvaildParams, gin.H{
			"data": "found no refresh token",
		})
		c.Abort()
		return
	}
	at := c.GetHeader("aToken")
	if at == "" {
		ResponseErrorWithData(c, CodeInvaildParams, gin.H{
			"data": "found no access token",
		})
		c.Abort()
		return
	}
	atSAf := strings.SplitN(at, " ", 2)
	if len(atSAf) != 2 || atSAf[0] != "Bearer" {
		ResponseErrorWithData(c, CodeInvaildParams, gin.H{
			"data": "access token parse error",
		})
		c.Abort()
		return
	}
	aToken, err := generate.RefreshToken(atSAf[1], rtbefore[1])
	if err != nil {
		ResponseErrorWithData(c, CodeInvaildParams, err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"aToken": aToken,
		})
}
