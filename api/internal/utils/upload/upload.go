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
}

func NewUpload(ctx context.Context) Upload {
	platformConfigColumns := daoPlatform.Config.Columns()
	uploadType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `uploadType`).Value(platformConfigColumns.ConfigValue)
	switch uploadType.String() {
	case `local`:
		return &Local{}
	case `aliyunOss`:
		return &AliyunOss{}
	default:
		return &Local{}
	}
}
