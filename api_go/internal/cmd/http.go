package cmd

import (
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gvalid"

	"api/internal/controller"
	"api/internal/corn"
	"api/internal/middleware"
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/router"
	"api/internal/utils"
)

func HttpFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	/**--------时区设置 开始--------**/
	gtime.SetTimeZone(`Asia/Shanghai`)
	/**--------时区设置 结束--------**/

	/**--------多语言设置 开始--------**/
	g.I18n().SetPath(g.Cfg().MustGet(ctx, `i18n.path`).String())         //设置资源目录
	g.I18n().SetLanguage(g.Cfg().MustGet(ctx, `i18n.language`).String()) //设置默认为中文（原默认为英文en）
	/**--------多语言设置 结束--------**/

	/**--------设置当前服务器IP并记录 开始--------**/
	serverNetworkIp := utils.GetServerNetworkIp()
	serverLocalIp := utils.GetServerLocalIp()
	genv.Set(`SERVER_NETWORK_IP`, serverNetworkIp) //设置服务器外网ip（key必须由大写和_组成，才能用g.Cfg().MustGetWithEnv()方法读取）
	genv.Set(`SERVER_LOCAL_IP`, serverLocalIp)     //设置服务器内网ip（key必须由大写和_组成，才能用g.Cfg().MustGetWithEnv()方法读取）
	daoPlatform.Server.ParseDbCtx(ctx).Data(g.Map{`networkIp`: serverNetworkIp, `localIp`: serverLocalIp}).Save()
	/**--------设置当前服务器IP并记录 结束--------**/

	/**--------定时器设置 开始--------**/
	corn.LogRequestPartition(ctx) //先执行一次请求日志分区

	corn.InitCorn(ctx) //启动定时器
	/**--------定时器设置 结束--------**/

	/**--------自定义校验规则注册 开始--------**/
	gvalid.RegisterRule(`distinct`, func(ctx context.Context, in gvalid.RuleFuncInput) (err error) {
		val := in.Value.Array()
		if len(val) != garray.NewFrom(val).Unique().Len() {
			err = gerror.Newf(`%s字段具有重复值`, in.Field)
			return
		}
		return
	})
	/**--------自定义校验规则注册 结束--------**/

	/*--------启动http服务 开始--------*/
	s := g.Server()
	//首页
	s.BindHandler(`/`, func(r *ghttp.Request) {
		r.Response.RedirectTo(`/view/platform`)
	})
	//上传回调
	s.Group(`/upload`, func(group *ghttp.RouterGroup) {
		group.ALL(`/notify`, controller.NewUpload().Notify)
	})
	// 测试使用
	s.Group(``, func(group *ghttp.RouterGroup) {
		//group.Middleware(middleware.HandlerResponse) // 现在没啥用！如果cotroller方法是用规范路由写的才有用
		group.Middleware(middleware.Cross, middleware.I18n)
		controllerThis := controller.NewTest()
		group.ALL(`/test`, controllerThis.Test)
		group.Bind(
			controllerThis.TestMeta,
			//controllerThis,
			//controller.NewTest().Test, //这样不会根据方法名自动设置路由
		)
	})

	router.InitRouterPlatform(s) //平台后台接口注册
	s.Run()
	/*--------启动http服务 结束--------*/
	return nil
}
