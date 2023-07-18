package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取加密盐 开始--------*/
type AdminEncryptStrReq struct {
	g.Meta  `path:"/encryptStr" method:"post" tags:"平台后台/登录" sm:"获取加密盐"`
	Account string `json:"account"  v:"required|length:4,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type AdminLoginReq struct {
	g.Meta   `path:"" method:"post" tags:"平台后台/登录" sm:"登录"`
	Account  string `json:"account"  v:"required|length:4,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Password string `json:"password"  v:"required|size:32|regex:^[\\p{L}\\p{N}]+$" dc:"密码。加密后发送，公式：md5(md5(用户输入密码)+加密盐)"`
}

/*--------登录 结束--------*/
