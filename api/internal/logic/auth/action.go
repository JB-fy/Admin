package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type sAction struct{}

func NewAction() *sAction {
	return &sAction{}
}

func init() {
	service.RegisterAction(NewAction())
}

// 总数
func (logicAction *sAction) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoAction := daoAuth.Action
	joinCodeArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		count, err = model.Handler(daoAction.ParseGroup([]string{"id"}, &joinCodeArr)).Distinct().Count(daoAction.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicAction *sAction) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoAction := daoAuth.Action
	joinCodeArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoAction.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicAction *sAction) Info(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string) (info gdb.Record, err error) {
	daoAction := daoAuth.Action
	joinCodeArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoAction.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicAction *sAction) Create(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
	daoAction := daoAuth.Action
	model := daoAction.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoAction.ParseInsert(data))
	}
	if len(data) == 1 {
		id, err = model.InsertAndGetId()
		return
	}
	result, err := model.Insert()
	if err != nil {
		return
	}
	id, err = result.RowsAffected()
	return
}

// 更新
func (logicAction *sAction) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoAction := daoAuth.Action
	joinCodeArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoAction.ParseUpdate(data))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinCodeArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	result, err := model.Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicAction *sAction) Delete(ctx context.Context, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoAction := daoAuth.Action
	joinCodeArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinCodeArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	result, err := model.Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}
