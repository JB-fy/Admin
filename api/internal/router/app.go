package router

import (
	"api/internal/controller"
	controllerIndex "api/internal/controller/app"
	"api/internal/controller/app/my"
	"api/internal/controller/app/platform"
	"api/internal/middleware"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterApp(ctx context.Context, s *ghttp.Server) {
	s.Group(`/app`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Scene())

		// 无需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Group(`/login`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerIndex.NewLogin())
			})
		})

		// 无需验证登录身份（但存在token时，会做解析，且忽视错误）
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfApp(false))

			group.Group(`/code`, func(group *ghttp.RouterGroup) {
				group.Bind(controllerIndex.NewCode())
			})

			group.Group(`/platform`, func(group *ghttp.RouterGroup) {
				group.Bind(platform.NewConfig())
			})
		})

		// 需验证登录身份
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.SceneLoginOfApp(true))

			group.Group(`/upload`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewUpload()
				group.Bind(controllerThis.Sign)
				group.Bind(controllerThis.Config)
			})

			group.Group(`/pay`, func(group *ghttp.RouterGroup) {
				controllerThis := controller.NewPay()
				group.Bind(
					controllerThis.List,
					controllerThis.Pay,
				)
			})

			group.Group(`/my`, func(group *ghttp.RouterGroup) {
				group.Bind(my.NewProfile())
			})

			/*--------后端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/
		})
	})
}
