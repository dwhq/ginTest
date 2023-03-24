package shop

import (
	"gintest/app/models/product"
	"gintest/pkg/app"
	"gintest/pkg/database"
	"gintest/pkg/paginator"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context,perPage int,categoryId int, curType int) map[string]any  {
	var products []product.Product
	query :=database.DB.Select("id", "cur_type","name", "category_id", "category_name", "main_picture", "is_new", "original_price", "price", "cur_type", "freight", "stock", "sales", "grade").
		Where("cur_status=1 and cur_type=2")
	if categoryId >0 {
		query.Where("category_id=?", categoryId)
	}
	if curType != 1 {
		curType =2
	}
	query.Where("cur_type =?", curType)
	query.Order("sort desc")
	data :=paginator.Paginate(
		c,
		query.Model(product.Product{}),
		&products,
		app.V1URL(database.TableName(&product.Product{})),
		perPage,
	)
	return map[string]any{
		"data":products,
		"perPage":data,
	}
}
