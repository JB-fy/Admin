package wx

import (
	"context"
	"errors"
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/util/gconv"
)

type OneClick struct {
	Host   string `json:"host"`
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
	client *gclient.Client
}

func NewOneClick(ctx context.Context, config map[string]any) *OneClick {
	obj := &OneClick{}
	gconv.Struct(config, obj)
	if obj.Host == `` || obj.AppId == `` || obj.Secret == `` {
		panic(`缺少插件配置：一键登录-微信`)
	}
	obj.client = g.Client()
	return obj
}

// 获取用户同意授权地址（也可以让前端自己处理。坏处：前端需要存appId，更新appId需同步修改前端）
// scope		应用授权作用域。 snsapi_base：不弹出授权页面，直接跳转，只能获取用户openid；snsapi_userinfo：弹出授权页面，可通过openid拿到昵称、性别、所在地
// state		重定向后会带上state参数。开发者可以填写a-zA-Z0-9的参数值，最多128字节
// forcePopup	强制此次授权需要用户弹窗确认。默认为false
func (oneClickThis *OneClick) CodeUrl(redirectUri string, scope string, state string, forcePopup bool) (codeUrl string, err error) {
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
func (oneClickThis *OneClick) AccessToken(ctx context.Context, code string) (accessToken AccessToken, err error) {
	res, err := oneClickThis.client.Get(ctx, oneClickThis.Host+`/sns/oauth2/access_token`, g.Map{
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
func (oneClickThis *OneClick) UserInfo(ctx context.Context, openid, accessToken string) (userInfo UserInfo, err error) {
	res, err := oneClickThis.client.Get(ctx, oneClickThis.Host+`/sns/userinfo`, g.Map{
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
func (oneClickThis *OneClick) RefreshToken(ctx context.Context, refreshTokenStr string) (refreshToken RefreshToken, err error) {
	res, err := oneClickThis.client.Get(ctx, oneClickThis.Host+`/sns/oauth2/refresh_token`, g.Map{
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
func (oneClickThis *OneClick) Auth(ctx context.Context, openid, accessToken string) (err error) {
	res, err := oneClickThis.client.Get(ctx, oneClickThis.Host+`/sns/auth`, g.Map{
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
