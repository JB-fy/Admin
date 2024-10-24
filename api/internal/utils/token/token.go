package token

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type TokenInfo struct {
	LoginId string `json:"login_id"`
	IP      string `json:"ip"` // 要做验证IP时才用
}

type Token interface {
	Create(tokenInfo TokenInfo) (token string, err error) // 生成
	Parse(token string) (tokenInfo TokenInfo, err error)  // 解析
	GetExpireTime() (expireTime int64)                    // 获取生成token时设置的过期时间，多少秒后token失效
}

func NewToken(ctx context.Context, config map[string]any) Token {
	switch gconv.Uint(config[`token_type`]) {
	// case 0:	//JWT
	default:
		return NewTokenOfJwt(ctx, config)
	}
}
