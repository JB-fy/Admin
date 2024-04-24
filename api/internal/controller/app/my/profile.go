package controller

import (
	"api/api"
	apiMy "api/api/app/my"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	id_card "api/internal/utils/id-card"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
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

	userColumns := daoUser.User.Columns()
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	for k, v := range data {
		switch k {
		/* case `account`: //前端太懒，可能把个人信息全部传回来，导致account有值，故不能用required-with:Account直接验证
		if gconv.String(v) != loginInfo[userColumns.Account].String() && g.Validator().Rules(`required`).Data(req.PasswordToCheck).Run(ctx) != nil {
			err = utils.NewErrorCode(ctx, 89999999, ``)
			return
		} */
		case `password`:
			if g.Validator().Rules(`required`).Data(req.PasswordToCheck).Run(ctx) != nil && g.Validator().Rules(`required`).Data(req.SmsCodeToPassword).Run(ctx) != nil {
				err = utils.NewErrorCode(ctx, 89999999, ``)
				return
			}
		case `password_to_check`:
			if gmd5.MustEncrypt(gconv.String(v)+loginInfo[userColumns.Salt].String()) != loginInfo[userColumns.Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
			delete(data, k)
		case `smsCodeToPassword`:
			phone := loginInfo[userColumns.Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39990007, ``)
				return
			}

			smsCode, _ := cache.NewSms(ctx, sceneCode, phone, 3).Get() //使用场景：3密码修改
			if smsCode == `` || smsCode != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39990008, ``)
				return
			}
			delete(data, k)
		case `smsCodeToBindPhone`:
			if loginInfo[userColumns.Phone].String() != `` {
				err = utils.NewErrorCode(ctx, 39990005, ``)
				return
			}

			phone := gconv.String(data[`phone`])
			smsCode, _ := cache.NewSms(ctx, sceneCode, phone, 4).Get() //使用场景：4绑定手机
			if smsCode == `` || smsCode != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39990008, ``)
				return
			}
			delete(data, k)
		case `smsCodeToUnbingPhone`:
			phone := loginInfo[userColumns.Phone].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39990007, ``)
				return
			}

			smsCode, _ := cache.NewSms(ctx, sceneCode, phone, 5).Get() //使用场景：5解绑手机
			if smsCode == `` || smsCode != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39990008, ``)
				return
			}
			delete(data, k)
			data[userColumns.Phone] = nil
		case `id_card_no`:
			if loginInfo[userColumns.IdCardNo].String() != `` {
				err = utils.NewErrorCode(ctx, 39990009, ``)
				return
			}

			idCardInfo, errTmp := id_card.NewIdCard(ctx).Auth(gconv.String(data[`id_card_name`]), gconv.String(data[`id_card_no`]))
			if errTmp != nil {
				err = errTmp
				return
			}
			if idCardInfo.Gender != 0 {
				data[userColumns.Gender] = idCardInfo.Gender
			}
			if idCardInfo.Address != `` {
				data[userColumns.Address] = idCardInfo.Address
			}
			if idCardInfo.Birthday != `` {
				data[userColumns.Birthday] = idCardInfo.Birthday
			}
		}
	}

	filter := map[string]interface{}{`id`: loginInfo[`login_id`]}
	/**--------参数处理 结束--------**/

	_, err = service.UserUser().Update(ctx, filter, data)
	return
}
