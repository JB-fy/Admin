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
	daoModelThis := daoThis.CtxDaoModel(ctx)

	if _, ok := data[`menu_id_arr`]; ok {
		menuIdArr := gconv.SliceUint(data[`menu_id_arr`])
		filterTmp := g.Map{daoAuth.Menu.PrimaryKey(): menuIdArr, daoAuth.Menu.Columns().SceneId: data[`scene_id`]}
		count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(filterTmp).Count()
		if len(menuIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}
	if _, ok := data[`action_id_arr`]; ok {
		actionIdArr := gconv.SliceUint(data[`action_id_arr`])
		filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: data[`scene_id`]}
		count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(filterTmp).Count()
		if len(actionIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if _, ok := data[`menu_id_arr`]; ok {
		menuIdArr := gconv.SliceUint(data[`menu_id_arr`])
		for _, id := range daoModelThis.IdArr {
			oldInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{daoAuth.Menu.PrimaryKey(): menuIdArr, daoAuth.Menu.Columns().SceneId: oldInfo[`scene_id`]}
			if _, ok := data[`scene_id`]; ok {
				filterTmp[daoAuth.Menu.Columns().SceneId] = data[`scene_id`]
			}
			count, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(filterTmp).Count()
			if len(menuIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	if _, ok := data[`action_id_arr`]; ok {
		actionIdArr := gconv.SliceUint(data[`action_id_arr`])
		for _, id := range daoModelThis.IdArr {
			oldInfo, _ := daoModelThis.CloneNew().Filter(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{daoAuth.ActionRelToScene.Columns().ActionId: actionIdArr, daoAuth.ActionRelToScene.Columns().SceneId: oldInfo[`scene_id`]}
			if _, ok := data[`scene_id`]; ok {
				filterTmp[daoAuth.ActionRelToScene.Columns().SceneId] = data[`scene_id`]
			}
			count, _ := daoAuth.ActionRelToScene.CtxDaoModel(ctx).Filters(filterTmp).Count()
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
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
