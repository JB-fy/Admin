package upload

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/utils"
	"api/internal/utils/upload/internal"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type AliyunOss struct{}

func (*AliyunOss) Sign(ctx context.Context, uploadFileType string) (signInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssCallbackUrl`})
	upload := internal.NewAliyunOss(ctx, config)

	bucketHost := upload.GetBucketHost()
	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	expire := time.Now().Unix() + 15*60
	signInfo = map[string]interface{}{
		`uploadUrl`: bucketHost,
		// `uploadData`:  map[string]interface{}{},
		`host`:   bucketHost,
		`dir`:    dir,
		`expire`: expire,
		`isRes`:  0,
	}

	policyBase64 := upload.CreatePolicyBase64(internal.AliyunOssSignOption{
		Expire:  expire,
		Dir:     dir,
		MinSize: 0,
		MaxSize: 1024 * 1024 * 1024,
	})
	uploadData := map[string]interface{}{
		`OSSAccessKeyId`:        upload.AccessKeyId,
		`policy`:                string(policyBase64),
		`signature`:             upload.CreateSign(policyBase64),
		`success_action_status`: `200`, //让服务端返回200,不然，默认会返回204
	}
	//是否回调
	if g.Cfg().MustGet(ctx, `upload.callbackEnable`).Bool() {
		callback := internal.AliyunOssCallback{
			Url:      utils.GetRequestUrl(ctx, 0) + `/upload/notify`,
			Body:     `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`,
			BodyType: `application/x-www-form-urlencoded`,
		}
		if utils.IsDev(ctx) {
			callback.Url = g.Cfg().MustGet(ctx, `upload.callbackUrl`).String()
		}
		uploadData[`callback`] = upload.CreateCallbackStr(callback)
		signInfo[`isRes`] = 1
	}

	signInfo[`uploadData`] = uploadData
	return
}

func (*AliyunOss) Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error) {
	config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssRoleArn`, `aliyunOssCallbackUrl`})
	dir := fmt.Sprintf(`common/%s/`, gtime.Now().Format(`Ymd`))
	option := internal.AliyunOssStsOption{
		SessionName: `oss_app_sts_token`,
		ExpireTime:  15 * 60,
		Policy:      `{"Statement": [{"Action": ["oss:PutObject","oss:ListParts","oss:AbortMultipartUpload"],"Effect": "Allow","Resource": ["acs:oss:*:*:` + gconv.String(config[`aliyunOssBucket`]) + `/` + dir + `*"]}],"Version": "1"}`,
	}

	//App端的SDK需设置一个地址来获取Sts Token，且必须按要求格式返回，该地址不验证登录token
	if g.RequestFromCtx(ctx).URL.Path == `/upload/sts` {
		upload := internal.NewAliyunOss(ctx, config)
		stsInfo, _ = upload.GetStsToken(option)
		return
	}

	stsInfo = map[string]interface{}{}
	//App端实际上传时需用到的字段，但必须验证登录token后才能拿到
	stsInfo[`endpoint`] = gconv.String(config[`aliyunOssHost`])
	stsInfo[`bucket`] = gconv.String(config[`aliyunOssBucket`])
	stsInfo[`dir`] = dir

	//是否回调
	if g.Cfg().MustGet(ctx, `upload.callbackEnable`).Bool() {
		stsInfo[`callbackUrl`] = utils.GetRequestUrl(ctx, 0) + `/upload/notify`
		stsInfo[`callbackBody`] = `filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}`
		stsInfo[`callbackBodyType`] = `application/x-www-form-urlencoded`
		if utils.IsDev(ctx) {
			stsInfo[`callbackUrl`] = g.Cfg().MustGet(ctx, `upload.callbackUrl`).String()
		}
	}
	return
}

func (*AliyunOss) Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error) {
	r := g.RequestFromCtx(ctx)
	filename := r.Get(`filename`).String()
	width := r.Get(`width`).String()
	height := r.Get(`height`).String()

	config, _ := daoPlatform.Config.Get(r.GetCtx(), []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`})
	upload := internal.NewAliyunOss(r.GetCtx(), config)
	err = upload.Notify(r)
	if err != nil {
		return
	}

	notifyInfo = map[string]interface{}{}
	notifyInfo[`url`] = upload.GetBucketHost() + `/` + filename + `?w=` + width + `&h=` + height //需要记录宽高，ios显示瀑布流必须知道宽高。直接存在query内
	return
}
