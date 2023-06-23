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
	"github.com/gogf/gf/v2/util/gconv"
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
			ExpireTime: 15 * 60,
			Dir:        fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`)),
			MinSize:    0,
			MaxSize:    100 * 1024 * 1024,
		}
		if g.Cfg().MustGet(r.GetCtx(), `uploadCallbackEnable`).Bool() {
			option.Callback = utils.AliyunOssCallback{
				CallbackUrl:      gstr.Replace(r.GetUrl(), r.URL.Path, `/upload/notify`, 1),
				CallbackBody:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
				CallbackBodyType: `application/x-www-form-urlencoded`,
			}
		}

		config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`})
		upload := utils.NewAliyunOss(r.GetCtx(), config)
		signInfo, _ := upload.CreateSign(option)
		utils.HttpSuccessJson(r, signInfo, 0)
	}
}

// 获取STS
func (c *Upload) Sts(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *api.UploadSignReq
	err := r.Parse(&param)
	if err != nil {
		utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
		return
	}
	/**--------参数处理 结束--------**/

	config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssRoleArn`})
	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	option := utils.AliyunOssStsOption{
		SessionName: `oss_app_sts_token`,
		ExpireTime:  15 * 60,
		Policy:      `{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:` + gconv.String(config[`aliyunOssBucket`]) + `/` + dir + `*"]}],"Version": "1"}`,
	}
	if g.Cfg().MustGet(r.GetCtx(), `uploadCallbackEnable`).Bool() {
		option.Callback = utils.AliyunOssCallback{
			CallbackUrl:      gstr.Replace(r.GetUrl(), r.URL.Path, `/upload/notify`, 1),
			CallbackBody:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			CallbackBodyType: `application/x-www-form-urlencoded`,
		}
	}

	upload := utils.NewAliyunOss(r.GetCtx(), config)
	stsInfo, _ := upload.GetStsToken(option)
	stsInfo[`endpoint`] = config[`aliyunOssHost`]
	stsInfo[`bucket`] = config[`aliyunOssBucket`]
	stsInfo[`callback`] = option.Callback
	stsInfo[`dir`] = dir
	r.Response.WriteJsonExit(stsInfo)
}

// 回调
func (c *Upload) Notify(r *ghttp.Request) {
	filename := r.Get(`filename`).String()
	width := r.Get(`width`).String()
	height := r.Get(`height`).String()

	config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssHost`, `aliyunOssBucket`})
	upload := utils.NewAliyunOss(r.GetCtx(), config)
	err := upload.Notify(r)
	if err != nil {
		utils.HttpFailJson(r, err)
		return
	}

	resData := map[string]interface{}{
		`url`: upload.GetBucketHost() + `/` + filename + `?w=` + width + `&h=` + height, //需要记录宽高，ios显示瀑布流必须知道宽高。直接存在query内
	}
	utils.HttpSuccessJson(r, resData, 0)
}
