package upload

import (
	"api/internal/utils/upload/aliyun_oss"
	"api/internal/utils/upload/local"
	"api/internal/utils/upload/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	uploadMap     = map[string]model.Upload{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	uploadMuMap   sync.Map
	uploadFuncMap = map[uint]model.UploadFunc{
		0: local.NewUpload,
		1: aliyun_oss.NewUpload,
	}
	uploadTypeDef uint = 0
)

func NewUpload(ctx context.Context, uploadType uint, config map[string]any) (upload model.Upload) {
	uploadKey := gconv.String(uploadType) + gmd5.MustEncrypt(config)
	ok := false
	if upload, ok = uploadMap[uploadKey]; ok { //先读一次（不加锁）
		return
	}
	muTmp, _ := uploadMuMap.LoadOrStore(uploadKey, &sync.Mutex{})
	mu := muTmp.(*sync.Mutex)
	mu.Lock()
	defer func() {
		mu.Unlock()
		uploadMuMap.Delete(uploadKey)
	}()
	if upload, ok = uploadMap[uploadKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = uploadFuncMap[uploadType]; !ok {
		uploadType = uploadTypeDef
	}
	upload = uploadFuncMap[uploadType](ctx, config)
	uploadMap[uploadKey] = upload
	return
}
