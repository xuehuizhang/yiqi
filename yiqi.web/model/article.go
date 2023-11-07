package model

type ArticleModel struct {
	BaseModel
	Title        string `json:"title"`
	Intro        string `json:"intro"`
	Cover        string `json:"cover"`
	ViewCount    int    `json:"view_count"`
	CollectCount int    `json:"collect_count"`
	CategoryId   int64  `json:"category_id"`
	UserId       int64  `json:"user_id"`
	Status       int    `json:"status"`
	TagId        int64  `json:"tag_id"`
}

func (*ArticleModel) TableName() string {
	return "article"
}
