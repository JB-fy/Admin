// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId     uint        `json:"userId"     ` // 用户ID
	Phone      string      `json:"phone"      ` // 手机
	Account    string      `json:"account"    ` // 账号
	Password   string      `json:"password"   ` // 密码。md5保存
	Salt       string      `json:"salt"       ` // 加密盐
	Nickname   string      `json:"nickname"   ` // 昵称
	Avatar     string      `json:"avatar"     ` // 头像
	Gender     uint        `json:"gender"     ` // 性别：0未设置 1男 2女
	Birthday   *gtime.Time `json:"birthday"   ` // 生日
	Address    string      `json:"address"    ` // 详细地址
	IdCardName string      `json:"idCardName" ` // 身份证姓名
	IdCardNo   string      `json:"idCardNo"   ` // 身份证号码
	IsStop     uint        `json:"isStop"     ` // 停用：0否 1是
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
}
