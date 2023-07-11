package controller

import (
	"api/api"
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

// 获取签名(web端直传用)
func (controllerThis *Upload) Sign(ctx context.Context, req *api.UploadSignReq) (res *api.UploadSignRes, err error) {
	request := g.RequestFromCtx(ctx)
	option := utils.AliyunOssSignOption{
		ExpireTime: 15 * 60,
		Dir:        fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`)),
		MinSize:    0,
		MaxSize:    100 * 1024 * 1024,
	}
	config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`})
	upload := utils.NewAliyunOss(ctx, config)
	signInfo, _ := upload.CreateSign(option)

	//是否回调
	if g.Cfg().MustGet(ctx, `upload.callbackEnable`).Bool() {
		callback := utils.AliyunOssCallback{
			Url:      gstr.Replace(request.GetUrl(), request.URL.String(), `/upload/notify`, 1),
			Body:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			BodyType: `application/x-www-form-urlencoded`,
		}
		if g.Cfg().MustGet(ctx, `dev`).Bool() {
			callback.Url = g.Cfg().MustGet(ctx, `upload.callbackUrl`).String()
		}
		signInfo[`callback`] = upload.CreateCallbackStr(callback)
	}
	res = &api.UploadSignRes{}
	gconv.Struct(signInfo, &res)
	return
}

// 获取Sts Token(App端直传用)
func (controllerThis *Upload) Sts(ctx context.Context, req *api.UploadStsReq) (res *api.UploadStsRes, err error) {
	request := g.RequestFromCtx(ctx)
	config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssRoleArn`})
	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	option := utils.AliyunOssStsOption{
		SessionName: `oss_app_sts_token`,
		ExpireTime:  15 * 60,
		Policy:      `{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:` + gconv.String(config[`aliyunOssBucket`]) + `/` + dir + `*"]}],"Version": "1"}`,
	}

	//App端的SDK需设置一个地址来获取Sts Token，且必须按要求格式返回，该地址不验证权限
	if request.URL.Path == `/upload/sts` {
		upload := utils.NewAliyunOss(ctx, config)
		stsInfo, _ := upload.GetStsToken(option)
		request.Response.WriteJson(stsInfo)
		return
	}

	//App端实际上传时需用到的字段，但必须验证权限后才能拿到
	res = &api.UploadStsRes{
		Endpoint: gconv.String(config[`aliyunOssHost`]),
		Bucket:   gconv.String(config[`aliyunOssBucket`]),
		Dir:      dir,
	}

	//是否回调
	if g.Cfg().MustGet(ctx, `upload.callbackEnable`).Bool() {
		res.CallbackUrl = gstr.Replace(request.GetUrl(), request.URL.String(), `/upload/notify`, 1)
		res.CallbackBody = `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
		res.CallbackBodyType = `application/x-www-form-urlencoded`
		if g.Cfg().MustGet(ctx, `dev`).Bool() {
			res.CallbackUrl = g.Cfg().MustGet(ctx, `upload.callbackUrl`).String()
		}
	}
	return
}

// 回调
func (controllerThis *Upload) Notify(ctx context.Context, req *api.UploadNotifyReq) (res *api.UploadNotifyRes, err error) {
	r := g.RequestFromCtx(ctx)
	filename := r.Get(`filename`).String()
	width := r.Get(`width`).String()
	height := r.Get(`height`).String()

	config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssHost`, `aliyunOssBucket`})
	upload := utils.NewAliyunOss(r.GetCtx(), config)
	err = upload.Notify(r)
	if err != nil {
		return
	}

	res = &api.UploadNotifyRes{
		Url: upload.GetBucketHost() + `/` + filename + `?w=` + width + `&h=` + height, //需要记录宽高，ios显示瀑布流必须知道宽高。直接存在query内
	}
	return
}
