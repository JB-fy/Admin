package controller

import (
	apiMy "api/api/platform/my"
	"api/internal/dao"
	daoAuth "api/internal/dao/auth"
	"api/internal/utils"
	"context"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表（树状）
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

	list, err := dao.NewDaoHandler(ctx, &daoAuth.Menu).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().All()
	if err != nil {
		return
	}
	tree := utils.Tree(list, 0, `menuId`, `pid`)

	utils.HttpWriteJson(ctx, map[string]interface{}{
		`tree`: tree,
	}, 0, ``)
	return
}
