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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `platformConfigRead`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for _, configKey := range *req.ConfigKeyArr {
			switch configKey {
			case `hot_search`, `user_agreement`, `privacy_agreement`:
				actionIdSet.Add(`platformConfigCommonRead`)
			case `smsType`, `smsOfAliyun`:
				actionIdSet.Add(`platformConfigSmsRead`)
			case `email_code`, `email_type`, `email_of_common`:
				actionIdSet.Add(`platformConfigEmailRead`)
			case `idCardType`, `idCardOfAliyun`:
				actionIdSet.Add(`platformConfigIdCardRead`)
			case `oneClickOfWx`, `oneClickOfYidun`:
				actionIdSet.Add(`platformConfigOneClickRead`)
			case `pushType`, `pushOfTx`:
				actionIdSet.Add(`platformConfigPushRead`)
			case `vodType`, `vodOfAliyun`:
				actionIdSet.Add(`platformConfigVodRead`)
			case `wxGzh`:
				actionIdSet.Add(`platformConfigWxRead`)
			}
		}
		_, err = service.AuthAction().CheckAuth(ctx, actionIdSet.Slice()...)
		if err != nil {
			return
		}
	}
	/**--------权限验证 结束--------**/

	config, err := daoPlatform.Config.Get(ctx, *req.ConfigKeyArr...)
	if err != nil {
		return
	}

	res = &apiPlatform.ConfigGetRes{}
	config.Struct(&res.Config)
	return
}

// 保存
func (controllerThis *Config) Save(ctx context.Context, req *apiPlatform.ConfigSaveReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	config := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(config) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `platformConfigSave`)
	if !isAuth {
		actionIdSet := gset.NewStrSet()
		for configKey := range config {
			switch configKey {
			case `hot_search`, `user_agreement`, `privacy_agreement`:
				actionIdSet.Add(`platformConfigCommonSave`)
			case `smsType`, `smsOfAliyun`:
				actionIdSet.Add(`platformConfigSmsSave`)
			case `email_code`, `email_type`, `email_of_common`:
				actionIdSet.Add(`platformConfigEmailSave`)
			case `idCardType`, `idCardOfAliyun`:
				actionIdSet.Add(`platformConfigIdCardSave`)
			case `oneClickOfWx`, `oneClickOfYidun`:
				actionIdSet.Add(`platformConfigOneClickSave`)
			case `pushType`, `pushOfTx`:
				actionIdSet.Add(`platformConfigPushSave`)
			case `vodType`, `vodOfAliyun`:
				actionIdSet.Add(`platformConfigVodSave`)
			case `wxGzh`:
				actionIdSet.Add(`platformConfigWxSave`)
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
