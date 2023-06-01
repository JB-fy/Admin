package hello

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	v1 "api/api/hello/v1"
	daoAuth "api/internal/dao/auth"
	daoLog "api/internal/dao/log"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	//fmt.Println(ctx)
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

func (c *Controller) Test(r *ghttp.Request) {
	//fmt.Println(r.GetCtx())
	//fmt.Println(r.Context())

	var req *v1.TestReq
	err := r.Parse(&req)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	var req1 *v1.TestReq
	r.Parse(&req1)
	fmt.Println(req)
	fmt.Println(req1)
	fmt.Println(r.GetCtxVar("a"))
	r.SetCtxVar("a", "aaa")
	fmt.Println(r.GetCtxVar("a").String())

	fmt.Println(daoLog.Request.Info(r.GetCtx(), []string{"logId", "createTime"}, g.Map{"logId": 6}, [][2]string{}))

	joinCodeArr := []string{}
	daoAuth.Menu.Ctx(r.GetCtx()).Handler(daoAuth.Menu.ParseField([]string{"id", "createTime"}, &joinCodeArr), daoAuth.Menu.ParseFilter(g.Map{"id": 2, "menuId > ?": 22}, &joinCodeArr)).All()

	r.SetError(gerror.New("aaaa"))
	/* r.SetError(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
	code := gerror.Code(r.GetError())
	fmt.Println(code)
	type DefaultHandlerResponse struct {
		Code    int         `json:"code"    dc:"Error code"`
		Message string      `json:"message" dc:"Error message"`
		Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
	}
	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code.Code(),
		Message: code.Message(),
		Data:    code.Detail(),
	}) */
	//r.Response.Writeln("Test")
}
