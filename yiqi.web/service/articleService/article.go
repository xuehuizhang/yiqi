package articleService

import (
	"com.yiqi/dao/articleDao"
	"com.yiqi/dao/articleDetailDao"
	"com.yiqi/viewModel"
	"strings"
	"time"
)

func ArticleList(pageIndex, pageSize, cate int) []viewModel.ArticleList {
	viewList := make([]viewModel.ArticleList, 0)

	if pageIndex == 0 {
		pageIndex = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	selectFields := []string{"id", "title", "intro", "cover", "view_count", "collect_count", "create_time"}
	querySql := "status=? and category_id=?"
	orderSql := "id desc"
	params := []interface{}{1, cate}
	list := articleDao.Find(selectFields, querySql, orderSql, params, pageIndex, pageSize)

	for _, a := range list {
		unix := time.Unix(a.CreateTime, 0).String()

		vo := viewModel.ArticleList{
			Id:           a.Id,
			Title:        a.Title,
			Intro:        a.Intro,
			Cover:        a.Cover,
			ViewCount:    a.ViewCount,
			CollectCount: a.CollectCount,
			CreateTime:   unix,
		}
		viewList = append(viewList, vo)
	}
	return viewList
}

func ArticleDetail(aId int64) viewModel.ArticleDetail {
	vo := viewModel.ArticleDetail{}

	article := articleDao.GetById(aId)
	articleDetail := articleDetailDao.GetByArticleId(aId)
	if article != nil && articleDetail != nil {
		vo.Id = article.Id
		vo.Title = article.Title
		vo.ViewCount = article.ViewCount
		vo.CollectCount = article.ViewCount
		vo.CreeateTime = time.Unix(article.CreateTime, 0).String()

		vo.Content = articleDetail.Content

		UpdateViewCount(article.Id, article.ViewCount+1)
	}

	return vo
}

func TotalArticle(cate int) int64 {
	return articleDao.Count("status=? and category_id=?", []interface{}{1, cate})
}

func ArticleListByTagId(tagId int64, text string) []viewModel.ArticleList {
	viewList := make([]viewModel.ArticleList, 0)

	pageIndex := 0
	pageSize := 9999

	selectFields := []string{"id", "title", "intro", "cover", "view_count", "collect_count", "create_time"}
	querySql := "status=? "
	orderSql := "id desc"
	params := []interface{}{1}
	if strings.TrimSpace(text) != "" {
		querySql += " and title like ?"
		params = append(params, "%"+text+"%")
	}

	if tagId != 0 {
		querySql += " and tag_id=? "
		params = append(params, tagId)
	}

	list := articleDao.Find(selectFields, querySql, orderSql, params, pageIndex, pageSize)

	for _, a := range list {
		unix := time.Unix(a.CreateTime, 0).String()

		vo := viewModel.ArticleList{
			Id:           a.Id,
			Title:        a.Title,
			Intro:        a.Intro,
			Cover:        a.Cover,
			ViewCount:    a.ViewCount,
			CollectCount: a.CollectCount,
			CreateTime:   unix,
		}
		viewList = append(viewList, vo)
	}
	return viewList
}

func UpdateViewCount(id int64, count int) {
	articleDao.Update("id=?", []interface{}{id}, map[string]interface{}{"view_count": count})
}
