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
	IHttp interface {
		Count(ctx context.Context, filter map[string]interface{}) (count int, err error)
		List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error)
	}
)

var (
	localHttp IHttp
)

func Http() IHttp {
	if localHttp == nil {
		panic("implement not found for interface IHttp, forgot register?")
	}
	return localHttp
}

func RegisterHttp(i IHttp) {
	localHttp = i
}
