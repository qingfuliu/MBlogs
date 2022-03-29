package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseData struct {
	Code    MStatus     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseError(ctx *gin.Context, code MStatus) {
	res := &ResponseData{
		Code:    code,
		Message: code.msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, res)
}

func ResponseErrorWithData(ctx *gin.Context, code MStatus, data interface{}) {
	res := &ResponseData{
		Code:    code,
		Message: code.msg(),
		Data:    data,
	}
	if _, ok := data.(error); ok {
		res.Data = data.(error).Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	res := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, res)
}
