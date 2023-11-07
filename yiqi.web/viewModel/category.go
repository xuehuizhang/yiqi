package viewModel

type Category struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	IsCur bool   `json:"is_cur"`
}
