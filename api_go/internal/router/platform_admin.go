package router

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerAuth "api/internal/controller/auth"
	controllerLog "api/internal/controller/log"
	controllerPlatform "api/internal/controller/platform"
	"api/internal/middleware"
)

func InitRouterPlatform(s *ghttp.Server) {
	s.Group("/platform", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.HandlerResponse) // 现在没啥用！如果cotroller方法是用规范路由写的才有用
		group.Middleware(middleware.Cross, middleware.I18n)
		//不做日志记录
		group.Group("", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Scene)
			//需验证登录身份
			group.Group("", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.SceneLoginOfPlatformAdmin)
				group.Group("/log/request", func(group *ghttp.RouterGroup) {
					controllerThis := controllerLog.NewRequest()
					group.ALLMap(g.Map{
						"/list": controllerThis.List,
					})
				})
			})
		})

		//做日志记录
		group.Group("", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Log, middleware.Scene)
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
					controllerThis := controller.NewUpload()
					group.ALLMap(g.Map{
						"/sign": controllerThis.Sign,
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
					controllerThis := controllerPlatform.NewAdmin()
					group.ALLMap(g.Map{
						"/list":   controllerThis.List,
						"/info":   controllerThis.Info,
						"/create": controllerThis.Create,
						"/update": controllerThis.Update,
						"/del":    controllerThis.Delete,
					})
				})

				group.Group("/platform/config", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewConfig()
					group.ALLMap(g.Map{
						"/get":  controllerThis.Get,
						"/save": controllerThis.Save,
					})
				})

				group.Group("/platform/server", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewServer()
					group.ALLMap(g.Map{
						"/list": controllerThis.List,
					})
				})

				group.Group("/platform/corn", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewCorn()
					group.ALLMap(g.Map{
						"/list":   controllerThis.List,
						"/info":   controllerThis.Info,
						"/create": controllerThis.Create,
						"/update": controllerThis.Update,
						"/del":    controllerThis.Delete,
					})
				})

				/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
			})
		})
	})
}
