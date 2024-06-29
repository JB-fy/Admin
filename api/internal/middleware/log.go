package middleware

import (
	"api/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func Log(r *ghttp.Request) {
	if !g.Cfg().MustGet(r.GetCtx(), `logger.http.isRecord`).Bool() {
		r.Middleware.Next()
		return
	}

	startTime := gtime.Now().UnixMicro()
	r.Middleware.Next()
	endTime := gtime.Now().UnixMicro()

	data := map[string]any{
		`url`:        r.GetUrl(),
		`header`:     r.Header,
		`req_data`:   r.GetMap(),
		`res_status`: r.Response.Status,
		`run_time`:   float64(endTime-startTime) / 1000,
		`client_ip`:  r.GetClientIp(),
	}
	if maxResBufferLength := g.Cfg().MustGet(r.GetCtx(), `logger.http.maxResBufferLength`).Int(); maxResBufferLength > 0 && r.Response.BufferLength() <= maxResBufferLength {
		data[`res_data`] = r.Response.BufferString()
	}
	loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
	if !loginInfo.IsEmpty() {
		data[`login_id`] = loginInfo[`login_id`]
	}

	g.Log(`http`).Info(r.GetCtx(), data)
}
