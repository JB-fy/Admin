package router

import (
	"api/internal/controller"
	controllerIndex "api/internal/controller/platform"
	"api/internal/controller/platform/app"
	"api/internal/controller/platform/auth"
	"api/internal/controller/platform/my"
	"api/internal/controller/platform/org"
	"api/internal/controller/platform/pay"
	"api/internal/controller/platform/platform"
	"api/internal/controller/platform/upload"
	"api/internal/controller/platform/users"
	"api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterPlatform(s *ghttp.Server) {
	s.Group(`/platform`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		// 无需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Group(`/login`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerIndex.NewLogin())
			})
		})

		// 无需验证登录身份（但存在token时，会做解析，且忽视错误）
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfPlatform(false))

			group.Group(`/code`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerIndex.NewCode())
			})
		})

		// 需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfPlatform(true))

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(controllerThis.Sign)
			})

			group.Group(`/my`, func(group *ghttp.RouterGroup) {
				group.Bind(my.NewProfile())
				group.Bind(my.NewMenu())
				group.Bind(my.NewAction())
			})

			group.Group(`/auth`, func(group *ghttp.RouterGroup) {
				group.Bind(auth.NewAction())
				group.Bind(auth.NewMenu())
				group.Bind(auth.NewRole())
				group.Bind(auth.NewScene())
			})

			group.Group(`/platform`, func(group *ghttp.RouterGroup) {
				group.Bind(platform.NewAdmin())
				group.Bind(platform.NewConfig())
			})

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				group.Bind(upload.NewUpload())
			})

			group.Group(`/pay`, func(group *ghttp.RouterGroup) {
				group.Bind(pay.NewChannel())
				group.Bind(pay.NewScene())
				group.Bind(pay.NewPay())
			})

			group.Group(`/app`, func(group *ghttp.RouterGroup) {
				group.Bind(app.NewApp())
			})

			group.Group(`/users`, func(group *ghttp.RouterGroup) {
				group.Bind(users.NewUsers())
			})

			group.Group(`/org`, func(group *ghttp.RouterGroup) {
				group.Bind(org.NewAdmin())
				group.Bind(org.NewOrg())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
