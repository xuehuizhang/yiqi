package channelService

import (
	"com.yiqi/dao/channelDao"
	"com.yiqi/model"
	"com.yiqi/viewModel"
)

func GetAll(category int) []viewModel.Channel {
	res := make([]viewModel.Channel, 0)
	selectFields := []string{"id", "name"}
	querySql := "status=0 and display=0 and category=? and parent_id=0"
	orderSql := "sort asc"
	params := []interface{}{category}

	parentList := channelDao.Find(selectFields, querySql, orderSql, params, 0, 9999)

	querySql = "status=0 and display=0 and category=? and parent_id=?"
	for _, p := range parentList {
		vc := viewModel.Channel{
			Id:     p.Id,
			Name:   p.Name,
			Childs: make([]viewModel.ChannelChild, 0),
		}
		childs := channelDao.Find(selectFields, querySql, orderSql, []interface{}{category, p.Id}, 0, 9999)
		for _, c := range childs {
			vc.Childs = append(vc.Childs, viewModel.ChannelChild{
				Id:   c.Id,
				Name: c.Name,
			})
		}
		res = append(res, vc)
	}

	return res
}

func Save(name string, parentId int64) (int64, error) {
	c := &model.Channel{
		Category: 0,
		ParentId: parentId,
		Name:     name,
		Sort:     0,
		Display:  0,
		Status:   0,
	}
	err := channelDao.Save(c)
	if err != nil {
		return 0, err
	}
	return c.Id, nil
}
