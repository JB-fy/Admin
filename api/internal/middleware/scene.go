package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Scene(r *ghttp.Request) {

	/* sceneCode := "platformAdmin";
	sceneInfo := getConfig("inDb.authScene." . $sceneCode);
	if (empty($sceneInfo)) {
		throwFailJson(39999999);
	}
	if ($sceneInfo->isStop) {
		throwFailJson(39999998);
	}
	$logicAuthScene = getContainer()->get(\App\Module\Logic\Auth\Scene::class);
	$logicAuthScene->setCurrentSceneInfo($sceneInfo); */

	//dao.Request.Ctx(r.GetCtx()).Data(data).Insert()
	r.Middleware.Next()
}
