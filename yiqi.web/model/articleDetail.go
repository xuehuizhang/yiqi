package model

type ArticleDetailModel struct {
	BaseModel
	ArticleId int64  `json:"article_id"`
	Content   string `json:"content"`
}

func (*ArticleDetailModel) TableName() string {
	return "article_detail"
}
