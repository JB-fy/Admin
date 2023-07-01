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

func (c *Test) TestMeta(ctx context.Context, req *api.TestMetaReq) (res *api.TestMetaRes, err error) {
	// time.Sleep(10 * time.Second)
	// utils.HttpFailJson(g.RequestFromCtx(ctx), utils.NewErrorCode(ctx, 99999999, ``))
	// g.RequestFromCtx(ctx).Response.Writeln(`TestMeta`)
	// g.RequestFromCtx(ctx).Response.Status = http.StatusMultipleChoices
	res = &api.TestMetaRes{
		Test: `测试`,
	}
	return
}

func (c *Test) Test(r *ghttp.Request) {
	/* var req *api.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	} */

	// g.Log().Info(r.GetCtx(), `日志打印`)
	// fmt.Println(g.I18n().T(ctx, `code.99999999`))

	// fmt.Println(garray.NewStrArrayFrom([]string{`a`, `b`, `c`}).Contains(`a`))

	// fmt.Println(gset.NewIntSetFrom([]int{1, 2, 3}).Diff(gset.NewIntSetFrom([]int{1, 3})).Slice())

	// fmt.Println(grand.N(1000, 9999))
	// fmt.Println(grand.Intn(1))
	// fmt.Println(grand.Str(`abcdefg0123456789`, 8))
	// fmt.Println(grand.S(8))       //abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
	// fmt.Println(grand.Digits(8))  //0123456789
	// fmt.Println(grand.Letters(8)) //abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	// fmt.Println(grand.Symbols(8)) //!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~

	// fmt.Println(daoAuth.RoleRelToMenu.ParseDbCtx(r.GetCtx()).Where(`roleId`, 1).Array(`menuId`))

	// fmt.Println(genv.Set(`X_X`, `xx`))                     //key必须由大写和_组成
	// fmt.Println(g.Cfg().MustGetWithEnv(r.GetCtx(), `X_X`)) //X_X或x_x或x.x方法都可以读取到

	// fmt.Println(g.Cfg().MustGet(r.GetCtx(), `superPlatformAdminId`).Int())

	//fmt.Println(ghttp.RestartAllServer(r.GetCtx()))

	// r.SetError(utils.NewErrorCode(r.GetCtx(), 99999999, ``))
	r.Response.WriteJson(map[string]interface{}{
		`code`: 0,
		`msg`:  g.I18n().T(r.GetCtx(), `code.0`),
		`data`: map[string]interface{}{
			`list`: []map[string]interface{}{},
		},
	})
}
