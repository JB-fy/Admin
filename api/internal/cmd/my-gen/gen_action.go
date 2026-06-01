package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"context"

	"github.com/gogf/gf/v2/text/gstr"
)

// 自动生成操作权限
func genAction(ctx context.Context, tpl *myGenTpl) {
	if !tpl.Option.IsAuthAction {
		return
	}

	actionList := []map[string]any{}
	if tpl.Option.IsList || tpl.Option.IsInfo {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Read`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(tpl.Option.CommonName, `/`, `-`) + `-查看`,
		})
	}
	if tpl.Option.IsCreate {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Create`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(tpl.Option.CommonName, `/`, `-`) + `-新增`,
		})
	}
	if tpl.Option.IsUpdate {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Update`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(tpl.Option.CommonName, `/`, `-`) + `-编辑`,
		})
	}
	if tpl.Option.IsDelete {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Delete`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(tpl.Option.CommonName, `/`, `-`) + `-删除`,
		})
	}

	for _, v := range actionList {
		daoAuth.Action.CtxDaoModel(ctx).HookInsert(v).InsertIgnore()
		/* _, err := daoAuth.Action.CtxDaoModel(ctx).HookInsert(v).InsertIgnore()
		if err != nil { //报错就是操作权限已存在（主键冲突）
			// daoAuth.Action.CtxDaoModel(ctx).SetIdArr(v[daoAuth.Action.Columns().ActionId]).HookUpdate(v).Update() //已存在不再更新
		} */
		daoAuth.ActionRelToScene.CtxDaoModel(ctx).Data(map[string]any{
			daoAuth.ActionRelToScene.Columns().ActionId: v[daoAuth.Action.Columns().ActionId],
			daoAuth.ActionRelToScene.Columns().SceneId:  tpl.Option.SceneInfo[daoAuth.Scene.Columns().SceneId],
		}).OnConflict(daoAuth.ActionRelToScene.Columns().ActionId, daoAuth.ActionRelToScene.Columns().SceneId).Save()
	}
}
