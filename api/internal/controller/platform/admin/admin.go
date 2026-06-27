package admin

import (
	"api/api"
	apiAdmin "api/api/platform/admin"
	"api/internal/consts"
	daoAdmin "api/internal/dao/admin"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"slices"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Admin struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewAdmin() *Admin {
	field := slices.Clone(append(daoAdmin.Admin.ColumnArr(), `id`, `label`))
	appendFieldOfList := []string{daoAuth.Scene.Columns().SceneName, `rel_name`}
	appendFieldOfInfo := []string{`role_id_arr`}
	return &Admin{
		defaultFieldOfList: slices.Clone(append(field, appendFieldOfList...)),
		defaultFieldOfInfo: slices.Clone(append(field, appendFieldOfInfo...)),
		allowField:         slices.Clone(append(field, gset.NewStrSetFrom(slices.Concat(appendFieldOfList, appendFieldOfInfo)).Slice()...)),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Admin) List(ctx context.Context, req *apiAdmin.AdminListReq) (res *apiAdmin.AdminListRes, err error) {
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `adminRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoAdmin.Admin.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiAdmin.AdminListRes{Count: count, List: []apiAdmin.AdminInfo{}}
	gconv.Structs(list.List(), &res.List)
	return
}

// 详情
func (controllerThis *Admin) Info(ctx context.Context, req *apiAdmin.AdminInfoReq) (res *apiAdmin.AdminInfoRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `adminRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoAdmin.Admin.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiAdmin.AdminInfoRes{}
	gconv.Struct(info.Map(), &res.Info)
	return
}

// 新增
func (controllerThis *Admin) Create(ctx context.Context, req *apiAdmin.AdminCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req.AdminCreateData, gconv.MapOption{Deep: true, OmitEmpty: true})

	switch *req.SceneId {
	case consts.SCENE_ID_PLATFORM:
		data[daoAdmin.Admin.Columns().IsSuper] = 0 //不允许创建平台超级管理员
	case consts.SCENE_ID_ORG:
		data[daoAdmin.Admin.Columns().IsSuper] = 1 //只允许创建机构超级管理员
	default:
		err = utils.NewErrorCode(ctx, 39999998, ``)
		return
	}
	data[`rel_id_of_role`] = 0 //logic层用于验证role_id_arr是否合法
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `adminCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Admin().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Admin) Update(ctx context.Context, req *apiAdmin.AdminUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.AdminUpdateDeleteFilter, gconv.MapOption{Deep: true, OmitEmpty: true})
	data := gconv.Map(req.AdminUpdateData, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	filter[`platform_update_delete`] = `` //不允许修改平台超级管理员 或 只允许修改机构超级管理员
	data[`rel_id_of_role`] = 0            //logic层用于验证role_id_arr是否合法
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `adminUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Admin().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Admin) Delete(ctx context.Context, req *apiAdmin.AdminDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.AdminUpdateDeleteFilter, gconv.MapOption{Deep: true, OmitEmpty: true})
	filter[`platform_update_delete`] = `` //不允许删除平台超级管理员 或 只允许删除机构超级管理员
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `adminDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Admin().Delete(ctx, filter)
	return
}
