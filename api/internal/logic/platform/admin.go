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
	daoModelThis := daoThis.CtxDaoModel(ctx)

	if _, ok := data[`role_id_arr`]; ok {
		roleIdArr := gconv.SliceUint(data[`role_id_arr`])
		roleIdArrLen := len(roleIdArr)
		if roleIdArrLen > 0 {
			sceneId, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneCode, `platform`).Value(daoAuth.Scene.Columns().SceneId)
			filterTmp := g.Map{daoAuth.Role.Columns().RoleId: roleIdArr, daoAuth.Role.Columns().SceneId: sceneId}
			count, _ := daoAuth.Role.CtxDaoModel(ctx).Filters(filterTmp).Count()
			if roleIdArrLen != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sPlatformAdmin) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if _, ok := data[`role_id_arr`]; ok {
		roleIdArr := gconv.SliceUint(data[`role_id_arr`])
		roleIdArrLen := len(roleIdArr)
		if roleIdArrLen > 0 {
			sceneId, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneCode, `platform`).Value(daoAuth.Scene.Columns().SceneId)
			count, _ := daoAuth.Role.CtxDaoModel(ctx).Filters(g.Map{daoAuth.Role.Columns().RoleId: roleIdArr, daoAuth.Role.Columns().SceneId: sceneId}).Count()
			if roleIdArrLen != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPlatformAdmin) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
