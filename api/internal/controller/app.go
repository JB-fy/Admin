package controller

import (
	"api/api"
	daoApp "api/internal/dao/app"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type App struct {
	defaultFieldOfInfo []string
	allowField         []string
}

func NewApp() *App {
	field := daoApp.App.ColumnArr()
	defaultFieldOfInfo := []string{`id`, `label`}
	return &App{
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, defaultFieldOfInfo...),
	}
}

// 详情
func (controllerThis *App) Info(ctx context.Context, req *api.AppInfoReq) (res *api.AppInfoRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	filter[daoApp.App.Columns().IsStop] = 0

	field := daoApp.App.ColumnArr()
	field = append(field, `download_url_to_app`, `download_url_to_h5`)

	fieldWithParam := g.Map{}
	if req.CurrentVerNo != nil {
		fieldWithParam[`is_force`] = req.CurrentVerNo
	}
	/**--------参数处理 结束--------**/

	info, err := daoApp.App.CtxDaoModel(ctx).Filters(filter).Fields(field...).FieldsWithParam(fieldWithParam).OrderDesc(daoApp.App.Columns().VerNo).One()
	if err != nil {
		return
	}
	if info.IsEmpty() && req.CurrentVerNo == nil {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &api.AppInfoRes{}
	info.Struct(&res.Info)
	return
}
