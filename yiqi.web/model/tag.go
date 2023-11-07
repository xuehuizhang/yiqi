package model

type TagModel struct {
	BaseModel
	Name string `json:"name"`
}

func (*TagModel) TableName() string {
	return "tag"
}
