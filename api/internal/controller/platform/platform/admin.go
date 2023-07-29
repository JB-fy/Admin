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
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoPlatform.Admin.Columns()
	allowField := daoPlatform.Admin.ColumnArr()
	allowField = append(allowField, `id`)
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `platformAdminLook`)
	if !isAuth {
		field = []string{`id`, columnsThis.AdminId, columnsThis.Nickname}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Admin().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Admin().List(ctx, filter, field, order, page, limit)
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
func (controllerThis *Admin) Info(ctx context.Context, req *apiPlatform.AdminInfoReq) (res *apiPlatform.AdminInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoPlatform.Admin.ColumnArr()
	allowField = append(allowField, `id`, `roleIdArr`)
	columnsThis := daoPlatform.Admin.Columns()
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password})).Slice() //移除敏感字段
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
	_, err = service.Action().CheckAuth(ctx, `platformAdminLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Admin().Info(ctx, filter, field)
	if err != nil {
		return
	}
	utils.HttpWriteJson(ctx, map[string]interface{}{
		`info`: info,
	}, 0, ``)
	return
}

// 新增
func (controllerThis *Admin) Create(ctx context.Context, req *apiPlatform.AdminCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformAdminCreate`)
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
func (controllerThis *Admin) Update(ctx context.Context, req *apiPlatform.AdminUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------不能修改平台超级管理员 开始--------**/
	if garray.NewIntArrayFrom(gconv.SliceInt(filter[`id`])).Contains(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Int()) {
		err = utils.NewErrorCode(ctx, 30000000, ``)
		return
	}
	/**--------不能修改平台超级管理员 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformAdminUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = service.Admin().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Admin) Delete(ctx context.Context, req *apiPlatform.AdminDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------不能删除平台超级管理员 开始--------**/
	if garray.NewIntArrayFrom(gconv.SliceInt(filter[`id`])).Contains(g.Cfg().MustGet(ctx, `superPlatformAdminId`).Int()) {
		err = utils.NewErrorCode(ctx, 30000001, ``)
		return
	}
	/**--------不能删除平台超级管理员 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformAdminDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = service.Admin().Delete(ctx, filter)
	return
}
