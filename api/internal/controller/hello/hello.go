package hello

import (
	"context"

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
	daoLog.Request.Ctx(r.GetCtx()).Where("logId", 6).OrderAsc("logId").OrderAsc("createTime").All()
	//daoLog.Request.Ctx(r.GetCtx()).Data("runTime", 2, "requestUrl", "1").Insert()
	//daoLog.Request.Ctx(r.GetCtx()).Data(g.Map{"requestUrl": "1", "runTime": 2}).Insert()
	joinCode := []string{}
	daoAuth.Menu.Ctx(r.GetCtx()).Handler(daoAuth.Menu.Filter(g.MapStrAny{"id": 2}, &joinCode)).All()
	/* res, _ := daoLog.Request.Ctx(r.GetCtx()).Where("logId", 6).All()
	fmt.Println(res) */
	//fmt.Println(r.GetCtx())
	r.Response.Writeln("Test")
}
