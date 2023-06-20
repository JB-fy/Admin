package controller

import (
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// 列表
func (controllerThis *Server) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiPlatform.ServerListReq
	err := r.Parse(&param)
	if err != nil {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
		return
	}
	filter := gconv.Map(param.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{param.Sort}
	page := param.Page
	limit := param.Limit
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------权限验证 开始--------**/
		_, err := service.Action().CheckAuth(r.GetCtx(), `platformServerLook`)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		allowField := daoPlatform.Server.ColumnArr()
		allowField = append(allowField, `id`, `name`)
		//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.Server().Count(r.GetCtx(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		list, err := service.Server().List(r.GetCtx(), filter, field, order, page, limit)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{`count`: count, `list`: list}, 0)
	}
}
