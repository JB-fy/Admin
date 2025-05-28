package app

import (
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sApp struct{}

func NewApp() *sApp {
	return &sApp{}
}

func init() {
	service.RegisterApp(NewApp())
}

// 新增
func (logicThis *sApp) Create(ctx context.Context, data map[string]any) (id any, err error) {
	daoModelThis := daoApp.App.CtxDaoModel(ctx)

	id = data[daoApp.App.Columns().AppId]
	_, err = daoModelThis.HookInsert(data).Insert()
	return
}

// 修改
func (logicThis *sApp) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoModelThis := daoApp.App.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sApp) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoApp.App.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoApp.Pkg.CtxDaoModel(ctx).Filter(daoApp.Pkg.Columns().AppId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.app.app`), count, g.I18n().T(ctx, `name.app.pkg`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
