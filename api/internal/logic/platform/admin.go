package platform

import (
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPlatformAdmin struct{}

func NewPlatformAdmin() *sPlatformAdmin {
	return &sPlatformAdmin{}
}

func init() {
	service.RegisterPlatformAdmin(NewPlatformAdmin())
}

// 验证数据（create和update共用）
func (logicThis *sPlatformAdmin) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[`role_id_arr`]; ok && len(gconv.Uints(data[`role_id_arr`])) > 0 {
		roleIdArr := gconv.Uints(data[`role_id_arr`])
		if count, _ := daoAuth.Role.CtxDaoModel(ctx).FilterPri(roleIdArr).Filter(daoAuth.Role.Columns().SceneId, `platform`).Count(); count != len(roleIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.role`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sPlatformAdmin) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoPlatform.Admin.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sPlatformAdmin) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoPlatform.Admin.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPlatformAdmin) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoPlatform.Admin.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
