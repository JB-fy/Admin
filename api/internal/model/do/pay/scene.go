// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure of table pay_scene for DAO operations like Where/Data.
type Scene struct {
	g.Meta    `orm:"table:pay_scene, do:true"`
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	PayId     interface{} // 支付ID
	PayScene  interface{} // 支付场景：0APP 1H5 2扫码 10微信小程序 11微信公众号 20支付宝小程序
}
