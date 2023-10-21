package controller

import (
	apiMy "api/api/platform/my"
	"api/internal/dao"
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
		`selfAction`: map[string]interface{}{
			`sceneCode`: sceneInfo[daoAuth.Scene.Columns().SceneCode],
			`sceneId`:   sceneInfo[daoAuth.Scene.PrimaryKey()],
			`loginId`:   loginInfo[`loginId`],
		},
	}
	list, err := dao.NewDaoHandler(ctx, &daoAuth.Action).Field(field).Filter(filter).JoinGroupByPrimaryKey().GetModel().All()
	if err != nil {
		return
	}
	res = &apiMy.ActionListRes{}
	list.Structs(&res.List)
	return
}
