package auth

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthMenu struct{}

func NewAuthMenu() *sAuthMenu {
	return &sAuthMenu{}
}

func init() {
	service.RegisterAuthMenu(NewAuthMenu())
}

// 验证数据（create和update共用）
func (logicThis *sAuthMenu) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[daoAuth.Menu.Columns().SceneId]; ok && gconv.String(data[daoAuth.Menu.Columns().SceneId]) != `` {
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).FilterPri(data[daoAuth.Menu.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sAuthMenu) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Menu.CtxDaoModel(ctx)

	if _, ok := data[daoAuth.Menu.Columns().Pid]; ok && gconv.Uint(data[daoAuth.Menu.Columns().Pid]) > 0 {
		pInfo, _ := daoModelThis.CloneNew().FilterPri(data[daoAuth.Menu.Columns().Pid]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.pid`)}})
			return
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthMenu) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Menu.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if _, ok := data[daoAuth.Menu.Columns().Pid]; ok && gconv.Uint(data[daoAuth.Menu.Columns().Pid]) > 0 {
		if garray.NewArrayFrom(gconv.SliceAny(gconv.SliceUint(daoModelThis.IdArr))).Contains(gconv.Uint(data[daoAuth.Menu.Columns().Pid])) {
			err = utils.NewErrorCode(ctx, 29999996, ``)
			return
		}
		pInfo, _ := daoModelThis.CloneNew().FilterPri(data[daoAuth.Menu.Columns().Pid]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.pid`)}})
			return
		}
		for _, id := range daoModelThis.IdArr {
			if garray.NewStrArrayFrom(gstr.Split(pInfo[daoAuth.Menu.Columns().IdPath].String(), `-`)).Contains(gconv.String(id)) {
				err = utils.NewErrorCode(ctx, 29999995, ``)
				return
			}
		}
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthMenu) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoAuth.Menu.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoModelThis.CloneNew().Filter(daoAuth.Menu.Columns().Pid, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ``)
		return
	}

	if count, _ := daoAuth.RoleRelToMenu.CtxDaoModel(ctx).Filter(daoAuth.RoleRelToMenu.Columns().MenuId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.menu`), count, g.I18n().T(ctx, `name.auth.roleRelToMenu`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
