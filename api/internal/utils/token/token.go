package token

import (
	"context"
)

type TokenInfo struct {
	LoginId string `json:"login_id"`
}

type Token interface {
	Create(tokenInfo TokenInfo) (token string, err error) // 生成
	Parse(token string) (tokenInfo TokenInfo, err error)  // 解析
	GetExpireTime() (expireTime int64)                    // 获取过期时间
}

func NewToken(ctx context.Context, config map[string]any) Token {
	switch tokenType, _ := config[`tokenType`].(string); tokenType {
	// case `tokenOfJwt`:
	default:
		return NewTokenOfJwt(ctx, config)
	}
}
