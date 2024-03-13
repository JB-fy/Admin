package controller

import (
	"api/api"
	apiCurrent "api/api/app"
	"api/internal/cache"
	daoAuth "api/internal/dao/auth"
	daoUser "api/internal/dao/user"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
)

type Login struct{}

func NewLogin() *Login {
	return &Login{}
}

// 获取加密盐
func (controllerThis *Login) Salt(ctx context.Context, req *apiCurrent.LoginSaltReq) (res *api.CommonSaltRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`regex:^[\\p{L}][\\p{L}\\p{N}_]+$`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	userColumns := daoUser.User.Columns()
	info, _ := daoUser.User.CtxDaoModel(ctx).Filter(`loginName`, req.LoginName).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[userColumns.IsStop].Uint() == 1 {
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
	res = &api.CommonSaltRes{SaltStatic: info[userColumns.Salt].String(), SaltDynamic: saltDynamic}
	return
}

// 登录
func (controllerThis *Login) Login(ctx context.Context, req *apiCurrent.LoginLoginReq) (res *api.CommonTokenRes, err error) {
	if g.Validator().Rules(`phone`).Data(req.LoginName).Run(ctx) != nil && g.Validator().Rules(`regex:^[\\p{L}][\\p{L}\\p{N}_]+$`).Data(req.LoginName).Run(ctx) != nil {
		err = utils.NewErrorCode(ctx, 89990000, ``)
		return
	}

	userColumns := daoUser.User.Columns()
	info, _ := daoUser.User.CtxDaoModel(ctx).Filter(`loginName`, req.LoginName).One()
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 39990000, ``)
		return
	}
	if info[userColumns.IsStop].Uint() == 1 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}

	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	if req.Password != `` { //密码
		salt, _ := cache.NewSalt(ctx, sceneCode, req.LoginName).Get()
		if salt == `` || gmd5.MustEncrypt(info[userColumns.Password].String()+salt) != req.Password {
			err = utils.NewErrorCode(ctx, 39990001, ``)
			return
		}
	} else if req.SmsCode != `` { //短信验证码
		phone := info[userColumns.Phone].String()
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
	userColumns := daoUser.User.Columns()
	data := g.Map{}
	if req.Account != `` {
		info, _ := daoUser.User.CtxDaoModel(ctx).Filter(userColumns.Account, req.Account).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[userColumns.Account] = req.Account
		data[userColumns.Nickname] = req.Account
	}
	if req.Password != `` {
		data[userColumns.Password] = req.Password
	}
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[daoAuth.Scene.Columns().SceneCode].String()
	if req.Phone != `` {
		smsCode, _ := cache.NewSms(ctx, sceneCode, req.Phone, 1).Get() //使用场景：1注册
		if smsCode == `` || smsCode != req.SmsCode {
			err = utils.NewErrorCode(ctx, 39990008, ``)
			return
		}

		info, _ := daoUser.User.CtxDaoModel(ctx).Filter(userColumns.Phone, req.Phone).One()
		if !info.IsEmpty() {
			err = utils.NewErrorCode(ctx, 39990004, ``)
			return
		}
		data[userColumns.Phone] = req.Phone
		data[userColumns.Nickname] = req.Phone[:3] + `****` + req.Phone[len(req.Phone)-4:]
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
