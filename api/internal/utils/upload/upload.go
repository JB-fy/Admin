package upload

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Upload interface {
	Upload(ctx context.Context) (uploadInfo map[string]interface{}, err error)
	Sign(ctx context.Context, uploadFileType string) (signInfo map[string]interface{}, err error)
	Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error)
	Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error)
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
