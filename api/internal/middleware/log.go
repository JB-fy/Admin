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
	data := map[string]interface{}{
		`url`:       r.GetUrl(),
		`header`:    r.Header,
		`reqData`:   r.GetMap(),
		`resData`:   r.Response.BufferString(),
		`resStatus`: r.Response.Status,
		`runTime`:   runTime,
	}
	data[`clientIp`] = r.GetClientIp()
	data[`loginId`] = 0
	loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
	if !loginInfo.IsEmpty() {
		data[`loginId`] = loginInfo[`loginId`]
	}
	g.Log(`loggerHttp`).Info(r.GetCtx(), data)
}
