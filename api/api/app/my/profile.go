package api

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------个人信息 开始--------*/
type ProfileInfoReq struct {
	g.Meta `path:"/profile/info" method:"post" tags:"APP/我的" sm:"个人信息"`
}

type ProfileInfoRes struct {
	Info ProfileInfo `json:"info" dc:"详情"`
}

type ProfileInfo struct {
	UserId     uint   `json:"userId" dc:"用户ID"`
	Phone      string `json:"phone" dc:"手机"`
	Account    string `json:"account" dc:"账号"`
	Nickname   string `json:"nickname" dc:"昵称"`
	Avatar     string `json:"avatar" dc:"头像"`
	Gender     uint   `json:"gender" dc:"性别：0未设置 1男 2女"`
	Birthday   string `json:"birthday" dc:"生日"`
	Address    string `json:"address" dc:"详细地址"`
	IdCardName string `json:"idCardName" dc:"身份证姓名"`
	IdCardNo   string `json:"idCardNo" dc:"身份证号码"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateReq struct {
	g.Meta            `path:"/profile/update" method:"post" tags:"APP/我的" sm:"修改个人信息"`
	Phone             *string     `json:"phone,omitempty" v:"length:1,30|phone" dc:"手机"`
	Nickname          *string     `json:"nickname,omitempty" v:"length:1,30" dc:"昵称"`
	Avatar            *string     `json:"avatar,omitempty" v:"length:1,120|url" dc:"头像"`
	Gender            *uint       `json:"gender,omitempty" v:"integer|in:0,1,2" dc:"性别：0未设置 1男 2女"`
	Birthday          *gtime.Time `json:"birthday,omitempty" v:"date-format:Y-m-d" dc:"生日"`
	Address           *string     `json:"address,omitempty" v:"length:1,60" dc:"详细地址"`
	IdCardName        *string     `json:"idCardName,omitempty" v:"length:1,30|regex:^[\\p{L}\\p{M}\\p{N}_-]+$" dc:"身份证姓名"`
	IdCardNo          *string     `json:"idCardNo,omitempty" v:"length:1,30" dc:"身份证号码"`
	Password          *string     `json:"password,omitempty" v:"size:32" dc:"新密码。加密后发送，公式：md5(新密码)"`
	CheckPassword     *string     `json:"checkPassword,omitempty" v:"required-with:Password|size:32|different:Password" dc:"旧密码。加密后发送，公式：md5(新密码)。修改密码时checkPassword和smsCodeToPassword必传其一"`
	SmsCodeToPassword *string     `json:"smsCodeToPassword,omitempty" v:"required-with:Password|size:4" dc:"短信验证码。修改密码时checkPassword和smsCodeToPassword必传其一"`
}

/*--------修改个人信息 结束--------*/
