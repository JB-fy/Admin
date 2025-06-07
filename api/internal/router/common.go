package router

import (
	"api/internal/controller"
	"api/internal/middleware"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
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
	//APP更新检测
	s.Group(`/app`, func(group *ghttp.RouterGroup) {
		group.Bind(controller.NewAppPkg())
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
			group.Bind(controller.NewTest())
			group.GET(`/swaggerNew`, func(r *ghttp.Request) {
				// r.Response.Write(`<!DOCTYPE html><html><head><title>API Reference</title><meta charset="utf-8"/><meta name="viewport" content="width=device-width, initial-scale=1"></head><body style="margin:  0; padding: 0;"><redoc spec-url="/api.json" show-object-schema-examples="true"></redoc><script>` + string(gres.GetContent(`/goframe/swaggerui/redoc.standalone.js`)) + `</script></body></html>`) // 搜索功能不支持中文，且框架文档使用的https://unpkg.com/redoc@2.0.0-rc.70/bundles/redoc.standalone.js有时无法访问（国内屏蔽）

				// r.Response.Write(`<!DOCTYPE html><html lang="en"><head><meta charset="utf-8" /><meta name="viewport" content="width=device-width, initial-scale=1" /><meta name="description" content="SwaggerUI" /><title>SwaggerUI</title><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" /></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script><script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-standalone-preset.js" crossorigin></script><script>window.onload = () => { window.ui = SwaggerUIBundle({ url: '/api.json', dom_id: '#swagger-ui', filter: true }) }</script></body></html>`)	// SwaggerUI。没有菜单栏

				r.Response.Write(`<!doctype html><html><head><meta charset="utf-8"><script type="module" src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"></script></head><body><rapi-doc spec-url="/api.json"> </rapi-doc></body></html>`) //rapidoc。功能最齐全

				// r.Response.Write(`<html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"><title>Elements in HTML</title><script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script><link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css"></head><body><elements-api apiDescriptionUrl="/api.json" router="hash"/></body></html>`)// stoplight elements。没有搜索功能，不然最好用
			})
		})
	}
}
