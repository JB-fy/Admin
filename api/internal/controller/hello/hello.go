package hello

import (
	"context"
	"fmt"

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

	var req *v1.Req
	err := r.Parse(&req)
	fmt.Println(err)

	fmt.Println(daoLog.Request.Info(r.GetCtx(), []string{"logId", "createTime"}, g.Map{"logId": 6}, [][2]string{}))

	joinCodeArr := []string{}
	daoAuth.Menu.Ctx(r.GetCtx()).Handler(daoAuth.Menu.ParseField([]string{"id", "createTime"}, &joinCodeArr), daoAuth.Menu.ParseFilter(g.Map{"id": 2, "menuId > ?": 22}, &joinCodeArr)).All()

	/* res, _ := daoLog.Request.Ctx(r.GetCtx()).Where("logId", 6).All()
	fmt.Println(res) */
	//fmt.Println(r.GetCtx())
	r.Response.Writeln("Test")
}
