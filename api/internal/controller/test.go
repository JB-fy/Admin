package controller

import (
	"api/api"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (c *Test) Test(ctx context.Context, req *api.TestReq) (res *api.TestRes, err error) {
	// time.Sleep(10 * time.Second)

	// fmt.Println(g.RequestFromCtx(ctx).GetUrl())   // http://192.168.2.200:20080/testMeta?a=1&b=2
	// fmt.Println(g.RequestFromCtx(ctx).GetHost())  // 192.168.2.200
	// fmt.Println(g.RequestFromCtx(ctx).Host)       // 192.168.2.200:20080
	// fmt.Println(g.RequestFromCtx(ctx).RequestURI) // /testMeta?a=1&b=2

	// fmt.Println(g.RequestFromCtx(ctx).URL.String())     // /testMeta?a=1&b=2
	// fmt.Println(g.RequestFromCtx(ctx).URL)              // /testMeta?a=1&b=2
	// fmt.Println(g.RequestFromCtx(ctx).URL.Path)         // /testMeta
	// fmt.Println(g.RequestFromCtx(ctx).URL.RawQuery)     // a=1&b=2
	// fmt.Println(g.RequestFromCtx(ctx).URL.RequestURI()) // /testMeta?a=1&b=2

	// fmt.Println(g.RequestFromCtx(ctx).Router)     // &{/testMeta GET default ^/testMeta$ [] 0}
	// fmt.Println(g.RequestFromCtx(ctx).Router.Uri) // /testMeta

	// fmt.Println(grand.N(1000, 9999))
	// fmt.Println(grand.Intn(1))
	// fmt.Println(grand.Str(`abcdefg0123456789`, 8))
	// fmt.Println(grand.S(8))       //abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
	// fmt.Println(grand.Digits(8))  //0123456789
	// fmt.Println(grand.Letters(8)) //abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	// fmt.Println(grand.Symbols(8)) //!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~

	// g.Log().Info(ctx, `日志打印`)
	// fmt.Println(g.I18n().T(ctx, `code.99999999`))

	/* //合并Map
	rawMap := gmap.NewStrAnyMapFrom(g.Map{`a`: 1, `b`: 2})
	rawMap.Merge(gmap.NewStrAnyMapFrom(g.Map{`a`: 4, `c`: 3}))
	fmt.Println(rawMap.Map()) */

	// fmt.Println(garray.NewStrArrayFrom([]string{`a`, `b`, `c`}).Contains(`a`))	//是否含有元素

	// fmt.Println(gset.NewIntSetFrom([]int{1, 2, 3}).Diff(gset.NewIntSetFrom([]int{1, 3})).Slice())	//差集

	// g.Validator().Rules(`required|integer`).Data(`aaaa`).Run(ctx) //单独验证

	// fmt.Println(genv.Set(`X_X`, `xx`))              //key必须由大写和_组成
	// fmt.Println(g.Cfg().MustGetWithEnv(ctx, `X_X`)) //X_X或x_x或x.x方法都可以读取到

	// fmt.Println(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint())	//获取配置参数

	// fmt.Println(g.DB(`default`).Model(`tab_user_unsubscribe`).Safe().Ctx(ctx))	//数据库连接

	// fmt.Println(ghttp.RestartAllServer(ctx))	//重启服务

	// g.RequestFromCtx(ctx).Response.Status = http.StatusMultipleChoices

	// err = utils.NewErrorCode(ctx, 99999999, ``)

	/* utils.HttpWriteJson(ctx, map[string]interface{}{
		`info`: map[string]interface{}{},
	}, 0, ``) */
	res = &api.TestRes{}
	/* info, _ := g.DB().Model(`auth_scene`).Ctx(ctx).One()
	info.Struct(&res.Info) */
	return
}

func (c *Test) Test1(r *ghttp.Request) {
	/* var req *api.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	} */

	// r.SetError(utils.NewErrorCode(r.GetCtx(), 99999999, ``))
	r.Response.WriteJson(map[string]interface{}{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: map[string]interface{}{
			`list`: []map[string]interface{}{},
		},
	})
}
