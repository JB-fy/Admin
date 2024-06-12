package api

import (
	"github.com/gogf/gf/v2/frame/g"
)

/*--------获取密码盐 开始--------*/
type LoginSaltReq struct {
	g.Meta    `path:"/salt" method:"post" tags:"APP/登录" sm:"获取密码盐"`
	LoginName string `json:"login_name" v:"required|max-length:20" dc:"手机/邮箱/账号"`
}

/*--------获取密码盐 结束--------*/

/*--------登录 开始--------*/
type LoginLoginReq struct {
	g.Meta    `path:"/login" method:"post" tags:"APP/登录" sm:"登录"`
	LoginName string `json:"login_name" v:"required|max-length:20" dc:"手机/邮箱/账号"`
	Password  string `json:"password" v:"required-without-all:SmsCode,EmailCode|size:32" dc:"密码。加密后发送，公式：md5(md5(md5(密码)+静态密码盐)+动态密码盐)"`
	SmsCode   string `json:"sms_code" v:"required-without-all:EmailCode,Password|size:4" dc:"短信验证码"`
	EmailCode string `json:"email_code" v:"required-without-all:SmsCode,Password|size:4" dc:"邮箱验证码"`
}

/*--------登录 结束--------*/

/*--------注册 开始--------*/
type LoginRegisterReq struct {
	g.Meta    `path:"/register" method:"post" tags:"APP/登录" sm:"注册"`
	Phone     string `json:"phone,omitempty" v:"required-without-all:Email,Account|max-length:20|phone" dc:"手机"`
	Email     string `json:"email,omitempty" v:"required-without-all:Phone,Account|max-length:60|email" dc:"邮箱"`
	Account   string `json:"account,omitempty" v:"required-without-all:Phone,Email|max-length:20|regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$" dc:"账号"`
	SmsCode   string `json:"sms_code" v:"required-with:Phone|size:4" dc:"短信验证码"`
	EmailCode string `json:"email_code" v:"required-with:Email|size:4" dc:"邮箱验证码"`
	Password  string `json:"password" v:"required|size:32" dc:"密码。加密后发送，公式：md5(密码)"`
	// Password string `json:"password" v:"required-with:Account|size:32" dc:"密码。加密后发送，公式：md5(密码)"`
}

/*--------注册 结束--------*/

/*--------密码找回 开始--------*/
type LoginPasswordRecoveryReq struct {
	g.Meta   `path:"/password-recovery" method:"post" tags:"APP/登录" sm:"密码找回"`
	Phone    string `json:"phone,omitempty" v:"required|max-length:20|phone" dc:"手机"`
	SmsCode  string `json:"sms_code" v:"required|size:4" dc:"短信验证码"`
	Password string `json:"password" v:"required|size:32" dc:"密码。加密后发送，公式：md5(密码)"`
}

/*--------密码找回 结束--------*/

/*--------一键登录前置信息（如一些配置信息） 开始--------*/
type LoginOneClickPreInfoReq struct {
	g.Meta          `path:"/one-click-pre-info" method:"post" tags:"APP/登录" sm:"一键登录前置信息（如一些配置信息）"`
	OneClickType    string `json:"one_click_type" v:"required|in:oneClickOfWx,oneClickOfYidun" default:"oneClickOfWx" dc:"一键登录类型：oneClickOfWx微信 oneClickOfYidun易盾"`
	RedirectUriOfWx string `json:"redirect_uri_of_wx" v:"required-if:OneClickType,oneClickOfWx" dc:"重定向地址（微信用）"`
	ScopeOfWx       string `json:"scope_of_wx" v:"in:snsapi_base,snsapi_userinfo,snsapi_login" default:"snsapi_base" dc:"微信授权作用域（微信用）：snsapi_base用于公众号网页授权，静默授权；snsapi_userinfo用于公众号网页授权，弹出授权页面；snsapi_login用于开放平台网站应用"`
	StateOfWx       string `json:"state_of_wx" v:"max-length:128|regex:^[a-zA-Z0-9]*$" dc:"重定向后会带上state参数（微信用）。"`
	ForcePopupOfWx  bool   `json:"force_popup_of_wx" v:"in:0,1" dc:"强制此次授权需要用户弹窗确认（微信用）"`
}

type LoginOneClickPreInfoRes struct {
	CodeUrlOfWx string `json:"code_url_of_wx" dc:"微信授权地址"`
}

/*--------一键登录前置信息（如一些配置信息） 结束--------*/

/*--------一键登录 开始--------*/
type LoginOneClickReq struct {
	g.Meta             `path:"/one-click" method:"post" tags:"APP/登录" sm:"一键登录"`
	OneClickType       string `json:"one_click_type" v:"required|in:oneClickOfWx,oneClickOfYidun" default:"oneClickOfWx" dc:"一键登录类型：oneClickOfWx微信 oneClickOfYidun易盾"`
	CodeOfWx           string `json:"code_of_wx" v:"required-if:OneClickType,oneClickOfWx" dc:"微信Code（微信用）"`
	TokenOfYidun       string `json:"token_of_yidun"  v:"required-if:OneClickType,oneClickOfYidun" dc:"易盾Token（易盾用）"`
	AccessTokenOfYidun string `json:"access_token_of_yidun"  v:"required-if:OneClickType,oneClickOfYidun" dc:"易盾运营商授权码（易盾用）"`
}

/*--------一键登录 结束--------*/
