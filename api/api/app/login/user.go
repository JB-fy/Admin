package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
type UserSaltReq struct {
	g.Meta  `path:"/salt" method:"post" tags:"APP/登录" sm:"获取加密盐"`
	Account string `json:"account" v:"required|length:4,30|passport" dc:"账号"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type UserLoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"APP/登录" sm:"登录"`
	Phone    string `json:"account" v:"required-without:Account|length:4,30|phone" dc:"手机"`
	Account  string `json:"account" v:"required-without:Phone|length:4,30|passport" dc:"账号"`
	Password string `json:"password" v:"required|size:32|regex:^[\\p{L}\\p{N}]+$" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态加密盐)+动态加密盐)"`
}

/*--------登录 结束--------*/
