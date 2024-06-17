package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthRole struct{}

func NewAuthRole() *sAuthRole {
	return &sAuthRole{}
}

func init() {
	service.RegisterAuthRole(NewAuthRole())
}

// 验证数据（create和update共用）
func (logicThis *sAuthRole) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[daoAuth.Role.Columns().SceneId]; ok && gconv.Uint(data[daoAuth.Role.Columns().SceneId]) > 0 {
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneId, data[daoAuth.Role.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`)}})
			return
		}
	}

	if _, ok := data[`action_id_arr`]; ok && len(gconv.SliceUint(data[`action_id_arr`])) > 0 {
		actionIdArr := gconv.SliceUint(data[`action_id_arr`])
		if count, _ := daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.Columns().ActionId, actionIdArr).Count(); count != len(actionIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.action`)}})
			return
		}
	}

	if _, ok := data[`menu_id_arr`]; ok && len(gconv.SliceUint(data[`menu_id_arr`])) > 0 {
		menuIdArr := gconv.SliceUint(data[`menu_id_arr`])
		if count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filter(daoAuth.Menu.Columns().MenuId, menuIdArr).Count(); count != len(menuIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.menu`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sAuthRole) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Role.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Role.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthRole) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoAuth.Role.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoAuth.RoleRelOfOrgAdmin.CtxDaoModel(ctx).Filter(daoAuth.RoleRelOfOrgAdmin.Columns().RoleId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`), count, g.I18n().T(ctx, `name.auth.roleRelOfOrgAdmin`)}})
		return
	}

	if count, _ := daoAuth.RoleRelOfPlatformAdmin.CtxDaoModel(ctx).Filter(daoAuth.RoleRelOfPlatformAdmin.Columns().RoleId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`), count, g.I18n().T(ctx, `name.auth.roleRelOfPlatformAdmin`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
