package controller

import (
	"api/api"
	apiPlatform "api/api/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取加密盐
func (controllerThis *Login) EncryptStr(ctx context.Context, req *api.LoginEncryptReq) (res *api.LoginEncryptRes, err error) {
	encryptStr, err := service.Login().EncryptStr(ctx, `platform`, req.Account)
	if err != nil {
		return
	}
	res = &api.LoginEncryptRes{EncryptStr: encryptStr}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *api.LoginReq) (res *api.LoginRes, err error) {
	token, err := service.Login().Login(ctx, `platform`, req.Account, req.Password)
	if err != nil {
		return
	}
	res = &api.LoginRes{Token: token}
	return
}

// 登录用户详情
func (c *Login) Info(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
		utils.HttpSuccessJson(r, map[string]interface{}{`info`: loginInfo}, 0)
	}
}

// 修改个人信息
func (c *Login) Update(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------参数处理 开始--------**/
		var param *apiPlatform.AdminUpdateSelfReq
		err := r.Parse(&param)
		if err != nil {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, err.Error()))
			return
		}
		data := gconv.Map(param)
		if len(data) == 0 {
			utils.HttpFailJson(r, utils.NewErrorCode(r.GetCtx(), 89999999, ``))
			return
		}
		loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
		filter := map[string]interface{}{`id`: loginInfo[`adminId`]}
		/**--------参数处理 结束--------**/

		_, err = service.Admin().Update(r.GetCtx(), filter, data)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		utils.HttpSuccessJson(r, map[string]interface{}{}, 0)
	}
}

// 用户菜单树
func (c *Login) MenuTree(r *ghttp.Request) {
	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		loginInfo := utils.GetCtxLoginInfo(r.GetCtx())
		sceneInfo := utils.GetCtxSceneInfo(r.GetCtx())
		filter := map[string]interface{}{}
		filter[`selfMenu`] = map[string]interface{}{
			`sceneCode`: sceneCode,
			`sceneId`:   sceneInfo[`sceneId`].Int(),
			`loginId`:   loginInfo[`adminId`].Int(),
		}
		field := []string{`menuTree`, `showMenu`}

		list, err := service.Menu().List(r.GetCtx(), filter, field, []string{}, 0, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		tree := utils.Tree(list, 0, `menuId`, `pid`)
		utils.HttpSuccessJson(r, map[string]interface{}{`tree`: tree}, 0)
	}
}
