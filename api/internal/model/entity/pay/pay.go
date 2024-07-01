// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Pay is the golang structure for table pay.
type Pay struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	IsStop      uint        `json:"isStop"      orm:"is_stop"      ` // 停用：0否 1是
	PayId       uint        `json:"payId"       orm:"pay_id"       ` // 支付ID
	PayName     string      `json:"payName"     orm:"pay_name"     ` // 名称
	PayIcon     string      `json:"payIcon"     orm:"pay_icon"     ` // 图标
	PayType     uint        `json:"payType"     orm:"pay_type"     ` // 类型：0支付宝 1微信
	PayConfig   string      `json:"payConfig"   orm:"pay_config"   ` // 配置。根据pay_type类型设置
	PayRate     float64     `json:"payRate"     orm:"pay_rate"     ` // 费率
	TotalAmount float64     `json:"totalAmount" orm:"total_amount" ` // 总额
	Balance     float64     `json:"balance"     orm:"balance"      ` // 余额
	Sort        uint        `json:"sort"        orm:"sort"         ` // 排序值。从大到小排序
	Remark      string      `json:"remark"      orm:"remark"       ` // 备注
}
