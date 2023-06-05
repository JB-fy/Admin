package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type sRole struct{}

func NewRole() *sRole {
	return &sRole{}
}

func init() {
	service.RegisterRole(NewRole())
}

// 总数
func (logicRole *sRole) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoRole := daoAuth.Role
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		count, err = model.Handler(daoRole.ParseGroup([]string{"id"}, &joinCodeArr)).Distinct().Count(daoRole.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicRole *sRole) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoRole := daoAuth.Role
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRole.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder(order, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoRole.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicRole *sRole) Info(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string) (info gdb.Record, err error) {
	daoRole := daoAuth.Role
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRole.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder(order, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoRole.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicRole *sRole) Create(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
	daoRole := daoAuth.Role
	model := daoRole.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoRole.ParseInsert(data))
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
func (logicRole *sRole) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoRole := daoAuth.Role
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoRole.ParseUpdate(data))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder(order, &joinCodeArr))
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
func (logicRole *sRole) Delete(ctx context.Context, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoRole := daoAuth.Role
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder(order, &joinCodeArr))
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
