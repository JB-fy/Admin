package controller

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiPlatform.ConfigGetReq) (res *apiPlatform.ConfigGetRes, err error) {
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformConfigLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	config, err := daoPlatform.Config.Get(ctx, *req.ConfigKeyArr)
	if err != nil {
		return
	}

	utils.HttpWriteJson(ctx, map[string]interface{}{
		`config`: config,
	}, 0, ``)
	return
}

// 保存
func (controllerThis *Config) Save(ctx context.Context, req *apiPlatform.ConfigSaveReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	config := gconv.Map(req)
	if len(config) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `platformConfigSave`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = daoPlatform.Config.Save(ctx, config)
	return
}
