package product

import "gintest/app/models"

type Product struct {
	models.BaseModel
	AdminId       int       `gorm:"column:admin_id;default:0;NOT NULL" json:"admin_id,omitempty"`                // 管理员id
	CategoryId    int       `gorm:"column:category_id;default:0;NOT NULL" json:"category_id,omitempty"`          // 分类id
	CategoryName  string    `gorm:"column:category_name" json:"category_name,omitempty"`                         // 分类名称
	Name          string    `gorm:"column:name;NOT NULL" json:"name,omitempty"`                                  // 商品名
	MainPicture   string    `gorm:"column:main_picture" json:"main_picture,omitempty"`                           // 商品主图
	Picture       string    `gorm:"column:picture" json:"picture,omitempty"`                                     // 商品图片列表
	Explain       string    `gorm:"column:explain" json:"explain,omitempty"`                                     // 说明
	IsSpec        int       `gorm:"column:is_spec;default:0;NOT NULL" json:"is_spec,omitempty"`                  // 是否多规格，1/0
	OriginalPrice string    `gorm:"column:original_price;default:0.00;NOT NULL" json:"original_price,omitempty"` // 原价价格
	Price         string    `gorm:"column:price;default:0.00;NOT NULL" json:"price,omitempty"`                   // 现价
	Freight       string    `gorm:"column:freight;default:0.00" json:"freight,omitempty"`                        // 运费
	Fee           string    `gorm:"column:fee;default:0.00;NOT NULL" json:"fee,omitempty"`                       // 交易手续费
	Integral      string    `gorm:"column:integral;NOT NULL" json:"integral,omitempty"`                          // 赠送积分数量
	CurStatus     int       `gorm:"column:cur_status;default:0;NOT NULL" json:"cur_status,omitempty"`            // 上下架,1/0
	CurType       int       `gorm:"column:cur_type;default:0;NOT NULL" json:"cur_type,omitempty"`                // 1=>团购商品 2=>展示商品
	CostPrice     string    `gorm:"column:cost_price;default:0.00" json:"cost_price,omitempty"`                  // 成本价，计算利润
	Stock         int       `gorm:"column:stock" json:"stock,omitempty"`                                         // 库存
	Sort          int       `gorm:"column:sort;default:0;NOT NULL" json:"sort,omitempty"`                        // 排序
	VirtualSales  int       `gorm:"column:virtual_sales;default:0;NOT NULL" json:"virtual_sales,omitempty"`      // 虚拟销量
	Sales         int       `gorm:"column:sales;default:0;NOT NULL" json:"sales,omitempty"`                      // 销量
	Grade         string    `gorm:"column:grade;default:5.0;NOT NULL" json:"grade,omitempty"`                    // 评分
	UsdtNumber    string    `gorm:"column:usdt_number;default:0.00" json:"usdt_number,omitempty"`                // 资产包数量
	IsNew         int       `gorm:"column:is_new;default:0;NOT NULL" json:"is_new,omitempty"`                    // 是否新品
	models.CommonTimestampsField
	models.SoftDeletes
}
func (Product) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "ds_product"
}
