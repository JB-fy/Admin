package auth

import (
	"api/api"
	apiAuth "api/api/org/auth"
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	get_or_set_ctx "api/internal/utils/get-or-set-ctx"
	"context"
	"slices"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Role struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewRole() *Role {
	field := slices.Clone(append(daoAuth.Role.ColumnArr(), `id`, `label`))
	appendFieldOfList := []string{ /* daoAuth.Scene.Columns().SceneName */ }
	appendFieldOfInfo := []string{`action_id_arr`, `menu_id_arr`}
	return &Role{
		defaultFieldOfList: slices.Clone(append(field, appendFieldOfList...)),
		defaultFieldOfInfo: slices.Clone(append(field, appendFieldOfInfo...)),
		allowField:         slices.Clone(append(field, gset.NewStrSetFrom(slices.Concat(appendFieldOfList, appendFieldOfInfo)).Slice()...)),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Role) List(ctx context.Context, req *apiAuth.RoleListReq) (res *apiAuth.RoleListRes, err error) {
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

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	filter[daoAuth.Role.Columns().RelId] = loginInfo[daoOrg.Admin.Columns().OrgId]
	sceneInfo := get_or_set_ctx.GetCtxSceneInfo(ctx)
	filter[daoAuth.Role.Columns().SceneId] = sceneInfo[daoAuth.Scene.Columns().SceneId]
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authRoleRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoAuth.Role.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiAuth.RoleListRes{Count: count, List: []apiAuth.RoleInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Role) Info(ctx context.Context, req *apiAuth.RoleInfoReq) (res *apiAuth.RoleInfoRes, err error) {
	/**--------参数处理 开始--------**/
	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfInfo
	}
	filter := map[string]any{`id`: req.Id}

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	filter[daoAuth.Role.Columns().RelId] = loginInfo[daoOrg.Admin.Columns().OrgId]
	sceneInfo := get_or_set_ctx.GetCtxSceneInfo(ctx)
	filter[daoAuth.Role.Columns().SceneId] = sceneInfo[daoAuth.Scene.Columns().SceneId]
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authRoleRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoAuth.Role.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
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
	data := gconv.Map(req.RoleCreateData, gconv.MapOption{Deep: true, OmitEmpty: true})

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	data[daoAuth.Role.Columns().RelId] = loginInfo[daoOrg.Admin.Columns().OrgId]
	sceneInfo := get_or_set_ctx.GetCtxSceneInfo(ctx)
	data[daoAuth.Role.Columns().SceneId] = sceneInfo[daoAuth.Scene.Columns().SceneId]
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
	filter := gconv.Map(req.RoleUpdateDeleteFilter, gconv.MapOption{Deep: true, OmitEmpty: true})
	data := gconv.Map(req.RoleUpdateData, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	filter[daoAuth.Role.Columns().RelId] = loginInfo[daoOrg.Admin.Columns().OrgId]
	sceneInfo := get_or_set_ctx.GetCtxSceneInfo(ctx)
	filter[daoAuth.Role.Columns().SceneId] = sceneInfo[daoAuth.Scene.Columns().SceneId]
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
	filter := gconv.Map(req.RoleUpdateDeleteFilter, gconv.MapOption{Deep: true, OmitEmpty: true})

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	filter[daoAuth.Role.Columns().RelId] = loginInfo[daoOrg.Admin.Columns().OrgId]
	sceneInfo := get_or_set_ctx.GetCtxSceneInfo(ctx)
	filter[daoAuth.Role.Columns().SceneId] = sceneInfo[daoAuth.Scene.Columns().SceneId]
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
