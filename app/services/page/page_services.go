package page

import (
	"gintest/app/models/new"
	"gintest/app/models/product"
	"gintest/pkg/app"
	"gintest/pkg/database"
)

func Index()map[string]any  {
	var banners = [...]string{
		app.URL("assets/images/banner-1.png"),
		app.URL("assets/images/banner-2.png"),
		app.URL("assets/images/banner-3.png"),
	}
	var video = map[string]string{
		"cover_img":app.URL("assets/images/cover_img-2.png"),
		"video_url":app.URL("assets/video/cover.mp4"),
	}
	var newMode []new.New
	var products []*product.Product
	database.DB.Select("id,title,title_en").Where("cat_id=1 and cur_status=?","1").Order("id desc").Find(&newMode)
	database.DB.Select("id,name,main_picture,price").Where("cur_status=1 and cur_type=2").Order("sort desc").Find(&products)
	for _, v := range products {
		v.MainPicture = app.URL("assets/images" + v.MainPicture)
	}
	return map[string]any{
		"banners": banners,
		"video": video,
		"news": newMode,
		"products": products,
	}
}
