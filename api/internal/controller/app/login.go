package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUsers "api/internal/dao/users"
	"api/internal/utils"
	one_click "api/internal/utils/one-click"
	"api/internal/utils/token"
	"context"
	"slices"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取密码盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	filter := g.Map{}
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoUsers.Users.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoUsers.Users.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoUsers.Users.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoUsers.Users.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoUsers.Users.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	saltStatic, _ := daoUsers.Privacy.CtxDaoModel(ctx).FilterPri(info[daoUsers.Users.Columns().UserId]).ValueStr(daoUsers.Privacy.Columns().Salt)
	if saltStatic == `` {
		err = utils.NewErrorCode(ctx, 39990004, ``)
		return
	}
	saltDynamic := grand.S(8)
	err = cache.Salt.Set(ctx, utils.GetCtxSceneInfo(ctx)[daoAuth.Scene.Columns().SceneId].String(), req.LoginName, saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: saltStatic, SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	filter := g.Map{}
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoUsers.Users.Columns().Phone] = req.LoginName
	} else if g.Validator().Rules(`email`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoUsers.Users.Columns().Email] = req.LoginName
	} else if g.Validator().Rules(`regex:^[\p{L}][\p{L}\p{N}_]{3,}$`).Data(req.LoginName).Run(ctx) == nil {
		filter[daoUsers.Users.Columns().Account] = req.LoginName
	} else {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoUsers.Users.CtxDaoModel(ctx).Filters(filter).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoUsers.Users.Columns().IsStop].Uint8() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	if req.Password != `` { //密码
		password, _ := daoUsers.Privacy.CtxDaoModel(ctx).FilterPri(info[daoUsers.Users.Columns().UserId]).ValueStr(daoUsers.Privacy.Columns().Password)
		if password == `` {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		salt, _ := cache.Salt.Get(ctx, sceneId, req.LoginName)
		if salt == `` || gmd5.MustEncrypt(password+salt) != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	} else if req.SmsCode != `` { //短信验证码
		phone := info[daoUsers.Users.Columns().Phone].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39991003, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, phone, 0) //场景：0登录(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
	} else if req.EmailCode != `` { //邮箱验证码
		email := info[daoUsers.Users.Columns().Email].String()
		if email == `` {
			err = utils.NewErrorCode(ctx, 39991013, ``)
			return
		}
		code, _ := cache.Code.Get(ctx, sceneId, email, 10) //场景：10登录(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
	}

	token, err := token.NewHandler(ctx).Create(info[daoUsers.Users.Columns().UserId].String(), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}

// 注册
func (controllerThis *Login) Register(ctx context.Context, req *apiCurrent.LoginRegisterReq) (res *api.CommonTokenRes, err error) {
	data := g.Map{}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	if req.Phone != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Phone, 1) //场景：1注册(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Phone, req.Phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991000, ``)
			return
		}
		data[daoUsers.Users.Columns().Phone] = req.Phone
		data[daoUsers.Users.Columns().Nickname] = req.Phone[:3] + gstr.Repeat(`*`, len(req.Phone)-7) + req.Phone[len(req.Phone)-4:]
	}
	if req.Email != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Email, 11) //场景：11注册(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Email, req.Email).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991010, ``)
			return
		}
		data[daoUsers.Users.Columns().Email] = req.Email
		data[daoUsers.Users.Columns().Nickname] = gstr.Split(req.Email, `@`)[0]
	}
	if req.Account != `` {
		info, _ := daoUsers.Users.CtxDaoModel(ctx).Filter(daoUsers.Users.Columns().Account, req.Account).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39991020, ``)
			return
		}
		data[daoUsers.Users.Columns().Account] = req.Account
		accountRune := []rune(req.Account)
		accountRuneLen := len(accountRune)
		data[daoUsers.Users.Columns().Nickname] = string(accountRune[:1]) + gstr.Repeat(`*`, accountRuneLen-2) + string(accountRune[accountRuneLen-1:])
	}
	if req.Password != `` {
		data[daoUsers.Privacy.Columns().Password] = req.Password
	}

	userId, err := daoUsers.Users.CtxDaoModel(ctx).HookInsert(data).InsertAndGetId()
	if err != nil {
		return
	}

	token, err := token.NewHandler(ctx).Create(gconv.String(userId), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}

