package one_click

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"errors"
	"net/url"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type OneClickOfWx struct {
	Ctx    context.Context
	Host   string `json:"oneClickOfWxHost"`   //https://api.weixin.qq.com 或 https://api2.weixin.qq.com（备用）
	AppId  string `json:"oneClickOfWxAppId"`  //wxabe672da5799762e
	Secret string `json:"oneClickOfWxSecret"` //
}

func NewOneClickOfWx(ctx context.Context, configOpt ...map[string]interface{}) *OneClickOfWx {
	var config map[string]interface{}
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`oneClickOfWxHost`, `oneClickOfWxAppId`, `oneClickOfWxSecret`})
		config = configTmp.Map()
		/* config = g.Map{
			`oneClickOfWxHost`:   `https://api.weixin.qq.com`,
			`oneClickOfWxAppId`:  `wxabe672da5799762e`,
			`oneClickOfWxSecret`: `11111`,
		} */
	}

	obj := OneClickOfWx{Ctx: ctx}
	gconv.Struct(config, &obj)
	return &obj
}

type AccessToken struct {
	UnionId        string `json:"unionid"`         //用户统一标识（全局唯一），只有当scope为"snsapi_userinfo"时返回
	OpenId         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	AccessToken    string `json:"access_token"`    //网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn      int    `json:"expires_in"`      //access_token 接口调用凭证超时时间，单位（秒）
	RefreshToken   string `json:"refresh_token"`   //用户刷新access_token
	Scope          string `json:"scope"`           //用户授权的作用域，使用逗号（,）分隔
	IsSnapshotuser int    `json:"is_snapshotuser"` //是否为快照页模式虚拟账号，只有当用户是快照页模式虚拟账号时返回，值为1
}

type UserInfo struct {
	UnionId   string `json:"unionid"`    //用户统一标识（全局唯一）
	OpenId    string `json:"openid"`     //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	Nickname  string `json:"nickname"`   //昵称
	Gender    int    `json:"sex"`        //性别：0未知 1男 2女
	Avatar    string `json:"headimgurl"` //头像。最后一个数值代表正方形头像大小，有0、46、64、96、132数值可选，0代表640*640正方形头像
	Country   string `json:"country"`    //国家，如中国为CN
	Province  string `json:"province"`   //用户个人资料填写的省份
	City      string `json:"city"`       //用户个人资料填写的城市
	Privilege string `json:"privilege"`  //用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
}

type RefreshToken struct {
	OpenId      string
	AccessToken string //授权Token
	ExpireTime  int    //超时时间。单位：秒
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
func (oneClickThis *OneClickOfWx) AccessToken(code string) (accessToken AccessToken, err error) {
	// res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/oauth2/access_token?appid=`+oneClickThis.AppId+`&secret=`+oneClickThis.Secret+`&code=`+code+`&grant_type=authorization_code`)
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

	gconv.Struct(resData.Map(), &accessToken)
	return
}

// 拉取用户信息(需scope为 snsapi_userinfo)
func (oneClickThis *OneClickOfWx) UserInfo(openId, accessToken string) (userInfo UserInfo, err error) {
	// res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/userinfo?access_token=`+accessToken+`&openid=`+openId+`&lang=zh_CN`)
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/userinfo`, g.Map{
		`access_token`: accessToken,
		`openid`:       openId,
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

	gconv.Struct(resData.Map(), &userInfo)
	return
}

// 刷新access_token（需要时用）
func (oneClickThis *OneClickOfWx) RefreshToken(refreshToken string) (accessToken AccessToken, err error) {
	// res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/oauth2/refresh_token?appid=`+oneClickThis.AppId+`&grant_type=refresh_token&refresh_token=REFRESH_TOKEN`)
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/oauth2/refresh_token`, g.Map{
		`appid`:         oneClickThis.AppId,
		`grant_type`:    `refresh_token`,
		`refresh_token`: refreshToken,
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

	gconv.Struct(resData.Map(), &accessToken)
	return
}

// 授权验证（需要时用）
func (oneClickThis *OneClickOfWx) Auth(openId, accessToken string) (err error) {
	// res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/auth?access_token=`+accessToken+`&openid=`+openId)
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/sns/auth`, g.Map{
		`access_token`: accessToken,
		`openid`:       openId,
		`lang`:         `zh_CN`,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if !resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}
	return
}

// 下面是一些额外接口
type AccessTokenOfCgi struct {
	AccessToken string `json:"access_token"` //授权Token
	ExpiresIn   int    `json:"expires_in"`   //有效时间，单位：秒
}

type UserInfoOfCgi struct {
	UnionId        string `json:"unionid"`         //用户统一标识（全局唯一）。公众号绑定到微信开放平台账号后，才会出现该字段
	OpenId         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	Subscribe      int    `json:"subscribe"`       //是否关注公众号：0否 1是
	SubscribeTime  int    `json:"subscribe_time"`  //关注时间戳
	SubscribeScene string `json:"subscribe_scene"` //关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK	图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_REPRINT 他人转载，ADD_SCENE_LIVESTREAM 视频号直播，ADD_SCENE_CHANNELS 视频号，ADD_SCENE_WXA 小程序关注，ADD_SCENE_OTHERS 其他
	Language       string `json:"language"`        //语言，简体中文为zh_CN
	Remark         string `json:"remark"`          //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId        string `json:"groupid"`         //用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      string `json:"tagid_list"`      //用户被打上的标签ID列表
	QrScene        string `json:"qr_scene"`        //二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"`    //二维码扫码场景描述（开发者自定义）
}

// 获取access_token（注意：与上面通过code换取授权access_token不一样）
func (oneClickThis *OneClickOfWx) AccessTokenOfCgi() (accessToken AccessTokenOfCgi, err error) {
	// res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/cgi-bin/token?grant_type=client_credential&appid=`+oneClickThis.AppId+`&secret=`+oneClickThis.Secret)
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/cgi-bin/token`, g.Map{
		`grant_type`: `client_credential`,
		`appid`:      oneClickThis.AppId,
		`secret`:     oneClickThis.Secret,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if !resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}

	gconv.Struct(resData.Map(), &accessToken)
	return
}

// 获取用户基本信息
func (oneClickThis *OneClickOfWx) UserInfoOfCgi(openId, accessToken string) (userInfo UserInfoOfCgi, err error) {
	// res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/cgi-bin/user/info?access_token=`+accessToken+`&openid=`+openId+`&lang=zh_CN`)
	res, err := g.Client().Get(oneClickThis.Ctx, oneClickThis.Host+`/cgi-bin/user/info`, g.Map{
		`access_token`: accessToken,
		`openid`:       openId,
		`lang`:         `zh_CN`,
	})
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData := gjson.New(resStr)
	if !resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}

	gconv.Struct(resData.Map(), &userInfo)
	return
}
