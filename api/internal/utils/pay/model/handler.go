package model

import "github.com/gogf/gf/v2/net/ghttp"

type Handler interface {
	App(payReq PayReq) (payRes PayRes, err error)
	H5(payReq PayReq) (payRes PayRes, err error)
	QRCode(payReq PayReq) (payRes PayRes, err error)
	Jsapi(payReq PayReq) (payRes PayRes, err error)
	Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error)
	NotifyRes(r *ghttp.Request, failMsg string)
}
