package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sMenu struct{}

func NewMenu() *sMenu {
	return &sMenu{}
}

func init() {
	service.RegisterMenu(NewMenu())
}

// 总数
func (logicThis *sMenu) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Menu
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr)).Distinct().Fields(daoThis.Table() + `.` + daoThis.PrimaryKey())
	}
	count, err = model.Count()
	return
}

// 列表
func (logicThis *sMenu) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Menu
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
		model = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset((page - 1) * limit).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicThis *sMenu) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Menu
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	if len(field) > 0 && len(field[0]) > 0 {
		model = model.Handler(daoThis.ParseField(field[0], &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr))
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
func (logicThis *sMenu) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Menu

	_, okPid := data[`pid`]
	if okPid {
		pid := gconv.Int(data[`pid`])
		if pid > 0 {
			pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).Fields(`sceneId`, `idPath`, `level`).One()
			if len(pInfo) == 0 {
				err = utils.NewErrorCode(ctx, 29999998, ``)
				return
			}
			sceneId := gconv.Int(data[`sceneId`])
			if pInfo[`sceneId`].Int() != sceneId {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *sMenu) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}
	hookData := map[string]interface{}{}

	_, okPid := data[`pid`]
	if okPid {
		pInfo := gdb.Record{}
		pid := gconv.Int(data[`pid`])
		if pid > 0 {
			pInfo, _ = daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if len(pInfo) == 0 {
				err = utils.NewErrorCode(ctx, 29999998, ``)
				return
			}
		}
		updateChildIdPathAndLevelList := []map[string]interface{}{}
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			if pid == oldInfo[daoThis.PrimaryKey()].Int() { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			if pid != oldInfo[`pid`].Int() {
				pIdPath := `0`
				pLevel := 0
				if pid > 0 {
					sceneId := oldInfo[`sceneId`].Int()
					_, okSceneId := data[`sceneId`]
					if okSceneId {
						sceneId = gconv.Int(data[`sceneId`])
					}
					if pInfo[`sceneId`].Int() != sceneId {
						err = utils.NewErrorCode(ctx, 89999998, ``)
						return
					}
					if garray.NewStrArrayFrom(gstr.Split(pInfo[`idPath`].String(), `-`)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999996, ``)
						return
					}
					pIdPath = pInfo[`idPath`].String()
					pLevel = pInfo[`level`].Int()
				}
				updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
					`newIdPath`: pIdPath + `-` + id.String(),
					`oldIdPath`: oldInfo[`idPath`],
					`newLevel`:  pLevel + 1,
					`oldLevel`:  oldInfo[`level`],
				})
			}
		}

		if len(updateChildIdPathAndLevelList) > 0 {
			hookData[`updateChildIdPathAndLevelList`] = updateChildIdPathAndLevelList
		}
	}

	model := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data))
	if len(hookData) > 0 {
		model = model.Hook(daoThis.HookUpdate(hookData, gconv.SliceInt(idArr)...))
	}
	_, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sMenu) Delete(ctx context.Context, filter map[string]interface{}) (err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}

	count, _ := daoThis.ParseDbCtx(ctx).Where(`pid`, idArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999995, ``)
		return
	}

	_, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	return
}
