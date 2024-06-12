package api

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
	Account  string `json:"account" dc:"账号"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateReq struct {
	g.Meta               `path:"/profile/update" method:"post" tags:"平台后台/我的" sm:"修改个人信息"`
	Nickname             *string `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Avatar               *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	Phone                *string `json:"phone,omitempty" v:"phone" dc:"手机"`
	Email                *string `json:"email,omitempty" v:"email" dc:"邮箱"`
	Account              *string `json:"account,omitempty" v:"max-length:30|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	Password             *string `json:"password,omitempty" v:"size:32" dc:"新密码。加密后发送，公式：md5(新密码)"`
	PasswordToCheck      *string `json:"password_to_check,omitempty" v:"required-with:Account,Phone,Password|size:32|different:Password" dc:"旧密码。加密后发送，公式：md5(新密码)。修改账号，手机，密码时必填"`
	SmsCodeToBindPhone   *string `json:"sms_code_to_bind_phone,omitempty" v:"required-with:Phone|size:4" dc:"短信验证码。修改手机时必填"`
	EmailCodeToBindEmail *string `json:"email_code_to_bind_email,omitempty" v:"required-with:Email|size:4" dc:"邮箱验证码。修改邮箱时必填"`
}

/*--------修改个人信息 结束--------*/
