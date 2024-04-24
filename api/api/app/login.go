package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
type LoginSaltReq struct {
	g.Meta    `path:"/salt" method:"post" tags:"APP/登录" sm:"获取加密盐"`
	LoginName string `json:"login_name" v:"required|max-length:30" dc:"账号/手机"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta    `path:"/login" method:"post" tags:"APP/登录" sm:"登录"`
	LoginName string `json:"login_name" v:"required|max-length:30" dc:"账号/手机"`
	Password  string `json:"password" v:"required-without:SmsCode|size:32" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态加密盐)+动态加密盐)"`
	SmsCode   string `json:"smsCode" v:"required-without:Password|size:4" dc:"短信验证码"`
}

/*--------登录 结束--------*/

/*--------注册 开始--------*/
type LoginRegisterReq struct {
	g.Meta  `path:"/register" method:"post" tags:"APP/登录" sm:"注册"`
	Phone   string `json:"phone,omitempty" v:"required-without:Account|max-length:30|phone" dc:"手机"`
	Account string `json:"account,omitempty" v:"required-without:Phone|max-length:30|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	// Password string `json:"password" v:"required-with:Account|lsize:32" dc:"密码。加密后发送，公式：md5(密码)"`
	Password string `json:"password" v:"required|lsize:32" dc:"密码。加密后发送，公式：md5(密码)"`
	SmsCode  string `json:"smsCode" v:"required-with:Phone|size:4" dc:"短信验证码"`
}

/*--------注册 结束--------*/

/*--------密码找回 开始--------*/
type LoginPasswordRecoveryReq struct {
	g.Meta   `path:"/passwordRecovery" method:"post" tags:"APP/登录" sm:"密码找回"`
	Phone    string `json:"phone,omitempty" v:"required|max-length:30|phone" dc:"手机"`
	SmsCode  string `json:"smsCode" v:"required|size:4" dc:"短信验证码"`
	Password string `json:"password" v:"required|lsize:32" dc:"密码。加密后发送，公式：md5(密码)"`
}

/*--------密码找回 结束--------*/
