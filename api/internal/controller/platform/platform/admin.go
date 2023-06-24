package controller

import (
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// 列表
func (controllerThis *Admin) List(ctx context.Context, req *apiPlatform.AdminListReq) (res *apiPlatform.AdminListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoPlatform.Admin.Columns()
	allowField := daoPlatform.Admin.ColumnArr()
	allowField = append(allowField, `id`, `name`)
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
		field = []string{`id`, `name`, columnsThis.AdminId, columnsThis.Nickname}
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
	res = &apiPlatform.AdminListRes{
		Count: count,
	}
	list.Structs(&res.List)
	// utils.HttpSuccessJson(g.RequestFromCtx(ctx), map[string]interface{}{`count`: count, `list`: list}, 0)
	return
}

// 详情
func (controllerThis *Admin) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminInfoReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}

		allowField := daoPlatform.Admin.ColumnArr()
		allowField = append(allowField, `id`, `name`, `roleIdArr`)
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		filter := map[string]interface{}{`id`: param.Id}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminLook`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		info, err := service.Admin().Info(r.GetCtx(), filter, field)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`info`: info}, 0)
	}
}

// 创建
func (controllerThis *Admin) Create(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminCreateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminCreate`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		id, err := service.Admin().Create(r.GetCtx(), data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`id`: id}, 0)
	}
}

// 更新
func (controllerThis *Admin) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminUpdateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		delete(data, `idArr`)
		if len(data) == 0 {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, ``))
			return
		}
		filter := map[string]interface{}{`id`: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------不能修改平台超级管理员 开始--------**/
		if garray.NewIntArrayFrom(gconv.SliceInt(filter[`id`])).Contains(g.Cfg().MustGet(r.GetCtx(), `superPlatformAdminId`).Int()) {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39990004, ``))
			return
		}
		/**--------不能修改平台超级管理员 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminUpdate`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Admin().Update(r.GetCtx(), filter, data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 删除
func (controllerThis *Admin) Delete(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminDeleteReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		filter := map[string]interface{}{`id`: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------不能删除平台超级管理员 开始--------**/
		if garray.NewIntArrayFrom(gconv.SliceInt(filter[`id`])).Contains(g.Cfg().MustGet(r.GetCtx(), `superPlatformAdminId`).Int()) {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39990005, ``))
			return
		}
		/**--------不能删除平台超级管理员 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminDelete`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Admin().Delete(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}
