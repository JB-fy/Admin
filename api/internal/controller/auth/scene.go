package controller

import (
	api "api/api/auth"
	"api/internal/service"
	"api/internal/utils"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Scene struct{}

func NewScene() *Scene {
	return &Scene{}
}

func (c *Scene) List(r *ghttp.Request) {
	var param *api.ReqSceneList
	err := r.Parse(&param)
	if err != nil {
		fmt.Println(err)
		r.Response.Writeln(err.Error())
		return
	}
	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	fmt.Println(sceneCode)
	switch sceneCode {
	case "platformAdmin":
		// $isAuth = $this->checkAuth(__FUNCTION__, $sceneCode, false);
		// /**--------参数处理 开始--------**/
		// if ($isAuth) {
		//     $allowField = $this->getAllowField(AuthScene::class);
		// } else {
		//     $allowField = ['sceneId', 'sceneName', 'id'];
		// }
		// $data['field'] = empty($data['field']) ? $allowField : array_intersect($data['field'], $allowField);
		// /**--------参数处理 结束--------**/
		filter := gconv.Map(param.Filter) //条件过滤
		count, err := service.Scene().Count(r.Context(), filter)
		list, err := service.Scene().List(r.Context(), filter, param.Field, [2]string{}, int((param.Page-1)*param.Limit), int(param.Limit))
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
