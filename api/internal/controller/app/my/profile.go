package controller

import (
	"api/api"
	apiMy "api/api/app/my"
	"api/internal/consts"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"fmt"

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
	data := gconv.MapDeep(req)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	loginInfo := utils.GetCtxLoginInfo(ctx)
	for k, v := range data {
		switch k {
		case `checkPassword`:
			if gmd5.MustEncrypt(gconv.String(v)+loginInfo[`salt`].String()) != loginInfo[`password`].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
			delete(data, k)
		case `smsCodeToPassword`, `smsCodeToBindPhone`, `smsCodeToUnbingPhone`:
			phone := loginInfo[`phone`].String()
			if phone == `` {
				err = utils.NewErrorCode(ctx, 39990007, ``)
				return
			}
			useScene := 3 //使用场景：3密码修改
			if k == `smsCodeToBindPhone` {
				useScene = 4 //使用场景：4绑定手机
			} else if k == `smsCodeToUnbingPhone` {
				useScene = 5 //使用场景：5解绑手机
			}

			sceneInfo := utils.GetCtxSceneInfo(ctx)
			sceneCode := sceneInfo[`sceneCode`].String()
			smsKey := fmt.Sprintf(consts.CacheSmsFormat, sceneCode, phone, useScene)
			smsCode, _ := g.Redis().Get(ctx, smsKey)
			if smsCode.String() != gconv.String(v) {
				err = utils.NewErrorCode(ctx, 39990008, ``)
				return
			}
			delete(data, k)
		}
	}

	filter := map[string]interface{}{`id`: loginInfo[`userId`]}
	/**--------参数处理 结束--------**/

	_, err = service.User().Update(ctx, filter, data)
	return
}
