package controller

import (
	apiLog "api/api/platform/log"
	daoLog "api/internal/dao/log"
	"api/internal/service"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Http struct{}

func NewHttp() *Http {
	return &Http{}
}

// 列表
func (controllerThis *Http) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiLog.HttpListReq
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
		_, err := service.Action().CheckAuth(r.GetCtx(), `logHttpLook`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		allowField := daoLog.Http.ColumnArr()
		allowField = append(allowField, `id`)
		//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.Http().Count(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		list, err := service.Http().List(r.GetCtx(), filter, field, order, param.Page, limit)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`count`: count, `list`: list}, 0)
	}
}
