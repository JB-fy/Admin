package upload

import (
	daoPlatform "api/internal/dao/platform"
	"context"
)

type UploadOption struct {
	Dir        string //上传的文件目录
	Expire     int64  //签名有效时间戳。单位：秒
	ExpireTime int64  //签名有效时间。单位：秒
	MinSize    int    //限制上传的文件大小。单位：字节
	MaxSize    int    //限制上传的文件大小。单位：字节。本地上传（local.go）需要同时设置配置文件api/manifest/config/config.yaml中的server.clientMaxBodySize字段
}

type Upload interface {
	Upload() (uploadInfo map[string]interface{}, err error)                // 本地上传
	Sign(option UploadOption) (signInfo map[string]interface{}, err error) // 获取签名（H5直传用）
	Config(option UploadOption) (config map[string]interface{}, err error) // 获取配置信息（APP直传前调用）
	Sts(option UploadOption) (stsInfo map[string]interface{}, err error)   // 获取Sts Token（APP直传用）
	Notify() (notifyInfo map[string]interface{}, err error)                // 回调
}

func NewUpload(ctx context.Context) Upload {
	platformConfigColumns := daoPlatform.Config.Columns()
	uploadType, _ := daoPlatform.Config.ParseDbCtx(ctx).Where(platformConfigColumns.ConfigKey, `uploadType`).Value(platformConfigColumns.ConfigValue)
	switch uploadType.String() {
	case `aliyunOss`:
		config, _ := daoPlatform.Config.Get(ctx, []string{`aliyunOssHost`, `aliyunOssBucket`, `aliyunOssAccessKeyId`, `aliyunOssAccessKeySecret`, `aliyunOssCallbackUrl`, `aliyunOssRoleArn`, `aliyunOssEndpoint`})
		return NewAliyunOss(ctx, config)
	// case `local`:
	default:
		config, _ := daoPlatform.Config.Get(ctx, []string{`localUploadUrl`, `localUploadSignKey`, `localUploadFileSaveDir`, `localUploadFileUrlPrefix`})
		return NewLocal(ctx, config)
	}
}
