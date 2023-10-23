package logic

import (
	"api/internal/dao"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthAction struct{}

func NewAuthAction() *sAuthAction {
	return &sAuthAction{}
}

func init() {
	service.RegisterAuthAction(NewAuthAction())
}

// 新增
func (logicThis *sAuthAction) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Action
	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthAction) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	_, okSceneIdArr := data[`sceneIdArr`]
	if okSceneIdArr {
		hookData[`sceneIdArr`] = data[`sceneIdArr`]
		delete(data, `sceneIdArr`)
	}

	row, err = daoHandlerThis.Update(data).HookUpdate(hookData, gconv.SliceInt(idArr)...).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthAction) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoHandlerThis.HookDelete(gconv.SliceInt(idArr)...).GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}

// 判断操作权限
func (logicThis *sAuthAction) CheckAuth(ctx context.Context, actionCode string) (isAuth bool, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	//平台超级管理员，无权限限制
	if sceneInfo[daoAuth.Scene.Columns().SceneCode].String() == `platform` && loginInfo[`loginId`].Int() == g.Cfg().MustGet(ctx, `superPlatformAdminId`).Int() {
		isAuth = true
		return
	}

	filter := map[string]interface{}{
		daoAuth.Action.Columns().ActionCode: actionCode,
		`selfAction`: map[string]interface{}{
			`sceneCode`: sceneInfo[daoAuth.Scene.Columns().SceneCode],
			`sceneId`:   sceneInfo[daoAuth.Scene.PrimaryKey()],
			`loginId`:   loginInfo[`loginId`],
		},
	}
	count, err := dao.NewDaoHandler(ctx, &daoAuth.Action).Filter(filter).GetModel().Count()
	if count == 0 {
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	}
	isAuth = true
	return
}
