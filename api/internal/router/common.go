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
			// controllerThis.Sign, //建议放其他场景内验证权限后才可调用
			controllerThis.Sts, //App端的SDK需设置一个地址来获取Sts Token，该地址不验证权限。需同时在其他场景路由下设置需验证权限的路由，用于给App端获取实际上传时需用到的字段
			controllerThis.Notify,
		)
	})
	//测试
	s.Group(``, func(group *ghttp.RouterGroup) {
		controllerThis := controller.NewTest()
		group.Bind(controllerThis.Test)
		group.ALL(`/test1`, controllerThis.Test1)
	})
}
