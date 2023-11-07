package viewModel

const (
	index_keyword     = ""
	index_description = ""

	detail_keyword     = ""
	detail_description = ""
)

type Tdk struct {
	Title   string `json:"title"`
	Des     string `json:"des"`
	KeyWord string `json:"key_word"`
}
