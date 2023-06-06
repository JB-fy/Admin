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
		//group.Middleware(middleware.HandlerResponse) // 现在没啥用！如果cotroller方法是用规范路由写的才有用
		group.Middleware(middleware.Cross, middleware.I18n)
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
				controllerThis := controller.NewLogin()
				group.ALLMap(g.Map{
					"/encryptStr": controllerThis.EncryptStr,
					"/":           controllerThis.Login,
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
					controllerThis := controller.NewLogin()
					group.ALLMap(g.Map{
						"/info":     controllerThis.Info,
						"/update":   controllerThis.Update,
						"/menuTree": controllerThis.MenuTree,
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
						"/tree":   controllerThis.Tree,
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
