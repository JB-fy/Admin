// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Pay is the golang structure for table pay.
type Pay struct {
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   ` // 更新时间
	PayId       uint        `json:"payId"       orm:"pay_id"       ` // 支付ID
	PayName     string      `json:"payName"     orm:"pay_name"     ` // 名称
	PayType     uint        `json:"payType"     orm:"pay_type"     ` // 类型：0支付宝 1微信
	PayConfig   string      `json:"payConfig"   orm:"pay_config"   ` // 配置。JSON格式，根据类型设置
	PayRate     float64     `json:"payRate"     orm:"pay_rate"     ` // 费率
	TotalAmount float64     `json:"totalAmount" orm:"total_amount" ` // 总额
	Balance     float64     `json:"balance"     orm:"balance"      ` // 余额
	Remark      string      `json:"remark"      orm:"remark"       ` // 备注
}
