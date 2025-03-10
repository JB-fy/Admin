package org

import (
	"api/api"
	apiOrg "api/api/platform/org"
	daoOrg "api/internal/dao/org"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Org struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewOrg() *Org {
	field := daoOrg.Org.ColumnArr()
	defaultFieldOfList := []string{`id`, `label`}
	defaultFieldOfInfo := []string{`id`, `label`}
	return &Org{
		defaultFieldOfList: append(field, defaultFieldOfList...),
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, gset.NewStrSetFrom(defaultFieldOfList).Merge(gset.NewStrSetFrom(defaultFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Org) List(ctx context.Context, req *apiOrg.OrgListReq) (res *apiOrg.OrgListRes, err error) {
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `orgRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoOrg.Org.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiOrg.OrgListRes{Count: count, List: []apiOrg.OrgInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Org) Info(ctx context.Context, req *apiOrg.OrgInfoReq) (res *apiOrg.OrgInfoRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `orgRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoOrg.Org.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiOrg.OrgInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Org) Create(ctx context.Context, req *apiOrg.OrgCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Org().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Org) Update(ctx context.Context, req *apiOrg.OrgUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`data`}})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Org().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Org) Delete(ctx context.Context, req *apiOrg.OrgDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `orgDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Org().Delete(ctx, filter)
	return
}
