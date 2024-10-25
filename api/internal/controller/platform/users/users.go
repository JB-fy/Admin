package users

import (
	"api/api"
	apiUsers "api/api/platform/users"
	daoUsers "api/internal/dao/users"
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Users struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewUsers() *Users {
	field := daoUsers.Users.ColumnArr().Slice()
	defaultFieldOfList := []string{`id`, `label`, daoUsers.Privacy.Columns().IdCardNo, daoUsers.Privacy.Columns().IdCardName, daoUsers.Privacy.Columns().IdCardGender, daoUsers.Privacy.Columns().IdCardBirthday, daoUsers.Privacy.Columns().IdCardAddress}
	defaultFieldOfInfo := []string{`id`, `label`, daoUsers.Privacy.Columns().IdCardNo, daoUsers.Privacy.Columns().IdCardName, daoUsers.Privacy.Columns().IdCardGender, daoUsers.Privacy.Columns().IdCardBirthday, daoUsers.Privacy.Columns().IdCardAddress}
	return &Users{
		defaultFieldOfList: append(field, defaultFieldOfList...),
		defaultFieldOfInfo: append(field, defaultFieldOfInfo...),
		allowField:         append(field, gset.NewStrSetFrom(defaultFieldOfList).Merge(gset.NewStrSetFrom(defaultFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Users) List(ctx context.Context, req *apiUsers.UsersListReq) (res *apiUsers.UsersListRes, err error) {
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `usersRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoUsers.Users.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiUsers.UsersListRes{Count: count, List: []apiUsers.UsersInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Users) Info(ctx context.Context, req *apiUsers.UsersInfoReq) (res *apiUsers.UsersInfoRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `usersRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoUsers.Users.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiUsers.UsersInfoRes{}
	info.Struct(&res.Info)
	return
}

// 修改
func (controllerThis *Users) Update(ctx context.Context, req *apiUsers.UsersUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `usersUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.Users().Update(ctx, filter, data)
	return
}
