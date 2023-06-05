package controller

import (
	"api/api"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取登录加密字符串(前端登录操作用于加密密码后提交)
func (c *Login) EncryptStr(r *ghttp.Request) {
	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.LoginEncryptReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		encryptStr, err := service.Login().EncryptStr(r.Context(), sceneCode, param.Account)
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"encryptStr": encryptStr}, 0, "")
	}
}

// 登录
func (c *Login) Login(r *ghttp.Request) {
	sceneCode := r.GetCtxVar("sceneInfo").Val().(gdb.Record)["sceneCode"].String()
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.LoginLoginReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		token, err := service.Login().Login(r.Context(), sceneCode, param.Account, param.Password)
		if err != nil {
			utils.HttpFailJson(r, 99999999, "", map[string]interface{}{})
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"token": token}, 0, "")
	}
}
