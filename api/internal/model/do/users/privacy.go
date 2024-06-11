// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Privacy is the golang structure of table users_privacy for DAO operations like Where/Data.
type Privacy struct {
	g.Meta         `orm:"table:users_privacy, do:true"`
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	UserId         interface{} // 用户ID
	Password       interface{} // 密码。md5保存
	Salt           interface{} // 密码盐
	IdCardNo       interface{} // 身份证号码
	IdCardName     interface{} // 身份证姓名
	IdCardGender   interface{} // 身份证性别：0未设置 1男 2女
	IdCardBirthday *gtime.Time // 身份证生日
	IdCardAddress  interface{} // 身份证地址
}
