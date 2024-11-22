package my

import (
	"api/api"
	apiMy "api/api/org/my"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

type Profile struct{}

func NewProfile() *Profile {
	return &Profile{}
}

// 个人信息
func (controllerThis *Profile) Info(ctx context.Context, req *apiMy.ProfileInfoReq) (res *apiMy.ProfileInfoRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	res = &apiMy.ProfileInfoRes{}
	loginInfo.Struct(&res.Info)
	return
}

// 修改个人信息
func (controllerThis *Profile) Update(ctx context.Context, req *apiMy.ProfileUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	loginInfo := utils.GetCtxLoginInfo(ctx)
	for k, v := range data {
		switch k {
		case `password_to_check`:
			if gmd5.MustEncrypt(gconv.String(v)+loginInfo[daoOrg.Admin.Columns().Salt].String()) != loginInfo[daoOrg.Admin.Columns().Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
			delete(data, k)
		case `sms_code_to_bind_phone`:
			phone := gconv.String(data[`phone`])
			sceneInfo := utils.GetCtxSceneInfo(ctx)
			sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
			code, _ := cache.Code.Get(ctx, sceneId, phone, 4) //场景：4绑定(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			delete(data, k)
		case `email_code_to_bind_email`:
			email := gconv.String(data[`email`])
			sceneInfo := utils.GetCtxSceneInfo(ctx)
			sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
			code, _ := cache.Code.Get(ctx, sceneId, email, 14) //场景：14绑定(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			delete(data, k)
		}
	}

	filter := map[string]any{`id`: loginInfo[`login_id`]}
	/**--------参数处理 结束--------**/

	_, err = service.OrgAdmin().Update(ctx, filter, data)
	return
}
