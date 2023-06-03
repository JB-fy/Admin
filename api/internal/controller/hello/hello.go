package hello

import (
	"context"
	"fmt"

	"github.com/gogf/gf/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	v1 "api/api/hello/v1"
	daoAuth "api/internal/model/dao/auth"
	daoLog "api/internal/model/dao/log"
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
	panic(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
	fmt.Println(1)
	//fmt.Println(r.GetCtx())
	//fmt.Println(r.Context())
	var req *v1.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	fmt.Println(req)

	r.SetCtxVar("a", "aaa")
	fmt.Println(r.GetCtxVar("a"))
	fmt.Println(r.Context().Value("a"))

	r.SetError(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
	r.Response.WriteJson(map[string]interface{}{
		"code": 0,
		"msg":  "成功",
		"data": map[string]interface{}{},
	})

	daoLog.Request.Info(r.GetCtx(), []string{"logId", "createTime"}, g.Map{"logId": 6}, [][2]string{})
	joinCodeArr := []string{}
	daoAuth.Menu.Ctx(r.GetCtx()).Handler(daoAuth.Menu.ParseField([]string{"id", "createTime"}, &joinCodeArr), daoAuth.Menu.ParseFilter(g.Map{"id": 2, "menuId > ?": 22}, &joinCodeArr)).All()
}
