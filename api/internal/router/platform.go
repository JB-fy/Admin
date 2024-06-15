package router

import (
	"api/internal/controller"
	controllerCurrent "api/internal/controller/platform"
	controllerAuth "api/internal/controller/platform/auth"
	controllerMy "api/internal/controller/platform/my"
	controllerOrg "api/internal/controller/platform/org"
	controllerPlatform "api/internal/controller/platform/platform"
	controllerUsers "api/internal/controller/platform/users"
	"api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterPlatform(s *ghttp.Server) {
	s.Group(`/platform`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		// 无需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Group(`/login`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerCurrent.NewLogin())
			})
		})

		// 无需验证登录身份（但存在token时，会做解析，且忽视错误）
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfPlatform(false))

			group.Group(`/code`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerCurrent.NewCode())
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
				group.Bind(controllerMy.NewProfile())
				group.Bind(controllerMy.NewMenu())
				group.Bind(controllerMy.NewAction())
			})

			group.Group(`/auth`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerAuth.NewAction())
				group.Bind(controllerAuth.NewMenu())
				group.Bind(controllerAuth.NewRole())
				group.Bind(controllerAuth.NewScene())
			})

			group.Group(`/platform`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerPlatform.NewAdmin())
				group.Bind(controllerPlatform.NewConfig())
			})

			group.Group(`/users`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerUsers.NewUsers())
			})

			group.Group(`/org`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerOrg.NewAdmin())
				group.Bind(controllerOrg.NewOrg())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
