package detail

import (
	"com.yiqi/service/articleService"
	"com.yiqi/viewModel"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Detail(ctx *gin.Context) {
	id := ctx.Param("id")

	idInt, _ := strconv.ParseInt(id, 10, 64)
	vo := articleService.ArticleDetail(idInt)

	tdk := viewModel.Tdk{
		Title:   viewModel.Title,
		Des:     viewModel.Detail_Description,
		KeyWord: viewModel.Detail_Keyword,
	}

	ctx.HTML(http.StatusOK, "detail/detail.html", gin.H{"detail": vo, "tdk": tdk})
}

func ArticleDetail(ctx *gin.Context) {
	id := ctx.Query("id")

	idInt, _ := strconv.ParseInt(id, 10, 64)
	vo := articleService.ArticleDetail(idInt)

	ctx.JSON(http.StatusOK, gin.H{"c": vo.Content})
}
