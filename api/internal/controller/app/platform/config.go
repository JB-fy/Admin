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
	config, err := daoPlatform.Config.Get(ctx, *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	config.Struct(&res.Config)
	return
}
