package my

import (
	apiMy "api/api/platform/my"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
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

	/* // 表数据很小，无需这样做，且会导致数据修改无法立即生效。确实需要减轻数据库压力时可以使用
	var list gdb.Result
	if loginInfo[daoPlatform.Admin.Columns().IsSuper].Bool() {
		list, err = daoAuth.Action.CacheGetListOfNoStop(ctx, sceneInfo[daoAuth.Scene.Columns().SceneId].String())
	} else {
		list, err = daoAuth.Action.CacheGetListOfSelf(ctx, sceneInfo[daoAuth.Scene.Columns().SceneId].String(), loginInfo[`login_id`])
	} */
	field := []string{`id`, `label`}
	filter := map[string]any{
		`self_action`: map[string]any{
			`scene_id`: sceneInfo[daoAuth.Scene.Columns().SceneId],
			`login_id`: loginInfo[`login_id`],
			`is_super`: loginInfo[daoPlatform.Admin.Columns().IsSuper].Uint8(),
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
