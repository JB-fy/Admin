package controller

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

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表
func (controllerThis *Menu) List(ctx context.Context, req *apiAuth.MenuListReq) (res *apiAuth.MenuListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoAuth.Menu.Columns()
	allowField := daoAuth.Menu.ColumnArr()
	allowField = append(allowField, `id`, `label`, `sceneName`, `pMenuName`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authMenuLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.MenuId, columnsThis.MenuName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.AuthMenu().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.AuthMenu().List(ctx, filter, field, order, page, limit)
	if err != nil {
		return
	}

	res = &apiAuth.MenuListRes{
		Count: count,
	}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Menu) Info(ctx context.Context, req *apiAuth.MenuInfoReq) (res *apiAuth.MenuInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoAuth.Menu.ColumnArr()
	allowField = append(allowField, `id`, `label`)
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
	_, err = service.AuthAction().CheckAuth(ctx, `authMenuLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.AuthMenu().Info(ctx, filter, field)
	if err != nil {
		return
	}

	res = &apiAuth.MenuInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Menu) Create(ctx context.Context, req *apiAuth.MenuCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.MapDeep(req)
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
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
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
	filter := map[string]interface{}{`id`: req.IdArr}
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

// 树状列表
func (controllerThis *Menu) Tree(ctx context.Context, req *apiAuth.MenuTreeReq) (res *apiAuth.MenuTreeRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}

	columnsThis := daoAuth.Menu.Columns()
	allowField := daoAuth.Menu.ColumnArr()
	allowField = append(allowField, `id`, `label`, `sceneName`, `pMenuName`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `authMenuLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.MenuId, columnsThis.MenuName}
	}
	/**--------权限验证 结束--------**/

	filter[`isStop`] = 0          //补充条件
	field = append(field, `tree`) //补充字段（树状列表所需）

	list, err := service.AuthMenu().List(ctx, filter, field, []string{}, 0, 0)
	if err != nil {
		return
	}
	tree := utils.Tree(list, 0, `menuId`, `pid`)

	utils.HttpWriteJson(ctx, map[string]interface{}{
		`tree`: tree,
	}, 0, ``)
	return
}
