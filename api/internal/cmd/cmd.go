package cmd

import (
	myGen "api/internal/cmd/my-gen"
	"api/internal/middleware"
	"api/internal/router"
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
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

			// 域名绑定：只建议在不想使用nginx且想要多域名绑定同一个端口时才使用
			//	1、单域名绑定时，因域名访问还必须带上端口号，且必须做DNS解析到该服务器，所以用处不大
			//	2、多域名绑定时，下方s.Xx等操作全得修改，且这样做还不如使用nginx做代理 或 直接使用两个端口 更方便
			// d1 := s.Domain(`admin.xx.com`)
			// d2 := s.Domain(`api.xx.com`)

			// 开启静态文件服务时设置
			serverRoot := g.Cfg().MustGet(ctx, `server.serverRoot`).String()
			if serverRoot != `` {
				// 上传文件目录设置
				s.BindHookHandler(`/upload/*`, ghttp.HookBeforeServe, func(r *ghttp.Request) {
					if r.IsFileRequest() {
						// r.Response.CORSDefault()
						r.Response.Header().Set(`Content-Disposition`, `attachment`) // 浏览器打开文件地址时，变为下载而不是显示
					}
				})
				// 前端文件处理，无法做到nginx一样的效果：try_files $uri @backend;
				s.BindHookHandler(`/admin/*`, ghttp.HookBeforeServe, func(r *ghttp.Request) {
					pathArr := strings.Split(r.URL.Path, `/`)
					path := `/` + pathArr[1] + `/` + pathArr[2]
					if len(pathArr) > 3 && gfile.IsFile(serverRoot+path+`/index.html`) {
						r.Response.RedirectTo(path)
					}
				})
			}

			s.BindMiddlewareDefault(middleware.Cross, middleware.I18n)
			if g.Cfg().MustGet(ctx, `logger.http.isRecord`).Bool() {
				s.BindMiddlewareDefault(middleware.Log)
			}
			s.BindMiddlewareDefault(middleware.HandlerResponse) // 不用规范路由方式可去掉。但如果是规范路由时则必须，且有用log中间件时，必须放在其后面，才能读取到响应数据

			router.InitRouterCommon(ctx, s)   //公共接口注册
			router.InitRouterPlatform(ctx, s) //平台后台接口注册
			router.InitRouterOrg(ctx, s)      //机构后台接口注册
			router.InitRouterApp(ctx, s)      //APP接口注册

			// router.InitRouterWebSocket(ctx, s) //WebScoket注册（如需使用，建议把部分全局中间件移到对应接口分组内）

			s.Run()
			return
		},
	}
)
