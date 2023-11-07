package viewModel

type ArticleList struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Intro        string `json:"intro"`
	Cover        string `json:"cover"`
	ViewCount    int    `json:"view_count"`
	CollectCount int    `json:"collect_count"`
	CreateTime   string `json:"create_time"`
}

type ArticleDetail struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	ViewCount    int    `json:"view_count"`
	CollectCount int    `json:"collect_count"`
	CreeateTime  string `json:"creeate_time"`
}
