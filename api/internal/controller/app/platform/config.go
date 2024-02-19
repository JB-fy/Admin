package controller

import (
	apiPlatform "api/api/app/platform"
	daoPlatform "api/internal/dao/platform"
	"context"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiPlatform.ConfigGetReq) (res *apiPlatform.ConfigGetRes, err error) {
	config, err := daoPlatform.Config.HandlerCtx(ctx).Filter(daoPlatform.Config.Columns().ConfigKey, *req.ConfigKeyArr).Pluck(daoPlatform.Config.Columns().ConfigKey, daoPlatform.Config.Columns().ConfigValue)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	config.Struct(&res.Config)
	return
}
