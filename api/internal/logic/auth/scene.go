package logic

import (
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type sAuthScene struct{}

func NewAuthScene() *sAuthScene {
	return &sAuthScene{}
}

func init() {
	service.RegisterAuthScene(NewAuthScene())
}

// 新增
func (logicThis *sAuthScene) Create(ctx context.Context, data map[string]interface{}) (id int64, err error) {
	daoThis := daoAuth.Scene
	id, err = daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseInsert(data)).InsertAndGetId()
	return
}

// 修改
func (logicThis *sAuthScene) Update(ctx context.Context, filter map[string]interface{}, data map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Scene
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}
	hookData := map[string]interface{}{}

	model := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{}), daoThis.ParseUpdate(data))
	if len(hookData) > 0 {
		model = model.Hook(daoThis.HookUpdate(hookData, gconv.SliceInt(idArr)...))
	}
	row, err = model.UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sAuthScene) Delete(ctx context.Context, filter map[string]interface{}) (row int64, err error) {
	daoThis := daoAuth.Scene
	idArr, _ := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Array(daoThis.PrimaryKey())
	if len(idArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	result, err := daoThis.ParseDbCtx(ctx).Handler(daoThis.ParseFilter(filter, &[]string{})).Hook(daoThis.HookDelete(gconv.SliceInt(idArr)...)).Delete()
	row, _ = result.RowsAffected()
	return
}
