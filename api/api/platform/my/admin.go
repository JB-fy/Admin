package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------个人信息 开始--------*/
type AdminInfoReq struct {
	g.Meta `path:"/info" method:"post" tags:"平台后台/我的" sm:"个人信息"`
}

type AdminInfoRes struct {
	Info AdminInfo `json:"info" dc:"详情"`
}

type AdminInfo struct {
	AdminId   uint        `json:"adminId" dc:"管理员ID"`
	Account   string      `json:"account" dc:"账号"`
	Phone     string      `json:"phone" dc:"手机号"`
	Nickname  string      `json:"nickname" dc:"昵称"`
	Avatar    string      `json:"avatar" dc:"头像"`
	IsStop    uint        `json:"isStop" dc:"是否停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type AdminUpdateReq struct {
	g.Meta        `path:"/update" method:"post" tags:"平台后台/我的" sm:"修改个人信息"`
	Account       *string `c:"account,omitempty" json:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Phone         *string `c:"phone,omitempty" json:"phone" v:"phone" dc:"手机号"`
	Nickname      *string `c:"nickname,omitempty" json:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"昵称"`
	Avatar        *string `c:"avatar,omitempty" json:"avatar" v:"url|length:1,120" dc:"头像"`
	Password      *string `c:"password,omitempty" json:"password" v:"size:32|regex:^[\\p{L}\\p{N}_-]+$|different:CheckPassword" dc:"新密码"`
	CheckPassword *string `c:"checkPassword,omitempty" json:"checkPassword" v:"required-with:account,phone,password|size:32|regex:^[\\p{L}\\p{N}_-]+$" dc:"旧密码。修改账号，手机号，密码时必填"`
}

/*--------修改个人信息 结束--------*/
