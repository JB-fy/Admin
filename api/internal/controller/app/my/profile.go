package my

import (
	"api/api"
	apiMy "api/api/app/my"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUsers "api/internal/dao/users"
	"api/internal/service"
	"api/internal/utils"
	id_card "api/internal/utils/id-card"
	"api/internal/utils/jbctx"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type Profile struct{}

func NewProfile() *Profile {
	return &Profile{}
}

// 个人信息
func (controllerThis *Profile) Info(ctx context.Context, req *apiMy.ProfileInfoReq) (res *apiMy.ProfileInfoRes, err error) {
	loginInfo := jbctx.GetCtxLoginInfo(ctx)
	res = &apiMy.ProfileInfoRes{}
	loginInfo.Struct(&res.Info)
	return
}

// 修改个人信息
func (controllerThis *Profile) Update(ctx context.Context, req *apiMy.ProfileUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req.ProfileUpdateData, gconv.MapOption{Deep: true, OmitEmpty: true})

	loginInfo := jbctx.GetCtxLoginInfo(ctx)
	var isGetPrivacy bool
	var privacyInfo gdb.Record
	initPrivacyInfo := func() {
		if !isGetPrivacy {
			isGetPrivacy = true
			privacyInfo, _ = daoUsers.Privacy.CtxDaoModel(ctx).FilterPri(loginInfo[`login_id`]).Fields(daoUsers.Privacy.Columns().Password, daoUsers.Privacy.Columns().Salt).One()
		}
	}
	for k, v := range data {
		switch k {
		/* case `account`: //前端太懒，可能把个人信息全部传回来，导致account有值，故不能用required-with:Account直接验证
		if gconv.String(v) != loginInfo[daoUsers.Users.Columns().Account].String() && g.Validator().Rules(`required`).Data(req.PasswordToCheck).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89999999, ``)
			return
		} */
		case `password`:
			if g.Validator().Rules(`required`).Data(req.PasswordToCheck).Run(ctx) != nil && g.Validator().Rules(`required`).Data(req.SmsCodeToPassword).Run(ctx) != nil && g.Validator().Rules(`required`).Data(req.EmailCodeToPassword).Run(ctx) != nil {
				err = utils.NewErrorCode(ctx, 89999999, ``)
				return
			}
		case `password_to_check`:
			delete(data, k)
			initPrivacyInfo()
			if privacyInfo[daoUsers.Privacy.Columns().Password].String() == `` {
				err = utils.NewErrorCode(ctx, 39990004, ``)
				return
			}
			if gmd5.MustEncrypt(gconv.String(v)+privacyInfo[daoUsers.Privacy.Columns().Salt].String()) != privacyInfo[daoUsers.Privacy.Columns().Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
		case `sms_code_to_password`:
			delete(data, k)
			phone := loginInfo[daoUsers.Users.Columns().Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39991003, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), phone, 3) //场景：3密码修改(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `sms_code_to_bind_phone`:
			delete(data, k)
			if req.Phone == nil {
				continue
			}
			if loginInfo[daoUsers.Users.Columns().Phone].String() != `` {
				err = utils.NewErrorCode(ctx, 39991001, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), *req.Phone, 4) //场景：4绑定(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `sms_code_to_unbing_phone`:
			delete(data, k)
			phone := loginInfo[daoUsers.Users.Columns().Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39991003, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), phone, 5) //场景：5解绑(手机)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			data[daoUsers.Users.Columns().Phone] = nil
		case `email_code_to_password`:
			delete(data, k)
			email := loginInfo[daoUsers.Users.Columns().Email].String()
			if email == `` {
				err = utils.NewErrorCode(ctx, 39991013, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), email, 13) //场景：13密码修改(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `email_code_to_bind_email`:
			delete(data, k)
			if req.Email == nil {
				continue
			}
			if loginInfo[daoUsers.Users.Columns().Email].String() != `` {
				err = utils.NewErrorCode(ctx, 39991011, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), *req.Email, 14) //场景：14绑定(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
		case `email_code_to_unbing_email`:
			delete(data, k)
			email := loginInfo[daoUsers.Users.Columns().Email].String()
			if email == `` {
				err = utils.NewErrorCode(ctx, 39991013, ``)
				return
			}
			code, _ := cache.Code.Get(ctx, jbctx.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), email, 15) //场景：15解绑(邮箱)
			if code == `` || code != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39991999, ``)
				return
			}
			data[daoUsers.Users.Columns().Email] = nil
		case `id_card_no`:
			initPrivacyInfo()
			if privacyInfo[daoUsers.Privacy.Columns().IdCardNo].String() != `` {
				err = utils.NewErrorCode(ctx, 39992000, ``)
				return
			}

			idCardInfo, errTmp := id_card.NewHandler(ctx).Auth(gconv.String(data[`id_card_name`]), gconv.String(data[`id_card_no`]))
			if errTmp != nil {
				err = errTmp
				return
			}
			if idCardInfo.Gender != 0 {
				data[daoUsers.Privacy.Columns().IdCardGender] = idCardInfo.Gender
				if loginInfo[daoUsers.Users.Columns().Gender].Uint() == 0 {
					data[daoUsers.Users.Columns().Gender] = idCardInfo.Gender
				}
			}
			if !idCardInfo.Birthday.IsZero() {
				data[daoUsers.Privacy.Columns().IdCardBirthday] = idCardInfo.Birthday
				if loginInfo[daoUsers.Users.Columns().Birthday].GTime().IsZero() {
					data[daoUsers.Users.Columns().Birthday] = idCardInfo.Birthday
				}
			}
			if idCardInfo.Address != `` {
				data[daoUsers.Privacy.Columns().IdCardAddress] = idCardInfo.Address
				if loginInfo[daoUsers.Users.Columns().Address].String() == `` {
					data[daoUsers.Users.Columns().Address] = idCardInfo.Address
				}
			}
		}
	}
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}

	filter := map[string]any{`id`: loginInfo[`login_id`]}
	/**--------参数处理 结束--------**/

	_, err = service.Users().Update(ctx, filter, data)
	return
}
