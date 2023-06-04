package logic

import (
	dao "api/internal/model/dao/auth"
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

func (logic *sRole) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoRole := dao.Role
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

func (logic *sRole) List(ctx context.Context, filter map[string]interface{}, field []string, order [2]string, offset int, limit int) (list gdb.Result, err error) {
	daoRole := dao.Role
	joinCodeArr := []string{}
	model := daoRole.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoRole.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoRole.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoRole.ParseOrder([][2]string{order}, &joinCodeArr))
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
