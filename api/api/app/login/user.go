package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
type UserSaltReq struct {
	g.Meta `path:"/salt" method:"post" tags:"APP/登录" sm:"获取加密盐"`
	// Phone   string `json:"phone" v:"required-without:Account|length:1,30|phone" dc:"手机"`
	// Account string `json:"account" v:"required-without:Phone|length:1,30|passport" dc:"账号"`
	LoginName string `json:"loginName" v:"required|length:1,30" dc:"账号/手机"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type UserLoginReq struct {
	g.Meta `path:"/login" method:"post" tags:"APP/登录" sm:"登录"`
	// Phone    string `json:"phone" v:"required-without:Account|length:1,30|phone" dc:"手机"`
	// Account  string `json:"account" v:"required-without:Phone|length:1,30|passport" dc:"账号"`
	// Password  string `json:"password" v:"required-with:Account|required-without:Code|size:32|regex:^[\\p{L}\\p{N}]+$" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态加密盐)+动态加密盐)"`
	// Code      string `json:"code" v:"required-without-all:Account,Password|size:4" dc:"验证码"`
	LoginName string `json:"loginName" v:"required|length:1,30" dc:"账号/手机"`
	Password  string `json:"password" v:"required-without:Code|size:32|regex:^[\\p{L}\\p{N}]+$" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态加密盐)+动态加密盐)"`
	Code      string `json:"code" v:"required-without:Password|size:4" dc:"验证码"`
}

/*--------登录 结束--------*/
