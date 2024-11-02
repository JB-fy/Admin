package router

import (
	"api/internal/controller"
	"api/internal/middleware"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gres"
)

func InitRouterCommon(ctx context.Context, s *ghttp.Server) {
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
		)
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.BodyRepeatable(true))
			group.Bind(controllerThis.Notify)
		})
	})
	//支付
	s.Group(`/pay`, func(group *ghttp.RouterGroup) {
		controllerThis := controller.NewPay()
		/* group.Bind(
		// controllerThis.List, //建议在场景内验证登录token后才可调用
		// controllerThis.Pay, //建议在场景内验证登录token后才可调用
		) */
		group.Group(``, func(group *ghttp.RouterGroup) {
			group.Middleware(middleware.BodyRepeatable(true))
			group.Bind(controllerThis.Notify)
		})
	})
	//微信
	s.Group(`/wx`, func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.BodyRepeatable(true))
		controllerThis := controller.NewWx()
		group.Bind(controllerThis.GzhNotify)
	})
	//开发环境用
	if utils.IsDev(ctx) {
		s.Group(``, func(group *ghttp.RouterGroup) {
			group.Bind(controller.NewTest()) //测试
			// 新文档（框架文档使用的https://unpkg.com/redoc@2.0.0-rc.70/bundles/redoc.standalone.js文件可能被墙）
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
}
