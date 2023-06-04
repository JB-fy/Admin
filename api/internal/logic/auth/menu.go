package logic

import (
	dao "api/internal/model/dao/auth"
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

func (logic *sMenu) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoMenu := dao.Menu
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

func (logic *sMenu) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoMenu := dao.Menu
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
