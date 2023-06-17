package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/packed"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
)

type sScene struct{}

func NewScene() *sScene {
	return &sScene{}
}

func init() {
	service.RegisterScene(NewScene())
}

// 总数
func (logicThis *sScene) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Scene
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicThis *sScene) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Scene
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

// 详情
func (logicThis *sScene) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Scene
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr))
	}
	info, err = model.One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = packed.NewErrorCode(ctx, 29999999, ``)
		return
	}
	return
}

// 创建
func (logicThis *sScene) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Scene
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
		if len(match) > 0 {
			err = packed.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
			return
		}
		return
	}
	return
}

// 更新
func (logicThis *sScene) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Scene
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
		if len(match) > 0 {
			err = packed.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
			return
		}
		return
	}
	row, _ = result.RowsAffected()
	if row == 0 {
		err = packed.NewErrorCode(ctx, 99999999, ``)
		return
	}
	return
}

// 删除
func (logicThis *sScene) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Scene
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, _ = result.RowsAffected()
	if row == 0 {
		err = packed.NewErrorCode(ctx, 99999999, ``)
		return
	}
	return
}
