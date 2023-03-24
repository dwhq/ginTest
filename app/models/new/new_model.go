package new

import "gintest/app/models"

type New struct {
	models.BaseModel
	CatId        uint `json:"cat_id,omitempty"`
	Title        string `json:"title,omitempty"`
	TitleEn        string `json:"title_en,omitempty"`
	Content      string `json:"content,omitempty"`
	ContentEn      string `json:"content_en,omitempty"`
	CurStatus    int8 `json:"cur_status,omitempty"`
	models.CommonTimestampsField
}