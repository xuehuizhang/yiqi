package model

type Channel struct {
	BaseModel
	Category int    `json:"category"`
	ParentId int64  `json:"parent_id"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	Display  int    `json:"display"`
	Status   int    `json:"status"`
}

func (*Channel) TableName() string {
	return "channel"
}
