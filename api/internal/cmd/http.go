package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"

	"api/internal/initialize"
	"api/internal/router"
)

func HttpFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	initialize.InitI18n(ctx) //多语言设置
	initialize.InitCron(ctx) //定时器设置
	initialize.InitGenv(ctx) //设置当前服务器IP并记录

	/*--------启动http服务 开始--------*/
	s := g.Server()

	router.InitRouterCommon(s)   //公共接口注册
	router.InitRouterPlatform(s) //平台后台接口注册
	s.Run()
	/*--------启动http服务 结束--------*/
	return
}
