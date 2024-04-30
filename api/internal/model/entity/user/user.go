// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	UserId      uint        `json:"userId"      orm:"user_id"        ` // 用户ID
	Phone       string      `json:"phone"       orm:"phone"          ` // 手机
	Account     string      `json:"account"     orm:"account"        ` // 账号
	Password    string      `json:"password"    orm:"password"       ` // 密码。md5保存
	Salt        string      `json:"salt"        orm:"salt"           ` // 密码盐
	Nickname    string      `json:"nickname"    orm:"nickname"       ` // 昵称
	Avatar      string      `json:"avatar"      orm:"avatar"         ` // 头像
	Gender      uint        `json:"gender"      orm:"gender"         ` // 性别：0未设置 1男 2女
	Birthday    *gtime.Time `json:"birthday"    orm:"birthday"       ` // 生日
	Address     string      `json:"address"     orm:"address"        ` // 详细地址
	OpenIdOfWx  string      `json:"openIdOfWx"  orm:"open_id_of_wx"  ` // 微信openId
	UnionIdOfWx string      `json:"unionIdOfWx" orm:"union_id_of_wx" ` // 微信unionId
	IdCardName  string      `json:"idCardName"  orm:"id_card_name"   ` // 身份证姓名
	IdCardNo    string      `json:"idCardNo"    orm:"id_card_no"     ` // 身份证号码
	IsStop      uint        `json:"isStop"      orm:"is_stop"        ` // 停用：0否 1是
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"     ` // 更新时间
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"     ` // 创建时间
}
