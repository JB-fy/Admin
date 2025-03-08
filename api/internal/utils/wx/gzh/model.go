package gzh

import "encoding/xml"

type CDATAText struct {
	Text string `xml:",innerxml"`
}

type EncryptReqBody struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string
	Encrypt    string
}

type EncryptResBody struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      CDATAText
	MsgSignature CDATAText
	Nonce        CDATAText
	TimeStamp    string
}

type Notify struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   string
	MsgType      string
	//事件消息
	Event    string
	EventKey string
	Ticket   string
	// 地理位置消息
	LocationX string
	LocationY string
	Scale     string
	Label     string
	// 文本消息
	Content   string
	MsgId     int64
	MsgDataId int64
	Idx       int
	// 图片消息
	PicUrl  string
	MediaId string
	// 语音消息
	Format string
	// 视频消息
	ThumbMediaId string
	// 链接消息
	Title       string
	Description string
	Url         string
}

type AccessToken struct {
	AccessToken string `json:"access_token"` //授权Token
	ExpiresIn   int64  `json:"expires_in"`   //有效时间，单位：秒
}

type UserInfo struct {
	Unionid        string `json:"unionid"`         //用户统一标识（全局唯一）。公众号绑定到微信开放平台账号后，才会出现该字段（注意：还需要用户关注公众号。微信文档未说明这点）
	Openid         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
	Language       string `json:"language"`        //语言，简体中文为zh_CN
	Remark         string `json:"remark"`          //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupId        string `json:"groupid"`         //用户所在的分组ID（兼容旧的用户分组接口）
	TagidList      string `json:"tagid_list"`      //用户被打上的标签ID列表
	QrScene        string `json:"qr_scene"`        //二维码扫码场景（开发者自定义）
	QrSceneStr     string `json:"qr_scene_str"`    //二维码扫码场景描述（开发者自定义）
	SubscribeScene string `json:"subscribe_scene"` //关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK	图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_REPRINT 他人转载，ADD_SCENE_LIVESTREAM 视频号直播，ADD_SCENE_CHANNELS 视频号，ADD_SCENE_WXA 小程序关注，ADD_SCENE_OTHERS 其他
	SubscribeTime  int64  `json:"subscribe_time"`  //关注时间戳
	Subscribe      uint8  `json:"subscribe"`       //关注公众号：0否 1是
}

type UserGet struct {
	Total uint `json:"total"` //关注该公众账号的总用户数
	Count uint `json:"count"` //拉取的OPENID个数，最大值为10000
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"` //列表数据，OPENID的列表
	NextOpenid string `json:"next_openid"` //拉取列表的最后一个用户的OPENID
}

type MsgOfCommon struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	MsgType      CDATAText
	CreateTime   string
}

type MsgOfText struct {
	MsgOfCommon
	Content CDATAText
}
