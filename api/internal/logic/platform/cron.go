package logic

import (
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type sCron struct{}

func NewCron() *sCron {
	return &sCron{}
}

func init() {
	service.RegisterCron(NewCron())
}

// 总数
func (logicThis *sCron) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoPlatform.Cron
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + `.` + daoThis.PrimaryKey()).Distinct().Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	count, err = model.Count()
	return
}

// 列表
func (logicThis *sCron) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoPlatform.Cron
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
		model = model.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicThis *sCron) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoPlatform.Cron
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	info, err = model.One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}
	return
}

// 新增
func (logicThis *sCron) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoPlatform.Cron
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *sCron) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error) {
	daoThis := daoPlatform.Cron
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}
	hookData := map[string]interface{}{}

	model := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data))
	if len(hookData) > 0 {
		model = model.Hook(daoThis.HookUpdate(hookData, gconv.SliceInt(idArr)...))
	}
	_, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sCron) Delete(ctx context.Context, filter map[string]interface{}) (err error) {
	daoThis := daoPlatform.Cron
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}

	_, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	return
}
