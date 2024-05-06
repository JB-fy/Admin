package controller

import (
	"api/api"
	"api/internal/utils"
	"api/internal/utils/wx"
	"context"

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
		err = utils.NewErrorCode(ctx, 99999999, `签名错误`)
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
		err = utils.NewErrorCode(ctx, 99999999, `请设置消息加解密方式：安全模式`)
		return
	}
	if wxGzhObj.MsgSign(timestamp, nonce, encrypt) != msgSignature {
		g.Log().Error(ctx, `消息签名错误`)
		err = utils.NewErrorCode(ctx, 99999999, `消息签名错误`)
		return
	}
	decryptByte, err := wxGzhObj.AesDecrypt(encrypt)
	if err != nil {
		return
	}
	notifyByte, err := wxGzhObj.ParseDecrypt(decryptByte)
	if err != nil {
		return
	}

	// 微信这个死垃圾没有针对事件给唯一标识，得将数据对每个事件都解析一次才能知道是什么事件（根据自身业务需要增加事件处理）
	// 关注/取消关注事件
	notify := wxGzhObj.NotifyOfSubscribe(notifyByte)
	if notify.MsgType == `event` {
		switch notify.Event {
		case `subscribe`: //关注
			// 关注处理逻辑
			return
		case `unsubscribe`: //取消关注
			// 取消关注处理逻辑
			return
		}
	}
	return
}
