package router

import (
	"api/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func InitRouterCommon(s *ghttp.Server) {
	//首页
	s.BindHandler(`/`, func(r *ghttp.Request) {
		r.Response.RedirectTo(`/admin/platform`)
	})
	//上传
	s.Group(`/upload`, func(group *ghttp.RouterGroup) {
		controllerThis := controller.NewUpload()
		group.Bind(
			controllerThis.Upload,
			// controllerThis.Sign, //建议在场景内验证登录token后才可调用
			// controllerThis.Config, //建议在场景内验证登录token后才可调用
			controllerThis.Sts,
			controllerThis.Notify,
		)
	})
	//支付回调
	s.Group(`/pay`, func(group *ghttp.RouterGroup) {
		group.Bind(controller.NewPay())
	})
	//测试
	s.Group(``, func(group *ghttp.RouterGroup) {
		group.Bind(controller.NewTest())
	})
}
