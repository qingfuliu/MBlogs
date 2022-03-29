package logic

import (
	"goProject/dao"
	"goProject/models"
)

func CreateArticle(article *models.Article) (err error) {
	var ok bool
	if ok, err = dao.IfArticleExisted(article); ok || err == dao.ErrorQueryFailed {
		return
	}
	err = dao.InsertArticle(article)
	return
}
