package controller

import (
	"api/api"
	apiUser "api/api/platform/user"
	daoUser "api/internal/dao/user"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

// 列表
func (controllerThis *User) List(ctx context.Context, req *apiUser.UserListReq) (res *apiUser.UserListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]any{}
	}

	allowField := daoUser.User.ColumnArr().Slice()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `userUserRead`)
	if !isAuth {
		field = []string{`id`, `label`}
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoUser.User.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).HookSelect().Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiUser.UserListRes{Count: count, List: []apiUser.UserListItem{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *User) Info(ctx context.Context, req *apiUser.UserInfoReq) (res *apiUser.UserInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoUser.User.ColumnArr().Slice()
	allowField = append(allowField, `id`, `label`)
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]any{`id`: req.Id}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `userUserRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoUser.User.CtxDaoModel(ctx).Filters(filter).Fields(field...).HookSelect().InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiUser.UserInfoRes{}
	info.Struct(&res.Info)
	return
}

// 修改
func (controllerThis *User) Update(ctx context.Context, req *apiUser.UserUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	delete(data, `id_arr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]any{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `userUserUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.UserUser().Update(ctx, filter, data)
	return
}
