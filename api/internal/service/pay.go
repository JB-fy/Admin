// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IPayChannel interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id any, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
	IPay interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id any, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
	IPayScene interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id any, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
)

var (
	localPayChannel IPayChannel
	localPay        IPay
	localPayScene   IPayScene
)

func PayChannel() IPayChannel {
	if localPayChannel == nil {
		panic("implement not found for interface IPayChannel, forgot register?")
	}
	return localPayChannel
}

func RegisterPayChannel(i IPayChannel) {
	localPayChannel = i
}

func Pay() IPay {
	if localPay == nil {
		panic("implement not found for interface IPay, forgot register?")
	}
	return localPay
}

func RegisterPay(i IPay) {
	localPay = i
}

func PayScene() IPayScene {
	if localPayScene == nil {
		panic("implement not found for interface IPayScene, forgot register?")
	}
	return localPayScene
}

func RegisterPayScene(i IPayScene) {
	localPayScene = i
}
