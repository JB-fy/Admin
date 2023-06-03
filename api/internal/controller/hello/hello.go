package hello

import (
	"context"
	"fmt"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	v1 "api/api/hello/v1"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	//fmt.Println(ctx)
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	res = &v1.Res{
		UserName: "aaa",
	}
	err = gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"}))
	code := gerror.Code(err)
	fmt.Println(code)
	return
}

func (c *Controller) Test(r *ghttp.Request) {
	fmt.Println(g.I18n().T(r.GetCtx(), "0"))
	fmt.Println(g.I18n().T(r.GetCtx(), "99999999"))
	fmt.Println(g.I18n().Tf(r.GetCtx(), "29991063", "phone"))
	var req *v1.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	fmt.Println(req)
	param := r.GetMap()
	fmt.Println(param)

	/* r.SetError(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
	r.Response.WriteJson(map[string]interface{}{
		"code": 0,
		"msg":  "成功",
		"data": map[string]interface{}{},
	}) */
}
