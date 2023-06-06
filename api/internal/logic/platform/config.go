package logic

import (
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/service"
	"context"
)

type sConfig struct{}

func NewConfig() *sConfig {
	return &sConfig{}
}

func init() {
	service.RegisterConfig(NewConfig())
}

// 创建
func (logicThis *sConfig) Create(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
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
