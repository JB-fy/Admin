package my

import (
	"api/api"
	apiMy "api/api/platform/my"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
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

	loginInfo := utils.GetCtxLoginInfo(ctx)
	var isGetPrivacy bool
	var privacyInfo gdb.Record
	initPrivacyInfo := func() {
		if !isGetPrivacy {
			isGetPrivacy = true
			privacyInfo, _ = daoPlatform.AdminPrivacy.CtxDaoModel(ctx).FilterPri(loginInfo[`login_id`]).Fields(daoPlatform.AdminPrivacy.Columns().Password, daoPlatform.AdminPrivacy.Columns().Salt).One()
		}
	}
	for k, v := range data {
		switch k {
		case `password_to_check`:
			delete(data, k)
			initPrivacyInfo()
			if privacyInfo[daoPlatform.AdminPrivacy.Columns().Password].String() == `` {
				err = utils.NewErrorCode(ctx, 39990004, ``)
				return
			}
			if gmd5.MustEncrypt(gconv.String(v)+privacyInfo[daoPlatform.AdminPrivacy.Columns().Salt].String()) != privacyInfo[daoPlatform.AdminPrivacy.Columns().Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
		case `sms_code_to_password`:
			delete(data, k)
			phone := loginInfo[daoPlatform.Admin.Columns().Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39991003, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), phone, 3) //场景：3密码修改(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `sms_code_to_bind_phone`:
			delete(data, k)
			if req.Phone == nil {
				continue
			}
			code, _ := cache.Code.Get(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), *req.Phone, 4) //场景：4绑定(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `sms_code_to_unbing_phone`:
			delete(data, k)
			phone := loginInfo[daoPlatform.Admin.Columns().Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39991003, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), phone, 5) //场景：5解绑(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			data[daoPlatform.Admin.Columns().Phone] = nil
		case `email_code_to_password`:
			delete(data, k)
			email := loginInfo[daoPlatform.Admin.Columns().Email].String()
			if email == `` {
				err = utils.NewErrorCode(ctx, 39991013, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), email, 13) //场景：13密码修改(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `email_code_to_bind_email`:
			delete(data, k)
			if req.Email == nil {
				continue
			}
			code, _ := cache.Code.Get(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), *req.Email, 14) //场景：14绑定(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `email_code_to_unbing_email`:
			delete(data, k)
			email := loginInfo[daoPlatform.Admin.Columns().Email].String()
			if email == `` {
				err = utils.NewErrorCode(ctx, 39991013, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), email, 15) //场景：15解绑(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			data[daoPlatform.Admin.Columns().Email] = nil
		}
	}
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	filter := map[string]any{`id`: loginInfo[`login_id`]}
	/**--------参数处理 结束--------**/

	_, err = service.PlatformAdmin().Update(ctx, filter, data)
	return
}
