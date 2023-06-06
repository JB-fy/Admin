package controller

import (
	apiAuth "api/api/auth"
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Scene struct{}

func NewScene() *Scene {
	return &Scene{}
}

// 列表
func (controllerThis *Scene) List(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiAuth.SceneListReq
	err := r.Parse(&param)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	filter := gconv.Map(param.Filter)
	order := [][2]string{{"id", "DESC"}}
	if param.Sort.Key != "" {
		order[0][0] = param.Sort.Key
	}
	if param.Sort.Order != "" {
		order[0][1] = param.Sort.Order
	}
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.Limit <= 0 {
		param.Limit = 10
	}
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.Context(), "authSceneLook", sceneCode)
		allowField := []string{"sceneId", "sceneName", "id"}
		if isAuth {
			allowField = daoAuth.Scene.ColumnArr()
			allowField = append(allowField, "id")
			//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------权限验证 结束--------**/

		count, err := service.Scene().Count(r.Context(), filter)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		list, err := service.Scene().List(r.Context(), filter, field, order, int((param.Page-1)*param.Limit), int(param.Limit))
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"count": count, "list": list}, 0)
		/* r.SetError(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
		r.Response.WriteJson(map[string]interface{}{
			"code": 0,
			"msg":  g.I18n().Tf(r.GetCtx(), "0"),
			"data": map[string]interface{}{
				"list": list,
			},
		}) */
	}
}

// 详情
func (controllerThis *Scene) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.SceneInfoReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}

		allowField := daoAuth.Scene.ColumnArr()
		allowField = append(allowField, "id")
		//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		filter := map[string]interface{}{"id": param.Id}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authSceneLook", sceneCode)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------权限验证 结束--------**/

		info, err := service.Scene().Info(r.Context(), filter, field, [][2]string{})
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"info": info}, 0)
	}
}

// 创建
func (controllerThis *Scene) Create(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.SceneCreateReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		data := gconv.Map(param)
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authSceneCreate", sceneCode)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Scene().Create(r.Context(), []map[string]interface{}{data})
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 更新
func (controllerThis *Scene) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.SceneUpdateReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		data := gconv.Map(param)
		delete(data, "idArr")
		if len(data) == 0 {
			utils.HttpFailJson(r, err)
			return
		}
		filter := map[string]interface{}{"id": param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authSceneUpdate", sceneCode)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Scene().Update(r.Context(), data, filter, [][2]string{}, 0, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 删除
func (controllerThis *Scene) Delete(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *apiAuth.SceneDeleteReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		filter := map[string]interface{}{"id": param.IdArr}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.Context(), "authSceneDelete", sceneCode)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------权限验证 结束--------**/

		_, err = service.Scene().Delete(r.Context(), filter, [][2]string{}, 0, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}
