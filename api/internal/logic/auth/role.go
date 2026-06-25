package auth

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
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
	if _, ok := data[daoAuth.Role.Columns().SceneId]; ok && gconv.String(data[daoAuth.Role.Columns().SceneId]) != `` {
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).FilterPri(data[daoAuth.Role.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`)}})
			return
		}
	}

	/* if _, ok := data[`action_id_arr`]; ok && len(gconv.Strings(data[`action_id_arr`])) > 0 {
		actionIdArr := gconv.Strings(data[`action_id_arr`])
		if count, _ := daoAuth.Action.CtxDaoModel(ctx).FilterPri(actionIdArr).Count(); count != len(actionIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.action`)}})
			return
		}
	}

	if _, ok := data[`menu_id_arr`]; ok && len(gconv.Uints(data[`menu_id_arr`])) > 0 {
		menuIdArr := gconv.Uints(data[`menu_id_arr`])
		if count, _ := daoAuth.Menu.CtxDaoModel(ctx).FilterPri(menuIdArr).Count(); count != len(menuIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.menu`)}})
			return
		}
	} */
	return
}

// 新增
func (logicThis *sAuthRole) Create(ctx context.Context, data map[string]any) (id any, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Role.CtxDaoModel(ctx)

	if _, ok := data[`action_id_arr`]; ok && len(gconv.Strings(data[`action_id_arr`])) > 0 {
		actionIdArr := gconv.Strings(data[`action_id_arr`])
		if count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: data[daoAuth.Role.Columns().SceneId]}).Count(); count != len(actionIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.action`)}})
			return
		}
	}

	if _, ok := data[`menu_id_arr`]; ok && len(gconv.Uints(data[`menu_id_arr`])) > 0 {
		menuIdArr := gconv.Uints(data[`menu_id_arr`])
		if count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{daoAuth.Menu.Columns().MenuId: menuIdArr, daoAuth.Menu.Columns().SceneId: data[daoAuth.Role.Columns().SceneId]}).Count(); count != len(menuIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.menu`)}})
			return
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Role.CtxDaoModel(ctx)

	sceneIdArr := []string{}
	sceneId := ``
	var list gdb.Result
	if _, ok := data[`action_id_arr`]; ok && len(gconv.Strings(data[`action_id_arr`])) > 0 {
		if len(sceneIdArr) == 0 {
			if _, ok := data[daoAuth.Role.Columns().SceneId]; ok {
				sceneIdArr = append(sceneIdArr, gconv.String(data[daoAuth.Role.Columns().SceneId]))
			} else {
				if len(list) == 0 {
					list, _ = daoAuth.Role.CtxDaoModel(ctx).Filters(filter).Fields(daoAuth.Role.Columns().RelId, daoAuth.Role.Columns().SceneId).All()
					if len(list) == 0 {
						err = utils.NewErrorCode(ctx, 29999998, ``)
						return
					}
				}
				sceneIdSet := map[string]struct{}{}
				for _, info := range list {
					sceneId := info[daoAuth.Role.Columns().SceneId].String()
					if _, ok := sceneIdSet[sceneId]; !ok {
						sceneIdSet[sceneId] = struct{}{}
						sceneIdArr = append(sceneIdArr, sceneId)
					}
				}
				/* if len(sceneIdSet) != 1 {
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				} */
			}
		}
		actionIdArr := gconv.Strings(data[`action_id_arr`])
		for _, sceneId := range sceneIdArr {
			if count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(g.Map{daoAuth.ActionRelToScene.Columns().SceneId: sceneId, daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr}).Count(); count != len(actionIdArr) {
				err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.action`)}})
				return
			}
		}
	}

	if _, ok := data[`menu_id_arr`]; ok && len(gconv.Uints(data[`menu_id_arr`])) > 0 {
		if sceneId == `` {
			if _, ok := data[daoAuth.Role.Columns().SceneId]; ok {
				sceneId = gconv.String(data[daoAuth.Role.Columns().SceneId])
			} else {
				if len(list) == 0 {
					list, _ = daoAuth.Role.CtxDaoModel(ctx).Filters(filter).Fields(daoAuth.Role.Columns().RelId, daoAuth.Role.Columns().SceneId).All()
					if len(list) == 0 {
						err = utils.NewErrorCode(ctx, 29999998, ``)
						return
					}
				}
				sceneIdSet := map[string]struct{}{}
				for _, info := range list {
					sceneIdSet[info[daoAuth.Role.Columns().SceneId].String()] = struct{}{}
				}
				if len(sceneIdSet) != 1 { //菜单所属场景只能设置一个，故必须同一场景下的角色才能一起修改
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				}
				sceneId = list[0][daoAuth.Role.Columns().SceneId].String()
			}
		}
		menuIdArr := gconv.Uints(data[`menu_id_arr`])
		if count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{daoAuth.Menu.Columns().SceneId: sceneId, daoAuth.Menu.Columns().MenuId: menuIdArr}).Count(); count != len(menuIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.menu`)}})
			return
		}
	}

	if len(list) > 0 {
		idArr := []uint{}
		for _, info := range list {
			idArr = append(idArr, info[daoAuth.Role.Columns().RelId].Uint())
		}
		filter = map[string]any{`id`: idArr}
	}

	daoModelThis.SetIdArr(filter)
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

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoAuth.RoleRelOfAdmin.CtxDaoModel(ctx).Filter(daoAuth.RoleRelOfAdmin.Columns().RoleId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`), count, g.I18n().T(ctx, `name.auth.roleRelOfAdmin`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
