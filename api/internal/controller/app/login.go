package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUser "api/internal/dao/user"
	"api/internal/utils"
	one_click "api/internal/utils/one-click"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取密码盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoUser.User.CtxDaoModel(ctx).Filter(`login_name`, req.LoginName).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoUser.User.Columns().IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	saltDynamic := grand.S(8)
	err = cache.NewSalt(ctx, sceneCode, req.LoginName).Set(saltDynamic, 5)
	if err != nil {
		return
	}
	res = &api.CommonSaltRes{SaltStatic: info[daoUser.User.Columns().Salt].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`regex:^[\\p{L}][\\p{L}\\p{N}_]{3,}$`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	info, _ := daoUser.User.CtxDaoModel(ctx).Filter(`login_name`, req.LoginName).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[daoUser.User.Columns().IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	if req.Password != `` { //密码
		salt, _ := cache.NewSalt(ctx, sceneCode, req.LoginName).Get()
		if salt == `` || gmd5.MustEncrypt(info[daoUser.User.Columns().Password].String()+salt) != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	} else if req.SmsCode != `` { //短信验证码
		phone := info[daoUser.User.Columns().Phone].String()
		if phone == `` {
			err = utils.NewErrorCode(ctx, 39990007, ``)
			return
		}

		smsCode, _ := cache.NewSms(ctx, sceneCode, phone, 0).Get() //使用场景：0登录
		if smsCode == `` || smsCode != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39990008, ``)
			return
		}
	}

	claims := utils.CustomClaims{LoginId: info[daoUser.User.PrimaryKey()].Uint()}
	jwt := utils.NewJWT(ctx, sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, sceneCode, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}

// 注册
func (controllerThis *Login) Register(ctx context.Context, req *apiCurrent.LoginRegisterReq) (res *api.CommonTokenRes, err error) {
	data := g.Map{}
	if req.Account != `` {
		info, _ := daoUser.User.CtxDaoModel(ctx).Filter(daoUser.User.Columns().Account, req.Account).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[daoUser.User.Columns().Account] = req.Account
		data[daoUser.User.Columns().Nickname] = req.Account
	}
	if req.Password != `` {
		data[daoUser.User.Columns().Password] = req.Password
	}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	if req.Phone != `` {
		smsCode, _ := cache.NewSms(ctx, sceneCode, req.Phone, 1).Get() //使用场景：1注册
		if smsCode == `` || smsCode != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39990008, ``)
			return
		}

		info, _ := daoUser.User.CtxDaoModel(ctx).Filter(daoUser.User.Columns().Phone, req.Phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[daoUser.User.Columns().Phone] = req.Phone
		data[daoUser.User.Columns().Nickname] = req.Phone[:3] + `****` + req.Phone[len(req.Phone)-4:]
	}

	userId, err := daoUser.User.CtxDaoModel(ctx).HookInsert(data).InsertAndGetId()
	if err != nil {
		return
	}

	claims := utils.CustomClaims{LoginId: uint(userId)}
	jwt := utils.NewJWT(ctx, sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, sceneCode, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}

// 密码找回
func (controllerThis *Login) PasswordRecovery(ctx context.Context, req *apiCurrent.LoginPasswordRecoveryReq) (res *api.CommonNoDataRes, err error) {
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	smsCode, _ := cache.NewSms(ctx, sceneCode, req.Phone, 2).Get() //使用场景：2密码找回
	if smsCode == `` || smsCode != req.SmsCode {
		err = utils.NewErrorCode(ctx, 39990008, ``)
		return
	}

	row, err := daoUser.User.CtxDaoModel(ctx).Filter(daoUser.User.Columns().Phone, req.Phone).HookUpdate(g.Map{daoUser.User.Columns().Password: req.Password}).UpdateAndGetAffected()
	if err != nil {
		return
	}
	if row == 0 {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	return
}

// 一键登录前置信息
func (controllerThis *Login) OneClickPreInfo(ctx context.Context, req *apiCurrent.LoginOneClickPreInfoReq) (res *apiCurrent.LoginOneClickPreInfoRes, err error) {
	res = &apiCurrent.LoginOneClickPreInfoRes{}
	switch req.OneClickType {
	case `oneClickOfWx`: //微信
		res.CodeUrlOfWx, err = one_click.NewOneClickOfWx(ctx).CodeUrl(req.RedirectUriOfWx, req.ScopeOfWx, req.StateOfWx, req.ForcePopupOfWx)
	case `oneClickOfYidun`: //易盾
	}
	return
}

// 一键登录
func (controllerThis *Login) OneClick(ctx context.Context, req *apiCurrent.LoginOneClickReq) (res *api.CommonTokenRes, err error) {
	filter := g.Map{}
	saveData := g.Map{}
	switch req.OneClickType {
	case `oneClickOfWx`: //微信
		accessToken, errTmp := one_click.NewOneClickOfWx(ctx).AccessToken(req.CodeOfWx)
		if errTmp != nil {
			err = errTmp
			return
		}
		filter[daoUser.User.Columns().OpenIdOfWx] = accessToken.OpenId
		saveData[daoUser.User.Columns().OpenIdOfWx] = accessToken.OpenId
		if accessToken.Scope == `snsapi_userinfo` {
			userInfo, errTmp := one_click.NewOneClickOfWx(ctx).UserInfo(accessToken.OpenId, accessToken.AccessToken)
			if errTmp != nil {
				err = errTmp
				return
			}
			saveData[daoUser.User.Columns().Nickname] = userInfo.Nickname
			saveData[daoUser.User.Columns().Gender] = userInfo.Gender
			saveData[daoUser.User.Columns().Avatar] = userInfo.Avatar
		}
	case `oneClickOfYidun`: //易盾
		phone, errTmp := one_click.NewOneClickOfYidun(ctx).Check(req.TokenOfYidun, req.AccessTokenOfYidun)
		if errTmp != nil {
			err = errTmp
			return
		}
		filter[daoUser.User.Columns().Phone] = phone
		saveData[daoUser.User.Columns().Phone] = phone
	}

	userId, _ := daoUser.User.CtxDaoModel(ctx).Filters(filter).ValueUint(daoUser.User.PrimaryKey())
	if userId == 0 {
		userIdTmp, errTmp := daoUser.User.CtxDaoModel(ctx).HookInsert(saveData).InsertAndGetId()
		if errTmp != nil { //报错就是并发引起的唯一索引冲突，故再做一次查询
			userId, _ = daoUser.User.CtxDaoModel(ctx).Filters(filter).ValueUint(daoUser.User.PrimaryKey())
			// daoUser.User.CtxDaoModel(ctx).Filters(filter).Update(saveData)	//一般情况下系统用户昵称，性别等字段不会随微信变动而改动
		} else {
			userId = uint(userIdTmp)
		}
	} /*  else {
		daoUser.User.CtxDaoModel(ctx).Filters(filter).Update(saveData)	//一般情况下系统用户昵称，性别等字段不会随微信变动而改动
	} */

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	claims := utils.CustomClaims{LoginId: userId}
	jwt := utils.NewJWT(ctx, sceneInfo[daoAuth.Scene.Columns().SceneConfig].Map())
	token, err := jwt.CreateToken(claims)
	if err != nil {
		return
	}
	// cache.NewToken(ctx, sceneCode, claims.LoginId).Set(token, int64(jwt.ExpireTime)) //缓存token（限制多地登录，多设备登录等情况下用）

	res = &api.CommonTokenRes{Token: token}
	return
}
