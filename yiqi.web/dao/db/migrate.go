package db

import "com.yiqi/model"

func Migrate() {
	DB.AutoMigrate(&model.ArticleModel{}, &model.ArticleDetailModel{})
}
