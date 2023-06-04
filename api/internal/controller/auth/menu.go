package controller

import (
	apiAuth "api/api/auth"
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表
func (c *Menu) List(r *ghttp.Request) {
	var param *apiAuth.MenuListReq
	err := r.Parse(&param)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	filter := gconv.Map(param.Filter) //条件过滤
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

	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	switch sceneCode {
	case "platformAdmin":
		//isAuth, _ := $this->checkAuth(__FUNCTION__, $sceneCode, false);
		isAuth := true
		/**--------参数处理 开始--------**/
		allowField := []string{"menuId", "menuName", "id"}
		if isAuth {
			allowField = daoAuth.Menu.ColumnArr()
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
		/**--------参数处理 结束--------**/
		count, err := service.Menu().Count(r.Context(), filter)
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		list, err := service.Menu().List(r.Context(), filter, field, order, int((param.Page-1)*param.Limit), int(param.Limit))
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"count": count, "list": list}, 0, "")
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
func (c *Menu) Info(r *ghttp.Request) {
	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	switch sceneCode {
	case "platformAdmin":
		var param *apiAuth.MenuInfoReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}

		//isAuth, err := $this->checkAuth(__FUNCTION__, $sceneCode, false);
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 开始--------**/
		allowField := daoAuth.Menu.ColumnArr()
		allowField = append(allowField, "id")
		//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{"password"})).Slice() //移除敏感字段
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}
		/**--------参数处理 结束--------**/
		info, err := service.Menu().Info(r.Context(), map[string]interface{}{"id": param.Id}, field, [][2]string{})
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"info": info}, 0, "")
	}
}

// 创建
func (c *Menu) Create(r *ghttp.Request) {
	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	switch sceneCode {
	case "platformAdmin":
		var param *apiAuth.MenuCreateReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}

		//isAuth, err := $this->checkAuth(__FUNCTION__, $sceneCode, false);
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 开始--------**/
		insert := gconv.Map(param) //条件过滤
		/**--------参数处理 结束--------**/
		_, err = service.Menu().Create(r.Context(), []map[string]interface{}{insert})
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0, "")
	}
}
