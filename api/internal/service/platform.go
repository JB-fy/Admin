// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPlatformAdmin interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id any, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
)

var (
	localPlatformAdmin IPlatformAdmin
)

func PlatformAdmin() IPlatformAdmin {
	if localPlatformAdmin == nil {
		panic("implement not found for interface IPlatformAdmin, forgot register?")
	}
	return localPlatformAdmin
}

func RegisterPlatformAdmin(i IPlatformAdmin) {
	localPlatformAdmin = i
}
