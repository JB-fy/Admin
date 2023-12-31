// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta     `orm:"table:user, do:true"`
	UserId     interface{} // 用户ID
	Phone      interface{} // 手机
	Account    interface{} // 账号
	Password   interface{} // 密码。md5保存
	Salt       interface{} // 加密盐
	Nickname   interface{} // 昵称
	Avatar     interface{} // 头像
	Gender     interface{} // 性别：0未设置 1男 2女
	Birthday   *gtime.Time // 生日
	Address    interface{} // 详细地址
	IdCardName interface{} // 身份证姓名
	IdCardNo   interface{} // 身份证号码
	IsStop     interface{} // 停用：0否 1是
	UpdatedAt  *gtime.Time // 更新时间
	CreatedAt  *gtime.Time // 创建时间
}
