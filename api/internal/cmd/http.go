package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"api/internal/middleware"
	"api/internal/router"
)

func HttpFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()

	s.BindMiddlewareDefault(middleware.Cross, middleware.I18n, middleware.Log)
	s.BindMiddlewareDefault(middleware.HandlerResponse) // 不用规范路由方式可去掉。但如果是规范路由时则必须，且有用log中间件时，必须放在其后面，才能读取到响应数据

	router.InitRouterCommon(s)   //公共接口注册
	router.InitRouterPlatform(s) //平台后台接口注册

	// router.InitRouterWebSocket(s) //WebScoket注册（如需使用，建议把部分全局中间件移到对应接口分组内）

	s.Run()
	return
}
