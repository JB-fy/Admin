package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"context"

	"github.com/gogf/gf/v2/text/gstr"
)

// 自动生成操作权限
func genAction(ctx context.Context, option myGenOption, tpl myGenTpl) {
	if !option.IsAuthAction {
		return
	}

	actionList := []map[string]interface{}{}
	if option.IsList || option.IsInfo {
		actionList = append(actionList, map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: gstr.CaseCamelLower(tpl.LogicStructName) + `Look`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-查看`,
		})
	}
	if option.IsCreate {
		actionList = append(actionList, map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: gstr.CaseCamelLower(tpl.LogicStructName) + `Create`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-新增`,
		})
	}
	if option.IsUpdate {
		actionList = append(actionList, map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: gstr.CaseCamelLower(tpl.LogicStructName) + `Update`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-编辑`,
		})
	}
	if option.IsDelete {
		actionList = append(actionList, map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: gstr.CaseCamelLower(tpl.LogicStructName) + `Delete`,
			daoAuth.Action.Columns().ActionName: gstr.Replace(option.CommonName, `/`, `-`) + `-删除`,
		})
	}

	for _, v := range actionList {
		id, err := daoAuth.Action.CtxDaoModel(ctx).HookInsert(v).InsertAndGetId()
		if err != nil { //报错就是操作权限已存在（唯一索引冲突）
			idVar, _ := daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.Columns().ActionCode, v[daoAuth.Action.Columns().ActionCode]).Value(daoAuth.Action.PrimaryKey())
			id = idVar.Int64()
			daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.Columns().ActionCode, v[daoAuth.Action.Columns().ActionCode]).HookUpdate(v).Update()
		}
		daoAuth.ActionRelToScene.CtxDaoModel(ctx).Data(map[string]interface{}{
			daoAuth.ActionRelToScene.Columns().ActionId: id,
			daoAuth.ActionRelToScene.Columns().SceneId:  option.SceneInfo[daoAuth.Scene.Columns().SceneId],
		}).Save()
	}
}
