package logic

import (
	daoAuth "api/internal/model/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

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
func (logicThis *sAction) Count(ctx context.Context, filter map[string]interface{}) (count int, err error) {
	daoThis := daoAuth.Action
	joinTableArr := []string{}
	model := daoThis.ParseDbCtx(ctx)
	if len(filter) > 0 {
		model = model.Handler(daoThis.ParseFilter(filter, &joinTableArr))
	}
	if len(joinTableArr) > 0 {
		count, err = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr)).Distinct().Count(daoThis.PrimaryKey())
	} else {
		count, err = model.Count()
	}
	return
}

// 列表
func (logicThis *sAction) List(ctx context.Context, filter map[string]interface{}, field []string, order [][2]string, page int, limit int) (list gdb.Result, err error) {
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
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
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
		model = model.Handler(daoThis.ParseGroup([]string{"id"}, &joinTableArr))
	}
	info, err = model.One()
	return
}

// 创建
func (logicThis *sAction) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Action
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert([]map[string]interface{}{data})).InsertAndGetId()
	if err != nil {
		return
	}

	_, okSceneIdArr := data["sceneIdArr"]
	if okSceneIdArr {
		daoThis.SaveRelScene(ctx, gconv.SliceInt(data["sceneIdArr"]), int(id))
	}
	return
}

// 更新
func (logicThis *sAction) Update(ctx context.Context, data map[string]interface{}, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	_, okSceneIdArr := data["sceneIdArr"]
	if okSceneIdArr {
		idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
		for _, v := range idArr {
			daoThis.SaveRelScene(ctx, gconv.SliceInt(data["sceneIdArr"]), v.Int())
		}
		daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update() //有可能只改sceneIdArr
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseUpdate(data), daoThis.ParseFilter(filter, &[]string{})).Update()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	return
}

// 删除
func (logicThis *sAction) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Action
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Delete()
	if err != nil {
		return
	}
	row, err = result.RowsAffected()
	if row > 0 {
		daoAuth.ActionRelToScene.ParseDbCtx(ctx).Where("actionId", idArr).Delete()
	}
	return
}

// 判断操作权限
func (logicAction *sAction) CheckAuth(ctx context.Context, actionCode string) (isAuth bool, err error) {
	loginInfo := utils.GetCtxLoginInfo(ctx)
	sceneInfo := utils.GetCtxSceneInfo(ctx)
	sceneCode := sceneInfo["sceneCode"].String()
	filter := map[string]interface{}{
		"actionCode": actionCode,
	}
	filter["selfAction"] = map[string]interface{}{
		"sceneCode": sceneCode,
		"sceneId":   sceneInfo["sceneId"].Int(),
		"loginId":   loginInfo["adminId"].Int(),
	}

	switch sceneCode {
	case "platformAdmin":
		if loginInfo["adminId"].Int() == g.Cfg().MustGet(ctx, "superPlatformAdminId").Int() { //平台超级管理员，不再需要其他条件
			isAuth = true
			return
		}
		//filter["selfAction"].(map[string]interface{})["loginId"] = loginInfo["adminId"]
	}
	daoAction := daoAuth.Action
	count, err := daoAction.ParseDbCtx(ctx).Handler(daoAction.ParseFilter(filter, &[]string{})).Count()
	if count == 0 {
		err = utils.NewErrorCode(ctx, 39990002, "")
		return
	}
	isAuth = true
	return
}
