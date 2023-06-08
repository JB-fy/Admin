package controller

import (
	"api/api"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (c *Test) TestMeta(ctx context.Context, req *api.TestMetaReq) (res *api.TestMetaRes, err error) {
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

func (c *Test) Test(r *ghttp.Request) {
	/* var req *api.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	} */

	// fmt.Println(garray.NewStrArrayFrom([]string{"a", "b", "c"}).Contains("a"))

	// fmt.Println(gset.NewIntSetFrom([]int{1, 2, 3}).Diff(gset.NewIntSetFrom([]int{1, 3})).Slice())

	// fmt.Println(grand.N(1000, 9999))
	// fmt.Println(grand.Intn(1))
	// fmt.Println(grand.Str("abcdefg0123456789", 8))
	// fmt.Println(grand.S(8))       //abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789
	// fmt.Println(grand.Digits(8))  //0123456789
	// fmt.Println(grand.Letters(8)) //abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
	// fmt.Println(grand.Symbols(8)) //!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~

	// fmt.Println(daoAuth.RoleRelToMenu.ParseDbCtx(r.GetCtx()).Where("roleId", 1).Array("menuId"))

	// fmt.Println(g.Cfg().MustGet(r.GetCtx(), "superPlatformAdminId").Int())

	/* r.Response.WriteJson(map[string]interface{}{
		"code": 0,
		"msg":  "成功",
		"data": map[string]interface{}{},
	}) */

	utils.HttpSuccessJson(r, map[string]interface{}{
		"list": []map[string]interface{}{},
	}, 0)

	utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 99999999, ""))
}
