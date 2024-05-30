package logic

import (
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sUserUser struct{}

func NewUserUser() *sUserUser {
	return &sUserUser{}
}

func init() {
	service.RegisterUserUser(NewUserUser())
}

// 新增
func (logicThis *sUserUser) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoThis := daoUser.User
	daoModelThis := daoThis.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sUserUser) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoThis := daoUser.User
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sUserUser) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoThis := daoUser.User
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
