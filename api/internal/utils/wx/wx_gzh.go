package wx

import (
	"bytes"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"io"
	"sort"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type WxGzh struct {
	Host           string `json:"host"`
	AppId          string `json:"appId"`
	Secret         string `json:"secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encodingAESKey"`
	AESKey         []byte
}

func NewWxGzh(config map[string]any) *WxGzh {
	wxGzhObj := &WxGzh{}
	gconv.Struct(config, wxGzhObj)
	if wxGzhObj.AppId == `` || wxGzhObj.Secret == `` || wxGzhObj.Token == `` || wxGzhObj.EncodingAESKey == `` {
		panic(`缺少插件配置：微信-公众号`)
	}
	wxGzhObj.AESKey, _ = base64.StdEncoding.DecodeString(wxGzhObj.EncodingAESKey + `=`)
	return wxGzhObj
}

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
	TimeStamp    string
	Nonce        CDATAText
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

// 获取签名
func (wxGzhThis *WxGzh) Sign(timestamp, nonce string) (sign string) {
	arr := []string{wxGzhThis.Token, timestamp, nonce}
	sort.Strings(arr)
	sha := sha1.New()
	sha.Write([]byte(gstr.Join(arr, ``)))
	sign = hex.EncodeToString(sha.Sum(nil))
	return
}

// 获取Msg签名
func (wxGzhThis *WxGzh) MsgSign(timestamp, nonce, encrypt string) (sign string) {
	arr := []string{wxGzhThis.Token, timestamp, nonce, encrypt}
	sort.Strings(arr)
	sha := sha1.New()
	sha.Write([]byte(gstr.Join(arr, ``)))
	sign = hex.EncodeToString(sha.Sum(nil))
	return
}

func (wxGzhThis *WxGzh) GetEncryptReqBody(r *ghttp.Request) (encryptReqBody *EncryptReqBody) {
	/* body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	} */
	body := r.GetBody()
	encryptReqBody = &EncryptReqBody{}
	xml.Unmarshal(body, encryptReqBody)
	return
}

// aes解密
func (wxGzhThis *WxGzh) AesDecrypt(encrypt string) (msgByte []byte, err error) {
	cipherData, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return
	}
	if len(cipherData)%len(wxGzhThis.AESKey) != 0 {
		err = errors.New(`crypto/cipher: ciphertext size is not multiple of aes key length`)
		return
	}

	block, err := aes.NewCipher(wxGzhThis.AESKey)
	if err != nil {
		return
	}

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	msgByte = make([]byte, len(cipherData))
	blockMode.CryptBlocks(msgByte, cipherData)
	return
}

// 加密消息体
func (wxGzhThis *WxGzh) EncryptMsg(fromUserName, toUserName, timestamp, msgType string, msg any) (encrypt string, err error) {
	type MsgOfCommon struct {
		XMLName      xml.Name `xml:"xml"`
		ToUserName   CDATAText
		FromUserName CDATAText
		CreateTime   string
		MsgType      CDATAText
	}
	msgOfCommon := MsgOfCommon{
		FromUserName: wxGzhThis.value2CDATA(fromUserName),
		ToUserName:   wxGzhThis.value2CDATA(toUserName),
		CreateTime:   timestamp,
		MsgType:      wxGzhThis.value2CDATA(msgType),
	}

	msgBody := []byte{}
	switch msgType {
	case `text`: //回复文本消息
		type MsgOfText struct {
			MsgOfCommon
			Content CDATAText
		}
		msgBody, err = xml.MarshalIndent(&MsgOfText{
			MsgOfCommon: msgOfCommon,
			Content:     wxGzhThis.value2CDATA(gconv.String(msg)),
		}, ``, `    `)
		if err != nil {
			return
		}
	}

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, int32(len(msgBody)))
	if err != nil {
		return
	}

	plainData := bytes.Join([][]byte{[]byte(grand.S(16)), buf.Bytes(), msgBody, []byte(wxGzhThis.AppId)}, nil)
	encrypt, err = wxGzhThis.aesEncrypt(plainData)
	return
}

// 回调
func (wxGzhThis *WxGzh) Notify(msgByte []byte) (notify *Notify, err error) {
	buf := bytes.NewBuffer(msgByte[16:20])
	var msgLength int32
	err = binary.Read(buf, binary.BigEndian, &msgLength)
	if err != nil {
		return
	}

	appIdPos := 20 + int(msgLength)
	if string(msgByte[appIdPos:appIdPos+len(wxGzhThis.AppId)]) != wxGzhThis.AppId {
		err = errors.New(`AppId无效`)
		return
	}

	notify = &Notify{}
	xml.Unmarshal(msgByte[20:20+msgLength], notify)
	return
}

