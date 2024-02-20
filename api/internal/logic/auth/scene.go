package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sAuthScene struct{}

func NewAuthScene() *sAuthScene {
	return &sAuthScene{}
}

func init() {
	service.RegisterAuthScene(NewAuthScene())
}

// 新增
func (logicThis *sAuthScene) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Scene
	id, err = daoThis.DaoModel(ctx).HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthScene) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Scene
	daoModelThis := daoThis.DaoModel(ctx).Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthScene) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Scene
	daoModelThis := daoThis.DaoModel(ctx).Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
