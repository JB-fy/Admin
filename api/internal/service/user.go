// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IUserUser interface {
		// 新增
		Create(ctx context.Context, data map[string]any) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]any) (row int64, err error)
	}
)

var (
	localUserUser IUserUser
)

func UserUser() IUserUser {
	if localUserUser == nil {
		panic("implement not found for interface IUserUser, forgot register?")
	}
	return localUserUser
}

func RegisterUserUser(i IUserUser) {
	localUserUser = i
}
