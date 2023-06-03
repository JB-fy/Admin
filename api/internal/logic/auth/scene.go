package logic

import (
	dao "api/internal/model/dao/auth"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type sScene struct{}

func NewScene() *sScene {
	return &sScene{}
}

func init() {
	service.RegisterScene(NewScene())
}

func (logic *sScene) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoScene := dao.Scene
	joinCodeArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		count, err = model.Handler(daoScene.ParseGroup([]string{"id"}, &joinCodeArr)).Distinct().Count(daoScene.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

func (logic *sScene) List(ctx context.Context, filter map[string]interface{}, field []string, order [2]string, offset int, limit int) (list gdb.Result, err error) {
	daoScene := dao.Scene
	joinCodeArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoScene.ParseField(field, &joinCodeArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinCodeArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoScene.ParseOrder([][2]string{order}, &joinCodeArr))
	}
	if len(joinCodeArr) > 0 {
		model = model.Handler(daoScene.ParseGroup([]string{"id"}, &joinCodeArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}
