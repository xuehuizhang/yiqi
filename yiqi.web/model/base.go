package model

type BaseModel struct {
	Id         int64 `json:"id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
	DeleteTime int64 `json:"delete_time"`
}
