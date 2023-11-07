package model

type UserModel struct {
	Nick     string `json:"nick"`
	FaceUrl  string `json:"face_url"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Status   int    `json:"status"`
}
