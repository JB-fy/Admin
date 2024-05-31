package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sAuthScene struct{}

func NewAuthScene() *sAuthScene {
	return &sAuthScene{}
}

func init() {
	service.RegisterAuthScene(NewAuthScene())
}

// 新增
func (logicThis *sAuthScene) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoThis := daoAuth.Scene
	daoModelThis := daoThis.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthScene) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoThis := daoAuth.Scene
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthScene) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoThis := daoAuth.Scene
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	if count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filter(daoAuth.Menu.Columns().SceneId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`), count, g.I18n().T(ctx, `name.auth.menu`)}})
		return
	}
	if count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filter(daoAuth.ActionRelToScene.Columns().SceneId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`), count, g.I18n().T(ctx, `name.auth.action`)}})
		return
	}
	if count, _ := daoAuth.Role.CtxDaoModel(ctx).Filter(daoAuth.Role.Columns().SceneId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`), count, g.I18n().T(ctx, `name.auth.role`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
