package controller

import (
	"api/api"
	apiPlatform "api/api/platform/platform"
	daoPlatform "api/internal/dao/platform"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Cron struct{}

func NewCron() *Cron {
	return &Cron{}
}

// 列表
func (controllerThis *Cron) List(ctx context.Context, req *apiPlatform.CronListReq) (res *apiPlatform.CronListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoPlatform.Cron.Columns()
	allowField := daoPlatform.Cron.ColumnArr()
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
	isAuth, _ := service.Action().CheckAuth(ctx, `platformCronLook`)
	if !isAuth {
		field = []string{`id`, `name`, columnsThis.CronId, columnsThis.CronName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Cron().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Cron().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}
	/* //不建议用这个返回，指定字段获取时，返回时其他字段也会返回，但都是空
	res = &apiPlatform.CronListRes{
		Count: count,
	}
	list.Structs(&res.List) */
	utils.HttpWriteJson(ctx, map[string]interface{}{
		`count`: count,
		`list`:  list,
	}, 0, ``)
	return
}

// 详情
func (controllerThis *Cron) Info(ctx context.Context, req *apiPlatform.CronInfoReq) (res *apiPlatform.CronInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoPlatform.Cron.ColumnArr()
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
	_, err = service.Action().CheckAuth(ctx, `platformCronLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Cron().Info(ctx, filter, field)
	if err != nil {
		return
	}
	/* //不建议用这个返回，指定字段获取时，返回时其他字段也会返回，但都是空
	res = &apiPlatform.CronInfoRes{}
	info.Struct(&res.Info) */
	utils.HttpWriteJson(ctx, map[string]interface{}{
		`info`: info,
	}, 0, ``)
	return
}

// 新增
func (controllerThis *Cron) Create(ctx context.Context, req *apiPlatform.CronCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCronCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Cron().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Cron) Update(ctx context.Context, req *apiPlatform.CronUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCronUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = service.Cron().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Cron) Delete(ctx context.Context, req *apiPlatform.CronDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `platformCronDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	err = service.Cron().Delete(ctx, filter)
	return
}
