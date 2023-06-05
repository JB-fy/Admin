package logic

import (
	daoAuth "api/internal/model/dao/auth"
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

// 总数
func (logicScene *sScene) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoScene := daoAuth.Scene
	joinTableArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoScene.ParseGroup([]string{"id"}, &joinTableArr)).Distinct().Count(daoScene.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicScene *sScene) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoScene := daoAuth.Scene
	joinTableArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoScene.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoScene.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoScene.ParseGroup([]string{"id"}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicScene *sScene) Info(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string) (info gdb.Record, err error) {
	daoScene := daoAuth.Scene
	joinTableArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoScene.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoScene.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoScene.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicScene *sScene) Create(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
	daoScene := daoAuth.Scene
	model := daoScene.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoScene.ParseInsert(data))
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
func (logicScene *sScene) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoScene := daoAuth.Scene
	joinTableArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoScene.ParseUpdate(data))
	}
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoScene.ParseOrder(order, &joinTableArr))
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
func (logicScene *sScene) Delete(ctx context.Context, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoScene := daoAuth.Scene
	joinTableArr := []string{}
	model := daoScene.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoScene.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoScene.ParseOrder(order, &joinTableArr))
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
