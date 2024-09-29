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
	GetExpireTime() (expireTime int64)                    // 获取生成token时设置的过期时间，多少秒后token失效
}

func NewToken(ctx context.Context, config map[string]any) Token {
	switch tokenType, _ := config[`token_type`].(string); tokenType {
	// case `tokenOfJwt`:
	default:
		return NewTokenOfJwt(ctx, config)
	}
}
