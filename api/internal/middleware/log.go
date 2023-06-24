package middleware

import (
	dao "api/internal/dao/log"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
)

func Log(r *ghttp.Request) {
	startTime := gtime.Now().UnixMicro()

	r.Middleware.Next()

	endTime := gtime.Now().UnixMicro()
	runTime := (float64(endTime) - float64(startTime)) / 1000
	data := map[string]interface{}{
		`url`:     r.GetUrl(),
		`header`:  r.Header,
		`reqData`: r.GetMap(),
		`resData`: r.Response.BufferString(),
		`runTime`: runTime,
	}
	dao.Http.ParseDbCtx(r.GetCtx()).Data(data).Insert()
}
