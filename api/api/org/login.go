package api

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取密码盐 开始--------*/
type LoginSaltReq struct {
	g.Meta `path:"/salt" method:"post" tags:"机构后台/登录" sm:"获取密码盐"`
	api.CommonHeaderReq
	LoginName string `json:"login_name,omitempty" v:"required|max-length:60" dc:"手机/邮箱/账号"`
	AdminType uint8  `json:"admin_type" v:"in:10" d:"10" dc:"类型：10机构"`
}

/*--------获取密码盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"机构后台/登录" sm:"登录"`
	api.CommonHeaderReq
	LoginName string `json:"login_name,omitempty" v:"required|max-length:60" dc:"手机/邮箱/账号"`
	Password  string `json:"password,omitempty" v:"required-without-all:SmsCode,EmailCode|size:32" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态密码盐)+动态密码盐)"`
	SmsCode   string `json:"sms_code,omitempty" v:"required-without-all:EmailCode,Password|size:4" dc:"短信验证码"`
	EmailCode string `json:"email_code,omitempty" v:"required-without-all:SmsCode,Password|size:4" dc:"邮箱验证码"`
	AdminType uint8  `json:"admin_type" v:"in:10" d:"10" dc:"类型：10机构"`
}

/*--------登录 结束--------*/

/*--------注册 开始--------*/
type LoginRegisterReq struct {
	g.Meta `path:"/register" method:"post" tags:"机构后台/登录" sm:"注册"`
	api.CommonHeaderReq
	Phone     string `json:"phone,omitempty" v:"required-without-all:Email,Account|max-length:20|phone" dc:"手机"`
	Email     string `json:"email,omitempty" v:"required-without-all:Phone,Account|max-length:60|email" dc:"邮箱"`
	Account   string `json:"account,omitempty" v:"required-without-all:Phone,Email|max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	SmsCode   string `json:"sms_code,omitempty" v:"required-with:Phone|size:4" dc:"短信验证码"`
	EmailCode string `json:"email_code,omitempty" v:"required-with:Email|size:4" dc:"邮箱验证码"`
	Password  string `json:"password,omitempty" v:"required|size:32" dc:"密码。加密后发送，公式：md5(密码)"`
	AdminType uint8  `json:"admin_type" v:"in:10" d:"10" dc:"类型：10机构"`
}

/*--------注册 结束--------*/

/*--------密码找回 开始--------*/
type LoginPasswordRecoveryReq struct {
	g.Meta `path:"/password-recovery" method:"post" tags:"机构后台/登录" sm:"密码找回"`
	api.CommonHeaderReq
	Phone     string `json:"phone,omitempty" v:"required-without:Email|max-length:30" dc:"手机"`
	Email     string `json:"email,omitempty" v:"required-without:Phone|max-length:60" dc:"邮箱"`
	SmsCode   string `json:"sms_code,omitempty" v:"required-with:Phone|size:4" dc:"短信验证码"`
	EmailCode string `json:"email_code,omitempty" v:"required-with:Email|size:4" dc:"邮箱验证码"`
	Password  string `json:"password,omitempty" v:"required|size:32" dc:"密码。加密后发送，公式：md5(密码)"`
	AdminType uint8  `json:"admin_type" v:"in:10" d:"10" dc:"类型：10机构"`
}

/*--------密码找回 结束--------*/
