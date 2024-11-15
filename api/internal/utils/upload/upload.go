package upload

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type UploadParam struct {
	Dir        string //上传的文件目录
	Expire     int64  //签名有效时间戳。单位：秒
	ExpireTime int64  //签名有效时间。单位：秒
	MinSize    int    //限制上传的文件大小。单位：字节
	MaxSize    int    //限制上传的文件大小。单位：字节。本地上传（uploadOfLocal.go）需要同时设置配置文件api/manifest/config/config.yaml中的server.clientMaxBodySize字段
}

type SignInfo struct {
	UploadUrl  string         //上传地址
	UploadData map[string]any //上传数据
	Host       string         //站点域名（当上传无响应数据，前端组件用于与上传目录拼接形成文件访问地址）
	Dir        string         //上传目录
	Expire     uint           //过期时间。单位：秒
	IsRes      uint           //是否有响应信息。0否 1是
}

type NotifyInfo struct {
	Url      string //地址
	Width    uint   //宽度
	Height   uint   //高度
	Size     uint   //大小。单位：比特
	MimeType string //文件类型
}

type Upload interface {
	Upload(ctx context.Context, r *ghttp.Request) (notifyInfo NotifyInfo, err error)  // 本地上传
	Sign(ctx context.Context, param UploadParam) (signInfo SignInfo, err error)       // 获取签名（H5直传用）
	Config(ctx context.Context, param UploadParam) (config map[string]any, err error) // 获取配置信息（APP直传前调用）
	Sts(ctx context.Context, param UploadParam) (stsInfo map[string]any, err error)   // 获取Sts Token（APP直传用）
	Notify(ctx context.Context, r *ghttp.Request) (notifyInfo NotifyInfo, err error)  // 回调
}

var (
	uploadMap = map[string]Upload{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	uploadMu  sync.Mutex
)

func NewUpload(config map[string]any) (upload Upload) {
	uploadKey := gmd5.MustEncrypt(config)

	ok := false
	if upload, ok = uploadMap[uploadKey]; ok { //先读一次（不加锁）
		return
	}
	uploadMu.Lock()
	defer uploadMu.Unlock()
	if upload, ok = uploadMap[uploadKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}

	switch gconv.Uint(config[`uploadType`]) {
	case 1: //阿里云OSS
		upload = NewUploadOfAliyunOss(config)
	// case 0: //本地
	default:
		upload = NewUploadOfLocal(config)
	}
	uploadMap[uploadKey] = upload
	return
}
