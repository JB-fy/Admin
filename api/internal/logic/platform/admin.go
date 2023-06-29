package logic

import (
	daoAuth "api/internal/dao/auth"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdmin struct{}

func NewAdmin() *sAdmin {
	return &sAdmin{}
}

func init() {
	service.RegisterAdmin(NewAdmin())
}

// 总数
func (logicThis *sAdmin) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
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
func (logicThis *sAdmin) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
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
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicThis *sAdmin) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
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
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}
	return
}

// 新增
func (logicThis *sAdmin) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
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
func (logicThis *sAdmin) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error) {
	daoThis := daoPlatform.Admin
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}
	hookData := map[string]interface{}{}

	_, okCheckPassword := data[`checkPassword`]
	if okCheckPassword {
		if len(idArr) > 1 { //该字段只支持单个用户更新
			err = utils.NewErrorCode(ctx, 89999996, ``, map[string]interface{}{`errField`: `checkPassword`})
			return
		}
		password, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr[0]).Value(`password`)
		if gconv.String(data[`checkPassword`]) != password.String() {
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
	_, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAdmin) Delete(ctx context.Context, filter map[string]interface{}) (err error) {
	daoThis := daoPlatform.Admin
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}

	_, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	return
}
