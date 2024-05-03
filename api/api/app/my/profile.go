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
	IdCardName string `json:"id_card_name" dc:"身份证姓名"`
	IdCardNo   string `json:"id_card_no" dc:"身份证号码"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateReq struct {
	g.Meta               `path:"/profile/update" method:"post" tags:"APP/我的" sm:"修改个人信息"`
	Phone                *string     `json:"phone,omitempty" v:"max-length:30|phone" dc:"手机"`
	Account              *string     `json:"account,omitempty" v:"max-length:30|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	Nickname             *string     `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Avatar               *string     `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	Gender               *uint       `json:"gender,omitempty" v:"in:0,1,2" dc:"性别：0未设置 1男 2女"`
	Birthday             *gtime.Time `json:"birthday,omitempty" v:"date-format:Y-m-d" dc:"生日"`
	Address              *string     `json:"address,omitempty" v:"max-length:60" dc:"详细地址"`
	IdCardName           *string     `json:"id_card_name,omitempty" v:"required-with:IdCardNo|max-length:30" dc:"身份证姓名"`
	IdCardNo             *string     `json:"id_card_no,omitempty" v:"required-with:IdCardName|max-length:30" dc:"身份证号码"`
	Password             *string     `json:"password,omitempty" v:"size:32" dc:"新密码。加密后发送，公式：md5(新密码)"`
	PasswordToCheck      *string     `json:"password_to_check,omitempty" v:"required-with:Account|size:32|different:Password" dc:"旧密码。加密后发送，公式：md5(新密码)。修改账号|密码用，password_to_check和sms_code_to_password传一个即可"`
	SmsCodeToPassword    *string     `json:"sms_code_to_password,omitempty" v:"size:4" dc:"短信验证码。修改密码用，password_to_check和sms_code_to_password传一个即可"`
	SmsCodeToBindPhone   *string     `json:"sms_code_to_bind_phone,omitempty" v:"required-with:Phone|size:4" dc:"短信验证码。绑定手机用"`
	SmsCodeToUnbingPhone *string     `json:"sms_code_to_unbing_phone,omitempty" v:"size:4" dc:"短信验证码。解绑手机用"`
}

/*--------修改个人信息 结束--------*/

/*--------关注信息（微信公众号） 开始--------*/
type ProfileFollowInfoOfWxReq struct {
	g.Meta `path:"/profile/follow-info-of-wx" method:"post" tags:"APP/我的" sm:"关注信息（微信公众号）"`
}

type ProfileFollowInfoOfWxRes struct {
	Info ProfileFollowInfoOfWx `json:"info" dc:"详情"`
}

type ProfileFollowInfoOfWx struct {
	IsFollow    int    `json:"is_follow" dc:"关注公众号：0否 1是"`
	FollowTime  int    `json:"follow_time" dc:"关注时间戳"`
	FollowScene string `json:"follow_scene" dc:"关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK	图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_REPRINT 他人转载，ADD_SCENE_LIVESTREAM 视频号直播，ADD_SCENE_CHANNELS 视频号，ADD_SCENE_WXA 小程序关注，ADD_SCENE_OTHERS 其他"`
}

/*--------关注信息（微信公众号） 结束--------*/
