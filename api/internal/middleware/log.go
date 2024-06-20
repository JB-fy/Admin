package middleware

import (
	"api/internal/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func Log(r *ghttp.Request) {
	startTime := gtime.Now().UnixMicro()

	r.Middleware.Next()

	endTime := gtime.Now().UnixMicro()
	runTime := (float64(endTime) - float64(startTime)) / 1000
	data := map[string]any{
		`url`:        r.GetUrl(),
		`header`:     r.Header,
		`req_data`:   r.GetMap(),
		`res_data`:   r.Response.BufferString(),
		`res_status`: r.Response.Status,
		`run_time`:   runTime,
	}
	data[`client_ip`] = r.GetClientIp()
	data[`login_id`] = 0
	loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
	if !loginInfo.IsEmpty() {
		data[`login_id`] = loginInfo[`login_id`]
	}
	g.Log(`http`).Info(r.GetCtx(), data)
}
