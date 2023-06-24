package controller

import (
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Corn struct{}

func NewCorn() *Corn {
	return &Corn{}
}

// 列表
func (controllerThis *Corn) List(ctx context.Context, req *apiPlatform.CornListReq) (res *apiPlatform.CornListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoPlatform.Corn.Columns()
	allowField := daoPlatform.Corn.ColumnArr()
	allowField = append(allowField, `id`, `name`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `platformCornLook`)
	if !isAuth {
		field = []string{`id`, `name`, columnsThis.CornId, columnsThis.CornName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Corn().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Corn().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	res = &apiPlatform.CornListRes{
		Count: count,
	}
	list.Structs(&res.List)
	// utils.HttpSuccessJson(g.RequestFromCtx(ctx), map[string]interface{}{`count`: count, `list`: list}, 0)
	return
}

// 详情
func (controllerThis *Corn) Info(ctx context.Context, req *apiPlatform.CornInfoReq) (res *apiPlatform.CornInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoPlatform.Corn.ColumnArr()
	allowField = append(allowField, `id`, `name`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{`id`: req.Id}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCornLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Corn().Info(ctx, filter, field)
	if err != nil {
		return
	}
	res = &apiPlatform.CornInfoRes{}
	info.Struct(&res.Info)
	// utils.HttpSuccessJson(g.RequestFromCtx(ctx), map[string]interface{}{`info`: info}, 0)
	return
}

// 新增
func (controllerThis *Corn) Create(ctx context.Context, req *apiPlatform.CornCreateReq) (res *apiPlatform.CornCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCornCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Corn().Create(ctx, data)
	if err != nil {
		return
	}
	res = &apiPlatform.CornCreateRes{
		Id: id,
	}
	return
}

// 修改
func (controllerThis *Corn) Update(ctx context.Context, req *apiPlatform.CornUpdateReq) (res *apiPlatform.CornUpdateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCornUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Corn().Update(ctx, filter, data)
	if err != nil {
		return
	}
	return
}

// 删除
func (controllerThis *Corn) Delete(ctx context.Context, req *apiPlatform.CornDeleteReq) (res *apiPlatform.CornDeleteRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCornDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Corn().Delete(ctx, filter)
	if err != nil {
		return
	}
	return
}
