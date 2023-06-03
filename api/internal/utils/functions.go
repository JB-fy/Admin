package utils

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

func HttpFailJson(r *ghttp.Request, code int, msg string, data map[string]interface{}) {
	resData := map[string]interface{}{
		"code": code,
		"data": data,
	}
	if msg == "" {
		resData["msg"] = g.I18n().Tf(r.GetCtx(), gconv.String(code))
	} else {
		resData["msg"] = msg
	}
	r.Response.WriteJson(resData)
}

func HttpSuccessJson(r *ghttp.Request, data map[string]interface{}, code int, msg string) {
	resData := map[string]interface{}{
		"code": code,
		"data": data,
	}
	if msg == "" {
		resData["msg"] = g.I18n().Tf(r.GetCtx(), gconv.String(code))
	} else {
		resData["msg"] = msg
	}
	r.Response.WriteJson(resData)
}
