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

type Role struct{}

func NewRole() *Role {
	return &Role{}
}

// 列表
func (controllerThis *Role) List(ctx context.Context, req *apiAuth.RoleListReq) (res *apiAuth.RoleListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := daoAuth.Role.ColumnArr().Slice()
	allowField = append(allowField, `id`, `label`, daoAuth.Scene.Columns().SceneName, `tableName`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authRoleLook`)
	if !isAuth {
		field = []string{`id`, `label`, daoAuth.Role.Columns().RoleId, daoAuth.Role.Columns().RoleName}
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoAuth.Role.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field).HookSelect().Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiAuth.RoleListRes{Count: count, List: []apiAuth.RoleListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Role) Info(ctx context.Context, req *apiAuth.RoleInfoReq) (res *apiAuth.RoleInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoAuth.Role.ColumnArr().Slice()
	allowField = append(allowField, `id`, `label`, `menuIdArr`, `actionIdArr`)
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
	_, err = service.AuthAction().CheckAuth(ctx, `authRoleLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoAuth.Role.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiAuth.RoleInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Role) Create(ctx context.Context, req *apiAuth.RoleCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authRoleCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AuthRole().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Role) Update(ctx context.Context, req *apiAuth.RoleUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `authRoleUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AuthRole().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Role) Delete(ctx context.Context, req *apiAuth.RoleDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authRoleDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AuthRole().Delete(ctx, filter)
	return
}
