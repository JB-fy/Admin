package router

import (
	"api/internal/controller"

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
			// controllerThis.Sign, //建议在场景内验证登录token后才可调用
			controllerThis.Sts, //App端的SDK需设置一个地址来获取Sts Token，该地址不验证登录token。需同时在场景内设置需验证登录token的路由，用于给App端获取实际上传时需用到的字段
			controllerThis.Notify,
			controllerThis.Upload,
		)
	})
	//测试
	s.Group(``, func(group *ghttp.RouterGroup) {
		group.Bind(controller.NewTest())
	})
}
