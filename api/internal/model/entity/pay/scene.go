// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Scene is the golang structure for table scene.
type Scene struct {
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	PayId     uint        `json:"payId"     orm:"pay_id"     ` // 支付ID
	PayScene  uint        `json:"payScene"  orm:"pay_scene"  ` // 支付场景：0APP 1H5 2扫码 10微信小程序 11微信公众号 20支付宝小程序
}
