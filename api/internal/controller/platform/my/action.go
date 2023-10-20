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
	filter := map[string]interface{}{}
	filter[`selfAction`] = map[string]interface{}{
		`sceneCode`: sceneInfo[`sceneCode`].String(),
		`sceneId`:   sceneInfo[`sceneId`].Int(),
		`loginId`:   loginInfo[`adminId`].Int(),
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
