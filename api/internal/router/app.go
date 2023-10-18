package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerLogin "api/internal/controller/app/login"
	controllerSms "api/internal/controller/app/sms"
	"api/internal/middleware"
)

func InitRouterApp(s *ghttp.Server) {
	s.Group(`/app`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		// 无需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Group(`/login`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerLogin.NewUser())
			})
		})

		// 无需验证登录身份（但存在token时，会做解析，且忽视错误）
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfApp(false))

			group.Group(`/sms`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerSms.NewSms())
			})
		})

		// 需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfApp(true))

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(controllerThis.Config)
			})

			/* group.Group(`/my`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewProfile())
			}) */

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
