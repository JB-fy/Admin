package controller

import (
	"api/api"
	"api/internal/service"
	"api/internal/utils"
	"fmt"

	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 获取签名
func (c *Upload) Sign(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	//case "platformAdmin":
	default:
		/**--------参数处理 开始--------**/
		var param *api.UploadSignReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/
		option := utils.AliyunOssSignOption{
			CallbackUrl: "",
			ExpireTime:  15 * 60,
			Dir:         fmt.Sprintf("common/%s_%d_", gtime.Now().Format("Y/m/d/His"), grand.N(1000, 9999)),
			MinSize:     0,
			MaxSize:     100 * 1024 * 1024,
		}
		switch param.UploadType {
		default:
			if g.Cfg().MustGet(r.GetCtx(), "uploadCallbackEnable").Bool() {
				option.CallbackUrl = gstr.Replace(r.GetUrl(), r.URL.Path, "/upload/notify", 1)
			}
		}

		filter := map[string]interface{}{
			"configKey": []string{"aliyunOssAccessKeyId", "aliyunOssAccessKeySecret", "aliyunOssHost", "aliyunOssBucket"},
		}
		config, _ := service.Config().Get(r.Context(), filter)
		upload := utils.NewAliyunOss(r.GetCtx(), config)
		signInfo, _ := upload.CreateSign(option)
		utils.HttpSuccessJson(r, signInfo, 0)
	}
}

// 回调
func (c *Upload) Notify(r *ghttp.Request) {
	filter := map[string]interface{}{
		"configKey": []string{"aliyunOssAccessKeyId", "aliyunOssAccessKeySecret", "aliyunOssHost", "aliyunOssBucket"},
	}
	config, _ := service.Config().Get(r.Context(), filter)
	upload := utils.NewAliyunOss(r.GetCtx(), config)
	err := upload.Notify(r)
	if err != nil {
		utils.HttpFailJson(r, err)
		return
	}
	utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
}
