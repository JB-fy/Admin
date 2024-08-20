package logic

import (
	daoPay "api/internal/dao/pay"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPayChannel struct{}

func NewPayChannel() *sPayChannel {
	return &sPayChannel{}
}

func init() {
	service.RegisterPayChannel(NewPayChannel())
}

// 验证数据（create和update共用）
func (logicThis *sPayChannel) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[daoPay.Channel.Columns().SceneId]; ok && gconv.Uint(data[daoPay.Channel.Columns().SceneId]) > 0 {
		if count, _ := daoPay.Scene.CtxDaoModel(ctx).Filter(daoPay.Scene.Columns().SceneId, data[daoPay.Channel.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.pay.scene`)}})
			return
		}
	}

	if _, ok := data[daoPay.Channel.Columns().PayId]; ok && gconv.Uint(data[daoPay.Channel.Columns().PayId]) > 0 {
		if count, _ := daoPay.Pay.CtxDaoModel(ctx).Filter(daoPay.Pay.Columns().PayId, data[daoPay.Channel.Columns().PayId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.pay.pay`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sPayChannel) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoPay.Channel.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sPayChannel) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoPay.Channel.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPayChannel) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoPay.Channel.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
