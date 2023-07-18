package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerAuth "api/internal/controller/platform/auth"
	controllerLogin "api/internal/controller/platform/login"
	controllerMy "api/internal/controller/platform/my"
	controllerPlatform "api/internal/controller/platform/platform"
	"api/internal/middleware"
)

func InitRouterPlatform(s *ghttp.Server) {
	s.Group(`/platform`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		//无需验证登录身份
		group.Group(`/login`, func(group *ghttp.RouterGroup) {
			group.Bind(controllerLogin.NewAdmin())
		})

		//需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfPlatform)

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(
					controllerThis.Sign,
					controllerThis.Sts,
				)
			})

			group.Group(`/my/admin`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewAdmin())
			})

			group.Group(`/my/menu`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewMenu())
			})

			group.Group(`/my/action`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewAction())
			})

			group.Group(`/auth/action`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerAuth.NewAction())
			})

			group.Group(`/auth/menu`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerAuth.NewMenu())
			})

			group.Group(`/auth/role`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerAuth.NewRole())
			})

			group.Group(`/auth/scene`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerAuth.NewScene())
			})

			group.Group(`/platform/admin`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerPlatform.NewAdmin())
			})

			group.Group(`/platform/config`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerPlatform.NewConfig())
			})

			/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
