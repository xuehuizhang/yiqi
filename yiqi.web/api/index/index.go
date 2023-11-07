package index

import (
	"com.yiqi/service/articleService"
	"com.yiqi/service/bannerService"
	"com.yiqi/service/categoryService"
	"com.yiqi/service/tagService"
	"com.yiqi/utils"
	"com.yiqi/viewModel"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func Index(ctx *gin.Context) {
	pageIndexStr := ctx.Query("page")
	cateStr := ctx.Query("cate")
	pageIndex, _ := strconv.Atoi(pageIndexStr)
	if pageIndex == 0 {
		pageIndex = 1
	}
	pageSize := 10

	cate, _ := strconv.Atoi(cateStr)
	if cate == 0 {
		cate = 1
	}

	if pageIndex == 0 {
		pageIndex = 1
	}

	total := articleService.TotalArticle(cate)

	totalPage := int(math.Ceil(float64(total) / float64(pageSize)))
	if pageIndex > totalPage {
		pageIndex = totalPage
	}

	paginator := utils.CreatePaginator(pageIndex, pageSize, int(total))

	channelList := categoryService.GetAll(int64(cate))

	list := articleService.ArticleList(pageIndex, pageSize, cate)

	tdk := viewModel.Tdk{
		Title:   viewModel.Title,
		Des:     viewModel.Index_Description,
		KeyWord: viewModel.Index_Keyword,
	}

	//srcByte, _ := ioutil.ReadFile(`./static/img/yiqibishe_logo.png`)

	//logoData := base64.StdEncoding.EncodeToString(srcByte)

	topBannerList := bannerService.GetBannerByPosition(0)
	leftBannerList := bannerService.GetBannerByPosition(1)
	personalInfo := viewModel.BannerPersonal{}
	if len(leftBannerList) > 0 {
		personalInfo.ImgUrl = leftBannerList[0].ImgUrl
		personalInfo.Intro = leftBannerList[0].Intro
	}

	tagList := tagService.GetAll()
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{"list": list, "paginator": paginator, "categoryList": channelList, "curCate": cate, "tdk": tdk, "topBannerList": topBannerList, "personalInfo": personalInfo, "tagList": tagList})
}

func Tag(ctx *gin.Context) {
	strTid := ctx.Query("tid")
	ser := ctx.Query("ser")
	tagId, _ := strconv.Atoi(strTid)

	list := articleService.ArticleListByTagId(int64(tagId), ser)
	leftBannerList := bannerService.GetBannerByPosition(1)
	personalInfo := viewModel.BannerPersonal{}
	if len(leftBannerList) > 0 {
		personalInfo.ImgUrl = leftBannerList[0].ImgUrl
		personalInfo.Intro = leftBannerList[0].Intro
	}
	tagList := tagService.GetAll()
	ctx.HTML(http.StatusOK, "index/list.html", gin.H{"list": list, "personalInfo": personalInfo, "tagList": tagList, "tagId": tagId, "ser": ser})
}
