package controller

import (
	"api/api"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 获取签名
func (c *Upload) Sign(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	//case `platform`:
	default:
		/**--------参数处理 开始--------**/
		var param *api.UploadSignReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		/**--------参数处理 结束--------**/
		option := utils.AliyunOssSignOption{
			CallbackUrl: ``,
			ExpireTime:  15 * 60,
			Dir:         fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`)),
			MinSize:     0,
			MaxSize:     100 * 1024 * 1024,
		}
		switch param.UploadType {
		default:
			if g.Cfg().MustGet(r.GetCtx(), `uploadCallbackEnable`).Bool() {
				option.CallbackUrl = gstr.Replace(r.GetUrl(), r.URL.Path, `/upload/notify`, 1)
			}
		}

		config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssHost`, `aliyunOssBucket`})
		upload := utils.NewAliyunOss(r.GetCtx(), config)
		signInfo, _ := upload.CreateSign(option)
		utils.HttpSuccessJson(r, signInfo, 0)
	}
}

// 回调
func (c *Upload) Notify(r *ghttp.Request) {
	filename := r.Get(`filename`).String()

	config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssHost`, `aliyunOssBucket`})
	upload := utils.NewAliyunOss(r.GetCtx(), config)
	err := upload.Notify(r)
	if err != nil {
		utils.HttpFailJson(r, err)
		return
	}

	resData := map[string]interface{}{
		`url`: upload.GetBucketHost() + `/` + filename,
	}
	utils.HttpSuccessJson(r, resData, 0)
}
