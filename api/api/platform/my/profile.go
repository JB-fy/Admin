package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------个人信息 开始--------*/
type ProfileInfoReq struct {
	g.Meta `path:"/profile/info" method:"post" tags:"平台后台/我的" sm:"个人信息"`
}

type ProfileInfoRes struct {
	Info ProfileInfo `json:"info" dc:"详情"`
}

type ProfileInfo struct {
	AdminId   *uint       `json:"adminId,omitempty" dc:"管理员ID"`
	Account   *string     `json:"account,omitempty" dc:"账号"`
	Phone     *string     `json:"phone,omitempty" dc:"手机号"`
	Nickname  *string     `json:"nickname,omitempty" dc:"昵称"`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	UpdatedAt *gtime.Time `json:"updatedAt,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt,omitempty" dc:"创建时间"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateReq struct {
	g.Meta        `path:"/profile/update" method:"post" tags:"平台后台/我的" sm:"修改个人信息"`
	Account       *string `c:"account,omitempty" json:"account" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"账号"`
	Phone         *string `c:"phone,omitempty" json:"phone" v:"phone" dc:"手机号"`
	Nickname      *string `c:"nickname,omitempty" json:"nickname" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"昵称"`
	Avatar        *string `c:"avatar,omitempty" json:"avatar" v:"url|length:1,120" dc:"头像"`
	Password      *string `c:"password,omitempty" json:"password" v:"size:32|regex:^[\\p{L}\\p{N}_-]+$|different:CheckPassword" dc:"新密码"`
	CheckPassword *string `c:"checkPassword,omitempty" json:"checkPassword" v:"required-with:Account,Phone,Password|size:32|regex:^[\\p{L}\\p{N}_-]+$" dc:"旧密码。修改账号，手机号，密码时必填"`
}

/*--------修改个人信息 结束--------*/
