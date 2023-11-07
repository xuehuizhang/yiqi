package categoryDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"gorm.io/gorm"
	"log"
)

func Save(channel *model.Category) error {
	return db.DB.Save(channel).Error
}

func GetById(id int64) *model.Category {
	var category model.Category
	err := db.DB.First(&category, id).Error
	if err != nil {
		log.Print("根据ID查询Category失败：", id, err.Error())
		return nil
	}
	return &category
}

func Find(selectFields []string, querySql string, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.Category {
	var list []*model.Category
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询Channel失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}
