package router

import (
	"github.com/gogf/gf/v2/net/ghttp"

	"api/internal/controller"
	"api/internal/middleware"
)

func InitRouterApp(s *ghttp.Server) {
	s.Group(`/app`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene)

		//无需验证登录身份
		group.Group(`/login`, func(group *ghttp.RouterGroup) {
			// group.Bind(controllerLogin.NewAdmin())
		})

		/* // 无需验证登录身份，但带token时，需解析token
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfApp(false))
		}) */

		//需验证登录身份
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
