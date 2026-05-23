package id_card

import (
	"api/internal/utils/id-card/aliyun"
	"api/internal/utils/id-card/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"golang.org/x/sync/singleflight"
)

var (
	idCardMap     sync.Map
	idCardSfg     singleflight.Group
	idCardTypeDef = `id_card_of_aliyun`
	idCardFuncMap = map[string]model.IdCardFunc{
		`id_card_of_aliyun`: aliyun.NewIdCard,
	}
)

func NewIdCard(ctx context.Context, idCardType string, config map[string]any) (obj model.IdCard) {
	if _, ok := idCardFuncMap[idCardType]; !ok {
		idCardType = idCardTypeDef
	}
	key := idCardType + gmd5.MustEncrypt(config)
	objTmp, ok := idCardMap.Load(key)
	if !ok {
		objTmp, _, _ = idCardSfg.Do(key, func() (obj any, err error) {
			obj = idCardFuncMap[idCardType](ctx, config)
			idCardMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.IdCard)
	return

}
