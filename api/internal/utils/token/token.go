package token

import (
	"context"
	"sync"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

type TokenInfo struct {
	LoginId string `json:"login_id"`
	IP      string `json:"ip"` // 要做验证IP时才用
}

type Token interface {
	Create(ctx context.Context, tokenInfo TokenInfo) (token string, err error) // 生成
	Parse(ctx context.Context, token string) (tokenInfo TokenInfo, err error)  // 解析
	GetExpireTime() (expireTime int64)                                         // 获取生成token时设置的过期时间，多少秒后token失效
}

var (
	// tokenMap sync.Map
	tokenMap = map[string]Token{} //存放不同配置实例。因初始化只有一次，故重要的是读性能，普通map比sync.Map的读性能好
	tokenMu  sync.Mutex
)

func NewToken(config map[string]any) (token Token) {
	tokenKey := gmd5.MustEncrypt(config)

	/* tokenTmp, _ := tokenMap.LoadOrStore(tokenKey, func() (token Token) {
		switch gconv.Uint(config[`token_type`]) {
		// case 0:	//JWT
		default:
			token = NewTokenOfJwt(config)
		}
		return
	})
	token = tokenTmp.(Token) */

	ok := false
	if token, ok = tokenMap[tokenKey]; ok { //先读一次（不加锁）
		return
	}
	tokenMu.Lock()
	defer tokenMu.Unlock()
	if token, ok = tokenMap[tokenKey]; ok { // 再读一次（加锁），防止重复初始化
		return
	}

	switch gconv.Uint(config[`token_type`]) {
	// case 0:	//JWT
	default:
		token = NewTokenOfJwt(config)
	}
	tokenMap[tokenKey] = token
	return

}
