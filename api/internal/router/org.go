package router

import (
	"api/internal/controller"
	controllerCurrent "api/internal/controller/org"
	controllerMy "api/internal/controller/org/my"
	"api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterOrg(s *ghttp.Server) {
	s.Group(`/org`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		// 无需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Group(`/login`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerCurrent.NewLogin())
			})
		})

		// 无需验证登录身份（但存在token时，会做解析，且忽视错误）
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfOrg(false))

			group.Group(`/code`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerCurrent.NewCode())
			})
		})

		// 需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfOrg(true))

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(controllerThis.Sign)
			})

			group.Group(`/my`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewProfile())
				group.Bind(controllerMy.NewMenu())
				group.Bind(controllerMy.NewAction())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
