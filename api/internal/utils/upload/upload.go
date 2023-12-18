package upload

import (
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/os/gtime"
)

type UploadParam struct {
	Dir        string //上传的文件目录
	Expire     int64  //签名有效时间戳。单位：秒
	ExpireTime int64  //签名有效时间。单位：秒
	MinSize    int    //限制上传的文件大小。单位：字节
	MaxSize    int    //限制上传的文件大小。单位：字节。本地上传（local.go）需要同时设置配置文件api/manifest/config/config.yaml中的server.clientMaxBodySize字段
}

type Upload interface {
	Upload() (uploadInfo map[string]interface{}, err error)              // 本地上传
	Sign(param UploadParam) (signInfo map[string]interface{}, err error) // 获取签名（H5直传用）
	Config(param UploadParam) (config map[string]interface{}, err error) // 获取配置信息（APP直传前调用）
	Sts(param UploadParam) (stsInfo map[string]interface{}, err error)   // 获取Sts Token（APP直传用）
	Notify() (notifyInfo map[string]interface{}, err error)              // 回调
}

func CreateUploadParam(uploadType string) (param UploadParam) {
	param = UploadParam{
		Dir:        `common/` + gtime.Now().Format(`Ymd`) + `/`,
		Expire:     gtime.Now().Unix() + 15*60,
		ExpireTime: 15 * 60,
		MinSize:    0,
		MaxSize:    1024 * 1024 * 1024,
	}
	return
}

func NewUpload(ctx context.Context, uploadTypeTmp ...string) Upload {
	uploadType := ``
	if len(uploadTypeTmp) > 0 {
		uploadType = uploadTypeTmp[0]
	} else {
		uploadTypeVar, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(daoPlatform.Config.Columns().ConfigKey, `uploadType`).Value(daoPlatform.Config.Columns().ConfigValue)
		uploadType = uploadTypeVar.String()
	}

	switch uploadType {
	case `aliyunOss`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssCallbackUrl`, `aliyunOssEndpoint`, `aliyunOssRoleArn`})
		return NewAliyunOss(ctx, config)
	// case `local`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileSaveDir`, `localUploadFileUrlPrefix`})
		return NewLocal(ctx, config)
	}
}
