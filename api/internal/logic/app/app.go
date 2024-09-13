package logic

import (
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sApp struct{}

func NewApp() *sApp {
	return &sApp{}
}

func init() {
	service.RegisterApp(NewApp())
}

// 新增
func (logicThis *sApp) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoModelThis := daoApp.App.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
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

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
