package cmd

import (
	myGen "api/internal/cmd/my-gen"
	"api/internal/middleware"
	"api/internal/router"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  `main`,
		Brief: `通过这个命令启动`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) { //未指定执行哪个子命令时，默认运行该方法。比如：gf run main.go -a "--gf.gcfg.file=config.yaml"
			go Http.Run(ctx)

			// 等待中断信号来优雅地关闭服务
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
			<-ch
			return
		},
	}

	MyGen = gcmd.Command{
		Name:  `myGen`,
		Usage: `myGen`,
		Brief: `代码自动生成`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			myGen.Run(ctx, parser)
			return
		},
	}

	Http = gcmd.Command{
		Name:  `http`,
		Usage: `http`,
		Brief: `http服务`,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.BindMiddlewareDefault(middleware.Cross, middleware.I18n)
			if g.Cfg().MustGet(ctx, `logger.http.isRecord`).Bool() {
				s.BindMiddlewareDefault(middleware.Log)
			}
			s.BindMiddlewareDefault(middleware.HandlerResponse) // 不用规范路由方式可去掉。但如果是规范路由时则必须，且有用log中间件时，必须放在其后面，才能读取到响应数据

			router.InitRouterCommon(s)   //公共接口注册
			router.InitRouterPlatform(s) //平台后台接口注册
			router.InitRouterOrg(s)      //机构后台接口注册
			router.InitRouterApp(s)      //APP接口注册

			// router.InitRouterWebSocket(s) //WebScoket注册（如需使用，建议把部分全局中间件移到对应接口分组内）

			s.Run()
			return
		},
	}
)
