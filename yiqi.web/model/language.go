package model

type Language struct {
	BaseModel
	Name string `json:"name"`
}

func (*Language) TableName() string {
	return "language"
}
