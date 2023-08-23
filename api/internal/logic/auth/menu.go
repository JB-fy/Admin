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

type sAuthMenu struct{}

func NewAuthMenu() *sAuthMenu {
	return &sAuthMenu{}
}

func init() {
	service.RegisterAuthMenu(NewAuthMenu())
}

// 新增
func (logicThis *sAuthMenu) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Menu

	_, okPid := data[daoThis.Columns().Pid]
	if okPid {
		pid := gconv.Int(data[daoThis.Columns().Pid])
		if pid > 0 {
			pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).Fields(daoThis.Columns().SceneId, daoThis.Columns().IdPath, daoThis.Columns().Level).One()
			if len(pInfo) == 0 {
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			sceneId := gconv.Int(data[daoThis.Columns().SceneId])
			if pInfo[daoThis.Columns().SceneId].Int() != sceneId {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthMenu) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	_, okPid := data[daoThis.Columns().Pid]
	if okPid {
		pInfo := gdb.Record{}
		pid := gconv.Int(data[daoThis.Columns().Pid])
		if pid > 0 {
			pInfo, _ = daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if len(pInfo) == 0 {
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
		}
		updateChildIdPathAndLevelList := []map[string]interface{}{}
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			if pid == oldInfo[daoThis.PrimaryKey()].Int() { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999996, ``)
				return
			}
			if pid != oldInfo[daoThis.Columns().Pid].Int() {
				pIdPath := `0`
				pLevel := 0
				if pid > 0 {
					sceneId := oldInfo[daoThis.Columns().SceneId].Int()
					_, okSceneId := data[daoThis.Columns().SceneId]
					if okSceneId {
						sceneId = gconv.Int(data[daoThis.Columns().SceneId])
					}
					if pInfo[daoThis.Columns().SceneId].Int() != sceneId {
						err = utils.NewErrorCode(ctx, 89999998, ``)
						return
					}
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().IdPath].String(), `-`)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ``)
						return
					}
					pIdPath = pInfo[daoThis.Columns().IdPath].String()
					pLevel = pInfo[daoThis.Columns().Level].Int()
				}
				updateChildIdPathAndLevelList = append(updateChildIdPathAndLevelList, map[string]interface{}{
					`newIdPath`: pIdPath + `-` + id.String(),
					`oldIdPath`: oldInfo[daoThis.Columns().IdPath],
					`newLevel`:  pLevel + 1,
					`oldLevel`:  oldInfo[daoThis.Columns().Level],
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
	row, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthMenu) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	count, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.Columns().Pid, idArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ``)
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	row, _ = result.RowsAffected()
	return
}
