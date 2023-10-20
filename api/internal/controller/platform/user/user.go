package controller

import (
	"api/api"
	apiUser "api/api/platform/user"
	"api/internal/dao"
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
	filter := gconv.MapDeep(req.Filter)
	if filter == nil {
		filter = map[string]interface{}{}
	}
	order := []string{req.Sort}
	page := req.Page
	limit := req.Limit

	columnsThis := daoUser.User.Columns()
	allowField := daoUser.User.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password, columnsThis.Salt})).Slice() //移除敏感字段
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `userLook`)
	if !isAuth {
		field = []string{`id`, `label`, columnsThis.Phone, columnsThis.UserId}
	}
	/**--------权限验证 结束--------**/

	daoHandlerThis := dao.NewDaoHandler(ctx, &daoUser.User)
	daoHandlerThis.Filter(filter)
	count, err := daoHandlerThis.Count()
	if err != nil {
		return
	}
	list, err := daoHandlerThis.Field(field).Order(order).JoinGroupByPrimaryKey().GetModel().Page(page, limit).All()
	if err != nil {
		return
	}

	res = &apiUser.UserListRes{
		Count: count,
	}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *User) Info(ctx context.Context, req *apiUser.UserInfoReq) (res *apiUser.UserInfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := daoUser.User.ColumnArr()
	allowField = append(allowField, `id`, `label`)
	columnsThis := daoUser.User.Columns()
	allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{columnsThis.Password, columnsThis.Salt})).Slice() //移除敏感字段
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
	_, err = service.AuthAction().CheckAuth(ctx, `userLook`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := dao.NewDaoHandler(ctx, &daoUser.User).Filter(filter).Field(field).JoinGroupByPrimaryKey().GetModel().One()
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
	data := gconv.MapDeep(req)
	delete(data, `idArr`)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	filter := map[string]interface{}{`id`: req.IdArr}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `userUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.User().Update(ctx, filter, data)
	return
}
