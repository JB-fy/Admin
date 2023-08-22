package logic

import (
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
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

// 总数
func (logicThis *sPlatformAdmin) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoPlatform.Admin
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + `.` + daoThis.PrimaryKey()).Distinct().Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	count, err = model.Count()
	return
}

// 列表
func (logicThis *sPlatformAdmin) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoPlatform.Admin
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(field) > 0 {
		model = model.Handler(daoThis.ParseField(field, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoThis.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	list, err = model.Page(page, limit).All()
	return
}

// 详情
func (logicThis *sPlatformAdmin) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoPlatform.Admin
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Group(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	info, err = model.One()
	if err != nil {
		return
	}
	if len(info) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	return
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

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *sPlatformAdmin) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
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

	model := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data))
	if len(hookData) > 0 {
		model = model.Hook(daoThis.HookUpdate(hookData, gconv.SliceInt(idArr)...))
	}
	row, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sPlatformAdmin) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoPlatform.Admin
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	row, _ = result.RowsAffected()
	return
}
