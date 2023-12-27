package router

import (
	"api/internal/controller"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gres"
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
			// controllerThis.Sign,   //建议在场景内验证登录token后才可调用
			// controllerThis.Config, //建议在场景内验证登录token后才可调用
			controllerThis.Sts,
			controllerThis.Notify,
		)
	})
	//支付
	s.Group(`/pay`, func(group *ghttp.RouterGroup) {
		controllerThis := controller.NewPay()
		group.Bind(
			// controllerThis.Pay, //建议在场景内验证登录token后才可调用
			controllerThis.Notify,
		)
	})
	//测试
	s.Group(``, func(group *ghttp.RouterGroup) {
		group.Bind(controller.NewTest())
	})
	//新文档（框架文档使用的https://unpkg.com/redoc@2.0.0-rc.70/bundles/redoc.standalone.js文件可能被墙）
	s.Group(``, func(group *ghttp.RouterGroup) {
		group.GET(`/swaggerNew`, func(r *ghttp.Request) {
			r.Response.Write(`<!DOCTYPE html>
<html>
	<head>
	<title>API Reference</title>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body {
			margin:  0;
			padding: 0;
		}
	</style>
	</head>
	<body>
		<redoc spec-url="/api.json" show-object-schema-examples="true"></redoc>
		<script>` + string(gres.GetContent(`/goframe/swaggerui/redoc.standalone.js`)) + `</script>
	</body>
</html>`)
		})
	})
}
