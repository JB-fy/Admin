package upload

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Upload interface {
	// UploadFile(file *multipart.FileHeader) (string, string, error)
	// DeleteFile(key string) error
	Sign(ctx context.Context, uploadFileType string) (signInfo map[string]interface{}, err error)
	Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error)
	Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error)
	Upload(ctx context.Context) (uploadInfo map[string]interface{}, err error)
}

func NewUpload(ctx context.Context) Upload {
	/* config, _ := daoPlatform.Config.Get(ctx, []string{`uploadType`})
	uploadType := gconv.String(config[`uploadType`]) */
	platformConfigColumns := daoPlatform.Config.Columns()
	uploadTypeTmp, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `uploadType`).Value(platformConfigColumns.ConfigValue)
	uploadType := uploadTypeTmp.String()
	switch uploadType {
	case `aliyunOss`:
		return &AliyunOss{}
	case `local`:
		return &Local{}
	default:
		return &Local{}
	}
}
