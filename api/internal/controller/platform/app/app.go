package controller

import (
	"api/api"
	apiApp "api/api/platform/app"
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type App struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewApp() *App {
	field := daoApp.App.ColumnArr().Slice()
	defaultFieldOfList := []string{`id`, `label`}
	defaultFieldOfInfo := []string{`id`, `label`}
	return &App{
		defaultFieldOfList: append(field, defaultFieldOfList...),
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, gset.NewStrSetFrom(defaultFieldOfList).Merge(gset.NewStrSetFrom(defaultFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *App) List(ctx context.Context, req *apiApp.AppListReq) (res *apiApp.AppListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]any{}
	}

	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfList
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoApp.App.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiApp.AppListRes{Count: count, List: []apiApp.AppInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *App) Info(ctx context.Context, req *apiApp.AppInfoReq) (res *apiApp.AppInfoRes, err error) {
	/**--------参数处理 开始--------**/
	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfInfo
	}
	filter := map[string]any{`id`: req.Id}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoApp.App.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApp.AppInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *App) Create(ctx context.Context, req *apiApp.AppCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.App().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *App) Update(ctx context.Context, req *apiApp.AppUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`data`}})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.App().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *App) Delete(ctx context.Context, req *apiApp.AppDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.App().Delete(ctx, filter)
	return
}
