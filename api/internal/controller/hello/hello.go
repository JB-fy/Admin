package hello

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"

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
	daoLog.Request.Ctx(r.GetCtx()).Fields("logId").Where("logId", 6).Order("logId", "Desc").OrderAsc("createTime").All()
	//daoLog.Request.Ctx(r.GetCtx()).Data("runTime", 2, "requestUrl", "1").Insert()
	//daoLog.Request.Ctx(r.GetCtx()).Data(g.Map{"requestUrl": "1", "runTime": 2}).Where("logId", 6).Update()
	joinCode := []string{}
	daoAuth.Menu.Ctx(r.GetCtx()).Handler(daoAuth.Menu.ParseField([]string{"id", "createTime"}, &[]string{}, &joinCode), daoAuth.Menu.ParseFilter(g.MapStrAny{"id": 2, "menuId > ?": 22}, &joinCode)).All()
	fmt.Println(daoAuth.Menu.ColumnArr())
	fmt.Println(daoAuth.Menu.Columns())
	fmt.Println(gconv.Map(daoAuth.Menu.Columns()))
	fmt.Println(gmap.NewStrAnyMapFrom(gconv.Map(daoAuth.Menu.Columns())).Values())
	/* res, _ := daoLog.Request.Ctx(r.GetCtx()).Where("logId", 6).All()
	fmt.Println(res) */
	//fmt.Println(r.GetCtx())
	r.Response.Writeln("Test")
}
