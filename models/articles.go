package models

type ArticleFrom struct {
	//ArticlesId int `json:"articles_id"`
	ArticleTitle    string `json:"article_title" binding:"required" gorm:"title"`
	ArticleContent  string `json:"article_content" binding:"required" gorm:"content"`
	ArticleAuthor   string `json:"article_author" binding:"required" gorm:"author"`
	ArticleCategory string `json:"article_category" binding:"required" gorm:"category"`
}

func (a ArticleFrom) TableName() string {
	return "article"
}

type ArticleForm struct {
}
