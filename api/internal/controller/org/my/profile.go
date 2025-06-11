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
	loginInfo := utils.GetCtxLoginInfo(ctx)
	if loginInfo[daoOrg.Admin.Columns().IsSuper].Uint8() == 0 {
		orgId := loginInfo[daoOrg.Admin.Columns().OrgId].Uint()
		if req.Phone != nil {
			*req.Phone = daoOrg.Admin.JoinLoginName(orgId, *req.Phone)
		}
		if req.Email != nil {
			*req.Email = daoOrg.Admin.JoinLoginName(orgId, *req.Email)
		}
		if req.Account != nil {
			*req.Account = daoOrg.Admin.JoinLoginName(orgId, *req.Account)
		}
	}
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})

	var isGetPrivacy bool
	var privacyInfo gdb.Record
	initPrivacyInfo := func() {
		if !isGetPrivacy {
			isGetPrivacy = true
			privacyInfo, _ = daoOrg.Admin.CtxDaoModel(ctx).FilterPri(loginInfo[`login_id`]).Fields(daoOrg.Admin.Columns().Password, daoOrg.Admin.Columns().Salt).One()
		}
	}
	for k, v := range data {
		switch k {
		case `password_to_check`:
			delete(data, k)
			initPrivacyInfo()
			if privacyInfo[daoOrg.Admin.Columns().Password].String() == `` {
				err = utils.NewErrorCode(ctx, 39990004, ``)
				return
			}
			if gmd5.MustEncrypt(gconv.String(v)+privacyInfo[daoOrg.Admin.Columns().Salt].String()) != privacyInfo[daoOrg.Admin.Columns().Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
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
		}
	}
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	filter := map[string]any{`id`: loginInfo[`login_id`]}
	/**--------参数处理 结束--------**/

	_, err = service.OrgAdmin().Update(ctx, filter, data)
	return
}
