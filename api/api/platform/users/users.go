package users

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 共用详情。list,info,tree等接口返回时用，但返回默认字段有差异。可根据需要在controller对应的defaultField中补充所需字段
type UsersInfo struct {
	Id             *uint       `json:"id,omitempty" dc:"ID"`
	Label          *string     `json:"label,omitempty" dc:"标签。常用于前端组件"`
	UserId         *uint       `json:"user_id,omitempty" dc:"用户ID"`
	Nickname       *string     `json:"nickname,omitempty" dc:"昵称"`
	Avatar         *string     `json:"avatar,omitempty" dc:"头像"`
	Gender         *uint       `json:"gender,omitempty" dc:"性别：0未设置 1男 2女"`
	Birthday       *string     `json:"birthday,omitempty" dc:"生日"`
	Address        *string     `json:"address,omitempty" dc:"地址"`
	Phone          *string     `json:"phone,omitempty" dc:"手机"`
	Email          *string     `json:"email,omitempty" dc:"邮箱"`
	Account        *string     `json:"account,omitempty" dc:"账号"`
	WxOpenid       *string     `json:"wx_openid,omitempty" dc:"微信openid"`
	WxUnionid      *string     `json:"wx_unionid,omitempty" dc:"微信unionid"`
	IdCardNo       *string     `json:"id_card_no,omitempty" dc:"身份证号码"`
	IdCardName     *string     `json:"id_card_name,omitempty" dc:"身份证姓名"`
	IdCardGender   *uint       `json:"id_card_gender,omitempty" dc:"身份证性别：0未设置 1男 2女"`
	IdCardBirthday *string     `json:"id_card_birthday,omitempty" dc:"身份证生日"`
	IdCardAddress  *string     `json:"id_card_address,omitempty" dc:"身份证地址"`
	IsStop         *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt      *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt      *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

type UsersListFilter struct {
	Id             *uint       `json:"id,omitempty" v:"between:1,4294967295" dc:"ID"`
	IdArr          []uint      `json:"id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"ID数组"`
	ExcId          *uint       `json:"exc_id,omitempty" v:"between:1,4294967295" dc:"排除ID"`
	ExcIdArr       []uint      `json:"exc_id_arr,omitempty" v:"distinct|foreach|between:1,4294967295" dc:"排除ID数组"`
	Label          string      `json:"label,omitempty" v:"max-length:30|regex:^[\\p{L}\\p{N}_-]+$" dc:"搜索关键词。常用于前端组件"`
	TimeRangeStart *gtime.Time `json:"time_range_start,omitempty" v:"date-format:Y-m-d H:i:s" dc:"开始时间：YYYY-mm-dd HH:ii:ss"`
	TimeRangeEnd   *gtime.Time `json:"time_range_end,omitempty" v:"date-format:Y-m-d H:i:s|after-equal:TimeRangeStart" dc:"结束时间：YYYY-mm-dd HH:ii:ss"`
	UserId         *uint       `json:"user_id,omitempty" v:"between:1,4294967295" dc:"用户ID"`
	Nickname       string      `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Gender         *uint       `json:"gender,omitempty" v:"in:0,1,2" dc:"性别：0未设置 1男 2女"`
	Birthday       *gtime.Time `json:"birthday,omitempty" v:"date-format:Y-m-d" dc:"生日"`
	Phone          string      `json:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	Email          string      `json:"email,omitempty" v:"max-length:60|email" dc:"邮箱"`
	Account        string      `json:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	WxOpenid       string      `json:"wx_openid,omitempty" v:"max-length:128" dc:"微信openid"`
	WxUnionid      string      `json:"wx_unionid,omitempty" v:"max-length:64" dc:"微信unionid"`
	IdCardNo       string      `json:"id_card_no,omitempty" v:"max-length:30" dc:"身份证号码"`
	IdCardName     string      `json:"id_card_name,omitempty" v:"max-length:30" dc:"身份证姓名"`
	IdCardGender   *uint       `json:"id_card_gender,omitempty" v:"in:0,1,2" dc:"身份证性别：0未设置 1男 2女"`
	IdCardBirthday *gtime.Time `json:"id_card_birthday,omitempty" v:"date-format:Y-m-d" dc:"身份证生日"`
	IsStop         *uint       `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

type UsersUpdateDeleteFilter struct {
	Id    uint   `json:"id,omitempty" v:"required-without:IdArr|between:1,4294967295" dc:"ID"`
	IdArr []uint `json:"id_arr,omitempty" v:"required-without:Id|distinct|foreach|between:1,4294967295" dc:"ID数组"`
}

/*--------列表 开始--------*/
type UsersListReq struct {
	g.Meta `path:"/users/list" method:"post" tags:"平台后台/用户管理/用户" sm:"列表"`
	api.CommonPlatformHeaderReq
	api.CommonListReq
	Filter UsersListFilter `json:"filter" dc:"过滤条件"`
}

type UsersListRes struct {
	Count int         `json:"count" dc:"总数"`
	List  []UsersInfo `json:"list" dc:"列表"`
}

/*--------列表 结束--------*/

/*--------详情 开始--------*/
type UsersInfoReq struct {
	g.Meta `path:"/users/info" method:"post" tags:"平台后台/用户管理/用户" sm:"详情"`
	api.CommonPlatformHeaderReq
	api.CommonInfoReq
	Id uint `json:"id" v:"required|between:1,4294967295" dc:"ID"`
}

type UsersInfoRes struct {
	Info UsersInfo `json:"info" dc:"详情"`
}

/*--------详情 结束--------*/

/*--------修改 开始--------*/
type UsersUpdateData struct {
	// Nickname       *string `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	// Avatar         *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	// Gender         *uint   `json:"gender,omitempty" v:"in:0,1,2" dc:"性别：0未设置 1男 2女"`
	// Birthday       *string `json:"birthday,omitempty" v:"date-format:Y-m-d" dc:"生日"`
	// Address        *string `json:"address,omitempty" v:"max-length:120" dc:"地址"`
	// Phone          *string `json:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	// Email          *string `json:"email,omitempty" v:"max-length:60|email" dc:"邮箱"`
	// Account        *string `json:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	// WxOpenid       *string `json:"wx_openid,omitempty" v:"max-length:128" dc:"微信openid"`
	// WxUnionid      *string `json:"wx_unionid,omitempty" v:"max-length:64" dc:"微信unionid"`
	// Password       *string `json:"password,omitempty" v:"size:32" dc:"密码。md5保存"`
	// IdCardNo       *string `json:"id_card_no,omitempty" v:"max-length:30" dc:"身份证号码"`
	// IdCardName     *string `json:"id_card_name,omitempty" v:"max-length:30" dc:"身份证姓名"`
	// IdCardGender   *uint   `json:"id_card_gender,omitempty" v:"in:0,1,2" dc:"身份证性别：0未设置 1男 2女"`
	// IdCardBirthday *string `json:"id_card_birthday,omitempty" v:"date-format:Y-m-d" dc:"身份证生日"`
	// IdCardAddress  *string `json:"id_card_address,omitempty" v:"max-length:120" dc:"身份证地址"`
	IsStop *uint `json:"is_stop,omitempty" v:"in:0,1" dc:"停用：0否 1是"`
}

type UsersUpdateReq struct {
	g.Meta `path:"/users/update" method:"post" tags:"平台后台/用户管理/用户" sm:"修改"`
	api.CommonPlatformHeaderReq
	UsersUpdateDeleteFilter
	UsersUpdateData
}

/*--------修改 结束--------*/
