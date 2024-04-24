// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IAuthAction interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
		// 判断操作权限
		CheckAuth(ctx context.Context, actionCode string) (isAuth bool, err error)
	}
	IAuthMenu interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IAuthRole interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IAuthScene interface {
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localAuthAction IAuthAction
	localAuthMenu   IAuthMenu
	localAuthRole   IAuthRole
	localAuthScene  IAuthScene
)

func AuthAction() IAuthAction {
	if localAuthAction == nil {
		panic("implement not found for interface IAuthAction, forgot register?")
	}
	return localAuthAction
}

func RegisterAuthAction(i IAuthAction) {
	localAuthAction = i
}

func AuthMenu() IAuthMenu {
	if localAuthMenu == nil {
		panic("implement not found for interface IAuthMenu, forgot register?")
	}
	return localAuthMenu
}

func RegisterAuthMenu(i IAuthMenu) {
	localAuthMenu = i
}

func AuthRole() IAuthRole {
	if localAuthRole == nil {
		panic("implement not found for interface IAuthRole, forgot register?")
	}
	return localAuthRole
}

func RegisterAuthRole(i IAuthRole) {
	localAuthRole = i
}

func AuthScene() IAuthScene {
	if localAuthScene == nil {
		panic("implement not found for interface IAuthScene, forgot register?")
	}
	return localAuthScene
}

func RegisterAuthScene(i IAuthScene) {
	localAuthScene = i
}
