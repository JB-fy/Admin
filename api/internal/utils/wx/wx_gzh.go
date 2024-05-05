package wx

import (
	daoPlatform "api/internal/dao/platform"
	"context"
	"errors"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type WxGzh struct {
	Ctx            context.Context
	Host           string `json:"wxGzhHost"`
	AppId          string `json:"wxGzhAppId"`
	Secret         string `json:"wxGzhSecret"`
	Token          string `json:"wxGzhToken"`
	EncodingAESKey string `json:"wxGzhEncodingAESKey"`
}

func NewWxGzh(ctx context.Context, configOpt ...map[string]interface{}) *WxGzh {
	var config map[string]interface{}
	if len(configOpt) > 0 && len(configOpt[0]) > 0 {
		config = configOpt[0]
	} else {
		configTmp, _ := daoPlatform.Config.Get(ctx, []string{`wxGzhHost`, `wxGzhAppId`, `wxGzhSecret`, `wxGzhToken`, `wxGzhEncodingAESKey`})
		config = configTmp.Map()
	}

	obj := WxGzh{Ctx: ctx}
	gconv.Struct(config, &obj)
	return &obj
}

type WxGzhAccessToken struct {
	AccessToken string `json:"access_token"` //授权Token
	ExpiresIn   int    `json:"expires_in"`   //有效时间，单位：秒
}

type WxGzhUserInfo struct {
	UnionId        string `json:"unionid"`         //用户统一标识（全局唯一）。公众号绑定到微信开放平台账号后，才会出现该字段
	OpenId         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	Subscribe      int    `json:"subscribe"`       //关注公众号：0否 1是
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
func (wxGzhThis *WxGzh) AccessToken() (accessToken WxGzhAccessToken, err error) {
	res, err := g.Client().Get(wxGzhThis.Ctx, wxGzhThis.Host+`/cgi-bin/token`, g.Map{
		`grant_type`: `client_credential`,
		`appid`:      wxGzhThis.AppId,
		`secret`:     wxGzhThis.Secret,
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

// 获取用户基本信息
func (wxGzhThis *WxGzh) UserInfo(openId, accessToken string) (userInfo WxGzhUserInfo, err error) {
	res, err := g.Client().Get(wxGzhThis.Ctx, wxGzhThis.Host+`/cgi-bin/user/info`, g.Map{
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
