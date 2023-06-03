package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"api/internal/controller/hello"
	"api/internal/middleware"
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
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(middleware.HandlerResponse)
				group.Middleware(middleware.Cross)
				group.Middleware(middleware.I18n)
				group.Middleware(middleware.Log)
				group.Bind(
					hello.New(),
				)
			})
			s.Run()
			return nil
		},
	}
)
