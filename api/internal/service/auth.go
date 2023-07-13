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
	IAction interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error)
		Delete(ctx context.Context, filter map[string]interface{}) (err error)
		CheckAuth(ctx context.Context, actionCode string) (isAuth bool, err error)
	}
	IMenu interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error)
		Delete(ctx context.Context, filter map[string]interface{}) (err error)
	}
	IRole interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error)
		Delete(ctx context.Context, filter map[string]interface{}) (err error)
	}
	IScene interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error)
		Delete(ctx context.Context, filter map[string]interface{}) (err error)
	}
)

var (
	localAction IAction
	localMenu   IMenu
	localRole   IRole
	localScene  IScene
)

func Action() IAction {
	if localAction == nil {
		panic("implement not found for interface IAction, forgot register?")
	}
	return localAction
}

func RegisterAction(i IAction) {
	localAction = i
}

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}

func Scene() IScene {
	if localScene == nil {
		panic("implement not found for interface IScene, forgot register?")
	}
	return localScene
}

func RegisterScene(i IScene) {
	localScene = i
}
