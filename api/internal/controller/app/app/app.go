package controller

import (
	apiApp "api/api/app/app"
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
	field := daoApp.App.ColumnArr().Slice()
	defaultFieldOfInfo := []string{`id`, `label`}
	return &App{
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, defaultFieldOfInfo...),
	}
}

// 详情
func (controllerThis *App) Info(ctx context.Context, req *apiApp.AppInfoReq) (res *apiApp.AppInfoRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	filter[daoApp.App.Columns().IsStop] = 0

	field := daoApp.App.ColumnArr().Slice()
	fieldWithParam := g.Map{}
	if req.CurrentVerNo != nil {
		fieldWithParam[`is_force`] = req.CurrentVerNo
	}
	/**--------参数处理 结束--------**/

	info, err := daoApp.App.CtxDaoModel(ctx).Filters(filter).Fields(field...).FieldsWithParam(fieldWithParam).OrderDesc(daoApp.App.Columns().VerNo).One()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApp.AppInfoRes{}
	info.Struct(&res.Info)
	return
}
