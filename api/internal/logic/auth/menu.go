package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gregex"
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
		count, err = model.Handler(daoThis.ParseGroup([]string{`id`}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
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

	pInfo := gdb.Record{}
	_, okPid := data[`pid`]
	if okPid {
		pid := gconv.Int(data[`pid`])
		if pid > 0 {
			pInfo, _ = daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).Fields(`sceneId`, `pidPath`, `level`).One()
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

	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
		if len(match) > 0 {
			err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
			return
		}
		return
	}

	update := map[string]interface{}{
		`pidPath`: `0-` + gconv.String(id),
		`level`:   1,
	}
	if len(pInfo) > 0 {
		update = map[string]interface{}{
			`pidPath`: pInfo[`pidPath`].String() + `-` + gconv.String(id),
			`level`:   pInfo[`level`].Int() + 1,
		}
	}
	daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(update).Update()
	return
}

// 修改
func (logicThis *sMenu) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}

	updateList := map[int]map[string]interface{}{}
	updateChildList := map[string]map[string]interface{}{}
	_, okPid := data[`pid`]
	if okPid {
		pInfo := gdb.Record{}
		pid := gconv.Int(data[`pid`])
		if pid > 0 {
			pInfo, _ = daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), pid).Fields(`sceneId`, `pidPath`, `level`).One()
			if len(pInfo) == 0 {
				err = utils.NewErrorCode(ctx, 29999998, ``)
				return
			}
		}
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			if pid == oldInfo[daoThis.PrimaryKey()].Int() { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			if pid != oldInfo[`pid`].Int() {
				idKey := id.Int()
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
					if garray.NewStrArrayFrom(gstr.Split(pInfo[`pidPath`].String(), `-`)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999996, ``)
						return
					}
					updateList[idKey] = map[string]interface{}{
						`pidPath`: pInfo[`pidPath`].String() + `-` + oldInfo[daoThis.PrimaryKey()].String(),
						`level`:   pInfo[`level`].Int() + 1,
					}
				} else {
					updateList[idKey] = map[string]interface{}{
						`pidPath`: `0-` + oldInfo[daoThis.PrimaryKey()].String(),
						`level`:   1,
					}
				}
				updateChildList[oldInfo[`pidPath`].String()] = map[string]interface{}{
					`pidPathOfChild`: map[string]interface{}{
						`newVal`: updateList[idKey][`pidPath`],
						`oldVal`: oldInfo[`pidPath`],
					},
					`levelOfChild`: map[string]interface{}{
						`newVal`: updateList[idKey][`level`],
						`oldVal`: oldInfo[`level`],
					},
				}
			}
		}
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
		if len(match) > 0 {
			err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
			return
		}
		return
	}
	row, _ = result.RowsAffected()

	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ``)
		return
	}

	if len(updateList) > 0 {
		for id, update := range updateList {
			daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).Data(update).Update()
		}
		//修改pid时，更新所有子孙级的pidPath和level
		for pidPath, update := range updateChildList {
			daoThis.ParseDbCtx(ctx).WhereLike(`pidPath`, pidPath+`%`).Handler(daoThis.ParseUpdate(update)).Update()
		}
	}
	return
}

// 删除
func (logicThis *sMenu) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
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

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, _ = result.RowsAffected()

	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ``)
		return
	}
	return
}
