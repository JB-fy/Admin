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
	IAdmin interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error)
		Delete(ctx context.Context, filter map[string]interface{}) (err error)
	}
	ICorn interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
		Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error)
		Create(ctx context.Context, data map[string]interface{}) (id int64, err error)
		Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error)
		Delete(ctx context.Context, filter map[string]interface{}) (err error)
	}
	IServer interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
	}
)

var (
	localAdmin  IAdmin
	localCorn   ICorn
	localServer IServer
)

func Admin() IAdmin {
	if localAdmin == nil {
		panic("implement not found for interface IAdmin, forgot register?")
	}
	return localAdmin
}

func RegisterAdmin(i IAdmin) {
	localAdmin = i
}

func Corn() ICorn {
	if localCorn == nil {
		panic("implement not found for interface ICorn, forgot register?")
	}
	return localCorn
}

func RegisterCorn(i ICorn) {
	localCorn = i
}

func Server() IServer {
	if localServer == nil {
		panic("implement not found for interface IServer, forgot register?")
	}
	return localServer
}

func RegisterServer(i IServer) {
	localServer = i
}
