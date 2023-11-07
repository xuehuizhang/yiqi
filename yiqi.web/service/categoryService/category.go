package categoryService

import (
	"com.yiqi/dao/categoryDao"
	"com.yiqi/viewModel"
)

func GetAll(category int64) []viewModel.Category {
	res := make([]viewModel.Category, 0)
	selectFields := []string{"id", "name"}
	querySql := "status=0 "
	orderSql := "sort asc"

	parentList := categoryDao.Find(selectFields, querySql, orderSql, nil, 0, 9999)

	for index, p := range parentList {
		vc := viewModel.Category{
			Id:    p.Id,
			Name:  p.Name,
			IsCur: false,
		}
		if p.Id == category {
			vc.IsCur = true
		}
		if category == 0 && index == 0 {
			vc.IsCur = true
		}
		res = append(res, vc)
	}

	return res
}
