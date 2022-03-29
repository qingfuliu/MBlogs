package models

import "time"

type ArticleForm struct {
	//ArticlesId int `json:"articles_id"`
	ArticleTitle    string `json:"article_title" binding:"required" gorm:"column:title"`
	ArticleContent  string `json:"article_content" binding:"required" gorm:"column:content"`
	ArticleAuthor   string `json:"article_author,omitempty"  gorm:"column:author"`
	ArticleCategory string `json:"article_category" binding:"required" gorm:"column:category"`
}

func (a *ArticleForm) TableName() string {
	return "articles"
}

type Article struct {
	*ArticleForm `gorm:"embedded"`
	CreatedAt    time.Time `gorm:"column:create_date"`
}
