package platform

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
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
func (controllerThis *Config) Get(ctx context.Context, req *apiPlatform.ConfigGetReq) (res *apiPlatform.ConfigGetRes, err error) {
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `pltCfgRead`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for _, configKey := range *req.ConfigKeyArr {
			switch configKey {
			case `hot_search`, `user_agreement`, `privacy_agreement`:
				actionIdSet.Add(`pltCfgCommonRead`)
			case `role_id_arr_of_platform_def`:
				actionIdSet.Add(`pltCfgPlatformRead`)
			case `role_id_arr_of_org_def`:
				actionIdSet.Add(`pltCfgOrgRead`)
			case `sms_type`, `sms_of_aliyun`:
				actionIdSet.Add(`pltCfgSmsRead`)
			case `email_code`, `email_type`, `email_of_common`:
				actionIdSet.Add(`pltCfgEmailRead`)
			case `id_card_type`, `id_card_of_aliyun`:
				actionIdSet.Add(`pltCfgIdCardRead`)
			case `one_click_of_wx`, `one_click_of_yidun`:
				actionIdSet.Add(`pltCfgOneClickRead`)
			case `push_type`, `push_of_tx`:
				actionIdSet.Add(`pltCfgPushRead`)
			case `vod_type`, `vod_of_aliyun`:
				actionIdSet.Add(`pltCfgVodRead`)
			case `wx_gzh`:
				actionIdSet.Add(`pltCfgWxRead`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	config, err := daoPlatform.Config.GetPluck(ctx, *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	gconv.Structs(config.Map(), &res.Config)
	return
}

// 保存
func (controllerThis *Config) Save(ctx context.Context, req *apiPlatform.ConfigSaveReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	config := gconv.Map(req.ConfigSaveData, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(config) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `pltCfgSave`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for configKey := range config {
			switch configKey {
			case `hot_search`, `user_agreement`, `privacy_agreement`:
				actionIdSet.Add(`pltCfgCommonSave`)
			case `role_id_arr_of_platform_def`:
				actionIdSet.Add(`pltCfgPlatformSave`)
			case `role_id_arr_of_org_def`:
				actionIdSet.Add(`pltCfgOrgSave`)
			case `sms_type`, `sms_of_aliyun`:
				actionIdSet.Add(`pltCfgSmsSave`)
			case `email_code`, `email_type`, `email_of_common`:
				actionIdSet.Add(`pltCfgEmailSave`)
			case `id_card_type`, `id_card_of_aliyun`:
				actionIdSet.Add(`pltCfgIdCardSave`)
			case `one_click_of_wx`, `one_click_of_yidun`:
				actionIdSet.Add(`pltCfgOneClickSave`)
			case `push_type`, `push_of_tx`:
				actionIdSet.Add(`pltCfgPushSave`)
			case `vod_type`, `vod_of_aliyun`:
				actionIdSet.Add(`pltCfgVodSave`)
			case `wx_gzh`:
				actionIdSet.Add(`pltCfgWxSave`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	err = daoPlatform.Config.Save(ctx, config)
	return
}
