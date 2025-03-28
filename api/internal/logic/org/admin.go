package org

import (
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOrgAdmin struct{}

func NewOrgAdmin() *sOrgAdmin {
	return &sOrgAdmin{}
}

func init() {
	service.RegisterOrgAdmin(NewOrgAdmin())
}

// 验证数据（create和update共用）
func (logicThis *sOrgAdmin) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[daoOrg.Admin.Columns().OrgId]; ok && gconv.Uint(data[daoOrg.Admin.Columns().OrgId]) > 0 {
		if count, _ := daoOrg.Org.CtxDaoModel(ctx).FilterPri(data[daoOrg.Admin.Columns().OrgId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.org.org`)}})
			return
		}
	}

	if _, ok := data[`role_id_arr`]; ok && len(gconv.Uints(data[`role_id_arr`])) > 0 {
		roleIdArr := gconv.Uints(data[`role_id_arr`])
		if count, _ := daoAuth.Role.CtxDaoModel(ctx).FilterPri(roleIdArr).Filter(daoAuth.Role.Columns().SceneId, `org`).Count(); count != len(roleIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sOrgAdmin) Create(ctx context.Context, data map[string]any) (id any, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoOrg.Admin.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sOrgAdmin) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoOrg.Admin.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sOrgAdmin) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoOrg.Admin.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
