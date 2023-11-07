package model

type LibraryCategory struct {
	BaseModel
	Name string `json:"name"`
}

func (*LibraryCategory) TableName() string {
	return "library_category"
}
