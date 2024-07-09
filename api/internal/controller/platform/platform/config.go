package controller

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
		actionCodeSet := gset.NewStrSet()
		for _, configKey := range *req.ConfigKeyArr {
			switch configKey {
			case `hotSearch`, `userAgreement`, `privacyAgreement`:
				actionCodeSet.Add(`platformConfigCommonRead`)
			case `smsType`, `smsOfAliyun`:
				actionCodeSet.Add(`platformConfigSmsRead`)
			case `emailCodeSubject`, `emailCodeTemplate`, `emailType`, `emailOfCommon`:
				actionCodeSet.Add(`platformConfigEmailRead`)
			case `idCardType`, `idCardOfAliyun`:
				actionCodeSet.Add(`platformConfigIdCardRead`)
			case `oneClickOfWxHost`, `oneClickOfWxAppId`, `oneClickOfWxSecret`, `oneClickOfYidunSecretId`, `oneClickOfYidunSecretKey`, `oneClickOfYidunBusinessId`:
				actionCodeSet.Add(`platformConfigOneClickRead`)
			case `pushType`, `pushOfTxHost`, `pushOfTxAndroidAccessID`, `pushOfTxAndroidSecretKey`, `pushOfTxIosAccessID`, `pushOfTxIosSecretKey`, `pushOfTxMacOSAccessID`, `pushOfTxMacOSSecretKey`:
				actionCodeSet.Add(`platformConfigPushRead`)
			case `vodType`, `vodOfAliyunAccessKeyId`, `vodOfAliyunAccessKeySecret`, `vodOfAliyunEndpoint`, `vodOfAliyunRoleArn`:
				actionCodeSet.Add(`platformConfigVodRead`)
			case `wxGzhHost`, `wxGzhAppId`, `wxGzhSecret`, `wxGzhToken`, `wxGzhEncodingAESKey`:
				actionCodeSet.Add(`platformConfigWxRead`)
			}
			_, err = service.AuthAction().CheckAuth(ctx, actionCodeSet.Slice()...)
			if err != nil {
				return
			}
		}
		return
	}
	/**--------权限验证 结束--------**/

	config, err := daoPlatform.Config.Get(ctx, *req.ConfigKeyArr)
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
		actionCodeSet := gset.NewStrSet()
		for configKey := range config {
			switch configKey {
			case `hotSearch`, `userAgreement`, `privacyAgreement`:
				actionCodeSet.Add(`platformConfigCommonSave`)
			case `smsType`, `smsOfAliyun`:
				actionCodeSet.Add(`platformConfigSmsSave`)
			case `emailCodeSubject`, `emailCodeTemplate`, `emailType`, `emailOfCommon`:
				actionCodeSet.Add(`platformConfigEmailSave`)
			case `idCardType`, `idCardOfAliyun`:
				actionCodeSet.Add(`platformConfigIdCardSave`)
			case `oneClickOfWxHost`, `oneClickOfWxAppId`, `oneClickOfWxSecret`, `oneClickOfYidunSecretId`, `oneClickOfYidunSecretKey`, `oneClickOfYidunBusinessId`:
				actionCodeSet.Add(`platformConfigOneClickSave`)
			case `pushType`, `pushOfTxHost`, `pushOfTxAndroidAccessID`, `pushOfTxAndroidSecretKey`, `pushOfTxIosAccessID`, `pushOfTxIosSecretKey`, `pushOfTxMacOSAccessID`, `pushOfTxMacOSSecretKey`:
				actionCodeSet.Add(`platformConfigPushSave`)
			case `vodType`, `vodOfAliyunAccessKeyId`, `vodOfAliyunAccessKeySecret`, `vodOfAliyunEndpoint`, `vodOfAliyunRoleArn`:
				actionCodeSet.Add(`platformConfigVodSave`)
			case `wxGzhHost`, `wxGzhAppId`, `wxGzhSecret`, `wxGzhToken`, `wxGzhEncodingAESKey`:
				actionCodeSet.Add(`platformConfigWxSave`)
			}
			_, err = service.AuthAction().CheckAuth(ctx, actionCodeSet.Slice()...)
			if err != nil {
				return
			}
		}
		return
	}
	/**--------权限验证 结束--------**/

	err = daoPlatform.Config.Save(ctx, config)
	return
}
