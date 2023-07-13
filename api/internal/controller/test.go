package controller

import (
	"api/api"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/genv"
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

	// fmt.Println(garray.NewStrArrayFrom([]string{`a`, `b`, `c`}).Contains(`a`))	//是否含有元素

	// fmt.Println(gset.NewIntSetFrom([]int{1, 2, 3}).Diff(gset.NewIntSetFrom([]int{1, 3})).Slice())	//差集

	// fmt.Println(genv.Set(`X_X`, `xx`))              //key必须由大写和_组成
	// fmt.Println(g.Cfg().MustGetWithEnv(ctx, `X_X`)) //X_X或x_x或x.x方法都可以读取到
	fmt.Println(genv.All())

	// fmt.Println(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Int())	//获取配置参数

	//fmt.Println(ghttp.RestartAllServer(ctx))	//重启服务

	/* utils.HttpWriteJson(ctx, map[string]interface{}{
		`info`: map[string]interface{}{},
	}, 0, ``) */
	// g.RequestFromCtx(ctx).Response.Status = http.StatusMultipleChoices
	res = &api.TestRes{
		Test: `测试`,
	}
	return
}

func (c *Test) TestNoMeta(r *ghttp.Request) {
	/* var req *api.TestNoMetaReq
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
