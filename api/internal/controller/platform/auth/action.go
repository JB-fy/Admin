package auth

import (
	"api/api"
	apiAuth "api/api/platform/auth"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"slices"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Action struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewAction() *Action {
	field := append(daoAuth.Action.ColumnArr(), `id`, `label`)
	appendFieldOfList := []string{}
	appendFieldOfInfo := []string{`scene_id_arr`}
	return &Action{
		defaultFieldOfList: append(slices.Clone(field), appendFieldOfList...),
		defaultFieldOfInfo: append(slices.Clone(field), appendFieldOfInfo...),
		allowField:         append(slices.Clone(field), gset.NewStrSetFrom(slices.Concat(appendFieldOfList, appendFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Action) List(ctx context.Context, req *apiAuth.ActionListReq) (res *apiAuth.ActionListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]any{}
	}

	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfList
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authActionRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoAuth.Action.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiAuth.ActionListRes{Count: count, List: []apiAuth.ActionInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Action) Info(ctx context.Context, req *apiAuth.ActionInfoReq) (res *apiAuth.ActionInfoRes, err error) {
	/**--------参数处理 开始--------**/
	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfInfo
	}
	filter := map[string]any{`id`: req.Id}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authActionRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoAuth.Action.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
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
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`data`}})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
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
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
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
