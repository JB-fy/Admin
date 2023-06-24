package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/util/gconv"
)

type sRole struct{}

func NewRole() *sRole {
	return &sRole{}
}

func init() {
	service.RegisterRole(NewRole())
}

// 总数
func (logicThis *sRole) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Role
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
func (logicThis *sRole) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Role
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
func (logicThis *sRole) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Role
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
func (logicThis *sRole) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Role
	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data[`menuIdArr`])
		filterTmp := g.Map{`sceneId`: data[`sceneId`], `menuId`: menuIdArr}
		count, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(menuIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
			return
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data[`actionIdArr`])
		filterTmp := g.Map{`sceneId`: data[`sceneId`], `actionId`: actionIdArr}
		count, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(filterTmp).Count()
		if len(actionIdArr) != count {
			err = utils.NewErrorCode(ctx, 89999998, ``)
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

	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data[`menuIdArr`])
		daoThis.SaveRelMenu(ctx, menuIdArr, int(id))
	}
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data[`actionIdArr`])
		daoThis.SaveRelAction(ctx, actionIdArr, int(id))
	}
	return
}

// 修改
func (logicThis *sRole) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())

	_, okMenuIdArr := data[`menuIdArr`]
	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data[`menuIdArr`])
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{`sceneId`: oldInfo[`sceneId`], `menuId`: menuIdArr}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[`sceneId`] = data[`sceneId`]
			}
			count, _ := daoAuth.Menu.ParseDbCtx(ctx).Where(filterTmp).Count()
			if len(menuIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
			}
		}
	}
	_, okActionIdArr := data[`actionIdArr`]
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data[`actionIdArr`])
		for _, id := range idArr {
			oldInfo, _ := daoThis.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), id).One()
			filterTmp := g.Map{`sceneId`: oldInfo[`sceneId`], `actionId`: actionIdArr}
			_, okSceneId := data[`sceneId`]
			if okSceneId {
				filterTmp[`sceneId`] = data[`sceneId`]
			}
			count, _ := daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(filterTmp).Count()
			if len(actionIdArr) != count {
				err = utils.NewErrorCode(ctx, 89999998, ``)
				return
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
	if okMenuIdArr {
		menuIdArr := gconv.SliceInt(data[`menuIdArr`])
		for _, id := range idArr {
			daoThis.SaveRelMenu(ctx, menuIdArr, id.Int())
		}
		row = 1 //有可能只改menuIdArr
	}
	if okActionIdArr {
		actionIdArr := gconv.SliceInt(data[`actionIdArr`])
		for _, id := range idArr {
			daoThis.SaveRelAction(ctx, actionIdArr, id.Int())
		}
		row = 1 //有可能只改actionIdArr
	}

	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ``)
		return
	}
	return
}

// 删除
func (logicThis *sRole) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Role
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, _ = result.RowsAffected()

	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ``)
		return
	}
	daoAuth.RoleRelToMenu.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
	daoAuth.RoleRelToAction.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
	daoAuth.RoleRelOfPlatformAdmin.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
	return
}
