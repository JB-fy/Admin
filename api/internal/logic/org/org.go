package logic

import (
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sOrg struct{}

func NewOrg() *sOrg {
	return &sOrg{}
}

func init() {
	service.RegisterOrg(NewOrg())
}

// 新增
func (logicThis *sOrg) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoModelThis := daoOrg.Org.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sOrg) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoModelThis := daoOrg.Org.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sOrg) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoOrg.Org.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoOrg.Admin.CtxDaoModel(ctx).Filter(daoOrg.Admin.Columns().OrgId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.org.org`), count, g.I18n().T(ctx, `name.org.admin`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
