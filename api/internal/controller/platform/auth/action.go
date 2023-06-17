package controller

import (
	apiAuth "api/api/platform/auth"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Action struct{}

func NewAction() *Action {
	return &Action{}
}

// 列表
func (controllerThis *Action) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiAuth.ActionListReq
	err := r.Parse(&param)
	if err != nil {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
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

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.GetCtx(), `authActionLook`)
		allowField := []string{`id`, `name`, `actionId`, `actionName`}
		if isAuth {
			allowField = daoAuth.Action.ColumnArr()
			allowField = append(allowField, `id`, `name`)
			//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.Action().Count(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		list, err := service.Action().List(r.GetCtx(), filter, field, order, param.Page, limit)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`count`: count, `list`: list}, 0)
	}
}

// 详情
func (controllerThis *Action) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiAuth.ActionInfoReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}

		allowField := daoAuth.Action.ColumnArr()
		allowField = append(allowField, `id`, `name`, `sceneIdArr`)
		//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
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
		_, err = service.Action().CheckAuth(r.GetCtx(), `authActionLook`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		info, err := service.Action().Info(r.GetCtx(), filter, field)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`info`: info}, 0)
	}
}

// 创建
func (controllerThis *Action) Create(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiAuth.ActionCreateReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `authActionCreate`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		id, err := service.Action().Create(r.GetCtx(), data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`id`: id}, 0)
	}
}

// 更新
func (controllerThis *Action) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiAuth.ActionUpdateReq
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

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `authActionUpdate`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Action().Update(r.GetCtx(), data, filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 删除
func (controllerThis *Action) Delete(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiAuth.ActionDeleteReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		filter := map[string]interface{}{`id`: param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `authActionDelete`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Action().Delete(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}
