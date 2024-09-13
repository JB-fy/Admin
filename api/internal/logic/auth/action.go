package logic

import (
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthAction struct{}

func NewAuthAction() *sAuthAction {
	return &sAuthAction{}
}

func init() {
	service.RegisterAuthAction(NewAuthAction())
}

// 验证数据（create和update共用）
func (logicThis *sAuthAction) verifyData(ctx context.Context, data map[string]any) (err error) {
	if _, ok := data[`scene_id_arr`]; ok && len(gconv.SliceUint(data[`scene_id_arr`])) > 0 {
		sceneIdArr := gconv.SliceUint(data[`scene_id_arr`])
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneId, sceneIdArr).Count(); count != len(sceneIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sAuthAction) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Action.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthAction) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Action.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthAction) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoAuth.Action.CtxDaoModel(ctx)

	daoModelThis.SetIdArr(filter)
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if count, _ := daoAuth.RoleRelToAction.CtxDaoModel(ctx).Filter(daoAuth.RoleRelToAction.Columns().ActionId, daoModelThis.IdArr).Count(); count > 0 {
		err = utils.NewErrorCode(ctx, 30009999, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.action`), count, g.I18n().T(ctx, `name.auth.roleRelToAction`)}})
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}

// 判断操作权限
func (logicThis *sAuthAction) CheckAuth(ctx context.Context, actionCodeArr ...string) (isAuth bool, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	if sceneInfo[daoAuth.Scene.Columns().SceneCode].String() == `platform` && loginInfo[daoPlatform.Admin.Columns().IsSuper].Uint() == 1 { //平台超级管理员，无权限限制
		isAuth = true
		return
	}

	if len(actionCodeArr) == 0 {
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	}

	filter := map[string]any{
		daoAuth.Action.Columns().ActionCode: actionCodeArr,
		`self_action`: map[string]any{
			`scene_code`: sceneInfo[daoAuth.Scene.Columns().SceneCode],
			`login_id`:   loginInfo[`login_id`],
		},
	}
	count, err := daoAuth.Action.CtxDaoModel(ctx).Filters(filter).Count()
	if count != len(actionCodeArr) {
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	}
	isAuth = true
	return
}
