package config

import (
	"api/api"
	apiConfig "api/api/org/config"
	"api/internal/consts"
	daoAdmin "api/internal/dao/admin"
	daoConfig "api/internal/dao/config"
	"api/internal/service"
	"api/internal/utils"
	"api/internal/utils/jbctx"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiConfig.ConfigGetReq) (res *apiConfig.ConfigGetRes, err error) {
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `orgCfgRead`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for _, configKey := range *req.ConfigKeyArr {
			switch configKey {
			case `hot_search`:
				actionIdSet.Add(`orgCfgCommonRead`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	loginInfo := jbctx.GetLoginInfo(ctx)
	config, err := daoConfig.Config.GetPluck(ctx, consts.SceneId(loginInfo[daoAdmin.Admin.Columns().SceneId].String()), loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiConfig.ConfigGetRes{}
	gconv.Struct(config.Map(), &res.Config)
	return
}

// 保存
func (controllerThis *Config) Save(ctx context.Context, req *apiConfig.ConfigSaveReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	config := gconv.Map(req.ConfigSaveData, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(config) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `orgCfgSave`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for configKey := range config {
			switch configKey {
			case `hot_search`:
				actionIdSet.Add(`orgCfgCommonSave`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	loginInfo := jbctx.GetLoginInfo(ctx)
	err = daoConfig.Config.Save(ctx, consts.SceneId(loginInfo[daoAdmin.Admin.Columns().SceneId].String()), loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), config)
	return
}
