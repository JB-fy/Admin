package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

// 自动生成操作权限
func genAction(ctx context.Context, sceneId uint, actionCode string, actionName string) {
	actionName = gstr.Replace(actionName, `/`, `-`)

	idVar, _ := daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.Columns().ActionCode, actionCode).Value(daoAuth.Action.PrimaryKey())
	id := idVar.Int64()
	if id == 0 {
		id, _ = daoAuth.Action.CtxDaoModel(ctx).HookInsert(map[string]interface{}{
			daoAuth.Action.Columns().ActionCode: actionCode,
			daoAuth.Action.Columns().ActionName: actionName,
		}).InsertAndGetId()
	} else {
		daoAuth.Action.CtxDaoModel(ctx).Filter(daoAuth.Action.PrimaryKey(), id).HookUpdate(g.Map{daoAuth.Action.Columns().ActionName: actionName}).Update()
	}
	daoAuth.ActionRelToScene.CtxDaoModel(ctx).Data(map[string]interface{}{
		daoAuth.ActionRelToScene.Columns().ActionId: id,
		daoAuth.ActionRelToScene.Columns().SceneId:  sceneId,
	}).Save()
}
