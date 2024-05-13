package controller

import (
	apiMy "api/api/platform/my"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"context"
)

type Action struct{}

func NewAction() *Action {
	return &Action{}
}

// 操作列表
func (controllerThis *Action) List(ctx context.Context, req *apiMy.ActionListReq) (res *apiMy.ActionListRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)

	actionColumns := daoAuth.Action.Columns()
	field := []string{`id`, `label`, actionColumns.ActionId, actionColumns.ActionName}
	filter := map[string]interface{}{
		`self_action`: map[string]interface{}{
			`scene_code`: sceneInfo[daoAuth.Scene.Columns().SceneCode],
			`scene_id`:   sceneInfo[daoAuth.Scene.Columns().SceneId],
			`login_id`:   loginInfo[`login_id`],
		},
	}
	list, err := daoAuth.Action.CtxDaoModel(ctx).Filters(filter).Fields(field...).HookSelect().ListPri()
	if err != nil {
		return
	}
	res = &apiMy.ActionListRes{List: []apiMy.ActionListItem{}}
	list.Structs(&res.List)
	return
}
