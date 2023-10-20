package controller

import (
	apiMy "api/api/platform/my"
	"api/internal/dao"
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
	adminDao := daoPlatform.Admin
	sceneDao := daoAuth.Scene
	sceneColumns := sceneDao.Columns()
	filter := map[string]interface{}{}
	filter[`selfAction`] = map[string]interface{}{
		`sceneCode`: sceneInfo[sceneColumns.SceneCode].String(),
		`sceneId`:   sceneInfo[sceneDao.PrimaryKey()].Int(),
		`loginId`:   loginInfo[adminDao.PrimaryKey()].Int(),
	}
	columns := daoAuth.Action.Columns()
	field := []string{`id`, `label`, columns.ActionId, columns.ActionName}

	list, err := dao.NewDaoHandler(ctx, &daoAuth.Action).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().All()
	if err != nil {
		return
	}
	res = &apiMy.ActionListRes{}
	list.Structs(&res.List)
	return
}
