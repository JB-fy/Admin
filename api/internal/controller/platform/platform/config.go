package controller

import (
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/packed"
	"api/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

// 获取
func (controllerThis *Config) Get(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.ConfigGetReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformConfigLook`)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		config, err := daoPlatform.Config.Get(r.GetCtx(), *param.ConfigKeyArr)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{`config`: config}, 0)
	}
}

// 创建
func (controllerThis *Config) Save(r *ghttp.Request) {
	sceneCode := packed.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.ConfigSaveReq
		err := r.Parse(&param)
		if err != nil {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		config := gconv.Map(param)
		if len(config) == 0 {
			packed.HttpFailJson(r, packed.NewErrorCode(r.GetCtx(), 89999999, ``))
			return
		}
		/**--------参数处理 结束--------**/

		/**--------权限验证 开始--------**/
		_, err = service.Action().CheckAuth(r.GetCtx(), `platformConfigSave`)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		/**--------权限验证 结束--------**/

		err = daoPlatform.Config.Save(r.GetCtx(), config)
		if err != nil {
			packed.HttpFailJson(r, err)
			return
		}
		packed.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}
