package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"api/internal/controller"
	"api/internal/middleware"
	"api/internal/router"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.I18n().SetPath(g.Cfg().MustGet(ctx, "i18n.path").String())         //设置资源目录
			g.I18n().SetLanguage(g.Cfg().MustGet(ctx, "i18n.language").String()) //设置默认为中文（原默认为英文en）

			s := g.Server()
			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.RedirectTo("/view/admin/platform")
			})
			s.Group("", func(group *ghttp.RouterGroup) {
				//group.Middleware(middleware.HandlerResponse) // 现在没啥用！如果cotroller方法是用规范路由写的才有用
				group.Middleware(middleware.Cross, middleware.I18n)
				group.ALL("/test", controller.NewTest().Test)
				/* group.Bind(
					//controller.NewTest().Test, //这样不会根据方法名自动设置路由
					controller.NewTest(),
				) */
			})
			router.InitRouterPlatformAdmin(s) //平台后台接口注册
			s.Run()
			return nil
		},
	}
)
