package token

import (
	"api/internal/utils/token/jwt"
	"api/internal/utils/token/model"
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	// tokenMap     sync.Map
	tokenMap     = map[string]model.Token{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	tokenMu      sync.Mutex
	tokenTypeDef uint = 0
	tokenFuncMap      = map[uint]model.TokenFunc{
		0: jwt.NewToken,
	}
)

func NewToken(ctx context.Context, tokenType uint, config map[string]any) (token model.Token) {
	tokenKey := gconv.String(tokenType) + gmd5.MustEncrypt(config)
	/* token, _ = tokenMap.LoadOrStore(tokenKey, func() Token {
		if _, ok := tokenFuncMap[tokenType]; !ok {
			tokenType = tokenTypeDef
		}
		return tokenFuncMap[tokenType](ctx, config)
	}) */
	ok := false
	if token, ok = tokenMap[tokenKey]; ok { //先读一次（不加锁）
		return
	}
	tokenMu.Lock()
	defer tokenMu.Unlock()
	if token, ok = tokenMap[tokenKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}
	if _, ok = tokenFuncMap[tokenType]; !ok {
		tokenType = tokenTypeDef
	}
	token = tokenFuncMap[tokenType](ctx, config)
	tokenMap[tokenKey] = token
	return

}
