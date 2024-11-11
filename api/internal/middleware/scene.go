package middleware

import (
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {
	pathArr := strings.Split(r.URL.Path, `/`)
	sceneId := pathArr[1]
	if sceneId == `` {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
		return
	}
	value, _ := cache.NewDbData(r.GetCtx(), &daoAuth.Scene, sceneId).GetOrSet(daoAuth.Scene.Columns().SceneId, daoAuth.Scene.Columns().SceneConfig, daoAuth.Scene.Columns().IsStop)
	if value == `` {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999998, ``))
		return
	}
	var sceneInfo gdb.Record
	gjson.New(value).Scan(&sceneInfo)
	if sceneInfo[daoAuth.Scene.Columns().IsStop].Uint() == 1 {
		r.SetError(utils.NewErrorCode(r.GetCtx(), 39999997, ``))
		return
	}
	utils.SetCtxSceneInfo(r, sceneInfo)

	r.Middleware.Next()
}
