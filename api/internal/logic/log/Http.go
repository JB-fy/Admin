package logic

import (
	daoLog "api/internal/dao/log"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type sHttp struct{}

func NewHttp() *sHttp {
	return &sHttp{}
}

func init() {
	service.RegisterHttp(NewHttp())
}

// 总数
func (logicThis *sHttp) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoLog.Http
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr)).Distinct().Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	count, err = model.Count()
	return
}

// 列表
func (logicThis *sHttp) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoLog.Http
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(field) > 0 {
		model = model.Handler(daoThis.ParseField(field, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoThis.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}
