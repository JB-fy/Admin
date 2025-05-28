package app

import (
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAppPkg struct{}

func NewAppPkg() *sAppPkg {
	return &sAppPkg{}
}

func init() {
	service.RegisterAppPkg(NewAppPkg())
}

// 验证数据（create和update共用）
func (logicThis *sAppPkg) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[daoApp.Pkg.Columns().AppId]; ok && gconv.String(data[daoApp.Pkg.Columns().AppId]) != `` {
		if count, _ := daoApp.App.CtxDaoModel(ctx).FilterPri(data[daoApp.Pkg.Columns().AppId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.app.app`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sAppPkg) Create(ctx context.Context, data map[string]any) (id any, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoApp.Pkg.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAppPkg) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoApp.Pkg.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAppPkg) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoApp.Pkg.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
