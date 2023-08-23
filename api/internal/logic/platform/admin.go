package logic

import (
	"api/internal/dao"
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
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
		roleIdArr := gconv.SliceInt(data[`roleIdArr`])
		sceneId, _ := daoAuth.Scene.ParseDbCtx(ctx).Where(`sceneCode`, `platform`).Value(`sceneId`)
		filterTmp := g.Map{`sceneId`: sceneId, `roleId`: roleIdArr}
		count, _ := daoAuth.Role.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(roleIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = dao.NewDaoHandler(ctx, &daoThis).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sPlatformAdmin) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	daoHandlerThis := dao.NewDaoHandler(ctx, &daoThis).Filter(filter)
	idArr, _ := daoHandlerThis.GetModel(true).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	_, okCheckPassword := data[`checkPassword`]
	if okCheckPassword {
		if len(idArr) > 1 { //不支持批量修改
			err = utils.NewErrorCode(ctx, 89999996, ``, map[string]interface{}{`errField`: `checkPassword`})
			return
		}
		oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr[0]).One()
		if gmd5.MustEncrypt(gconv.String(data[`checkPassword`])+oldInfo[daoThis.Columns().Salt].String()) != oldInfo[daoThis.Columns().Password].String() {
			err = utils.NewErrorCode(ctx, 39990003, ``)
			return
		}
		delete(data, `checkPassword`)
	}
	_, okRoleIdArr := data[`roleIdArr`]
	if okRoleIdArr {
		roleIdArr := gconv.SliceInt(data[`roleIdArr`])
		sceneId, _ := daoAuth.Scene.ParseDbCtx(ctx).Where(`sceneCode`, `platform`).Value(`sceneId`)
		count, _ := daoAuth.Role.ParseDbCtx(ctx).Where(g.Map{`sceneId`: sceneId, `roleId`: roleIdArr}).Count()
		if len(roleIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
		hookData[`roleIdArr`] = data[`roleIdArr`]
		delete(data, `roleIdArr`)
	}

	row, err = daoHandlerThis.Update(data).HookUpdate(hookData, gconv.SliceInt(idArr)...).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPlatformAdmin) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
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
