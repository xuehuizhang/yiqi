package languageDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"gorm.io/gorm"
	"log"
)

func Save(Language *model.Language) error {
	return db.DB.Save(Language).Error
}

func GetById(id int64) *model.Language {
	var Language model.Language
	err := db.DB.First(&Language, id).Error
	if err != nil {
		log.Print("根据ID查询Language失败：", id, err.Error())
		return nil
	}
	return &Language
}

func Find(selectFields []string, querySql string, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.Language {
	var list []*model.Language
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询Language失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}
