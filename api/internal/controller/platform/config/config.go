package config

import (
	"api/api"
	apiConfig "api/api/platform/config"
	"api/internal/consts"
	daoConfig "api/internal/dao/config"
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
func (controllerThis *Config) Get(ctx context.Context, req *apiConfig.ConfigGetReq) (res *apiConfig.ConfigGetRes, err error) {
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `configRead`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for _, configKey := range *req.ConfigKeyArr {
			switch configKey {
			case `hot_search`, `user_agreement`, `privacy_agreement`:
				actionIdSet.Add(`configCommonRead`)
			case `role_id_arr_of_platform_def`:
				actionIdSet.Add(`configPlatformRead`)
			case `role_id_arr_of_org_def`:
				actionIdSet.Add(`configOrgRead`)
			case `sms_type`, `sms_of_aliyun`:
				actionIdSet.Add(`configSmsRead`)
			case `email_code`, `email_type`, `email_of_common`:
				actionIdSet.Add(`configEmailRead`)
			case `id_card_type`, `id_card_of_aliyun`:
				actionIdSet.Add(`configIdCardRead`)
			case `one_click_of_wx`, `one_click_of_yidun`:
				actionIdSet.Add(`configOneClickRead`)
			case `push_type`, `push_of_tx`:
				actionIdSet.Add(`configPushRead`)
			case `vod_type`, `vod_of_aliyun`:
				actionIdSet.Add(`configVodRead`)
			case `wx_gzh`:
				actionIdSet.Add(`configWxRead`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	config, err := daoConfig.Config.GetPluck(ctx, consts.SCENE_ID_PLATFORM, 0, *req.ConfigKeyArr...)
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `configSave`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for configKey := range config {
			switch configKey {
			case `hot_search`, `user_agreement`, `privacy_agreement`:
				actionIdSet.Add(`configCommonSave`)
			case `role_id_arr_of_platform_def`:
				actionIdSet.Add(`configPlatformSave`)
			case `role_id_arr_of_org_def`:
				actionIdSet.Add(`configOrgSave`)
			case `sms_type`, `sms_of_aliyun`:
				actionIdSet.Add(`configSmsSave`)
			case `email_code`, `email_type`, `email_of_common`:
				actionIdSet.Add(`configEmailSave`)
			case `id_card_type`, `id_card_of_aliyun`:
				actionIdSet.Add(`configIdCardSave`)
			case `one_click_of_wx`, `one_click_of_yidun`:
				actionIdSet.Add(`configOneClickSave`)
			case `push_type`, `push_of_tx`:
				actionIdSet.Add(`configPushSave`)
			case `vod_type`, `vod_of_aliyun`:
				actionIdSet.Add(`configVodSave`)
			case `wx_gzh`:
				actionIdSet.Add(`configWxSave`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	err = daoConfig.Config.Save(ctx, consts.SCENE_ID_PLATFORM, 0, config)
	return
}
