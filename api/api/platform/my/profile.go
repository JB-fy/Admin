package my

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------个人信息 开始--------*/
type ProfileInfoReq struct {
	g.Meta `path:"/profile/info" method:"post" tags:"平台后台/我的" sm:"个人信息"`
}

type ProfileInfoRes struct {
	Info ProfileInfo `json:"info" dc:"详情"`
}

type ProfileInfo struct {
	AdminId  uint   `json:"admin_id" dc:"管理员ID"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Phone    string `json:"phone" dc:"手机"`
	Email    string `json:"email" dc:"邮箱"`
	Account  string `json:"account" dc:"账号"`
	IsSuper  uint   `json:"is_super" dc:"超管：0否 1是"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateReq struct {
	g.Meta          `path:"/profile/update" method:"post" tags:"平台后台/我的" sm:"修改个人信息"`
	Nickname        *string `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Avatar          *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	Phone           *string `json:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	Email           *string `json:"email,omitempty" v:"max-length:60|email" dc:"邮箱"`
	Account         *string `json:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	Password        *string `json:"password,omitempty" v:"size:32" dc:"新密码。加密后发送，公式：md5(新密码)"`
	PasswordToCheck *string `json:"password_to_check,omitempty" v:"required-with:Phone,Email,Account,Password|size:32|different:Password" dc:"旧密码。加密后发送，公式：md5(新密码)。修改手机，邮箱，账号，密码用，password_to_check,sms_code_to_password,email_code_to_password传一个即可"`
	// SmsCodeToPassword      *string `json:"sms_code_to_password,omitempty" v:"size:4" dc:"短信验证码。修改密码用，password_to_check,sms_code_to_password,email_code_to_password传一个即可"`
	SmsCodeToBindPhone *string `json:"sms_code_to_bind_phone,omitempty" v:"required-with:Phone|size:4" dc:"短信验证码。绑定手机用"`
	// SmsCodeToUnbingPhone *string `json:"sms_code_to_unbing_phone,omitempty" v:"size:4" dc:"短信验证码。解绑手机用"`
	// EmailCodeToPassword    *string `json:"email_code_to_password,omitempty" v:"size:4" dc:"邮箱验证码。修改密码用，password_to_check,sms_code_to_password,email_code_to_password传一个即可"`
	EmailCodeToBindEmail *string `json:"email_code_to_bind_email,omitempty" v:"required-with:Email|size:4" dc:"邮箱验证码。绑定邮箱用"`
	// EmailCodeToUnbingEmail *string `json:"email_code_to_unbing_email,omitempty" v:"size:4" dc:"邮箱验证码。解绑邮箱用"`
}

/*--------修改个人信息 结束--------*/
