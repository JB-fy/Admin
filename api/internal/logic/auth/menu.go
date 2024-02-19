package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
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
		pid := gconv.Uint(data[daoThis.Columns().Pid])
		if pid > 0 {
			pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			sceneId := gconv.Uint(data[daoThis.Columns().SceneId])
			if pInfo[daoThis.Columns().SceneId].Uint() != sceneId {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	id, err = daoThis.HandlerCtx(ctx).Insert(data).GetModel().InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthMenu) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	_, okPid := data[daoThis.Columns().Pid]
	if okPid {
		pid := gconv.Uint(data[daoThis.Columns().Pid])
		if pid > 0 {
			pInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			oldList, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), daoHandlerThis.IdArr).All()
			for _, oldInfo := range oldList {
				if pid == oldInfo[daoThis.PrimaryKey()].Uint() { //父级不能是自身
					err = utils.NewErrorCode(ctx, 29999996, ``)
					return
				}
				if pid != oldInfo[daoThis.Columns().Pid].Uint() {
					sceneId := oldInfo[daoThis.Columns().SceneId].Uint()
					_, okSceneId := data[daoThis.Columns().SceneId]
					if okSceneId {
						sceneId = gconv.Uint(data[daoThis.Columns().SceneId])
					}
					if pInfo[daoThis.Columns().SceneId].Uint() != sceneId {
						err = utils.NewErrorCode(ctx, 89999998, ``)
						return
					}
					if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().IdPath].String(), `-`)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999995, ``)
						return
					}
				}
			}
		}
	}

	row, err = daoHandlerThis.Update(data).GetModel().UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthMenu) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	daoHandlerThis := daoThis.HandlerCtx(ctx).Filters(filter).SetIdArr()
	if len(daoHandlerThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	count, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.Columns().Pid, daoHandlerThis.IdArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ``)
		return
	}

	result, err := daoHandlerThis.Delete().GetModel().Delete()
	row, _ = result.RowsAffected()
	return
}
