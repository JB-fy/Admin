package admin

import (
	"api/internal/consts"
	daoAdmin "api/internal/dao/admin"
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdmin struct{}

func NewAdmin() *sAdmin {
	return &sAdmin{}
}

func init() {
	service.RegisterAdmin(NewAdmin())
}

// 验证数据（create和update共用）
func (logicThis *sAdmin) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[daoAdmin.Admin.Columns().SceneId]; ok && gconv.String(data[daoAdmin.Admin.Columns().SceneId]) != `` {
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).FilterPri(data[daoAdmin.Admin.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`)}})
			return
		}
	}

	/* if _, ok := data[`role_id_arr`]; ok && len(gconv.Uints(data[`role_id_arr`])) > 0 {
		roleIdArr := gconv.Uints(data[`role_id_arr`])
		if count, _ := daoAuth.Role.CtxDaoModel(ctx).FilterPri(roleIdArr).Count(); count != len(roleIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`)}})
			return
		}
	} */
	return
}

// 新增
func (logicThis *sAdmin) Create(ctx context.Context, data map[string]any) (id any, err error) {
	relIdOfRole := gconv.Uint(data[`rel_id_of_role`])
	delete(data, `rel_id_of_role`)
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAdmin.Admin.CtxDaoModel(ctx)

	relId := gconv.Uint(data[daoAdmin.Admin.Columns().RelId])
	switch gconv.String(data[daoAdmin.Admin.Columns().SceneId]) {
	case consts.SCENE_ID_PLATFORM:
		if relId > 0 {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
		data[daoAdmin.Admin.Columns().AdminType] = 0
	case consts.SCENE_ID_ORG:
		if relId == 0 {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
		orgInfo, _ := daoOrg.Org.CacheGetInfo(ctx, relId)
		if orgInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.org.org`)}})
			return
		}
		data[daoAdmin.Admin.Columns().AdminType] = orgInfo[daoOrg.Org.Columns().OrgType]
	}

	if _, ok := data[`role_id_arr`]; ok && len(gconv.Uints(data[`role_id_arr`])) > 0 {
		roleIdArr := gconv.Uints(data[`role_id_arr`])
		if count, _ := daoAuth.Role.CtxDaoModel(ctx).Filters(g.Map{daoAuth.Role.Columns().RelId: relIdOfRole, daoAuth.Role.Columns().SceneId: data[daoAdmin.Admin.Columns().SceneId], daoAuth.Role.Columns().RoleId: roleIdArr}).Count(); count != len(roleIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`)}})
			return
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAdmin) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	relIdOfRole := gconv.Uint(data[`rel_id_of_role`])
	delete(data, `rel_id_of_role`)
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAdmin.Admin.CtxDaoModel(ctx)

	sceneId := ``
	var list gdb.Result
	if _, ok := data[daoAuth.Role.Columns().RelId]; ok {
		if sceneId == `` {
			if _, ok := data[daoAdmin.Admin.Columns().SceneId]; ok {
				sceneId = gconv.String(data[daoAdmin.Admin.Columns().SceneId])
			} else {
				if len(list) == 0 {
					list, _ = daoAdmin.Admin.CtxDaoModel(ctx).Filters(filter).Fields(daoAdmin.Admin.Columns().AdminId, daoAdmin.Admin.Columns().SceneId).All()
					if len(list) == 0 {
						err = utils.NewErrorCode(ctx, 29999998, ``)
						return
					}
				}
				sceneIdSet := map[string]struct{}{}
				for _, info := range list {
					sceneIdSet[info[daoAdmin.Admin.Columns().SceneId].String()] = struct{}{}
				}
				if len(sceneIdSet) != 1 { //不同场景下rel_id只能对应一个，故必须同一场景下的管理员才能一起修改
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				}
				sceneId = list[0][daoAuth.Role.Columns().SceneId].String()
			}
		}
		relId := gconv.Uint(data[daoAdmin.Admin.Columns().RelId])
		switch sceneId {
		case consts.SCENE_ID_PLATFORM:
			if relId > 0 {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
			data[daoAdmin.Admin.Columns().AdminType] = 0
		case consts.SCENE_ID_ORG:
			if relId == 0 {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
			orgInfo, _ := daoOrg.Org.CacheGetInfo(ctx, relId)
			if orgInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.org.org`)}})
				return
			}
			data[daoAdmin.Admin.Columns().AdminType] = orgInfo[daoOrg.Org.Columns().OrgType]
		}
	}

	if _, ok := data[`role_id_arr`]; ok && len(gconv.Strings(data[`role_id_arr`])) > 0 {
		if sceneId == `` {
			if _, ok := data[daoAdmin.Admin.Columns().SceneId]; ok {
				sceneId = gconv.String(data[daoAdmin.Admin.Columns().SceneId])
			} else {
				if len(list) == 0 {
					list, _ = daoAdmin.Admin.CtxDaoModel(ctx).Filters(filter).Fields(daoAdmin.Admin.Columns().AdminId, daoAdmin.Admin.Columns().SceneId).All()
					if len(list) == 0 {
						err = utils.NewErrorCode(ctx, 29999998, ``)
						return
					}
				}
				sceneIdSet := map[string]struct{}{}
				for _, info := range list {
					sceneIdSet[info[daoAdmin.Admin.Columns().SceneId].String()] = struct{}{}
				}
				if len(sceneIdSet) != 1 { //角色所属场景只能设置一个，故必须同一场景下的管理员才能一起修改
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				}
				sceneId = list[0][daoAuth.Role.Columns().SceneId].String()
			}
		}
		roleIdArr := gconv.Strings(data[`role_id_arr`])
		if count, _ := daoAuth.Role.CtxDaoModel(ctx).Filters(g.Map{daoAuth.Role.Columns().RelId: relIdOfRole, daoAuth.Role.Columns().SceneId: sceneId, daoAuth.Role.Columns().RelId: roleIdArr}).Count(); count != len(roleIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`)}})
			return
		}
	}

	if len(list) > 0 {
		idArr := []uint{}
		for _, info := range list {
			idArr = append(idArr, info[daoAdmin.Admin.Columns().AdminId].Uint())
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
func (logicThis *sAdmin) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoAdmin.Admin.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
