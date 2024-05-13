package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
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
	daoModelThis := daoThis.CtxDaoModel(ctx)

	if _, ok := data[daoThis.Columns().Pid]; ok {
		pid := gconv.Uint(data[daoThis.Columns().Pid])
		if pid > 0 {
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().MenuId, pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			if pInfo[daoThis.Columns().SceneId].Uint() != gconv.Uint(data[daoThis.Columns().SceneId]) {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthMenu) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if _, ok := data[daoThis.Columns().Pid]; ok {
		pid := gconv.Uint(data[daoThis.Columns().Pid])
		if pid > 0 {
			if garray.NewArrayFrom(gconv.SliceAny(daoModelThis.IdArr)).Contains(pid) {
				err = utils.NewErrorCode(ctx, 29999996, ``)
				return
			}
			pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().MenuId, pid).One()
			if pInfo.IsEmpty() {
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			for _, id := range daoModelThis.IdArr {
				if garray.NewStrArrayFrom(gstr.Split(pInfo[daoThis.Columns().IdPath].String(), `-`)).Contains(gconv.String(id)) {
					err = utils.NewErrorCode(ctx, 29999995, ``)
					return
				}
			}
			if _, ok := data[daoThis.Columns().SceneId]; ok {
				if pInfo[daoThis.Columns().SceneId].Uint() != gconv.Uint(data[daoThis.Columns().SceneId]) {
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				}
			} else {
				count, _ := daoModelThis.CloneNew().Filters(g.Map{
					`id`:                      daoModelThis.IdArr,
					daoThis.Columns().SceneId: pInfo[daoThis.Columns().SceneId],
				}).Count()
				if count != len(daoModelThis.IdArr) {
					err = utils.NewErrorCode(ctx, 89999998, ``)
					return
				}
			}
		}
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthMenu) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	count, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().Pid, daoModelThis.IdArr).Count()
	if count > 0 {
		err = utils.NewErrorCode(ctx, 29999994, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
