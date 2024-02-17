package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
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
	id, err = daoThis.HandlerCtx(ctx).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthAction) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filter(filter, true)
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	_, okSceneIdArr := data[`sceneIdArr`]
	if okSceneIdArr {
		daoHandlerThis.AfterUpdate[`sceneIdArr`] = data[`sceneIdArr`]
		delete(data, `sceneIdArr`)
	}

	row, err = daoHandlerThis.Update(data).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthAction) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filter(filter, true)
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoHandlerThis.Delete().GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}

// 判断操作权限
func (logicThis *sAuthAction) CheckAuth(ctx context.Context, actionCode string) (isAuth bool, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	//平台超级管理员，无权限限制
	if sceneInfo[daoAuth.Scene.Columns().SceneCode].String() == `platform` && loginInfo[`loginId`].Uint() == g.Cfg().MustGet(ctx, `superPlatformAdminId`).Uint() {
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
	count, err := daoAuth.Action.HandlerCtx(ctx).Filter(filter).GetModel().Count()
	if count == 0 {
		err = utils.NewErrorCode(ctx, 39999996, ``)
		return
	}
	isAuth = true
	return
}
