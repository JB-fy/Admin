package pay

import (
	daoPay "api/internal/dao/pay"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
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

	daoModelThis.SetIdArr(filter)
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

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoPay.Channel.CtxDaoModel(ctx).Filter(daoPay.Channel.Columns().PayId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.pay.pay`), count, g.I18n().T(ctx, `name.pay.channel`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
