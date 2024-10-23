package users

import (
	daoUsers "api/internal/dao/users"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sUsers struct{}

func NewUsers() *sUsers {
	return &sUsers{}
}

func init() {
	service.RegisterUsers(NewUsers())
}

// 新增
func (logicThis *sUsers) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoModelThis := daoUsers.Users.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sUsers) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoModelThis := daoUsers.Users.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sUsers) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoUsers.Users.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
