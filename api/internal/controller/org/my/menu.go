package my

import (
	apiMy "api/api/org/my"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表（树状）
func (controllerThis *Menu) Tree(ctx context.Context, req *apiMy.MenuTreeReq) (res *apiMy.MenuTreeRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)

	field := []string{`id`, `label`, `tree`, `show_menu`}
	filter := map[string]any{
		`self_menu`: map[string]any{
			`scene_id`: sceneInfo[daoAuth.Scene.Columns().SceneId],
			`login_id`: loginInfo[`login_id`],
		},
	}
	list, err := daoAuth.Menu.CtxDaoModel(ctx).Filters(filter).Fields(field...).ListPri()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, daoAuth.Menu.Columns().MenuId, daoAuth.Menu.Columns().Pid)

	res = &apiMy.MenuTreeRes{Tree: []apiMy.MenuTreeItem{}}
	gconv.Structs(tree, &res.Tree)
	return
}
