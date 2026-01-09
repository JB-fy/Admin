package my

import (
	apiMy "api/api/org/my"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils/jbctx"
	"context"
)

type Action struct{}

func NewAction() *Action {
	return &Action{}
}

// 操作列表
func (controllerThis *Action) List(ctx context.Context, req *apiMy.ActionListReq) (res *apiMy.ActionListRes, err error) {
	loginInfo := jbctx.GetLoginInfo(ctx)
	sceneInfo := jbctx.GetSceneInfo(ctx)

	/* // 表数据很小，无需这样做，且会导致数据修改无法立即生效。确实需要减轻数据库压力时可以使用
	list, err := daoAuth.Action.CacheGetListOfSelf(ctx, sceneInfo[daoAuth.Scene.Columns().SceneId].String(), loginInfo[`login_id`]) */
	field := []string{`id`, `label`}
	filter := map[string]any{
		`self_action`: map[string]any{
			`scene_id`: sceneInfo[daoAuth.Scene.Columns().SceneId],
			`login_id`: loginInfo[`login_id`],
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
