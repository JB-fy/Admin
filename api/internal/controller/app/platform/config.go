package platform

import (
	apiPlatform "api/api/app/platform"
	daoPlatform "api/internal/dao/platform"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiPlatform.ConfigGetReq) (res *apiPlatform.ConfigGetRes, err error) {
	config, err := daoPlatform.Config.GetPluck(ctx, *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	gconv.Structs(config.Map(), &res.Config)
	return
}
