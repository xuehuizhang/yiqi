package model

type Library struct {
	BaseModel
	Name       string `json:"name"`
	Url        string `json:"url"`
	GitUrl     string `json:"git_url"`
	IconUrl    string `json:"icon_url"`
	Score      int    `json:"score"`
	Intro      string `json:"intro"`
	LibCateId  int64  `json:"lib_cate_id"`
	LanguageId int64  `json:"language_id"`
}

func (*Library) TableName() string {
	return "library"
}
