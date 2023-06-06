package controller

import (
	"api/api"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Test struct{}

func NewTest() *Test {
	return &Test{}
}

func (c *Test) TestMeta(ctx context.Context, req *api.TestMetaReq) (res *api.TestMetaRes, err error) {
	//fmt.Println(ctx)
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	res = &api.TestMetaRes{
		UserName: "aaa",
	}
	err = gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"}))
	code := gerror.Code(err)
	fmt.Println(code)
	return
}

func (c *Test) Test(r *ghttp.Request) {
	/* var req *api.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	fmt.Println(req) */
	/* r.Response.WriteJson(map[string]interface{}{
		"code": 0,
		"msg":  "成功",
		"data": map[string]interface{}{},
	}) */
	fmt.Println(g.Cfg().MustGet(r.GetCtx(), "superPlatformAdminId").Int())

	utils.HttpSuccessJson(r, map[string]interface{}{
		"list": []map[string]interface{}{},
	}, 0)
	err := utils.NewErrorCode(r.GetCtx(), 99999999, "")
	utils.HttpFailJson(r, err)
}
