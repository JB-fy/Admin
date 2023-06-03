package middleware

import (
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
)

func SceneLoginOfPlatformAdmin(r *ghttp.Request) {
	sceneCode := r.GetCtxVar("sceneInfo")
	fmt.Println(sceneCode)
	/* if !verifyToken(sceneCode) {
		return
	} */
	r.Middleware.Next()
}
