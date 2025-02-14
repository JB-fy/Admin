package model

import "github.com/gogf/gf/v2/net/ghttp"

type Handler interface {
	Upload(r *ghttp.Request) (notifyInfo NotifyInfo, err error)
	Sign() (signInfo SignInfo, err error)
	Config() (config map[string]any, err error)
	Sts() (stsInfo map[string]any, err error)
	Notify(r *ghttp.Request) (notifyInfo NotifyInfo, err error)
	// createUploadParam() (param UploadParam)
}
