package controller

import (
	"api/api"
	apiMy "api/api/platform/my"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Admin struct{}

func NewAdmin() *Admin {
	return &Admin{}
}

// 用户详情
func (controllerThis *Admin) Info(ctx context.Context, req *apiMy.AdminInfoReq) (res *apiMy.AdminInfoRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	res = &apiMy.AdminInfoRes{}
	loginInfo.Struct(&res.Info)
	return
}

// 修改个人信息
func (controllerThis *Admin) Update(ctx context.Context, req *apiMy.AdminUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	loginInfo := utils.GetCtxLoginInfo(ctx)
	filter := map[string]interface{}{`id`: loginInfo[`adminId`]}
	/**--------参数处理 结束--------**/

	err = service.Admin().Update(ctx, filter, data)
	return
}
