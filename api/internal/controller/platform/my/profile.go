package controller

import (
	"api/api"
	apiMy "api/api/platform/my"
	daoPlatform "api/internal/dao/platform"
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
	data := gconv.MapDeep(req)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	adminDao := daoPlatform.Admin
	adminColumns := adminDao.Columns()
	loginInfo := utils.GetCtxLoginInfo(ctx)
	for k, v := range data {
		switch k {
		case `passwordToCheck`:
			if gmd5.MustEncrypt(gconv.String(v)+loginInfo[adminColumns.Salt].String()) != loginInfo[adminColumns.Password].String() {
				err = utils.NewErrorCode(ctx, 39990003, ``)
				return
			}
			delete(data, k)
		}
	}

	filter := map[string]interface{}{`id`: loginInfo[adminDao.PrimaryKey()]}
	/**--------参数处理 结束--------**/

	_, err = service.PlatformAdmin().Update(ctx, filter, data)
	return
}
