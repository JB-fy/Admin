package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CommonCreateRes struct {
	Id int64 `json:"id" dc:"ID"`
}

type CommonNoDataRes struct {
}

/*--------获取加密盐 开始--------*/
type LoginEncryptReq struct {
	g.Meta  `path:"/encryptStr" method:"post" tags:"平台/登录" sm:"获取加密盐"`
	Account string `json:"account"  v:"required|length:4,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
}

type LoginEncryptRes struct {
	EncryptStr string `json:"encryptStr" dc:"加密盐"`
}

/*--------获取加密盐 结束--------*/

/*--------登录 开始--------*/
type LoginReq struct {
	g.Meta   `path:"/" method:"post" tags:"平台/登录" sm:"登录"`
	Account  string `json:"account"  v:"required|length:4,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$"`
	Password string `json:"password"  v:"required|size:32|regex:^[\\p{L}\\p{N}]+$"`
}

type LoginRes struct {
	Token string `json:"token" dc:"登录授权token"`
}

/*--------登录 结束--------*/
