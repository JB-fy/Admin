package logic

import (
	daoPay "api/internal/dao/pay"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sPay struct{}

func NewPay() *sPay {
	return &sPay{}
}

func init() {
	service.RegisterPay(NewPay())
}

// 新增
func (logicThis *sPay) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoModelThis := daoPay.Pay.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sPay) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoModelThis := daoPay.Pay.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPay) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoPay.Pay.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
