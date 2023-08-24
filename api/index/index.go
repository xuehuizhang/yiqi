package index

import (
	"com.yiqi/service/articleService"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(ctx *gin.Context) {
	pageIndexStr := ctx.Query("pageIndex")
	pageIndex, _ := strconv.Atoi(pageIndexStr)
	if pageIndex == 0 {
		pageIndex = 1
	}
	pageSize := 10

	list := articleService.ArticleList(pageIndex, pageSize)
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{"list": list, "curIndex": pageIndex})
}
