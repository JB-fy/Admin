package middleware

import (
	dao "api/internal/model/dao/auth"
	"api/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {
	pathArr := strings.Split(r.URL.Path, "/")
	sceneCode := pathArr[1]
	if sceneCode == "" {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39999999,
			"msg":  "成功",
			"data": map[string]interface{}{},
		})
		return
	}
	sceneInfo, _ := dao.Scene.CommonModel(r.GetCtx(), g.Map{"sceneCode": sceneCode}, []string{}).One()
	if sceneInfo.IsEmpty() {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39999999,
			"msg":  "成功",
			"data": map[string]interface{}{},
		})
		return
	}
	if sceneInfo["isStop"].Int() > 0 {
		r.Response.WriteJson(map[string]interface{}{
			"code": 39999998,
			"msg":  "成功",
			"data": map[string]interface{}{},
		})
		return
	}

	utils.SetCtxSceneInfo(r, sceneInfo)
	r.Middleware.Next()
}
