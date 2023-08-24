package articleService

import (
	"com.yiqi/dao/articleDao"
	"com.yiqi/dao/articleDetailDao"
	"com.yiqi/viewModel"
	"time"
)

func ArticleList(pageIndex, pageSize int) []viewModel.ArticleList {
	viewList := make([]viewModel.ArticleList, 0)

	if pageIndex == 0 {
		pageIndex = 1
	}

	if pageSize == 0 {
		pageSize = 10
	}

	selectFields := []string{"id", "title", "intro", "cover", "view_count", "collect_count", "create_time"}
	querySql := "status=?"
	params := []interface{}{1}
	list := articleDao.Find(selectFields, querySql, params, pageIndex, pageSize)

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
	}
	return vo
}
