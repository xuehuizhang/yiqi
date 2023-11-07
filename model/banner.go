package model

type Banner struct {
	BaseModel
	Position int    `json:"position"`
	JumpUrl  string `json:"jump_url"`
	Intro    string `json:"intro"`
	ImgUrl   string `json:"img_url"`
	Sort     int    `json:"sort"`
}

func (*Banner) TableName() string {
	return "banner"
}
