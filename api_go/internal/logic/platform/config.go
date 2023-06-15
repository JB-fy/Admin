package logic

import (
	daoPlatform "api/internal/model/dao/platform"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sConfig struct{}

func NewConfig() *sConfig {
	return &sConfig{}
}

func init() {
	service.RegisterConfig(NewConfig())
}

// 获取
func (logicThis *sConfig) Get(ctx context.Context, filter map[string]interface{}) (config map[string]interface{}, err error) {
	daoThis := daoPlatform.Config
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Fields(`configValue`, `configKey`).All()
	config = map[string]interface{}{}
	for _, v := range result {
		config[v[`configKey`].String()] = v[`configValue`]
	}
	return
}

// 保存
func (logicThis *sConfig) Save(ctx context.Context, data map[string]interface{}) (err error) {
	daoThis := daoPlatform.Config
	for k, v := range data {
		daoThis.ParseDbCtx(ctx).Data(g.Map{`configKey`: k, `configValue`: v}).Save()
	}
	return
}
