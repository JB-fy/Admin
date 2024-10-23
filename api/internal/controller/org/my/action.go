package my

import (
	apiMy "api/api/org/my"
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

	field := []string{`id`, `label`, daoAuth.Action.Columns().ActionCode}
	filter := map[string]any{
		`self_action`: map[string]any{
			`scene_code`: sceneInfo[daoAuth.Scene.Columns().SceneCode],
			`login_id`:   loginInfo[`login_id`],
		},
	}
	list, err := daoAuth.Action.CtxDaoModel(ctx).Filters(filter).Fields(field...).ListPri()
	if err != nil {
		return
	}
	res = &apiMy.ActionListRes{List: []apiMy.ActionListItem{}}
	list.Structs(&res.List)
	return
}
