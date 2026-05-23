package token

import (
	"api/internal/utils/token/jwt"
	"api/internal/utils/token/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/sync/singleflight"
)

var (
	tokenMap     sync.Map
	tokenSfg     singleflight.Group
	tokenFuncMap = map[uint]model.TokenFunc{
		0: jwt.NewToken,
	}
	tokenTypeDef uint = 0
)

func NewToken(ctx context.Context, tokenType uint, config map[string]any) (obj model.Token) {
	if _, ok := tokenFuncMap[tokenType]; !ok {
		tokenType = tokenTypeDef
	}
	key := gconv.String(tokenType) + gmd5.MustEncrypt(config)
	objTmp, ok := tokenMap.Load(key)
	if !ok {
		objTmp, _, _ = tokenSfg.Do(key, func() (obj any, err error) {
			obj = tokenFuncMap[tokenType](ctx, config)
			tokenMap.Store(key, obj)
			return
		})
	}
	obj = objTmp.(model.Token)
	return
}
