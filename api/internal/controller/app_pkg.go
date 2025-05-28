package controller

import (
	"api/api"
	daoApp "api/internal/dao/app"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type AppPkg struct {
	defaultFieldOfInfo []string
	allowField         []string
}

func NewAppPkg() *AppPkg {
	field := daoApp.Pkg.ColumnArr()
	defaultFieldOfInfo := []string{`id`, `label`, `download_url_to_app`, `download_url_to_h5`}
	return &AppPkg{
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, defaultFieldOfInfo...),
	}
}

// 详情
func (controllerThis *AppPkg) Info(ctx context.Context, req *api.AppPkgInfoReq) (res *api.AppPkgInfoRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	filter[daoApp.Pkg.Columns().IsStop] = 0

	field := controllerThis.defaultFieldOfInfo

	fieldWithParam := g.Map{}
	if req.VerNoOfCurrent != nil {
		fieldWithParam[`is_force`] = req.VerNoOfCurrent
	}
	/**--------参数处理 结束--------**/

	info, err := daoApp.Pkg.CtxDaoModel(ctx).Filters(filter).Fields(field...).FieldsWithParam(fieldWithParam).OrderDesc(daoApp.Pkg.Columns().VerNo).One()
	if err != nil {
		return
	}
	if info.IsEmpty() && req.VerNoOfCurrent == nil {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &api.AppPkgInfoRes{}
	info.Struct(&res.Info)
	return
}
