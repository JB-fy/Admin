// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IApp interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id any, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
)

var (
	localApp IApp
)

func App() IApp {
	if localApp == nil {
		panic("implement not found for interface IApp, forgot register?")
	}
	return localApp
}

func RegisterApp(i IApp) {
	localApp = i
}
