package controller

import (
	"api/api"
	apiAuth "api/api/platform/auth"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Action struct{}

func NewAction() *Action {
	return &Action{}
}

// 列表
func (controllerThis *Action) List(ctx context.Context, req *apiAuth.ActionListReq) (res *apiAuth.ActionListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoAuth.Action.Columns()
	allowField := daoAuth.Action.ColumnArr()
	allowField = append(allowField, `id`, `name`)
	// allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `authActionLook`)
	if !isAuth {
		field = []string{`id`, `name`, columnsThis.ActionId, columnsThis.ActionName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Action().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Action().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, map[string]interface{}{
		`count`: count,
		`list`:  list,
	}, 0, ``)
	return
}

// 详情
func (controllerThis *Action) Info(ctx context.Context, req *apiAuth.ActionInfoReq) (res *apiAuth.ActionInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoAuth.Action.ColumnArr()
	allowField = append(allowField, `id`, `name`, `sceneIdArr`)
	//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{`id`: req.Id}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authActionLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Action().Info(ctx, filter, field)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, map[string]interface{}{
		`info`: info,
	}, 0, ``)
	return
}

// 新增
func (controllerThis *Action) Create(ctx context.Context, req *apiAuth.ActionCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authActionCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Action().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Action) Update(ctx context.Context, req *apiAuth.ActionUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authActionUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = service.Action().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Action) Delete(ctx context.Context, req *apiAuth.ActionDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authActionDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = service.Action().Delete(ctx, filter)
	return
}
