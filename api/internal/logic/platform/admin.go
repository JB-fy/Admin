package logic

import (
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPlatformAdmin struct{}

func NewPlatformAdmin() *sPlatformAdmin {
	return &sPlatformAdmin{}
}

func init() {
	service.RegisterPlatformAdmin(NewPlatformAdmin())
}

// 新增
func (logicThis *sPlatformAdmin) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoPlatform.Admin

	_, okRoleIdArr := data[`roleIdArr`]
	if okRoleIdArr {
		roleIdArr := gconv.SliceUint(data[`roleIdArr`])
		sceneId, _ := daoAuth.Scene.HandlerCtx(ctx).Filter(daoAuth.Scene.Columns().SceneCode, `platform`).Value(daoAuth.Scene.PrimaryKey())
		filterTmp := g.Map{daoAuth.Role.PrimaryKey(): roleIdArr, daoAuth.Role.Columns().SceneId: sceneId}
		count, _ := daoAuth.Role.HandlerCtx(ctx).Filters(filterTmp).Count()
		if len(roleIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = daoThis.HandlerCtx(ctx).HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sPlatformAdmin) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	_, okRoleIdArr := data[`roleIdArr`]
	if okRoleIdArr {
		roleIdArr := gconv.SliceUint(data[`roleIdArr`])
		sceneId, _ := daoAuth.Scene.HandlerCtx(ctx).Filter(daoAuth.Scene.Columns().SceneCode, `platform`).Value(daoAuth.Scene.PrimaryKey())
		count, _ := daoAuth.Role.HandlerCtx(ctx).Filters(g.Map{daoAuth.Role.PrimaryKey(): roleIdArr, daoAuth.Role.Columns().SceneId: sceneId}).Count()
		if len(roleIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	row, err = daoHandlerThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPlatformAdmin) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoHandlerThis.HookSelect().DeleteAndGetAffected()
	return
}
