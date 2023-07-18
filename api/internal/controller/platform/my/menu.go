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

// 菜单树
func (controllerThis *Menu) MenuTree(ctx context.Context, req *apiMy.MenuTreeReq) (res *apiMy.MenuTreeRes, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	filter := map[string]interface{}{}
	filter[`selfMenu`] = map[string]interface{}{
		`sceneCode`: sceneInfo[`sceneCode`].String(),
		`sceneId`:   sceneInfo[`sceneId`].Int(),
		`loginId`:   loginInfo[`adminId`].Int(),
	}
	field := []string{`menuTree`, `showMenu`}

	list, err := service.Menu().List(ctx, filter, field, []string{}, 0, 0)
	if err != nil {
		return
	}
	tree := utils.Tree(list, 0, `menuId`, `pid`)
	res = &apiMy.MenuTreeRes{}
	tree.Structs(&res.Tree)
	return
}
