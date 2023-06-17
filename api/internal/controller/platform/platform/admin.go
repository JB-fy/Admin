package controller

import (
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/packed"
	"api/internal/service"

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
func (controllerThis *Admin) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiPlatform.AdminListReq
	err := r.Parse(&param)
	if err != nil {
		packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
		return
	}
	filter := gconv.Map(param.Filter)
	order := [][2]string{{`id`, `DESC`}}
	if param.Sort.Key != `` {
		order[0][0] = param.Sort.Key
	}
	if param.Sort.Order != `` {
		order[0][1] = param.Sort.Order
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	limit := 10
	if param.Limit != nil {
		limit = *param.Limit
	}
	/**--------参数处理 结束--------**/

	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.GetCtx(), `platformAdminLook`)
		allowField := []string{`id`, `name`, `adminId`, `nickname`}
		if isAuth {
			allowField = daoPlatform.Admin.ColumnArr()
			allowField = append(allowField, `id`, `name`)
			allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.Admin().Count(r.GetCtx(), filter)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		list, err := service.Admin().List(r.GetCtx(), filter, field, order, param.Page, limit)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{`count`: count, `list`: list}, 0)
	}
}

// 详情
func (controllerThis *Admin) Info(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminInfoReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
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
			packed.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		info, err := service.Admin().Info(r.GetCtx(), filter, field)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{`info`: info}, 0)
	}
}

// 创建
func (controllerThis *Admin) Create(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminCreateReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminCreate`)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		id, err := service.Admin().Create(r.GetCtx(), data)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{`id`: id}, 0)
	}
}

// 更新
func (controllerThis *Admin) Update(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminUpdateReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		delete(data, `idArr`)
		if len(data) == 0 {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, ``))
			return
		}
		filter := map[string]interface{}{`id`: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------不能修改平台超级管理员 开始--------**/
		if garray.NewIntArrayFrom(gconv.SliceInt(filter[`id`])).Contains(g.Cfg().MustGet(r.GetCtx(), `superPlatformAdminId`).Int()) {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 39990004, ``))
			return
		}
		/**--------不能修改平台超级管理员 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminUpdate`)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Admin().Update(r.GetCtx(), data, filter)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 删除
func (controllerThis *Admin) Delete(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminDeleteReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		filter := map[string]interface{}{`id`: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------不能删除平台超级管理员 开始--------**/
		if garray.NewIntArrayFrom(gconv.SliceInt(filter[`id`])).Contains(g.Cfg().MustGet(r.GetCtx(), `superPlatformAdminId`).Int()) {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 39990005, ``))
			return
		}
		/**--------不能删除平台超级管理员 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformAdminDelete`)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Admin().Delete(r.GetCtx(), filter)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}
