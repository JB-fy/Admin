package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type sMenu struct{}

func NewMenu() *sMenu {
	return &sMenu{}
}

func init() {
	service.RegisterMenu(NewMenu())
}

// 总数
func (logic *sMenu) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoMenu := daoAuth.Menu
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		count, err = model.Handler(daoMenu.ParseGroup([]string{"id"}, &joinCodeArr)).Distinct().Count(daoMenu.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logic *sMenu) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoMenu := daoAuth.Menu
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoMenu.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logic *sMenu) Info(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string) (info gdb.Record, err error) {
	daoMenu := daoAuth.Menu
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoMenu.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoMenu.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logic *sMenu) Create(ctx context.Context, insert []map[string]interface{}) (id int64, err error) {
	daoMenu := daoAuth.Menu
	model := daoMenu.Ctx(ctx)
	if len(insert) > 0 {
		model = model.Handler(daoMenu.ParseInsert(insert))
	}
	if len(insert) == 1 {
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
func (logic *sMenu) Update(ctx context.Context, update map[string]interface{}, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoMenu := daoAuth.Menu
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(update) > 0 {
		model = model.Handler(daoMenu.ParseUpdate(update))
	}
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinCodeArr))
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
func (logic *sMenu) Delete(ctx context.Context, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoMenu := daoAuth.Menu
	joinCodeArr := []string{}
	model := daoMenu.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoMenu.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoMenu.ParseOrder(order, &joinCodeArr))
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
