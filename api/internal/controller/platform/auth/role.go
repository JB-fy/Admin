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
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoAuth.Role.Columns()
	allowField := daoAuth.Role.ColumnArr()
	allowField = append(allowField, `id`, `label`, `sceneName`, `tableName`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `authRoleLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.RoleId, columnsThis.RoleName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Role().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Role().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}

	res = &apiAuth.RoleListRes{
		Count: count,
	}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Role) Info(ctx context.Context, req *apiAuth.RoleInfoReq) (res *apiAuth.RoleInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoAuth.Role.ColumnArr()
	allowField = append(allowField, `id`, `label`, `sceneName`, `menuIdArr`, `actionIdArr`)
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
	_, err = service.Action().CheckAuth(ctx, `authRoleLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Role().Info(ctx, filter, field)
	if err != nil {
		return
	}

	res = &apiAuth.RoleInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Role) Create(ctx context.Context, req *apiAuth.RoleCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authRoleCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Role().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Role) Update(ctx context.Context, req *apiAuth.RoleUpdateReq) (res *api.CommonNoDataRes, err error) {
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
	_, err = service.Action().CheckAuth(ctx, `authRoleUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Role().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Role) Delete(ctx context.Context, req *apiAuth.RoleDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authRoleDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Role().Delete(ctx, filter)
	return
}
