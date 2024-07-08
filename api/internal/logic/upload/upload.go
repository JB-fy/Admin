package logic

import (
	daoUpload "api/internal/dao/upload"
	"api/internal/service"
	"api/internal/utils"
	"context"
)

type sUpload struct{}

func NewUpload() *sUpload {
	return &sUpload{}
}

func init() {
	service.RegisterUpload(NewUpload())
}

// 新增
func (logicThis *sUpload) Create(ctx context.Context, data map[string]any) (id int64, err error) {
	daoModelThis := daoUpload.Upload.CtxDaoModel(ctx)

	id, err = daoModelThis.HookInsert(data).InsertAndGetId()
	return
}

// 修改
func (logicThis *sUpload) Update(ctx context.Context, filter map[string]any, data map[string]any) (row int64, err error) {
	daoModelThis := daoUpload.Upload.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookUpdate(data).UpdateAndGetAffected()
	return
}

// 删除
func (logicThis *sUpload) Delete(ctx context.Context, filter map[string]any) (row int64, err error) {
	daoModelThis := daoUpload.Upload.CtxDaoModel(ctx)

	daoModelThis.Filters(filter).SetIdArr()
	if len(daoModelThis.IdArr) == 0 {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	row, err = daoModelThis.HookDelete().DeleteAndGetAffected()
	return
}
