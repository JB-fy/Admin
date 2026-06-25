package my

import (
	"api/api"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

/*--------个人信息 开始--------*/
type ProfileInfoReq struct {
	g.Meta `path:"/profile/info" method:"post" tags:"平台后台/我的" sm:"个人信息"`
	api.CommonPlatformHeaderReq
}

type ProfileInfoRes struct {
	Info ProfileInfo `json:"info" dc:"详情"`
}

type ProfileInfo struct {
	AdminId   *uint       `json:"admin_id,omitempty" dc:"管理员ID"`
	SceneId   *string     `json:"scene_id,omitempty" dc:"场景ID"`
	RelId     *uint       `json:"rel_id,omitempty" dc:"关联ID。根据scene_id对应不同表"`
	AdminType *uint       `json:"admin_type,omitempty" dc:"类型：0平台 10机构"`
	Account   *string     `json:"account,omitempty" dc:"账号"`
	Phone     *string     `json:"phone,omitempty" dc:"手机"`
	Email     *string     `json:"email,omitempty" dc:"邮箱"`
	Nickname  *string     `json:"nickname,omitempty" dc:"昵称"`
	Avatar    *string     `json:"avatar,omitempty" dc:"头像"`
	IsSuper   *uint       `json:"is_super,omitempty" dc:"超管：0否 1是"`
	IsStop    *uint       `json:"is_stop,omitempty" dc:"停用：0否 1是"`
	UpdatedAt *gtime.Time `json:"updated_at,omitempty" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"created_at,omitempty" dc:"创建时间"`
}

/*--------个人信息 结束--------*/

/*--------修改个人信息 开始--------*/
type ProfileUpdateData struct {
	Nickname        *string `json:"nickname,omitempty" v:"max-length:30" dc:"昵称"`
	Avatar          *string `json:"avatar,omitempty" v:"max-length:200|url" dc:"头像"`
	Phone           *string `json:"phone,omitempty" v:"max-length:20|phone" dc:"手机"`
	Email           *string `json:"email,omitempty" v:"max-length:60|email" dc:"邮箱"`
	Account         *string `json:"account,omitempty" v:"max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	Password        *string `json:"password,omitempty" v:"size:32" dc:"新密码。加密后发送，公式：md5(新密码)"`
	PasswordToCheck *string `json:"password_to_check,omitempty" v:"required-with:Phone,Email,Account,Password|size:32|different:Password" dc:"旧密码。加密后发送，公式：md5(新密码)。修改手机，邮箱，账号，密码用，password_to_check,sms_code_to_password,email_code_to_password传一个即可"`
	// SmsCodeToPassword      *string `json:"sms_code_to_password,omitempty" v:"size:4" dc:"短信验证码。修改密码用，password_to_check,sms_code_to_password,email_code_to_password传一个即可"`
	SmsCodeToBindPhone *string `json:"sms_code_to_bind_phone,omitempty" v:"required-with:Phone|size:4" dc:"短信验证码。绑定手机用"`
	// SmsCodeToUnbingPhone   *string `json:"sms_code_to_unbing_phone,omitempty" v:"size:4" dc:"短信验证码。解绑手机用"`
	// EmailCodeToPassword    *string `json:"email_code_to_password,omitempty" v:"size:4" dc:"邮箱验证码。修改密码用，password_to_check,sms_code_to_password,email_code_to_password传一个即可"`
	EmailCodeToBindEmail *string `json:"email_code_to_bind_email,omitempty" v:"required-with:Email|size:4" dc:"邮箱验证码。绑定邮箱用"`
	// EmailCodeToUnbingEmail *string `json:"email_code_to_unbing_email,omitempty" v:"size:4" dc:"邮箱验证码。解绑邮箱用"`
}

type ProfileUpdateReq struct {
	g.Meta `path:"/profile/update" method:"post" tags:"平台后台/我的" sm:"修改个人信息"`
	api.CommonPlatformHeaderReq
	ProfileUpdateData
}

/*--------修改个人信息 结束--------*/
