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
	AdminId  uint   `json:"adminId" dc:"管理员ID"`
	Phone    string `json:"phone" dc:"手机"`
	Account  string `json:"account" dc:"账号"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateReq struct {
	g.Meta          `path:"/profile/update" method:"post" tags:"平台后台/我的" sm:"修改个人信息"`
	Phone           *string `json:"phone,omitempty" v:"phone" dc:"手机"`
	Account         *string `json:"account,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Nickname        *string `json:"nickname,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"昵称"`
	Avatar          *string `json:"avatar,omitempty" v:"url|length:1,200" dc:"头像"`
	Password        *string `json:"password,omitempty" v:"size:32" dc:"新密码。加密后发送，公式：md5(新密码)"`
	PasswordToCheck *string `json:"passwordToCheck,omitempty" v:"required-with:Account,Phone,Password|size:32|different:Password" dc:"旧密码。加密后发送，公式：md5(新密码)。修改账号，手机，密码时必填"`
}

/*--------修改个人信息 结束--------*/
