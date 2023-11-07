package tagService

import (
	"com.yiqi/dao/categoryDao"
	"com.yiqi/viewModel"
)

func GetAll() []viewModel.TagViewModel {
	res := make([]viewModel.TagViewModel, 0)
	selectFields := []string{"id", "name"}
	querySql := "status=0 "
	orderSql := ""

	parentList := categoryDao.Find(selectFields, querySql, orderSql, nil, 0, 9999)

	for _, p := range parentList {
		vc := viewModel.TagViewModel{
			Id:   p.Id,
			Name: p.Name,
		}
		res = append(res, vc)
	}

	return res
}
