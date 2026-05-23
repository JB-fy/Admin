package vod

import (
	"api/internal/utils/vod/aliyun"
	"api/internal/utils/vod/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	vodMap     sync.Map
	vodSfg     singleflight.Group
	vodFuncMap = map[string]model.VodFunc{
		`vod_of_aliyun`: aliyun.NewVod,
	}
	vodTypeDef = `vod_of_aliyun`
)

func NewVod(ctx context.Context, vodType string, config map[string]any) (obj model.Vod) {
	if _, ok := vodFuncMap[vodType]; !ok {
		vodType = vodTypeDef
	}
	key := vodType + gmd5.MustEncrypt(config)
	objTmp, ok := vodMap.Load(key)
	if !ok {
		objTmp, _, _ = vodSfg.Do(key, func() (obj any, err error) {
			obj = vodFuncMap[vodType](ctx, config)
			vodMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Vod)
	return
}
