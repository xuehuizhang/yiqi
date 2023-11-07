package articleDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func GetById(id int64) *model.ArticleModel {
	var article model.ArticleModel
	err := db.DB.First(&article, id).Error
	if err != nil {
		log.Print("根据ID查询文章失败：", id, err.Error())
		return nil
	}
	return &article
}

func Find(selectFields []string, querySql, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.ArticleModel {
	var list []*model.ArticleModel
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询文章失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}

func Count(querySql string, params []interface{}) int64 {
	fmt.Println(querySql, params)
	var total int64
	err := db.DB.Model(&model.ArticleModel{}).Where(querySql, params...).Count(&total).Error
	if err != nil {
		log.Print("查询数量失败,", err)
	}
	return total
}

func Update(querySql string, params []interface{}, fields map[string]interface{}) error {
	return db.DB.Model(&model.ArticleModel{}).Where(querySql, params).Updates(fields).Error
}
