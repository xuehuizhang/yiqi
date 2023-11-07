package viewModel

type BannerViewModel struct {
	Id      int64  `json:"id"`
	JumpUrl string `json:"jump_url"`
	ImgUrl  string `json:"img_url"`
	Intro   string `json:"intro"`
	BaseUrl string `json:"base_url"`
}

type BannerPersonal struct {
	ImgUrl string `json:"img_url"`
	Intro  string `json:"intro"`
}
