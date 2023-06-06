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

// 获取签名
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

		token, err := service.Upload().Sign(r.Context(), sceneCode, param.Type)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"token": token}, 0)
	}
}

// 回调
func (c *Upload) Notify(r *ghttp.Request) {

}
