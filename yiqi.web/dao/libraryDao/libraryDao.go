package libraryDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"gorm.io/gorm"
	"log"
)

func Save(Library *model.Library) error {
	return db.DB.Save(Library).Error
}

func GetById(id int64) *model.Library {
	var Library model.Library
	err := db.DB.First(&Library, id).Error
	if err != nil {
		log.Print("根据ID查询Library失败：", id, err.Error())
		return nil
	}
	return &Library
}

func Find(selectFields []string, querySql string, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.Library {
	var list []*model.Library
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询Library失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}
