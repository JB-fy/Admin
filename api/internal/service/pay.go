// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPay interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
)

var (
	localPay IPay
)

func Pay() IPay {
	if localPay == nil {
		panic("implement not found for interface IPay, forgot register?")
	}
	return localPay
}

func RegisterPay(i IPay) {
	localPay = i
}
