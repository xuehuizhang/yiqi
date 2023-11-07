package doc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/doc.html", gin.H{})
}

func Com(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/com.html", gin.H{})
}
