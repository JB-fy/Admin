package platform

import (
	apiPlatform "api/api/app/platform"
	"api/internal/consts"
	daoConfig "api/internal/dao/config"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiPlatform.ConfigGetReq) (res *apiPlatform.ConfigGetRes, err error) {
	config, err := daoConfig.Config.GetPluck(ctx, consts.SCENE_ID_PLATFORM, 0, *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	gconv.Struct(config.Map(), &res.Config)
	return
}
