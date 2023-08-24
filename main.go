package main

import (
	"com.yiqi/api/detail"
	"com.yiqi/api/index"
	"com.yiqi/initialize"
	"github.com/gin-gonic/gin"

	"net/http"
)

func init() {
	//初始化DB
	initialize.InitDb()
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("./static"))

	r.GET("/", index.Index)
	r.GET("/detail/:id", detail.Detail)
	r.GET("/article/detail", detail.ArticleDetail)
	r.Run(":9099")
}
