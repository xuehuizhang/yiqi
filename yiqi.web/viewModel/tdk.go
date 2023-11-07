package viewModel

const (
	Title string = "一起毕设"

	Index_Keyword     = "一起毕设,毕设,毕业设计,Go,go,golang,java,Java,.net,.Net,Asp.Net,asp.net,Java毕业设计,golang毕业设计,前端毕业设计,后端毕业设计,大学毕业设计,毕业季,毕业"
	Index_Description = "一起毕设，帮你解决毕业设计难题，迈出职业生涯第一步."

	Detail_Keyword = "一起毕设,毕设,毕业设计,Go,go,golang,java,Java,.net,.Net,Asp.Net,asp.net,Java毕业设计,golang毕业设计,前端毕业设计,后端毕业设计,大学毕业设计,毕业季,毕业"

	Detail_Description = "一起毕设，帮你解决毕业设计难题，迈出职业生涯第一步."
)

type Tdk struct {
	Title   string `json:"title"`
	Des     string `json:"des"`
	KeyWord string `json:"key_word"`
}
