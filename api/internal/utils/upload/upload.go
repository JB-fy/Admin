package upload

import "context"

type Upload interface {
	// UploadFile(file *multipart.FileHeader) (string, string, error)
	// DeleteFile(key string) error
	Sign(ctx context.Context, uploadFileType string) (signInfo map[string]interface{}, err error)
	Sts(ctx context.Context, uploadFileType string) (stsInfo map[string]interface{}, err error)
	Notify(ctx context.Context) (notifyInfo map[string]interface{}, err error)
}

func NewUpload(uploadType string) Upload {
	switch uploadType {
	case `local`:
		return &Local{}
	case `aliyun_oss`:
		return &AliyunOss{}
	default:
		return &Local{}
	}
}
