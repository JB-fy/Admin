package router

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerAuth "api/internal/controller/platform/auth"
	controllerLog "api/internal/controller/platform/log"
	controllerLogin "api/internal/controller/platform/login"
	controllerPlatform "api/internal/controller/platform/platform"
	"api/internal/middleware"
)

func InitRouterPlatform(s *ghttp.Server) {
	s.Group("/platform", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Cross, middleware.I18n)
		//不做日志记录
		group.Group("", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.HandlerResponse) // 不用规范路由方式可去掉。且如果有用log中间件，必须放在其后面，才能读取到响应数据
			group.Middleware(middleware.Scene)
			//需验证登录身份
			group.Group("", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.SceneLoginOfPlatformAdmin)
				group.Group("/log/http", func(group *ghttp.RouterGroup) {
					controllerThis := controllerLog.NewHttp()
					group.Bind(controllerThis)
				})
			})
		})

		//做日志记录
		group.Group("", func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.Log)
			group.Middleware(middleware.HandlerResponse) // 不用规范路由方式可去掉。且如果有用log中间件，必须放在其后面，才能读取到响应数据
			group.Middleware(middleware.Scene)
			//无需验证登录身份
			group.Group("/login", func(group *ghttp.RouterGroup) {
				controllerThis := controllerLogin.NewLogin()
				group.Bind(
					controllerThis.EncryptStr,
					controllerThis.Login,
				)
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
					controllerThis := controllerLogin.NewLogin()
					group.ALLMap(g.Map{
						"/info":     controllerThis.Info,
						"/update":   controllerThis.Update,
						"/menuTree": controllerThis.MenuTree,
					})
				})

				group.Group("/auth/action", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewAction()
					group.Bind(controllerThis)
				})

				group.Group("/auth/menu", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewMenu()
					group.Bind(controllerThis)
				})

				group.Group("/auth/role", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewRole()
					group.Bind(controllerThis)
				})

				group.Group("/auth/scene", func(group *ghttp.RouterGroup) {
					controllerThis := controllerAuth.NewScene()
					group.Bind(controllerThis)
				})

				group.Group("/platform/admin", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewAdmin()
					group.Bind(controllerThis)
				})

				group.Group("/platform/config", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewConfig()
					group.Bind(controllerThis)
				})

				group.Group("/platform/server", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewServer()
					group.Bind(controllerThis)
				})

				group.Group("/platform/corn", func(group *ghttp.RouterGroup) {
					controllerThis := controllerPlatform.NewCorn()
					group.Bind(controllerThis)
				})

				/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
			})
		})
	})
}
