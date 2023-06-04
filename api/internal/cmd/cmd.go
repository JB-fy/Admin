package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"api/internal/controller"
	controllerAuth "api/internal/controller/auth"
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
					//controller.NewTest().Test, //这样不会根据方法名自动设置路由
					controller.NewTest(),
				)
			})

			/**--------平台后台接口 开始--------**/
			s.Group("/platformAdmin", func(group *ghttp.RouterGroup) {
				group.ALL("/test", controller.NewTest().Test)
				//不做日志记录
				group.Group("", func(group *ghttp.RouterGroup) {
					group.Middleware(middleware.Scene)
					//需验证登录身份
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.SceneLoginOfPlatformAdmin)
						group.ALLMap(g.Map{
							"/log/request": controller.NewTest().Test,
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
							"/encryptStr": controller.NewTest().Test,
							"/":           controller.NewTest().Test,
						})
					})

					//需验证登录身份
					group.Group("", func(group *ghttp.RouterGroup) {
						group.Middleware(middleware.SceneLoginOfPlatformAdmin)

						group.Group("/upload", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/sign": controller.NewTest().Test,
							})
						})

						group.Group("/login", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/info":     controller.NewTest().Test,
								"/update":   controller.NewTest().Test,
								"/menuTree": controller.NewTest().Test,
							})
						})

						group.Group("/auth/action", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   controller.NewTest().Test,
								"/info":   controller.NewTest().Test,
								"/create": controller.NewTest().Test,
								"/update": controller.NewTest().Test,
								"/del":    controller.NewTest().Test,
							})
						})

						group.Group("/auth/menu", func(group *ghttp.RouterGroup) {
							controllerAuthMenu := controllerAuth.NewMenu()
							group.ALLMap(g.Map{
								"/list":   controllerAuthMenu.List,
								"/info":   controllerAuthMenu.Info,
								"/create": controllerAuthMenu.Create,
								"/update": controllerAuthMenu.Update,
								"/del":    controllerAuthMenu.Delete,
								"/tree":   controllerAuthMenu.List,
							})
						})
						group.Group("/auth/role", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   controller.NewTest().Test,
								"/info":   controller.NewTest().Test,
								"/create": controller.NewTest().Test,
								"/update": controller.NewTest().Test,
								"/del":    controller.NewTest().Test,
							})
						})

						group.Group("/auth/scene", func(group *ghttp.RouterGroup) {
							controllerAuthScene := controllerAuth.NewScene()
							group.ALLMap(g.Map{
								"/list":   controllerAuthScene.List,
								"/info":   controllerAuthScene.Info,
								"/create": controllerAuthScene.Create,
								"/update": controllerAuthScene.Update,
								"/del":    controllerAuthScene.Delete,
							})
						})

						group.Group("/auth/admin", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list":   controller.NewTest().Test,
								"/info":   controller.NewTest().Test,
								"/create": controller.NewTest().Test,
								"/update": controller.NewTest().Test,
								"/del":    controller.NewTest().Test,
							})
						})

						group.Group("/platform/config", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/get":  controller.NewTest().Test,
								"/save": controller.NewTest().Test,
							})
						})

						group.Group("/platform/server", func(group *ghttp.RouterGroup) {
							group.ALLMap(g.Map{
								"/list": controller.NewTest().Test,
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
