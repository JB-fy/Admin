package one_click

import (
	"context"
	"errors"
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type OneClickOfWx struct {
	Ctx    context.Context
	Host   string `json:"host"`
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
}

func NewOneClickOfWx(ctx context.Context, config map[string]any) *OneClickOfWx {
	oneClickObj := &OneClickOfWx{Ctx: ctx}
	gconv.Struct(config, oneClickObj)
	if oneClickObj.Host == `` || oneClickObj.AppId == `` || oneClickObj.Secret == `` {
		panic(`缺少插件配置：一键登录-微信`)
	}
	return oneClickObj
}

type AccessTokenOfWx struct {
	Unionid        string `json:"unionid"`         //用户统一标识（全局唯一），只有当scope为"snsapi_userinfo"时返回
	Openid         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	AccessToken    string `json:"access_token"`    //网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn      int    `json:"expires_in"`      //access_token 接口调用凭证超时时间，单位（秒）
	RefreshToken   string `json:"refresh_token"`   //用户刷新access_token
	Scope          string `json:"scope"`           //用户授权的作用域，使用逗号（,）分隔
	IsSnapshotuser int    `json:"is_snapshotuser"` //快照页模式虚拟账号：0否 1是。只有当用户是快照页模式虚拟账号时返回
}

type UserInfoOfWx struct {
	Unionid   string `json:"unionid"`    //用户统一标识（全局唯一）
	Openid    string `json:"openid"`     //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	Nickname  string `json:"nickname"`   //昵称
	Gender    int    `json:"sex"`        //性别：0未知 1男 2女
	Avatar    string `json:"headimgurl"` //头像。最后一个数值代表正方形头像大小，有0、46、64、96、132数值可选，0代表640*640正方形头像
	Country   string `json:"country"`    //国家，如中国为CN
	Province  string `json:"province"`   //用户个人资料填写的省份
	City      string `json:"city"`       //用户个人资料填写的城市
	Privilege string `json:"privilege"`  //用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
}

type RefreshTokenOfWx struct {
	Openid       string `json:"openid"`        //网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	AccessToken  string `json:"access_token"`  //授权Token
	ExpiresIn    int    `json:"expires_in"`    //access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` //用户刷新access_token
	Scope        string `json:"scope"`         //用户授权的作用域，使用逗号（,）分隔
}

// 获取用户同意授权地址（也可以让前端自己处理。坏处：前端需要存appId，更新appId需同步修改前端）
// scope		应用授权作用域。 snsapi_base：不弹出授权页面，直接跳转，只能获取用户openid；snsapi_userinfo：弹出授权页面，可通过openid拿到昵称、性别、所在地
// state		重定向后会带上state参数。开发者可以填写a-zA-Z0-9的参数值，最多128字节
// forcePopup	强制此次授权需要用户弹窗确认。默认为false
func (oneClickThis *OneClickOfWx) CodeUrl(redirectUri string, scope string, state string, forcePopup bool) (codeUrl string, err error) {
	codeUrl = `https://open.weixin.qq.com/connect/oauth2/authorize?appid=` + oneClickThis.AppId + `&redirect_uri=` + url.QueryEscape(redirectUri) + `&response_type=code&scope=` + scope
	if state != `` {
		codeUrl += `&state=` + state
	}
	if forcePopup {
		codeUrl += `&forcePopup=1`
	}
	codeUrl += `#wechat_redirect`
	return
}

// 通过code换取网页授权access_token（code由前端自己获取）
func (oneClickThis *OneClickOfWx) AccessToken(code string) (accessToken AccessTokenOfWx, err error) {
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/oauth2/access_token`, g.Map{
		`appid`:      oneClickThis.AppId,
		`secret`:     oneClickThis.Secret,
		`code`:       code,
		`grant_type`: `authorization_code`,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}

	resData.Var().Struct(&accessToken)
	return
}

// 拉取用户信息(需scope为 snsapi_userinfo)
func (oneClickThis *OneClickOfWx) UserInfo(openid, accessToken string) (userInfo UserInfoOfWx, err error) {
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/userinfo`, g.Map{
		`access_token`: accessToken,
		`openid`:       openid,
		`lang`:         `zh_CN`,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}

	resData.Var().Struct(&userInfo)
	return
}

// 刷新access_token（需要时用）
func (oneClickThis *OneClickOfWx) RefreshToken(refreshTokenStr string) (refreshToken RefreshTokenOfWx, err error) {
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/oauth2/refresh_token`, g.Map{
		`appid`:         oneClickThis.AppId,
		`grant_type`:    `refresh_token`,
		`refresh_token`: refreshTokenStr,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}

	resData.Var().Struct(&refreshToken)
	return
}

// 授权验证（需要时用）
func (oneClickThis *OneClickOfWx) Auth(openid, accessToken string) (err error) {
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/auth`, g.Map{
		`access_token`: accessToken,
		`openid`:       openid,
		`lang`:         `zh_CN`,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}
	return
}
