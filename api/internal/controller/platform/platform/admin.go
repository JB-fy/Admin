package controller

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// 列表
func (controllerThis *Admin) List(ctx context.Context, req *apiPlatform.AdminListReq) (res *apiPlatform.AdminListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := daoPlatform.Admin.ColumnArr()
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `platformAdminLook`)
	if !isAuth {
		field = []string{`id`, `label`, daoPlatform.Admin.Columns().Phone, daoPlatform.Admin.Columns().Account, daoPlatform.Admin.Columns().AdminId}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := daoPlatform.Admin.HandlerCtx(ctx).Filters(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Fields(field).Order(req.Sort).JoinGroupByPrimaryKey().GetModel().Page(req.Page, req.Limit).All()
	if err != nil {
		return
	}

	res = &apiPlatform.AdminListRes{Count: count, List: []apiPlatform.AdminListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Admin) Info(ctx context.Context, req *apiPlatform.AdminInfoReq) (res *apiPlatform.AdminInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoPlatform.Admin.ColumnArr()
	allowField = append(allowField, `id`, `label`, `roleIdArr`)
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
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoPlatform.Admin.HandlerCtx(ctx).Filters(filter).Fields(field).JoinGroupByPrimaryKey().GetModel().One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiPlatform.AdminInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Admin) Create(ctx context.Context, req *apiPlatform.AdminCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.PlatformAdmin().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Admin) Update(ctx context.Context, req *apiPlatform.AdminUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	if garray.NewFrom(gconv.SliceAny(req.IdArr)).Contains(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint()) { //不能修改平台超级管理员
		err = utils.NewErrorCode(ctx, 30000000, ``)
		return
	}

	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.PlatformAdmin().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Admin) Delete(ctx context.Context, req *apiPlatform.AdminDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	if garray.NewFrom(gconv.SliceAny(req.IdArr)).Contains(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint()) { //不能删除平台超级管理员
		err = utils.NewErrorCode(ctx, 30000001, ``)
		return
	}

	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformAdminDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.PlatformAdmin().Delete(ctx, filter)
	return
}
