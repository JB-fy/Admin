package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	controllerLogin "api/internal/controller/demo/login"
	controllerMy "api/internal/controller/demo/my"
	"api/internal/middleware"
)

func InitRouterDemo(s *ghttp.Server) {
	s.Group(`/demo`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		//无需验证登录身份
		group.Group(`/login`, func(group *ghttp.RouterGroup) {
			group.Bind(controllerLogin.NewAdmin())
		})

		//需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfDemo)

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(
					controllerThis.Sign,
					controllerThis.Sts,
				)
			})

			group.Group(`/my`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerMy.NewProfile())
				group.Bind(controllerMy.NewMenu())
				group.Bind(controllerMy.NewAction())
			})

			/*--------自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
