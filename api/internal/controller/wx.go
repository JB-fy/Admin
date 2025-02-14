package controller

import (
	"api/api"
	"api/internal/utils/wx"
	"context"
	"errors"

	"github.com/gogf/gf/v2/frame/g"
)

type Wx struct{}

func NewWx() *Wx {
	return &Wx{}
}

// 公众号回调
func (controllerThis *Wx) GzhNotify(ctx context.Context, req *api.WxGzhNotifyReq) (res *api.CommonNoDataRes, err error) {
	r := g.RequestFromCtx(ctx)
	wxGzhObj := wx.NewWxGzh(ctx)
	timestamp := r.Get(`timestamp`).String()
	nonce := r.Get(`nonce`).String()
	signature := r.Get(`signature`).String()

	if wxGzhObj.Sign(timestamp, nonce) != signature {
		g.Log().Error(ctx, `签名错误`)
		err = errors.New(`签名错误`)
		return
	}
	//接入验证：GET请求
	if r.Method == `GET` {
		r.Response.WriteJson(r.Get(`echostr`).String())
		return
	}

	//微信推送事件：POST请求
	encryptType := r.Get(`encrypt_type`).String()
	msgSignature := r.Get(`msg_signature`).String()

	encryptReqBody := wxGzhObj.GetEncryptReqBody(r)
	encrypt := encryptReqBody.Encrypt
	// encrypt := r.Get(`Encrypt`).String() //body是xml时，框架已经做了解析。故也可以直接取
	if encryptType != `aes` {
		g.Log().Error(ctx, `请设置消息加解密方式：安全模式`)
		err = errors.New(`请设置消息加解密方式：安全模式`)
		return
	}
	if wxGzhObj.MsgSign(timestamp, nonce, encrypt) != msgSignature {
		g.Log().Error(ctx, `消息签名错误`)
		err = errors.New(`消息签名错误`)
		return
	}
	msgByte, err := wxGzhObj.AesDecrypt(encrypt)
	if err != nil {
		return
	}

	notify, err := wxGzhObj.Notify(msgByte)
	if err != nil {
		return
	}
	encryptMsg := ``
	switch notify.MsgType {
	case `event`: //接收事件推送
		switch notify.Event {
		case `subscribe`: //关注
		case `unsubscribe`: //取消关注
		}
	case `text`: //接收普通消息（文本）
		/* //回复用户消息（文本）
		encryptMsg, err = wxGzhObj.EncryptMsg(notify.ToUserName, notify.FromUserName, timestamp, `text`, `测试`)
		if err != nil {
			return
		} */
	}

	wxGzhObj.NotifyRes(r, notify.FromUserName, notify.ToUserName, nonce, timestamp, encryptMsg)
	return
}
