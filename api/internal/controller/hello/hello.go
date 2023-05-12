package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	v1 "api/api/hello/v1"
	dao "api/internal/dao/auth"
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
	dao.Menu.PrimaryKey()
	//fmt.Println(r.GetCtx())
	r.Response.Writeln("Test")
}
