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
func (logicThis *sAuthMenu) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoThis := daoAuth.Menu
	daoModelThis := daoThis.CtxDaoModel(ctx)

	if _, ok := data[daoThis.Columns().SceneId]; ok && gconv.Uint(data[daoThis.Columns().SceneId]) > 0 {
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneId, data[daoThis.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999998, ``)
			return
		}
	}

	if _, ok := data[daoThis.Columns().Pid]; ok && gconv.Uint(data[daoThis.Columns().Pid]) > 0 {
		pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().MenuId, data[daoThis.Columns().Pid]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`errValues`: []any{g.I18n().T(ctx, `name.pid`)}})
			return
		}
		if pInfo[daoThis.Columns().SceneId].Uint() != gconv.Uint(data[daoThis.Columns().SceneId]) {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthMenu) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoThis := daoAuth.Menu
	daoModelThis := daoThis.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	if _, ok := data[daoThis.Columns().SceneId]; ok && gconv.Uint(data[daoThis.Columns().SceneId]) > 0 {
		if count, _ := daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneId, data[daoThis.Columns().SceneId]).Count(); count == 0 {
			err = utils.NewErrorCode(ctx, 29999998, ``)
			return
		}
	}

	if _, ok := data[daoThis.Columns().Pid]; ok && gconv.Uint(data[daoThis.Columns().Pid]) > 0 {
		if garray.NewArrayFrom(gconv.SliceAny(gconv.SliceUint(daoModelThis.IdArr))).Contains(gconv.Uint(data[daoThis.Columns().Pid])) {
			err = utils.NewErrorCode(ctx, 29999996, ``)
			return
		}
		pInfo, _ := daoModelThis.CloneNew().Filter(daoThis.Columns().MenuId, data[daoThis.Columns().Pid]).One()
		if pInfo.IsEmpty() {
			err = utils.NewErrorCode(ctx, 29999997, ``, g.Map{`errValues`: []any{g.I18n().T(ctx, `name.pid`)}})
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
			if count, _ := daoModelThis.CloneNew().Filters(g.Map{daoThis.Columns().SceneId: pInfo[daoThis.Columns().SceneId], `id`: daoModelThis.IdArr}).Count(); count != len(daoModelThis.IdArr) {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthMenu) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
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
