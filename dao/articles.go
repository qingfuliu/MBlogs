package dao

import "goProject/models"

func InsertArticle(article *models.Article) (err error) {
	err = dbConn.Create(article).Error
	return
}

func IfArticleExisted(article *models.Article) (bool, error) {
	temp := makeStruct(map[string]interface{}{
		"author":     article.ArticleAuthor,
		"title":      article.ArticleTitle,
		"table_name": article.TableName(),
	})
	return IfIsExisted(temp)
}
