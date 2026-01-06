package gzh

import (
	"api/internal/utils"
	"bytes"
	"context"
	"crypto/aes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"sort"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Wx struct {
	Host           string `json:"host"`
	AppId          string `json:"app_id"`
	Secret         string `json:"secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encoding_aes_key"`
	AESKey         []byte
	client         *utils.HttpClient
}

func NewWx(ctx context.Context, config map[string]any) *Wx {
	obj := &Wx{}
	gconv.Struct(config, obj)
	if obj.AppId == `` || obj.Secret == `` || obj.Token == `` || obj.EncodingAESKey == `` {
		panic(`缺少插件配置：微信-公众号`)
	}
	obj.AESKey, _ = base64.StdEncoding.DecodeString(obj.EncodingAESKey + `=`)
	obj.client = utils.NewHttpClient(ctx)
	return obj
}

// 获取签名
func (wxGzhThis *Wx) Sign(timestamp, nonce string) (sign string) {
	arr := []string{wxGzhThis.Token, timestamp, nonce}
	sort.Strings(arr)
	sha := sha1.New()
	sha.Write([]byte(gstr.Join(arr, ``)))
	sign = hex.EncodeToString(sha.Sum(nil))
	return
}

// 获取Msg签名
func (wxGzhThis *Wx) MsgSign(timestamp, nonce, encrypt string) (sign string) {
	arr := []string{wxGzhThis.Token, timestamp, nonce, encrypt}
	sort.Strings(arr)
	sha := sha1.New()
	sha.Write([]byte(gstr.Join(arr, ``)))
	sign = hex.EncodeToString(sha.Sum(nil))
	return
}

func (wxGzhThis *Wx) GetEncryptReqBody(r *ghttp.Request) (encryptReqBody *EncryptReqBody) {
	/* body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	} */
	body := r.GetBody()
	encryptReqBody = &EncryptReqBody{}
	xml.Unmarshal(body, encryptReqBody)
	return
}

// aes加密
func (wxGzhThis *Wx) AesEncrypt(msgByte []byte) (encrypt string, err error) {
	encrypt, err = utils.AesEncryptOfBase64(msgByte, wxGzhThis.AESKey, utils.AesPadTypeOfPKCS7, len(wxGzhThis.AESKey), utils.AesCipherTypeOfCBC, wxGzhThis.AESKey[:aes.BlockSize]...)
	return
}

// aes解密
func (wxGzhThis *Wx) AesDecrypt(encrypt string) (msgByte []byte, err error) {
	msgByte, err = utils.AesDecryptOfBase64(encrypt, wxGzhThis.AESKey, utils.AesPadTypeOfPKCS7, len(wxGzhThis.AESKey), utils.AesCipherTypeOfCBC, wxGzhThis.AESKey[:aes.BlockSize]...)
	return
}

// 加密消息体
func (wxGzhThis *Wx) EncryptMsg(fromUserName, toUserName, timestamp, msgType string, msg any) (encrypt string, err error) {
	msgOfCommon := MsgOfCommon{
		FromUserName: wxGzhThis.value2CDATA(fromUserName),
		ToUserName:   wxGzhThis.value2CDATA(toUserName),
		CreateTime:   timestamp,
		MsgType:      wxGzhThis.value2CDATA(msgType),
	}

	msgBody := []byte{}
	switch msgType {
	case `text`: //回复文本消息
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
	encrypt, err = wxGzhThis.AesEncrypt(plainData)
	return
}

// 回调
func (wxGzhThis *Wx) Notify(msgByte []byte) (notify *Notify, err error) {
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
func (wxGzhThis *Wx) NotifyRes(r *ghttp.Request, fromUserName, toUserName, nonce, timestamp, encrypt string) (err error) {
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

// 获取access_token（需要在公众号内设置IP白名单）
func (wxGzhThis *Wx) AccessToken(ctx context.Context) (accessToken AccessToken, err error) {
	resData, err := wxGzhThis.get(ctx, `/cgi-bin/token`, g.Map{
		`grant_type`: `client_credential`,
		`appid`:      wxGzhThis.AppId,
		`secret`:     wxGzhThis.Secret,
	})
	if err != nil {
		return
	}
	resData.Var().Struct(&accessToken)
	return
}

// 获取用户基本信息
func (wxGzhThis *Wx) UserInfo(ctx context.Context, accessToken, openid string) (userInfo UserInfo, err error) {
	resData, err := wxGzhThis.get(ctx, `/cgi-bin/user/info`, g.Map{
		`access_token`: accessToken,
		`openid`:       openid,
		`lang`:         `zh_CN`,
	})
	if err != nil {
		return
	}
	resData.Var().Struct(&userInfo)
	return
}

// 获取用户列表
func (wxGzhThis *Wx) UserGet(ctx context.Context, accessToken, nextOpenid string) (userGet UserGet, err error) {
	resData, err := wxGzhThis.get(ctx, `/cgi-bin/user/get`, g.Map{
		`access_token`: accessToken,
		`next_openid`:  nextOpenid,
	})
	if err != nil {
		return
	}
	resData.Var().Struct(&userGet)
	return
}

func (wxGzhThis *Wx) value2CDATA(v string) CDATAText {
	return CDATAText{`<![CDATA[` + v + `]]>`}
}

func (wxGzhThis *Wx) get(ctx context.Context, apiPath string, param g.Map) (resData *gjson.Json, err error) {
	res, err := wxGzhThis.client.Get(ctx, wxGzhThis.Host+apiPath, param)
	if err != nil {
		return
	}
	defer res.Close()
	resStr := res.ReadAllString()
	resData = gjson.New(resStr)
	if resData.Contains(`errcode`) && resData.Get(`errcode`).Int() != 0 {
		err = errors.New(resData.Get(`errmsg`).String())
		return
	}
	return
}
