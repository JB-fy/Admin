package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取密码盐 开始--------*/
type LoginSaltReq struct {
	g.Meta    `path:"/salt" method:"post" tags:"机构后台/登录" sm:"获取密码盐"`
	OrgId     uint   `json:"org_id" v:"required|between:1,4294967295" dc:"机构ID"`
	LoginName string `json:"login_name" v:"required|max-length:60" dc:"手机/邮箱/账号"`
}

/*--------获取密码盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta    `path:"/login" method:"post" tags:"机构后台/登录" sm:"登录"`
	OrgId     uint   `json:"org_id" v:"required|between:1,4294967295" dc:"机构ID"`
	LoginName string `json:"login_name" v:"required|max-length:60" dc:"手机/邮箱/账号"`
	Password  string `json:"password" v:"required|size:32" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态密码盐)+动态密码盐)"`
}

/*--------登录 结束--------*/
