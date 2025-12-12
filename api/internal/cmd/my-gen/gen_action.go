package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"context"

	"github.com/gogf/gf/v2/text/gstr"
)

// 自动生成操作权限
func genAction(ctx context.Context, option myGenOption, tpl *myGenTpl) {
	if !option.IsAuthAction {
		return
	}

	actionList := []map[string]any{}
	if option.IsList || option.IsInfo {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Read`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-查看`,
		})
	}
	if option.IsCreate {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Create`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-新增`,
		})
	}
	if option.IsUpdate {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Update`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-编辑`,
		})
	}
	if option.IsDelete {
		actionList = append(actionList, map[string]any{
			daoAuth.Action.Columns().ActionId:   gstr.CaseCamelLower(tpl.LogicStructName) + `Delete`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-删除`,
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
			daoAuth.ActionRelToScene.Columns().SceneId:  option.SceneInfo[daoAuth.Scene.Columns().SceneId],
		}).OnConflict(daoAuth.ActionRelToScene.Columns().ActionId, daoAuth.ActionRelToScene.Columns().SceneId).Save()
	}
}
