package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"context"
	"errors"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
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
func (logicAction *sAction) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoAction := daoAuth.Action
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoAction.ParseGroup([]string{"id"}, &joinTableArr)).Distinct().Count(daoAction.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicAction *sAction) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, offset int, limit int) (list gdb.Result, err error) {
	daoAction := daoAuth.Action
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoAction.ParseGroup([]string{"id"}, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	list, err = model.All()
	return
}

// 详情
func (logicAction *sAction) Info(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string) (info gdb.Record, err error) {
	daoAction := daoAuth.Action
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(field) > 0 {
		model = model.Handler(daoAction.ParseField(field, &joinTableArr))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		model = model.Handler(daoAction.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicAction *sAction) Create(ctx context.Context, data []map[string]interface{}) (id int64, err error) {
	daoAction := daoAuth.Action
	model := daoAction.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoAction.ParseInsert(data))
	}
	if len(data) == 1 {
		id, err = model.InsertAndGetId()
		return
	}
	result, err := model.Insert()
	if err != nil {
		return
	}
	id, err = result.RowsAffected()
	return
}

// 更新
func (logicAction *sAction) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoAction := daoAuth.Action
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(data) > 0 {
		model = model.Handler(daoAction.ParseUpdate(data))
	}
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	result, err := model.Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicAction *sAction) Delete(ctx context.Context, filter map[string]interface{}, order [][2]string, offset int, limit int) (row int64, err error) {
	daoAction := daoAuth.Action
	joinTableArr := []string{}
	model := daoAction.Ctx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoAction.ParseFilter(filter, &joinTableArr))
	}
	if len(order) > 0 {
		model = model.Handler(daoAction.ParseOrder(order, &joinTableArr))
	}
	if limit > 0 {
		model = model.Offset(offset).Limit(limit)
	}
	result, err := model.Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 判断操作权限
func (logicAction *sAction) CheckAuth(ctx context.Context, actionCode string, sceneCode string) (isAuth bool, err error) {
	infoTmp := ctx.Value(sceneCode + "Info")
	info := gvar.New(infoTmp).Map()
	filter := map[string]interface{}{
		"actionCode": actionCode,
	}
	filter["selfAction"] = map[string]interface{}{
		"sceneCode": sceneCode,
		"loginId":   info["adminId"],
	}

	switch sceneCode {
	case "platformAdmin":
		if gconv.Int(info["adminId"]) == g.Cfg().MustGet(ctx, "superPlatformAdminId").Int() { //平台超级管理员，不再需要其他条件
			isAuth = true
			return
		}
		//filter["selfAction"].(map[string]interface{})["loginId"] = info["adminId"]
	}
	daoAction := daoAuth.Action
	count, err := daoAction.Ctx(ctx).Handler(daoAction.ParseFilter(filter, &[]string{})).Count()
	if count == 0 {
		err = errors.New("39990002")
		return
	}
	isAuth = true
	return
}