// 回调响应处理
func (wxGzhThis *WxGzh) NotifyRes(r *ghttp.Request, fromUserName, toUserName, nonce, timestamp, encrypt string) (err error) {
	if encrypt == `` {
		r.Response.Write(`success`)
		return
	}
	encryptResBody := &EncryptResBody{}
	encryptResBody.Encrypt = wxGzhThis.value2CDATA(encrypt)
	encryptResBody.MsgSignature = wxGzhThis.value2CDATA(wxGzhThis.MsgSign(timestamp, nonce, encrypt))
	encryptResBody.TimeStamp = timestamp
	encryptResBody.Nonce = wxGzhThis.value2CDATA(nonce)
	notifyResBody, errTmp := xml.MarshalIndent(encryptResBody, ``, `    `)
	if errTmp != nil {
		err = errTmp
		return
	}
	r.Response.WriteXml(string(notifyResBody))
	return
}

type WxGzhAccessToken struct {
	AccessToken string `json:"access_token"` //授权Token
	ExpiresIn   int    `json:"expires_in"`   //有效时间，单位：秒
}

// 获取access_token（需要在公众号内设置IP白名单）
func (wxGzhThis *WxGzh) AccessToken(ctx context.Context) (accessToken WxGzhAccessToken, err error) {
	res, err := g.Client().Get(ctx, wxGzhThis.Host+`/cgi-bin/token`, g.Map{
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

	resData.Var().Struct(&accessToken)
	return
}

type WxGzhUserInfo struct {
	Unionid        string `json:"unionid"`         //用户统一标识（全局唯一）。公众号绑定到微信开放平台账号后，才会出现该字段（注意：还需要用户关注公众号。微信文档未说明这点）
	Openid         string `json:"openid"`          //用户唯一标识（相对于公众号、开放平台下的应用唯一）
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

// 获取用户基本信息
func (wxGzhThis *WxGzh) UserInfo(ctx context.Context, accessToken, openid string) (userInfo WxGzhUserInfo, err error) {
	res, err := g.Client().Get(ctx, wxGzhThis.Host+`/cgi-bin/user/info`, g.Map{
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

type WxGzhUserGet struct {
	Total uint `json:"total"` //关注该公众账号的总用户数
	Count uint `json:"count"` //拉取的OPENID个数，最大值为10000
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"` //列表数据，OPENID的列表
	NextOpenid string `json:"next_openid"` //拉取列表的最后一个用户的OPENID
}

// 获取用户列表
func (wxGzhThis *WxGzh) UserGet(ctx context.Context, accessToken, nextOpenid string) (userGet WxGzhUserGet, err error) {
	res, err := g.Client().Get(ctx, wxGzhThis.Host+`/cgi-bin/user/get`, g.Map{
		`access_token`: accessToken,
		`next_openid`:  nextOpenid,
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

	resData.Var().Struct(&userGet)
	return
}

func (wxGzhThis *WxGzh) value2CDATA(v string) CDATAText {
	return CDATAText{`<![CDATA[` + v + `]]>`}
}

func (wxGzhThis *WxGzh) pKCS7Pad(message []byte, blocksize int) (padded []byte, err error) {
	if blocksize < 2 {
		err = errors.New(`block size is too small(minimum is 2 bytes)`)
		return
	}
	if blocksize > 255 {
		err = errors.New(`block size is too long(maxmum is 255 bytes)`)
		return
	}

	// calculate padding length
	padlen := blocksize - len(message)%blocksize
	if padlen == 0 {
		padlen = blocksize
	}

	// define PKCS7 padding blockbody
	padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

	// apply padding
	padded = append(message, padding...)
	return
}

// aes加密
func (wxGzhThis *WxGzh) aesEncrypt(msgByte []byte) (encrypt string, err error) {
	k := len(wxGzhThis.AESKey)
	if len(msgByte)%k != 0 {
		msgByte, err = wxGzhThis.pKCS7Pad(msgByte, k)
		if err != nil {
			return
		}
	}

	block, err := aes.NewCipher(wxGzhThis.AESKey)
	if err != nil {
		return
	}

	iv := make([]byte, aes.BlockSize)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return
	}

	cipherData := make([]byte, len(msgByte))
	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cipherData, msgByte)
	encrypt = base64.StdEncoding.EncodeToString(cipherData)
	return
}
