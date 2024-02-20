package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthRole struct{}

func NewAuthRole() *sAuthRole {
	return &sAuthRole{}
}

func init() {
	service.RegisterAuthRole(NewAuthRole())
}

// 新增
func (logicThis *sAuthRole) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Role
	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceUint(data[`menuIdArr`])
		filterTmp := g.Map{daoAuth.Menu.PrimaryKey(): menuIdArr, daoAuth.Menu.Columns().SceneId: data[`sceneId`]}
		count, _ := daoAuth.Menu.HandlerCtx(ctx).Filters(filterTmp).Count()
		if len(menuIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceUint(data[`actionIdArr`])
		filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: data[`sceneId`]}
		count, _ := daoAuth.ActionRelToScene.HandlerCtx(ctx).Filters(filterTmp).Count()
		if len(actionIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = daoThis.HandlerCtx(ctx).HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	daoModelThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceUint(data[`menuIdArr`])
		for _, id := range daoModelThis.IdArr {
			oldInfo, _ := daoThis.HandlerCtx(ctx).Filter(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{daoAuth.Menu.PrimaryKey(): menuIdArr, daoAuth.Menu.Columns().SceneId: oldInfo[`sceneId`]}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[daoAuth.Menu.Columns().SceneId] = data[`sceneId`]
			}
			count, _ := daoAuth.Menu.HandlerCtx(ctx).Filters(filterTmp).Count()
			if len(menuIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceUint(data[`actionIdArr`])
		for _, id := range daoModelThis.IdArr {
			oldInfo, _ := daoThis.HandlerCtx(ctx).Filter(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: oldInfo[`sceneId`]}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[daoAuth.ActionRelToScene.Columns().SceneId] = data[`sceneId`]
			}
			count, _ := daoAuth.ActionRelToScene.HandlerCtx(ctx).Filters(filterTmp).Count()
			if len(actionIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthRole) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	daoModelThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookSelect().DeleteAndGetAffected()
	return
}
