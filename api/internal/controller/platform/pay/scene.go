package pay

import (
	"api/api"
	apiPay "api/api/platform/pay"
	daoPay "api/internal/dao/pay"
	"api/internal/service"
	"api/internal/utils"
	"context"
	"slices"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type Scene struct {
	defaultFieldOfList []string
	defaultFieldOfInfo []string
	allowField         []string
	noAuthField        []string
}

func NewScene() *Scene {
	field := append(daoPay.Scene.ColumnArr(), `id`, `label`)
	appendFieldOfList := []string{}
	appendFieldOfInfo := []string{}
	return &Scene{
		defaultFieldOfList: append(slices.Clone(field), appendFieldOfList...),
		defaultFieldOfInfo: append(slices.Clone(field), appendFieldOfInfo...),
		allowField:         append(slices.Clone(field), gset.NewStrSetFrom(slices.Concat(appendFieldOfList, appendFieldOfInfo)).Slice()...),
		noAuthField:        []string{`id`, `label`},
	}
}

// 列表
func (controllerThis *Scene) List(ctx context.Context, req *apiPay.SceneListReq) (res *apiPay.SceneListRes, err error) {
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
	isAuth, _ := service.AuthAction().CheckAuth(ctx, `paySceneRead`)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/

	daoModelThis := daoPay.Scene.CtxDaoModel(ctx).Filters(filter)
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &apiPay.SceneListRes{Count: count, List: []apiPay.SceneInfo{}}
	list.Structs(&res.List)
	return
}

// 详情
func (controllerThis *Scene) Info(ctx context.Context, req *apiPay.SceneInfoReq) (res *apiPay.SceneInfoRes, err error) {
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
	_, err = service.AuthAction().CheckAuth(ctx, `paySceneRead`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	info, err := daoPay.Scene.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ``)
		return
	}

	res = &apiPay.SceneInfoRes{}
	info.Struct(&res.Info)
	return
}

// 新增
func (controllerThis *Scene) Create(ctx context.Context, req *apiPay.SceneCreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `paySceneCreate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	id, err := service.PayScene().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}

// 修改
func (controllerThis *Scene) Update(ctx context.Context, req *apiPay.SceneUpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`filter`}})
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true, Tags: []string{`data`}})
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ``)
		return
	}
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `paySceneUpdate`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.PayScene().Update(ctx, filter, data)
	return
}

// 删除
func (controllerThis *Scene) Delete(ctx context.Context, req *apiPay.SceneDeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/

	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, `paySceneDelete`)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/

	_, err = service.PayScene().Delete(ctx, filter)
	return
}
