package controller

import (
	apiAuth "api/api/platform/auth"
	daoAuth "api/internal/dao/auth"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type Menu struct{}

func NewMenu() *Menu {
	return &Menu{}
}

// 列表
func (controllerThis *Menu) List(ctx context.Context, req *apiAuth.MenuListReq) (res *apiAuth.MenuListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter)
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoAuth.Menu.Columns()
	allowField := daoAuth.Menu.ColumnArr()
	allowField = append(allowField, `id`, `name`, `sceneName`, `pMenuName`)
	// allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.Action().CheckAuth(ctx, `authMenuLook`)
	if !isAuth {
		field = []string{`id`, `name`, columnsThis.MenuId, columnsThis.MenuName}
	}
	/**--------权限验证 结束--------**/

	count, err := service.Menu().Count(ctx, filter)
	if err != nil {
		return
	}
	list, err := service.Menu().List(ctx, filter, field, order, page, limit)
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
	allowField = append(allowField, `id`, `name`)
	//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
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
	_, err = service.Action().CheckAuth(ctx, `authMenuLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := service.Menu().Info(ctx, filter, field)
	if err != nil {
		return
	}
	res = &apiAuth.MenuInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Menu) Create(ctx context.Context, req *apiAuth.MenuCreateReq) (res *apiAuth.MenuCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req)
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authMenuCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.Menu().Create(ctx, data)
	if err != nil {
		return
	}
	res = &apiAuth.MenuCreateRes{
		Id: id,
	}
	return
}

// 修改
func (controllerThis *Menu) Update(ctx context.Context, req *apiAuth.MenuUpdateReq) (res *apiAuth.MenuUpdateRes, err error) {
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
	_, err = service.Action().CheckAuth(ctx, `authMenuUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Menu().Update(ctx, filter, data)
	if err != nil {
		return
	}
	return
}

// 删除
func (controllerThis *Menu) Delete(ctx context.Context, req *apiAuth.MenuDeleteReq) (res *apiAuth.MenuDeleteRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.Action().CheckAuth(ctx, `authMenuDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Menu().Delete(ctx, filter)
	if err != nil {
		return
	}
	return
}

// 菜单树
func (controllerThis *Menu) Tree(r *ghttp.Request) {
	/**--------参数处理 开始--------**/
	var param *apiAuth.MenuTreeReq
	err := r.Parse(&param)
	if err != nil {
		r.Response.Writeln(err.Error())
		return
	}
	filter := gconv.Map(param.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	/**--------参数处理 结束--------**/

	sceneCode := utils.GetCtxSceneCode(r.GetCtx())
	switch sceneCode {
	case `platform`:
		/**--------权限验证 开始--------**/
		isAuth, _ := service.Action().CheckAuth(r.GetCtx(), `authMenuLook`)
		allowField := []string{`id`, `name`, `menuId`, `menuName`}
		if isAuth {
			allowField = daoAuth.Menu.ColumnArr()
			allowField = append(allowField, `id`, `name`, `sceneName`, `pMenuName`)
			//allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{`password`})).Slice() //移除敏感字段
		}
		field := allowField
		if len(param.Field) > 0 {
			field = gset.NewStrSetFrom(param.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
			if len(field) == 0 {
				field = allowField
			}
		}

		filter[`isStop`] = 0              //补充条件
		field = append(field, `menuTree`) //补充字段（菜单树所需）
		/**--------权限验证 结束--------**/

		list, err := service.Menu().List(r.GetCtx(), filter, field, []string{}, 0, 0)
		if err != nil {
			utils.HttpFailJson(r, err)
			return
		}
		tree := utils.Tree(list, 0, `menuId`, `pid`)
		utils.HttpSuccessJson(r, map[string]interface{}{`tree`: tree}, 0)
	}
}
