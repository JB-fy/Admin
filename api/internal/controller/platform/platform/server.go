package controller

import (
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// 列表
func (controllerThis *Server) List(ctx context.Context, req *apiPlatform.ServerListReq) (res *apiPlatform.ServerListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoPlatform.Server.Columns()
	allowField := daoPlatform.Server.ColumnArr()
	allowField = append(allowField, `id`, `name`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `platformServerLook`)
	if !isAuth {
		field = []string{`id`, `name`, columnsThis.ServerId}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Server().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Server().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	res = &apiPlatform.ServerListRes{
		Count: count,
	}
	list.Structs(&res.List)
	return
}
