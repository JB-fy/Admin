package controller

import (
	api "api/api/auth"
	"api/internal/service"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Scene struct{}

func NewScene() *Scene {
	return &Scene{}
}

func (c *Scene) List(r *ghttp.Request) {
	var param *api.ReqSceneList
	err := r.Parse(&param)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	fmt.Println(param)

	service.Scene().List(r.Context())
	/* r.SetError(gerror.NewCode(gcode.New(1, "aaaa", g.Map{"a": "a"})))
	r.Response.WriteJson(map[string]interface{}{
		"code": 0,
		"msg":  "成功",
		"data": map[string]interface{}{},
	}) */
}
