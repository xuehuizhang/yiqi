package articleDetailDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"log"
)

func GetByArticleId(aid int64) *model.ArticleDetailModel {
	var article model.ArticleDetailModel
	err := db.DB.Where("article_id=?", aid).First(&article).Error
	if err != nil {
		log.Print("根据文章ID查询文章失败：", aid, err.Error())
		return nil
	}
	return &article
}
