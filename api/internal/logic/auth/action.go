package logic

import (
	dao "api/internal/model/dao/auth"
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

func (logic *sAction) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoAction := dao.Action
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

func (logic *sAction) List(ctx context.Context, filter map[string]interface{}, field []string, order [2]string, offset int, limit int) (list gdb.Result, err error) {
	daoAction := dao.Action
	joinCodeArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder([][2]string{order}, &joinCodeArr))
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
