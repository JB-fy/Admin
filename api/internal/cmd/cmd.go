package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	controller "api/internal/controller/auth"
	"api/internal/controller/hello"
	"api/internal/middleware"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.I18n().SetPath(g.Cfg().MustGet(ctx, "i18n.path").String())         //设置资源目录
			g.I18n().SetLanguage(g.Cfg().MustGet(ctx, "i18n.language").String()) //设置默认为中文（原默认为英文en）

			s := g.Server()
			s.BindMiddlewareDefault(middleware.HandlerResponse, middleware.Cross, middleware.I18n)
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					//hello.New().Test, //这样不会根据方法名自动设置路由
					hello.New(),
				)
			})

			/**--------平台后台接口 开始--------**/
			s.Group("/platformAdmin", func(group *ghttp.RouterGroup) {
				group.ALL("/test", hello.New().Test)
				//不做日志记录
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Scene)
					//需验证登录身份
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.SceneLoginOfPlatformAdmin)
						group.ALLMap(g.Map{
							"/log/request": hello.New().Test,
						})
					})
				})

				//做日志记录
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Log)
					group.Middleware(middleware.Scene)
					//无需验证登录身份
					group.Group("/login", func(group *ghttp.RouterGroup) {
						group.ALLMap(g.Map{
							"/encryptStr": hello.New().Test,
							"/":           hello.New().Test,
						})
					})

					//需验证登录身份
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.SceneLoginOfPlatformAdmin)

						group.Group("/upload", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/sign": hello.New().Test,
							})
						})

						group.Group("/login", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/info":     hello.New().Test,
								"/update":   hello.New().Test,
								"/menuTree": hello.New().Test,
							})
						})

						group.Group("/auth/action", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   hello.New().Test,
								"/info":   hello.New().Test,
								"/create": hello.New().Test,
								"/update": hello.New().Test,
								"/del":    hello.New().Test,
							})
						})

						group.Group("/auth/menu", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   hello.New().Test,
								"/info":   hello.New().Test,
								"/create": hello.New().Test,
								"/update": hello.New().Test,
								"/del":    hello.New().Test,
								"/tree":   hello.New().Test,
							})
						})
						group.Group("/auth/role", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   hello.New().Test,
								"/info":   hello.New().Test,
								"/create": hello.New().Test,
								"/update": hello.New().Test,
								"/del":    hello.New().Test,
							})
						})

						group.Group("/auth/scene", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   controller.NewScene().List,
								"/info":   hello.New().Test,
								"/create": hello.New().Test,
								"/update": hello.New().Test,
								"/del":    hello.New().Test,
							})
						})

						group.Group("/auth/admin", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   hello.New().Test,
								"/info":   hello.New().Test,
								"/create": hello.New().Test,
								"/update": hello.New().Test,
								"/del":    hello.New().Test,
							})
						})

						group.Group("/platform/config", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/get":  hello.New().Test,
								"/save": hello.New().Test,
							})
						})

						group.Group("/platform/server", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list": hello.New().Test,
							})
						})

					})
				})
			})
			/**--------平台后台接口 结束--------**/
			s.Run()
			return nil
		},
	}
)
