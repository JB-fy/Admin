package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"

	"api/internal/controller"
	"api/internal/middleware"
	daoLog "api/internal/model/dao/log"
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/router"
	"api/internal/utils"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			/**--------时区设置 开始--------**/
			gtime.SetTimeZone("Asia/Shanghai")
			/**--------时区设置 结束--------**/

			/**--------多语言设置 开始--------**/
			g.I18n().SetPath(g.Cfg().MustGet(ctx, "i18n.path").String())         //设置资源目录
			g.I18n().SetLanguage(g.Cfg().MustGet(ctx, "i18n.language").String()) //设置默认为中文（原默认为英文en）
			/**--------多语言设置 结束--------**/

			/**--------设置当前服务器IP并记录 开始--------**/
			serverNetworkIp := utils.GetServerNetworkIp()
			serverLocalIp := utils.GetServerLocalIp()
			// g.Cfg().Set("server.networkIp", serverNetworkIp);   //设置服务器外网ip
			// g.Cfg().Set("server.localIp", serverLocalIp);   //设置服务器内网ip
			daoPlatform.Server.ParseDbCtx(ctx).Data(g.Map{"networkIp": serverNetworkIp, "localIp": serverLocalIp}).Save()
			/**--------设置当前服务器IP并记录 结束--------**/

			/**--------数据库表分区 开始--------**/
			utils.DbTablePartition(ctx, daoLog.Request.Group(), daoLog.Request.Table(), 7, 24*60*60, `createAt`) //请求日志
			/**--------数据库表分区 结束--------**/

			/*--------启动http服务 开始--------*/
			s := g.Server()
			s.BindHandler("/", func(r *ghttp.Request) {
				r.Response.RedirectTo("/view/admin/platform")
			})
			s.Group("/upload", func(group *ghttp.RouterGroup) {
				group.ALL("/notify", controller.NewUpload().Notify)
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
			/*--------启动http服务 结束--------*/
			return nil
		},
	}
)
