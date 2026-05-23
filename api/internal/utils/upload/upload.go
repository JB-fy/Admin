package upload

import (
	"api/internal/utils/upload/aliyun_oss"
	"api/internal/utils/upload/local"
	"api/internal/utils/upload/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/sync/singleflight"
)

var (
	uploadMap     sync.Map
	uploadSfg     singleflight.Group
	uploadFuncMap = map[uint]model.UploadFunc{
		0: local.NewUpload,
		1: aliyun_oss.NewUpload,
	}
	uploadTypeDef uint = 0
)

func NewUpload(ctx context.Context, uploadType uint, config map[string]any) (obj model.Upload) {
	if _, ok := uploadFuncMap[uploadType]; !ok {
		uploadType = uploadTypeDef
	}
	key := gconv.String(uploadType) + gmd5.MustEncrypt(config)
	objTmp, ok := uploadMap.Load(key)
	if !ok {
		objTmp, _, _ = uploadSfg.Do(key, func() (obj any, err error) {
			obj = uploadFuncMap[uploadType](ctx, config)
			uploadMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Upload)
	return
}
