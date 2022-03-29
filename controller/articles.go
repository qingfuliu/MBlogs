package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"goProject/logic"
	"goProject/models"
)

func PushArticle(c *gin.Context) {
	articleFrom := &models.ArticleForm{}
	if err := c.ShouldBindJSON(articleFrom); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithData(c, CodeInvaildParams, errs)
			return
		}
		ResponseErrorWithData(c, CodeInvaildParams, err)
		return
	}
	article := &models.Article{
		ArticleForm: articleFrom,
	}
	article.ArticleAuthor = c.GetString("username")
	if err := logic.CreateArticle(article); err != nil {
		ResponseErrorWithData(c, CodeInvaildParams, err)
		return
	}
	ResponseSuccess(c, gin.H{
		"msg": "发表成功!",
	})
}