// 密码找回
func (controllerThis *Login) PasswordRecovery(ctx context.Context, req *apiCurrent.LoginPasswordRecoveryReq) (res *api.CommonNoDataRes, err error) {
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneId := sceneInfo[daoAuth.Scene.Columns().SceneId].String()
	filter := g.Map{}
	if req.Phone != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Phone, 2) //场景：2密码找回(手机)
		if code == `` || code != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoUsers.Users.Columns().Phone] = req.Phone
	} else if req.Email != `` {
		code, _ := cache.Code.Get(ctx, sceneId, req.Email, 12) //场景：12密码找回(邮箱)
		if code == `` || code != req.EmailCode {
			err = utils.NewErrorCode(ctx, 39991999, ``)
			return
		}
		filter[daoUsers.Users.Columns().Email] = req.Email
	}

	daoModelUsers := daoUsers.Users.CtxDaoModel(ctx).SetIdArr(filter)
	if len(daoModelUsers.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	_, err = daoModelUsers.HookUpdateOne(daoUsers.Privacy.Columns().Password, req.Password).Update()
	return
}

// 一键登录前置信息
func (controllerThis *Login) OneClickPreInfo(ctx context.Context, req *apiCurrent.LoginOneClickPreInfoReq) (res *apiCurrent.LoginOneClickPreInfoRes, err error) {
	res = &apiCurrent.LoginOneClickPreInfoRes{}
	switch req.OneClickType {
	case `one_click_of_wx`: //微信
		res.CodeUrlOfWx, err = one_click.NewOneClickOfWx(ctx).CodeUrl(req.RedirectUriOfWx, req.ScopeOfWx, req.ForcePopupOfWx, req.StateOfWx)
	case `one_click_of_yidun`: //易盾
	}
	return
}

// 一键登录
func (controllerThis *Login) OneClick(ctx context.Context, req *apiCurrent.LoginOneClickReq) (res *api.CommonTokenRes, err error) {
	filter := g.Map{}
	saveData := g.Map{}
	switch req.OneClickType {
	case `one_click_of_wx`: //微信
		accessToken, errTmp := one_click.NewOneClickOfWx(ctx).AccessToken(ctx, req.CodeOfWx)
		if errTmp != nil {
			err = errTmp
			return
		}
		filter[daoUsers.Users.Columns().WxOpenid] = accessToken.Openid
		saveData[daoUsers.Users.Columns().WxOpenid] = accessToken.Openid
		if accessToken.Unionid != `` {
			saveData[daoUsers.Users.Columns().WxUnionid] = accessToken.Unionid
		}
		if slices.Contains([]string{`snsapi_userinfo`, `snsapi_login`}, accessToken.Scope) {
			userInfo, errTmp := one_click.NewOneClickOfWx(ctx).UserInfo(ctx, accessToken.Openid, accessToken.AccessToken)
			if errTmp != nil {
				err = errTmp
				return
			}
			saveData[daoUsers.Users.Columns().WxUnionid] = userInfo.Unionid
			saveData[daoUsers.Users.Columns().Nickname] = userInfo.Nickname
			saveData[daoUsers.Users.Columns().Gender] = userInfo.Gender
			saveData[daoUsers.Users.Columns().Avatar] = userInfo.Avatar
		}
	case `one_click_of_yidun`: //易盾
		phone, errTmp := one_click.NewOneClickOfYidun(ctx).Check(ctx, req.TokenOfYidun, req.AccessTokenOfYidun)
		if errTmp != nil {
			err = errTmp
			return
		}
		filter[daoUsers.Users.Columns().Phone] = phone
		saveData[daoUsers.Users.Columns().Phone] = phone
	}

	userId, _ := daoUsers.Users.CtxDaoModel(ctx).Filters(filter).ValueUint(daoUsers.Users.Columns().UserId)
	if userId == 0 {
		userIdTmp, errTmp := daoUsers.Users.CtxDaoModel(ctx).HookInsert(saveData).InsertAndGetId()
		if errTmp != nil { //报错就是并发引起的唯一索引冲突，故再做一次查询
			userId, _ = daoUsers.Users.CtxDaoModel(ctx).Filters(filter).ValueUint(daoUsers.Users.Columns().UserId)
		} else {
			userId = uint(userIdTmp)
		}
	} /* else {
		daoUsers.Users.CtxDaoModel(ctx).SetIdArr(filter).HookUpdate(saveData).Update() //一般情况下用户昵称，性别等字段不会每次登录都随第三方变动
	} */

	token, err := token.NewHandler(ctx).Create(gconv.String(userId), nil)
	if err != nil {
		return
	}

	res = &api.CommonTokenRes{Token: token}
	return
}
