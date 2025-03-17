package auth

import (
	"api/api"
	apiAuth "api/api/platform/auth"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	defaultFieldOfTree []string
	allowField         []string
	noAuthField        []string
}

func NewMenu() *Menu {
	field := daoAuth.Menu.ColumnArr()
	defaultFieldOfList := []string{`id`, `label`, daoAuth.Scene.Columns().SceneName, `p_menu_name`, daoAuth.Menu.Columns().IsLeaf}
	defaultFieldOfInfo := []string{`id`, `label`}
	defaultFieldOfTree := []string{`id`, `label`}
	return &Menu{
		defaultFieldOfList: append(field, defaultFieldOfList...),
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		defaultFieldOfTree: append(field, defaultFieldOfTree...),
		allowField:         append(field, gset.NewStrSetFrom(defaultFieldOfList).Merge(gset.NewStrSetFrom(defaultFieldOfInfo)).Merge(gset.NewStrSetFrom(defaultFieldOfTree)).Slice()...),
		noAuthField:        []string{`id`, `label`, daoAuth.Menu.Columns().IsLeaf},
	}
}

// 列表
func (controllerThis *Menu) List(ctx context.Context, req *apiAuth.MenuListReq) (res *apiAuth.MenuListRes, err error) {
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authMenuRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoAuth.Menu.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiAuth.MenuListRes{Count: count, List: []apiAuth.MenuInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Menu) Info(ctx context.Context, req *apiAuth.MenuInfoReq) (res *apiAuth.MenuInfoRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `authMenuRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoAuth.Menu.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiAuth.MenuInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Menu) Create(ctx context.Context, req *apiAuth.MenuCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authMenuCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.AuthMenu().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Menu) Update(ctx context.Context, req *apiAuth.MenuUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`data`}})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authMenuUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AuthMenu().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Menu) Delete(ctx context.Context, req *apiAuth.MenuDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `authMenuDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.AuthMenu().Delete(ctx, filter)
	return
}

// 列表（树状）
func (controllerThis *Menu) Tree(ctx context.Context, req *apiAuth.MenuTreeReq) (res *apiAuth.MenuTreeRes, err error) {
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
		field = controllerThis.defaultFieldOfTree
	}
	field = append(field, `tree`)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authMenuRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	list, err := daoAuth.Menu.CtxDaoModel(ctx).Filters(filter).Fields(field...).ListPri()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, daoAuth.Menu.Columns().MenuId, daoAuth.Menu.Columns().Pid)

	res = &apiAuth.MenuTreeRes{Tree: []apiAuth.MenuInfo{}}
	gconv.Structs(tree, &res.Tree)
	return
}
