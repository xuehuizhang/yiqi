package bannerService

import (
	"com.yiqi/dao/bannerDao"
	"com.yiqi/viewModel"
)

func GetBannerByPosition(position int) []viewModel.BannerViewModel {
	res := make([]viewModel.BannerViewModel, 0)
	selectFields := []string{"id", "intro", "jump_url", "img_url"}
	querySql := "status=0 and position=?"
	orderSql := "sort asc"

	parentList := bannerDao.Find(selectFields, querySql, orderSql, []interface{}{position}, 0, 9999)

	for _, p := range parentList {
		//srcByte, _ := ioutil.ReadFile("/Users/chenxin/data" + p.ImgUrl)
		//logoData := base64.StdEncoding.EncodeToString(srcByte)
		vc := viewModel.BannerViewModel{
			Id:      p.Id,
			Intro:   p.Intro,
			JumpUrl: p.JumpUrl,
			ImgUrl:  p.ImgUrl,
			//BaseUrl: logoData,
		}
		res = append(res, vc)
	}
	return res
}
