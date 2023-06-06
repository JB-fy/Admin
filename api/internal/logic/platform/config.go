package logic

import (
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
)

type sConfig struct{}

func NewConfig() *sConfig {
	return &sConfig{}
}

func init() {
	service.RegisterConfig(NewConfig())
}

// 获取
func (logicThis *sConfig) Get(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoPlatform.Config
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 保存
func (logicThis *sConfig) Save(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
	daoThis := daoPlatform.Config
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseInsert(data))
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
