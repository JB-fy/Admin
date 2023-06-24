package router

import (
	"api/internal/controller"
	"api/internal/middleware"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterCommon(s *ghttp.Server) {
	//首页
	s.BindHandler(`/`, func(r *ghttp.Request) {
		r.Response.RedirectTo(`/view/platform`)
	})
	//上传
	s.Group(`/upload`, func(group *ghttp.RouterGroup) {
		controllerThis := controller.NewUpload()
		group.Bind(
			// controllerThis.Sign, //建议放其他场景内验证权限后才可调用
			controllerThis.Sts, //App端的SDK需设置一个地址来获取Sts Token，该地址不验证权限。需同时在其他场景路由下设置需验证权限的路由，用于给App端获取实际上传时需用到的字段
			controllerThis.Notify,
		)
	})
	//测试
	s.Group(``, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Cross, middleware.I18n)
		// group.Middleware(middleware.HandlerResponse) // 不用规范路由方式可去掉。且如果有用log中间件，必须放在其后面，才能读取到响应数据
		controllerThis := controller.NewTest()
		group.ALL(`/test`, controllerThis.Test)
		group.Bind(controllerThis.TestMeta)
	})
}
