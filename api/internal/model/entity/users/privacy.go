// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Privacy is the golang structure for table privacy.
type Privacy struct {
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       ` // 更新时间
	UserId         uint        `json:"userId"         orm:"user_id"          ` // 用户ID
	Password       string      `json:"password"       orm:"password"         ` // 密码。md5保存
	Salt           string      `json:"salt"           orm:"salt"             ` // 密码盐
	IdCardNo       string      `json:"idCardNo"       orm:"id_card_no"       ` // 身份证号码
	IdCardName     string      `json:"idCardName"     orm:"id_card_name"     ` // 身份证姓名
	IdCardGender   uint        `json:"idCardGender"   orm:"id_card_gender"   ` // 身份证性别：0未设置 1男 2女
	IdCardBirthday *gtime.Time `json:"idCardBirthday" orm:"id_card_birthday" ` // 身份证生日
	IdCardAddress  string      `json:"idCardAddress"  orm:"id_card_address"  ` // 身份证地址
}
