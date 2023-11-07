package channelDao

import (
	"com.yiqi/dao/db"
	"com.yiqi/model"
	"gorm.io/gorm"
	"log"
)

func Save(channel *model.Channel) error {
	return db.DB.Save(channel).Error
}

func GetById(id int64) *model.Channel {
	var channel model.Channel
	err := db.DB.First(&channel, id).Error
	if err != nil {
		log.Print("根据ID查询Channel失败：", id, err.Error())
		return nil
	}
	return &channel
}

func Find(selectFields []string, querySql string, orderSql string, params []interface{}, pageIndex, pageSize int) []*model.Channel {
	var list []*model.Channel
	err := db.DB.Select(selectFields).Where(querySql, params...).Order(orderSql).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print("分页查询Channel失败,", selectFields, querySql, params, pageIndex, pageSize)
		return nil
	}
	return list
}
