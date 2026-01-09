package auth

import (
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"api/internal/utils/jbctx"
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
	if _, ok := data[`scene_id_arr`]; ok && len(gconv.Strings(data[`scene_id_arr`])) > 0 {
		sceneIdArr := gconv.Strings(data[`scene_id_arr`])
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).FilterPri(sceneIdArr).Count(); count != len(sceneIdArr) {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`i18nValues`: []any{g.I18n().T(ctx, `name.auth.scene`)}})
			return
		}
	}
	return
}

// 新增
func (logicThis *sAuthAction) Create(ctx context.Context, data map[string]any) (id any, err error) {
	if err = logicThis.verifyData(ctx, data); err != nil {
		return
	}
	daoModelThis := daoAuth.Action.CtxDaoModel(ctx)

	id = data[daoAuth.Action.Columns().ActionId]
	_, err = daoModelThis.HookInsert(data).Insert()
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
func (logicThis *sAuthAction) CheckAuth(ctx context.Context, actionIdArr ...string) (isAuth bool, err error) {
	loginInfo := jbctx.GetLoginInfo(ctx)
	sceneInfo := jbctx.GetSceneInfo(ctx)
	if sceneInfo[daoAuth.Scene.Columns().SceneId].String() == `platform` && loginInfo[daoPlatform.Admin.Columns().IsSuper].Uint8() == 1 { //平台超级管理员，无权限限制
		isAuth = true
		return
	}

	if len(actionIdArr) == 0 {
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	}

	/* // 表数据很小，无需这样做，且会导致数据修改无法立即生效。确实需要减轻数据库压力时可以使用
	actionIdArrOfSelf, err := daoAuth.Action.CacheGetActionIdArrOfSelf(ctx, sceneInfo[daoAuth.Scene.Columns().SceneId].String(), loginInfo[`login_id`])
	if err != nil {
		return
	}
	actionIdArrOfSelf = gset.NewStrSetFrom(actionIdArrOfSelf).Intersect(gset.NewStrSetFrom(actionIdArr)).Slice() //交集
	if actionIdArrLen := len(actionIdArr); actionIdArrLen == 0 || actionIdArrLen != len(actionIdArrOfSelf) {     //因为是判断操作权限，所以actionIdArr和actionIdArrOfSelf必须一样，否则必定缺少权限
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	} */
	filter := map[string]any{
		`self_action`: map[string]any{
			`scene_id`:            sceneInfo[daoAuth.Scene.Columns().SceneId],
			`login_id`:            loginInfo[`login_id`],
			`check_action_id_arr`: actionIdArr,
		},
	}
	count, err := daoAuth.Action.CtxDaoModel(ctx).Filters(filter).Count()
	if count != len(actionIdArr) {
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	}
	isAuth = true
	return
}
