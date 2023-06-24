package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
	var pInfo gdb.Record
	pid := gconv.Int(data[`pid`])
	if pid > 0 {
		joinTableArr := []string{}
		field := []string{`pidPath`, `level`}
		filterTmp := g.Map{daoThis.PrimaryKey(): data[`pid`], `sceneId`: data[`sceneId`]}
		pInfo, _ = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filterTmp, &joinTableArr), daoThis.ParseField(field, &joinTableArr)).One()
		if len(pInfo) == 0 {
			err = utils.NewErrorCode(ctx, 29999998, ``)
			return
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
	if pid > 0 {
		update = map[string]interface{}{
			`pidPath`: pInfo[`pidPath`].String() + `-` + gconv.String(id),
			`level`:   pInfo[`level`].Int() + 1,
		}
	}
	daoThis.ParseDbCtx(ctx).Data(update).Where(daoThis.PrimaryKey(), id).Update()
	return
}

// 修改
func (logicThis *sMenu) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu

	_, okPid := data[`pid`]
	if okPid { //存在pid则只能一个个循环更新
		idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
		for _, id := range idArr {
			filterOne := map[string]interface{}{daoThis.PrimaryKey(): id}
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filterOne, &[]string{})).One()
			pid := gconv.Int(data[`pid`])
			if pid == oldInfo[daoThis.PrimaryKey()].Int() { //父级不能是自身
				err = utils.NewErrorCode(ctx, 29999997, ``)
				return
			}
			if pid != oldInfo[`pid`].Int() {
				if pid > 0 {
					joinTableArr := []string{}
					field := []string{`pidPath`, `level`}
					filterTmp := g.Map{daoThis.PrimaryKey(): data[`pid`], `sceneId`: oldInfo[`sceneId`]}
					_, okSceneId := data[`sceneId`]
					if okSceneId {
						filterTmp[`sceneId`] = data[`sceneId`]
					}
					pInfo, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filterTmp, &joinTableArr), daoThis.ParseField(field, &joinTableArr)).One()
					if len(pInfo) == 0 {
						err = utils.NewErrorCode(ctx, 29999998, ``)
						return
					}
					if garray.NewStrArrayFrom(gstr.Split(pInfo[`pidPath`].String(), `-`)).Contains(oldInfo[daoThis.PrimaryKey()].String()) { //父级不能是自身的子孙级
						err = utils.NewErrorCode(ctx, 29999996, ``)
						return
					}
					data[`pidPath`] = pInfo[`pidPath`].String() + `-` + oldInfo[daoThis.PrimaryKey()].String()
					data[`level`] = pInfo[`level`].Int() + 1
				} else {
					data[`pidPath`] = `0-` + oldInfo[daoThis.PrimaryKey()].String()
					data[`level`] = 1
				}
				//修改pid时，更新所有子孙级的pidPath和level
				update := map[string]interface{}{
					`pidPathOfChild`: map[string]interface{}{
						`newVal`: data[`pidPath`],
						`oldVal`: oldInfo[`pidPath`],
					},
					`levelOfChild`: map[string]interface{}{
						`newVal`: data[`level`],
						`oldVal`: oldInfo[`level`],
					},
				}
				filterPidPath := map[string]interface{}{`pidPath Like ?`: oldInfo[`pidPath`].String() + `%`}
				daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(update), daoThis.ParseFilter(filterPidPath, &[]string{})).Update()
			}
			_, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filterOne, &[]string{})).Update() //有可能只改menuIdArr或actionIdArr
			if err != nil {
				match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
				if len(match) > 0 {
					err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
					return
				}
				return
			}
		}
		return
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
	return
}

// 删除
func (logicThis *sMenu) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Menu
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
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
