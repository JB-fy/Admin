package my

import (
	"api/api"
	apiMy "api/api/org/my"
	"api/internal/cache"
	"api/internal/consts"
	daoAdmin "api/internal/dao/admin"
	"api/internal/service"
	"api/internal/utils"
	"api/internal/utils/jbctx"
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
	loginInfo := jbctx.GetLoginInfo(ctx)
	res = &apiMy.ProfileInfoRes{}
	gconv.Struct(loginInfo.Map(), &res.Info)
	return
}

// 修改个人信息
func (controllerThis *Profile) Update(ctx context.Context, req *apiMy.ProfileUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req.ProfileUpdateData, gconv.MapOption{Deep: true, OmitEmpty: true})

	loginInfo := jbctx.GetLoginInfo(ctx)
	if req.Phone != nil && *req.Phone != `` {
		data[daoAdmin.Admin.Columns().Phone] = daoAdmin.Admin.JoinLoginName(loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), loginInfo[daoAdmin.Admin.Columns().IsSuper].Uint8(), *req.Phone)
	}
	if req.Email != nil && *req.Email != `` {
		data[daoAdmin.Admin.Columns().Email] = daoAdmin.Admin.JoinLoginName(loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), loginInfo[daoAdmin.Admin.Columns().IsSuper].Uint8(), *req.Email)
	}
	if req.Account != nil && *req.Account != `` {
		data[daoAdmin.Admin.Columns().Account] = daoAdmin.Admin.JoinLoginName(loginInfo[daoAdmin.Admin.Columns().RelId].Uint(), loginInfo[daoAdmin.Admin.Columns().IsSuper].Uint8(), *req.Account)
	}
	var isGetPrivacy bool
	var privacyInfo gdb.Record
	initPrivacyInfo := func() {
		if !isGetPrivacy {
			isGetPrivacy = true
			privacyInfo, _ = daoAdmin.Privacy.CtxDaoModel(ctx).FilterPri(loginInfo[consts.CTX_LOGIN_ID_NAME]).Fields(daoAdmin.Privacy.Columns().Password, daoAdmin.Privacy.Columns().Salt).One()
		}
	}
	for k, v := range data {
		switch k {
		case `password_to_check`:
			delete(data, k)
			initPrivacyInfo()
			if privacyInfo[daoAdmin.Privacy.Columns().Password].String() == `` {
				err = utils.NewErrorCode(ctx, 39990004, ``)
				return
			}
			if gmd5.MustEncrypt(gconv.String(v)+privacyInfo[daoAdmin.Privacy.Columns().Salt].String()) != privacyInfo[daoAdmin.Privacy.Columns().Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
		case `sms_code_to_password`:
			delete(data, k)
			phone := loginInfo[daoAdmin.Admin.Columns().Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39991003, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), phone, loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8(), 3) //场景：3密码修改(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `sms_code_to_bind_phone`:
			delete(data, k)
			if req.Phone == nil {
				continue
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), *req.Phone, loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8(), 4) //场景：4绑定(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `sms_code_to_unbing_phone`:
			delete(data, k)
			phone := loginInfo[daoAdmin.Admin.Columns().Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39991003, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), phone, loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8(), 5) //场景：5解绑(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			data[daoAdmin.Admin.Columns().Phone] = nil
		case `email_code_to_password`:
			delete(data, k)
			email := loginInfo[daoAdmin.Admin.Columns().Email].String()
			if email == `` {
				err = utils.NewErrorCode(ctx, 39991013, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), email, loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8(), 13) //场景：13密码修改(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `email_code_to_bind_email`:
			delete(data, k)
			if req.Email == nil {
				continue
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), *req.Email, loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8(), 14) //场景：14绑定(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `email_code_to_unbing_email`:
			delete(data, k)
			email := loginInfo[daoAdmin.Admin.Columns().Email].String()
			if email == `` {
				err = utils.NewErrorCode(ctx, 39991013, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetSceneId(ctx).String(), email, loginInfo[daoAdmin.Admin.Columns().AdminType].Uint8(), 15) //场景：15解绑(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			data[daoAdmin.Admin.Columns().Email] = nil
		}
	}
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	filter := map[string]any{`id`: loginInfo[consts.CTX_LOGIN_ID_NAME]}
	/**--------参数处理 结束--------**/

	_, err = service.Admin().Update(ctx, filter, data)
	return
}
