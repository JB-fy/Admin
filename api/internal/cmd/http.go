package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"api/internal/initialize"
	"api/internal/middleware"
	"api/internal/router"
)

func HttpFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	initialize.InitI18n(ctx) //多语言设置
	initialize.InitCron(ctx) //定时器设置
	initialize.InitGenv(ctx) //环境变量设置。如：记录当前服务器IP

	/*--------启动http服务 开始--------*/
	s := g.Server()
	s.BindMiddlewareDefault(middleware.Cross, middleware.I18n, middleware.Log)
	s.BindMiddlewareDefault(middleware.HandlerResponse) // 不用规范路由方式可去掉。但如果是规范路由时则必须，且有用log中间件时，必须放在其后面，才能读取到响应数据
	router.InitRouterCommon(s)                          //公共接口注册
	router.InitRouterPlatform(s)                        //平台后台接口注册
	s.Run()
	/*--------启动http服务 结束--------*/
	return
}
