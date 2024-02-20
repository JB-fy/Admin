package controller

import (
	apiMy "api/api/platform/my"
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

	field := []string{`id`, `label`, `tree`, `showMenu`}
	filter := map[string]interface{}{
		`selfMenu`: map[string]interface{}{
			`sceneCode`: sceneInfo[daoAuth.Scene.Columns().SceneCode],
			`sceneId`:   sceneInfo[daoAuth.Scene.PrimaryKey()],
			`loginId`:   loginInfo[`loginId`],
		},
	}
	list, err := daoAuth.Menu.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().ListPri()
	if err != nil {
		return
	}
	menuColumns := daoAuth.Menu.Columns()
	tree := utils.Tree(list.List(), 0, menuColumns.MenuId, menuColumns.Pid)

	res = &apiMy.MenuTreeRes{}
	gconv.Structs(tree, &res.Tree)
	return
}
