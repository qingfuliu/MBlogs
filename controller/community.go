package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goProject/dao"
	"goProject/logic"
	"goProject/models"
	"net/http"
)

func CreateCommunityHandle(c *gin.Context) {
	newCommunity := models.CommunityDetail{}
	if err := c.ShouldBindJSON(&newCommunity); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithData(c, CodeInvaildParams, errs.Translate(translator))
			return
		}
		ResponseError(c, CodeInvaildParams)
		return
	}
	newCommunity.Creator = c.GetString("username")
	if err := logic.CreateCommunity(&newCommunity); err != nil {
		ResponseErrorWithData(c, CodeInvaildParams, err)
		return
	}
	return
}

func QueryBatchCommunityHandle(c *gin.Context) {
	newQuery := &models.BatchCommunities{}
	if err := c.ShouldBindJSON(newQuery); err != nil {
		errs, ok := err.(*validator.ValidationErrors)
		if ok {
			ResponseErrorWithData(c, CodeInvaildParams, errs)
			return
		}
		ResponseErrorWithData(c, CodeSeverBase, err)
		return
	}
	if rs, err := dao.BatchCommunityQuery(newQuery); err != nil {
		ResponseErrorWithData(c, CodeSeverBase, err)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"communities": rs,
		})
	}
	return
}

func ModifyCommunity(c *gin.Context) {
	community := &models.Community{}
	if err := c.BindJSON(community); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithData(c, CodeInvaildParams, errs.Translate(translator))
			return
		}
		ResponseErrorWithData(c, CodeInvaildParams, err)
		return
	}
	if err := logic.ModifyCommunity(community); err != nil {
		ResponseErrorWithData(c, CodeInvaildParams, err)
		return
	}
	ResponseSuccess(c, gin.H{
		"msg": "修改成功",
	})
	return
}
