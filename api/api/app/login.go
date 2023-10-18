package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
type LoginSaltReq struct {
	g.Meta    `path:"/salt" method:"post" tags:"APP/登录" sm:"获取加密盐"`
	LoginName string `json:"loginName" v:"required|length:1,30" dc:"账号/手机"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta    `path:"/login" method:"post" tags:"APP/登录" sm:"登录"`
	LoginName string `json:"loginName" v:"required|length:1,30" dc:"账号/手机"`
	Password  string `json:"password" v:"required-without:SmsCode|size:32|regex:^[\\p{L}\\p{N}]+$" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态加密盐)+动态加密盐)"`
	SmsCode   string `json:"smsCode" v:"required-without:Password|size:4" dc:"短信验证码"`
}

/*--------登录 结束--------*/
