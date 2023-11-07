package model

type Category struct {
	BaseModel
	Name   string `json:"name"`
	Sort   int    `json:"sort"`
	Status int    `json:"status"`
}

func (*Category) TableName() string {
	return "category"
}
