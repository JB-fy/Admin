package controller

import (
	apiMy "api/api/platform/my"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 树状列表
func (controllerThis *Menu) Tree(ctx context.Context, req *apiMy.MenuTreeReq) (res *apiMy.MenuTreeRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	filter := map[string]interface{}{}
	filter[`selfMenu`] = map[string]interface{}{
		`sceneCode`: sceneInfo[`sceneCode`].String(),
		`sceneId`:   sceneInfo[`sceneId`].Int(),
		`loginId`:   loginInfo[`adminId`].Int(),
	}
	field := []string{`id`, `label`, `tree`, `showMenu`}

	list, err := service.Menu().List(ctx, filter, field, []string{}, 0, 0)
	if err != nil {
		return
	}
	tree := utils.Tree(list, 0, `menuId`, `pid`)

	utils.HttpWriteJson(ctx, map[string]interface{}{
		`tree`: tree,
	}, 0, ``)
	return
}
