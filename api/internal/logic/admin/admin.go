package admin

import (
	daoAdmin "api/internal/dao/admin"
	daoAuth "api/internal/dao/auth"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

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
	if _, ok := data[daoAdmin.Admin.Columns().OrgId]; ok {
		if orgId := gconv.Uint(data[daoAdmin.Admin.Columns().OrgId]); orgId > 0 {
			info, _ := daoOrg.Org.CtxDaoModel(ctx).FilterPri(data[daoAdmin.Admin.Columns().OrgId]).One() //daoOrg.Org.CacheGetInfo(ctx, gconv.Uint(data[daoAdmin.Admin.Columns().OrgId]))
			if len(info) == 0 {
				err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.org.org`)}})
				return
			}
			data[daoAdmin.Admin.Columns().AdminType] = info[daoOrg.Org.Columns().OrgType]
		} else {
			data[daoAdmin.Admin.Columns().AdminType] = 0
		}
	}

	if _, ok := data[`role_id_arr`]; ok && len(gconv.Uints(data[`role_id_arr`])) > 0 {
		roleIdArr := gconv.Uints(data[`role_id_arr`])
		// TODO Filter(daoAuth.Role.Columns().SceneId, `platform 或 org`)
		if count, _ := daoAuth.Role.CtxDaoModel(ctx).FilterPri(roleIdArr).Count(); count != len(roleIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sAdmin) Create(ctx context.Context, data map[string]any) (id any, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAdmin.Admin.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAdmin) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAdmin.Admin.CtxDaoModel(ctx)

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
