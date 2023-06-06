package controller

import (
	"api/api"
	"api/internal/service"
	"api/internal/utils"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取登录加密字符串(前端登录操作用于加密密码后提交)
func (c *Login) EncryptStr(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.LoginEncryptReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		encryptStr, err := service.Login().EncryptStr(r.Context(), sceneCode, param.Account)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"encryptStr": encryptStr}, 0)
	}
}

// 登录
func (c *Login) Login(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		/**--------参数处理 开始--------**/
		var param *api.LoginLoginReq
		err := r.Parse(&param)
		if err != nil {
			r.Response.Writeln(err.Error())
			return
		}
		/**--------参数处理 结束--------**/

		token, err := service.Login().Login(r.Context(), sceneCode, param.Account, param.Password)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"token": token}, 0)
	}
}

// 登录用户详情
func (c *Login) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
		utils.HttpSuccessJson(r, map[string]interface{}{"info": loginInfo}, 0)
	}
}

// 修改个人信息
func (c *Login) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		// /**--------参数验证并处理 开始--------**/
		// $data = $this->request->all();
		// $data = $this->container->get(\App\Module\Validation\Platform\Admin::class)->make($data, 'updateSelf')->validate();
		// /**--------参数验证并处理 结束--------**/

		// $loginInfo = $this->container->get(\App\Module\Logic\Login::class)->getCurrentInfo($sceneCode);
		// $this->container->get(\App\Module\Service\Platform\Admin::class)->update($data, ['id' => $loginInfo->adminId]);
	}
}

// 用户菜单树
func (c *Login) MenuTree(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case "platformAdmin":
		loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
		sceneInfo := utils.GetCtxSceneInfo(r.GetCtx())
		filter := map[string]interface{}{}
		filter["selfMenu"] = map[string]interface{}{
			"sceneCode": sceneCode,
			"sceneId":   sceneInfo["sceneId"],
			"loginId":   loginInfo["adminId"],
		}
		field := []string{"menuTree", "showMenu"}

		list, err := service.Menu().List(r.Context(), filter, field, [][2]string{}, 0, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		tree, err := service.Menu().Tree(r.Context(), list, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{"tree": tree}, 0)
	}
}
