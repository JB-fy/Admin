package model

import (
	"context"
)

type TokenFunc func(ctx context.Context, config map[string]any) Token

type Token interface {
	Create(ctx context.Context, tokenInfo TokenInfo) (token string, err error) // 生成
	Parse(ctx context.Context, token string) (tokenInfo TokenInfo, err error)  // 解析
	GetExpireTime() (expireTime int64)                                         // 获取生成token时设置的过期时间，多少秒后token失效
}

type TokenInfo struct {
	LoginId string `json:"login_id"`
	IP      string `json:"ip"` // 验证IP时用
}
