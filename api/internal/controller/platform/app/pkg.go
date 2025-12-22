package app

import (
	"api/api"
	apiApp "api/api/platform/app"
	daoApp "api/internal/dao/app"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"slices"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Pkg struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewPkg() *Pkg {
	field := append(daoApp.Pkg.ColumnArr(), `id`, `label`)
	appendFieldOfList := []string{daoApp.App.Columns().AppName}
	appendFieldOfInfo := []string{}
	return &Pkg{
		defaultFieldOfList: append(slices.Clone(field), appendFieldOfList...),
		defaultFieldOfInfo: append(slices.Clone(field), appendFieldOfInfo...),
		allowField:         append(slices.Clone(field), gset.NewStrSetFrom(slices.Concat(appendFieldOfList, appendFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Pkg) List(ctx context.Context, req *apiApp.PkgListReq) (res *apiApp.PkgListRes, err error) {
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `appPkgRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoApp.Pkg.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiApp.PkgListRes{Count: count, List: []apiApp.PkgInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Pkg) Info(ctx context.Context, req *apiApp.PkgInfoReq) (res *apiApp.PkgInfoRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `appPkgRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoApp.Pkg.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiApp.PkgInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Pkg) Create(ctx context.Context, req *apiApp.PkgCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appPkgCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AppPkg().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Pkg) Update(ctx context.Context, req *apiApp.PkgUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`data`}})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appPkgUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppPkg().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Pkg) Delete(ctx context.Context, req *apiApp.PkgDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `appPkgDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AppPkg().Delete(ctx, filter)
	return
}
