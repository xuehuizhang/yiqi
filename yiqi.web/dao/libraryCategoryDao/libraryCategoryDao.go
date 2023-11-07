package LibraryCategoryDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"gorm.io/gorm"
	"log"
)

func Save(LibraryCategory *model.LibraryCategory) error {
	return db.DB.Save(LibraryCategory).Error
}

func GetById(id int64) *model.LibraryCategory {
	var LibraryCategory model.LibraryCategory
	err := db.DB.First(&LibraryCategory, id).Error
	if err != nil {
		log.Print("根据ID查询LibraryCategory失败：", id, err.Error())
		return nil
	}
	return &LibraryCategory
}

func Find(selectFields []string, querySql string, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.LibraryCategory {
	var list []*model.LibraryCategory
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询LibraryCategory失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}
