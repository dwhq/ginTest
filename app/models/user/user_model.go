package user

import (
	"gintest/app/models"
	"time"
)

type User struct {
	models.BaseModel
	Account         string    `gorm:"column:account;NOT NULL" json:"account,omitempty"`                                 // 会员编号
	Mobile          string    `gorm:"column:mobile;NOT NULL" json:"mobile,omitempty"`                                   // 手机号码或者邮箱
	Nickname        string    `gorm:"column:nickname;NOT NULL" json:"nickname,omitempty"`                               // 昵称
	AvatarUrl       string    `gorm:"column:avatar_url" json:"avatar_url,omitempty"`                                    // 头像
	Gender          int       `gorm:"column:gender;default:0;NOT NULL" json:"gender,omitempty"`                         // 性别
	Signature       string    `gorm:"column:signature" json:"signature,omitempty"`                                      // 用户签名
	Password        string    `gorm:"column:password;NOT NULL" json:"password,omitempty"`                               // 会员密码
	PayPass         string    `gorm:"column:pay_pass" json:"pay_pass,omitempty"`                                        // 安全密码
	ReId            int       `gorm:"column:re_id;default:0;NOT NULL" json:"re_id,omitempty"`                           // 推荐
	RePath          string    `gorm:"column:re_path;NOT NULL" json:"re_path,omitempty"`                                 // 关系树
	ReName          string    `gorm:"column:re_name;NOT NULL" json:"re_name,omitempty"`                                 // 推荐人编号
	ReLevel         int       `gorm:"column:re_level;default:0;NOT NULL" json:"re_level,omitempty"`                     // 层级
	IsDelete        int       `gorm:"column:is_delete;default:0;NOT NULL" json:"is_delete,omitempty"`                   // 是否删除
	ReNum           int       `gorm:"column:re_num;default:0;NOT NULL" json:"re_num,omitempty"`                         // 直推总人数
	RePerformance   string    `gorm:"column:re_performance;default:0.00;NOT NULL" json:"re_performance,omitempty"`      // 直推总业绩
	TeamNum         int       `gorm:"column:team_num;default:0;NOT NULL" json:"team_num,omitempty"`                     // 团队总人数
	TeamPerformance string    `gorm:"column:team_performance;default:0.00;NOT NULL" json:"team_performance,omitempty"`  // 团队业绩
	IsLock          int       `gorm:"column:is_lock;default:0;NOT NULL" json:"is_lock,omitempty"`                       // 是否锁定
	SumIncome       string    `gorm:"column:sum_income;default:0.00;NOT NULL" json:"sum_income,omitempty"`              // 总收益
	RunRegisterAt   time.Time `gorm:"column:run_register_at" json:"run_register_at,omitempty"`                          // 执行任务时间
	JobRegister     int       `gorm:"column:job_register;default:0;NOT NULL" json:"job_register,omitempty"`             // 注册任务
	CellHub         string    `gorm:"column:cell_hub;default:0.0000;NOT NULL" json:"cell_hub,omitempty"`                // 小区业绩
	LargeHub        string    `gorm:"column:large_hub;default:0.0000;NOT NULL" json:"large_hub,omitempty"`              // 大区业绩
	IsCertification int       `gorm:"column:is_certification;default:-1;NOT NULL" json:"is_certification,omitempty"`    // 是否实名认证-1未提交 0审核中 1=>成功 2失败
	BuyNumber       int       `gorm:"column:buy_number;default:1;NOT NULL" json:"buy_number,omitempty"`                 // 最大购买数量
	models.CommonTimestampsField
}

func (m *User) TableName() string {
	return "ds_user"
}