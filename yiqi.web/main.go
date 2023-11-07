package main

import (
	"com.yiqi/api/detail"
	"com.yiqi/api/doc"
	"com.yiqi/api/index"
	"com.yiqi/api/library"
	"com.yiqi/dao/db"
	"com.yiqi/initialize"
	"flag"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	//初始化DB
	initialize.InitDb()
}

func main() {
	mig := flag.String("mig", "", "是否迁移数据库")
	flag.Parse()
	if *mig == "true" {
		db.Migrate()
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("./static"))
	//r.StaticFS("/img", http.Dir("/data/work/zhangxuehui/data/img"))
	r.StaticFS("/img", http.Dir("/Users/chenxin/data/img"))

	r.GET("/", index.Index)
	r.GET("/tag", index.Tag)
	r.GET("/detail/:id", detail.Detail)
	r.GET("/article/detail", detail.ArticleDetail)
	r.GET("/lib", library.Index)
	r.GET("/doc", doc.Index)
	r.GET("/com", doc.Com)
	r.Run("0.0.0.0:7001")
}
