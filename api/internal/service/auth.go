// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IAuthAction interface {
		// 总数
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		// 列表
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		// 详情
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
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
		// 总数
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		// 列表
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		// 详情
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IAuthRole interface {
		// 总数
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		// 列表
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		// 详情
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
	IAuthScene interface {
		// 总数
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		// 列表
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		// 详情
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		// 新增
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		// 修改
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error)
		// 删除
		Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error)
	}
)

var (
	localAuthScene  IAuthScene
	localAuthAction IAuthAction
	localAuthMenu   IAuthMenu
	localAuthRole   IAuthRole
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
