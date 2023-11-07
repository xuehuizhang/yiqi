package bannerDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"gorm.io/gorm"
	"log"
)

func Save(banner *model.Banner) error {
	return db.DB.Save(banner).Error
}

func GetById(id int64) *model.Banner {
	var banner model.Banner
	err := db.DB.First(&banner, id).Error
	if err != nil {
		log.Print("根据ID查询Banner失败：", id, err.Error())
		return nil
	}
	return &banner
}

func Find(selectFields []string, querySql string, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.Banner {
	var list []*model.Banner
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询Banner失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}
