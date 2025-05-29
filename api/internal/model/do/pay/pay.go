// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package pay

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Pay is the golang structure of table pay for DAO operations like Where/Data.
type Pay struct {
	g.Meta      `orm:"table:pay, do:true"`
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	PayId       interface{} // 支付ID
	PayName     interface{} // 名称
	PayType     interface{} // 类型：0支付宝 1微信
	PayConfig   interface{} // 配置。JSON格式，根据类型设置
	PayRate     interface{} // 费率
	TotalAmount interface{} // 总额
	Balance     interface{} // 余额
	Remark      interface{} // 备注
}
