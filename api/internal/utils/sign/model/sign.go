package model

import (
	"context"
)

type SignFunc func(ctx context.Context, config map[string]any) Sign

type Sign interface {
	Create(ctx context.Context, data map[string]any) (sign string)            // 生成
	Verify(ctx context.Context, data map[string]any, sign string) (err error) // 验证
}
