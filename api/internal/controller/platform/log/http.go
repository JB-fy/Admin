package controller

import (
	apiLog "api/api/platform/log"
	daoLog "api/internal/dao/log"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Http struct{}

func NewHttp() *Http {
	return &Http{}
}

// 列表
func (controllerThis *Http) List(ctx context.Context, req *apiLog.HttpListReq) (res *apiLog.HttpListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	allowField := daoLog.Http.ColumnArr()
	allowField = append(allowField, `id`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `logHttpLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	count, err := service.Http().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Http().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, map[string]interface{}{
		`count`: count,
		`list`:  list,
	}, 0, ``)
	return
}
