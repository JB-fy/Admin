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
		menuIdArr := gconv.SliceInt(data[`menuIdArr`])
		filterTmp := g.Map{`sceneId`: data[`sceneId`], `menuId`: menuIdArr}
		count, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(menuIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data[`actionIdArr`])
		filterTmp := g.Map{`sceneId`: data[`sceneId`], `actionId`: actionIdArr}
		count, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(actionIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthRole) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data[`menuIdArr`])
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{`sceneId`: oldInfo[`sceneId`], `menuId`: menuIdArr}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[`sceneId`] = data[`sceneId`]
			}
			count, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(filterTmp).Count()
			if len(menuIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
		hookData[`menuIdArr`] = data[`menuIdArr`]
		delete(data, `menuIdArr`)
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data[`actionIdArr`])
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{`sceneId`: oldInfo[`sceneId`], `actionId`: actionIdArr}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[`sceneId`] = data[`sceneId`]
			}
			count, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(filterTmp).Count()
			if len(actionIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
		hookData[`actionIdArr`] = data[`actionIdArr`]
		delete(data, `actionIdArr`)
	}

	model := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data))
	if len(hookData) > 0 {
		model = model.Hook(daoThis.HookUpdate(hookData, gconv.SliceInt(idArr)...))
	}
	row, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthRole) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	row, _ = result.RowsAffected()
	return
}
