package upload

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Upload interface {
	Upload() (uploadInfo map[string]interface{}, err error)                  // 本地上传
	Sign(uploadFileType string) (signInfo map[string]interface{}, err error) // 获取签名（H5直传用）
	Config(uploadFileType string) (config map[string]interface{}, err error) // 获取配置信息（APP直传前调用，后期也可用在其它地方）
	Sts(uploadFileType string) (stsInfo map[string]interface{}, err error)   // 获取Sts Token（APP直传用）
	Notify() (notifyInfo map[string]interface{}, err error)                  // 回调
}

func NewUpload(ctx context.Context) Upload {
	platformConfigColumns := daoPlatform.Config.Columns()
	uploadType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `uploadType`).Value(platformConfigColumns.ConfigValue)
	switch uploadType.String() {
	case `aliyunOss`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssRoleArn`, `aliyunOssCallbackUrl`})
		return NewAliyunOss(ctx, config)
	case `local`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileSaveDir`, `localUploadFileUrlPrefix`})
		return NewLocal(ctx, config)
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileSaveDir`, `localUploadFileUrlPrefix`})
		return NewLocal(ctx, config)
	}
}
