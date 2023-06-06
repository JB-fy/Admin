package middleware

import (
	dao "api/internal/model/dao/auth"
	"api/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {
	pathArr := strings.Split(r.URL.Path, "/")
	sceneCode := pathArr[1]
	if sceneCode == "" {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39999999, ""))
		return
	}
	sceneInfo, _ := dao.Scene.ParseDbCtx(r.GetCtx()).Where("sceneCode", sceneCode).One()
	if sceneInfo.IsEmpty() {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39999999, ""))
		return
	}
	if sceneInfo["isStop"].Int() > 0 {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 39999998, ""))
		return
	}

	utils.SetCtxSceneInfo(r, sceneInfo)
	r.Middleware.Next()
}
