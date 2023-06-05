package router

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerAuth "api/internal/controller/auth"
	"api/internal/middleware"
)

func InitRouterPlatformAdmin(s *ghttp.Server) {
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
					controllerThis := controllerAuth.NewAction()
					group.ALLMap(g.Map{
						"/list":   controllerThis.List,
						"/info":   controllerThis.Info,
						"/create": controllerThis.Create,
						"/update": controllerThis.Update,
						"/del":    controllerThis.Delete,
					})
				})

				group.Group("/auth/menu", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewMenu()
					group.ALLMap(g.Map{
						"/list":   controllerThis.List,
						"/info":   controllerThis.Info,
						"/create": controllerThis.Create,
						"/update": controllerThis.Update,
						"/del":    controllerThis.Delete,
						"/tree":   controllerThis.List,
					})
				})
				group.Group("/auth/role", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewRole()
					group.ALLMap(g.Map{
						"/list":   controllerThis.List,
						"/info":   controllerThis.Info,
						"/create": controllerThis.Create,
						"/update": controllerThis.Update,
						"/del":    controllerThis.Delete,
					})
				})

				group.Group("/auth/scene", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewScene()
					group.ALLMap(g.Map{
						"/list":   controllerThis.List,
						"/info":   controllerThis.Info,
						"/create": controllerThis.Create,
						"/update": controllerThis.Update,
						"/del":    controllerThis.Delete,
					})
				})

				group.Group("/platform/admin", func(group *ghttp.RouterGroup) {
					controllerThis := controller.NewTest()
					group.ALLMap(g.Map{
						"/list":   controllerThis.Test,
						"/info":   controllerThis.Test,
						"/create": controllerThis.Test,
						"/update": controllerThis.Test,
						"/del":    controllerThis.Test,
					})
				})

				group.Group("/platform/config", func(group *ghttp.RouterGroup) {
					controllerThis := controller.NewTest()
					group.ALLMap(g.Map{
						"/get":  controllerThis.Test,
						"/save": controllerThis.Test,
					})
				})

				group.Group("/platform/server", func(group *ghttp.RouterGroup) {
					controllerThis := controller.NewTest()
					group.ALLMap(g.Map{
						"/list": controllerThis.Test,
					})
				})

			})
		})
	})
}
