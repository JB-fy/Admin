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
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := daoAuth.Action.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authActionLook`)
	if !isAuth {
		field = []string{`id`, `label`, daoAuth.Action.Columns().ActionName, daoAuth.Action.Columns().ActionId}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := daoAuth.Action.HandlerCtx(ctx).Filters(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order([]string{req.Sort}).JoinGroupByPrimaryKey().GetModel().Page(req.Page, req.Limit).All()
	if err != nil {
		return
	}

	res = &apiAuth.ActionListRes{Count: count, List: []apiAuth.ActionListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Action) Info(ctx context.Context, req *apiAuth.ActionInfoReq) (res *apiAuth.ActionInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoAuth.Action.ColumnArr()
	allowField = append(allowField, `id`, `label`, `sceneIdArr`)
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
	_, err = service.AuthAction().CheckAuth(ctx, `authActionLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoAuth.Action.HandlerCtx(ctx).Filters(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiAuth.ActionInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Action) Create(ctx context.Context, req *apiAuth.ActionCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authActionCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AuthAction().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Action) Update(ctx context.Context, req *apiAuth.ActionUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authActionUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AuthAction().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Action) Delete(ctx context.Context, req *apiAuth.ActionDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authActionDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AuthAction().Delete(ctx, filter)
	return
}
