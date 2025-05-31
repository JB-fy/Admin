package org

import (
	"api/api"
	apiOrg "api/api/org/org"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(ctx context.Context, req *apiOrg.ConfigGetReq) (res *apiOrg.ConfigGetRes, err error) {
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `orgConfigRead`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for _, configKey := range *req.ConfigKeyArr {
			switch configKey {
			case `hot_search`:
				actionIdSet.Add(`orgConfigCommonRead`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	loginInfo := utils.GetCtxLoginInfo(ctx)
	config, err := daoOrg.Config.GetPluck(ctx, loginInfo[daoOrg.Admin.Columns().OrgId].String(), *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiOrg.ConfigGetRes{}
	config.Struct(&res.Config)
	return
}

// 保存
func (controllerThis *Config) Save(ctx context.Context, req *apiOrg.ConfigSaveReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	config := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(config) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `orgConfigSave`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for configKey := range config {
			switch configKey {
			case `hot_search`:
				actionIdSet.Add(`orgConfigCommonSave`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	loginInfo := utils.GetCtxLoginInfo(ctx)
	err = daoOrg.Config.Save(ctx, loginInfo[daoOrg.Admin.Columns().OrgId].String(), config)
	return
}
