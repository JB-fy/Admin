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

type sAction struct{}

func NewAction() *sAction {
	return &sAction{}
}

func init() {
	service.RegisterAction(NewAction())
}

// 总数
func (logicThis *sAction) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Action
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
func (logicThis *sAction) List(ctx context.Context, filter map[string]interface{}, field []string, order []string, page int, limit int) (list gdb.Result, err error) {
	daoThis := daoAuth.Action
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
func (logicThis *sAction) Info(ctx context.Context, filter map[string]interface{}, field ...[]string) (info gdb.Record, err error) {
	daoThis := daoAuth.Action
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
func (logicThis *sAction) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Action
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	if err != nil {
		match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
		if len(match) > 0 {
			err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
			return
		}
		return
	}

	_, okSceneIdArr := data[`sceneIdArr`]
	if okSceneIdArr {
		sceneIdArr := gconv.SliceInt(data[`sceneIdArr`])
		daoThis.SaveRelScene(ctx, sceneIdArr, int(id))
	}
	return
}

// 修改
func (logicThis *sAction) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data)).Update()
	if err != nil {
		match, _ := gregex.MatchString(`1062.*Duplicate.*\.([^']*)'`, err.Error())
		if len(match) > 0 {
			err = utils.NewErrorCode(ctx, 29991062, ``, map[string]interface{}{`errField`: match[1]})
			return
		}
		return
	}
	row, _ = result.RowsAffected()

	_, okSceneIdArr := data[`sceneIdArr`]
	if okSceneIdArr {
		sceneIdArr := gconv.SliceInt(data[`sceneIdArr`])
		for _, id := range idArr {
			daoThis.SaveRelScene(ctx, sceneIdArr, id.Int())
		}
		row = 1 //有可能只改sceneIdArr
	}

	if row == 0 {
		err = utils.NewErrorCode(ctx, 99999999, ``)
		return
	}
	return
}

// 删除
func (logicThis *sAction) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999999, ``)
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
	daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where(daoThis.PrimaryKey(), idArr).Delete()
	return
}

// 判断操作权限
func (logicAction *sAction) CheckAuth(ctx context.Context, actionCode string) (isAuth bool, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo[`sceneCode`].String()
	filter := map[string]interface{}{
		`actionCode`: actionCode,
	}
	filter[`selfAction`] = map[string]interface{}{
		`sceneCode`: sceneCode,
		`sceneId`:   sceneInfo[`sceneId`].Int(),
		`loginId`:   loginInfo[`adminId`].Int(),
	}

	switch sceneCode {
	case `platform`:
		if loginInfo[`adminId`].Int() == g.Cfg().MustGet(ctx, `superPlatformAdminId`).Int() { //平台超级管理员，不再需要其他条件
			isAuth = true
			return
		}
		//filter[`selfAction`].(map[string]interface{})[`loginId`] = loginInfo[`adminId`]
	}
	daoAction := daoAuth.Action
	count, err := daoAction.ParseDbCtx(ctx).Handler(daoAction.ParseFilter(filter, &[]string{})).Count()
	if count == 0 {
		err = utils.NewErrorCode(ctx, 39990002, ``)
		return
	}
	isAuth = true
	return
}
