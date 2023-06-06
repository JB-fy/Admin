package controller

import (
	"api/api"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 登录
func (c *Upload) Sign(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.UploadSignReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		token, err := service.Upload().Sign(r.Context(), sceneCode, param.Account, param.Password)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"token": token}, 0)
	}
}

// 获取登录加密字符串(前端登录操作用于加密密码后提交)
func (c *Upload) Notify(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.UploadNotifyReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		encryptStr, err := service.Upload().Notify(r.Context(), sceneCode, param.Account)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"encryptStr": encryptStr}, 0)
	}
}
