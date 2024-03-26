package my_gen

import (
	"api/internal/utils"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenController struct {
	importDao []string
	list      []string
	info      []string
	tree      []string
	noAuth    []string
	// diff      []string // 可以不要。数据返回时，会根据API文件中的结构体做过滤
}

func (controllerThis *myGenController) Merge(controllerOther myGenController) {
	controllerThis.importDao = append(controllerThis.importDao, controllerOther.importDao...)
	controllerThis.list = append(controllerThis.list, controllerOther.list...)
	controllerThis.info = append(controllerThis.info, controllerOther.info...)
	controllerThis.tree = append(controllerThis.tree, controllerOther.tree...)
	controllerThis.noAuth = append(controllerThis.noAuth, controllerOther.noAuth...)
}

func (controllerThis *myGenController) Unique() {
	controllerThis.importDao = garray.NewStrArrayFrom(controllerThis.importDao).Unique().Slice()
	controllerThis.list = garray.NewStrArrayFrom(controllerThis.list).Unique().Slice()
	controllerThis.info = garray.NewStrArrayFrom(controllerThis.info).Unique().Slice()
	controllerThis.tree = garray.NewStrArrayFrom(controllerThis.tree).Unique().Slice()
	controllerThis.noAuth = garray.NewStrArrayFrom(controllerThis.noAuth).Unique().Slice()
}

// controller生成
func genController(option myGenOption, tpl myGenTpl) {
	controller := getControllerIdAndLabel(tpl)
	controller.Merge(getControllerFieldList(tpl))
	controller.Unique()

	tplController := `package controller

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `"
	dao` + tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + tpl.ModuleDirCaseKebab + `"` + gstr.Join(append([]string{``}, controller.importDao...), `
	`) + `
	"api/internal/service"
	"api/internal/utils"
	"context"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/util/gconv"
)

type ` + tpl.TableCaseCamel + ` struct{}

func New` + tpl.TableCaseCamel + `() *` + tpl.TableCaseCamel + ` {
	return &` + tpl.TableCaseCamel + `{}
}
`
	if option.IsList {
		tplController += `
// 列表
func (controllerThis *` + tpl.TableCaseCamel + `) List(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}
`
		tplController += `
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()`
		if len(controller.list) > 0 {
			tplController += `
	allowField = append(allowField` + gstr.Join(append([]string{``}, controller.list...), `, `) + `)`
		}
		/* if len(controller.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controller.diff, `, `) + `})).Slice() //移除敏感字段`
		} */
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Look` + "`" + `)
	if !isAuth {
		field = []string{` + gstr.Join(controller.noAuth, `, `) + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	daoModelThis := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter)`
		if option.IsCount {
			tplController += `
	count, err := daoModelThis.CountPri()
	if err != nil {
		return
	}`
		}
		tplController += `
	list, err := daoModelThis.Fields(field).HookSelect().Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes{`
		if option.IsCount {
			tplController += `Count: count, `
		}
		tplController += `List:  []api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListItem{}}
	list.Structs(&res.List)
	return
}
`
	}
	if option.IsInfo {
		tplController += `
// 详情
func (controllerThis *` + tpl.TableCaseCamel + `) Info(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoRes, err error) {
	/**--------参数处理 开始--------**/
	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()`
		if len(controller.info) > 0 {
			tplController += `
	allowField = append(allowField` + gstr.Join(append([]string{``}, controller.info...), `, `) + `)`
		}
		/* if len(controller.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controller.diff, `, `) + `})).Slice() //移除敏感字段`
		} */
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	filter := map[string]interface{}{` + "`id`" + `: req.Id}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Look` + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	info, err := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().InfoPri()
	if err != nil {
		return
	}
	if info.IsEmpty() {
		err = utils.NewErrorCode(ctx, 29999998, ` + "``" + `)
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `InfoRes{}
	info.Struct(&res.Info)
	return
}
`
	}
	if option.IsCreate {
		tplController += `
// 新增
func (controllerThis *` + tpl.TableCaseCamel + `) Create(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `CreateReq) (res *api.CommonCreateRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Create` + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	id, err := service.` + tpl.LogicStructName + `().Create(ctx, data)
	if err != nil {
		return
	}
	res = &api.CommonCreateRes{Id: id}
	return
}
`
	}

	if option.IsUpdate {
		tplController += `
// 修改
func (controllerThis *` + tpl.TableCaseCamel + `) Update(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `UpdateReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	data := gconv.Map(req, gconv.MapOption{Deep: true, OmitEmpty: true})
	delete(data, ` + "`idArr`" + `)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ` + "``" + `)
		return
	}
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Update` + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	_, err = service.` + tpl.LogicStructName + `().Update(ctx, filter, data)
	return
}
`
	}

	if option.IsDelete {
		tplController += `
// 删除
func (controllerThis *` + tpl.TableCaseCamel + `) Delete(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `DeleteReq) (res *api.CommonNoDataRes, err error) {
	/**--------参数处理 开始--------**/
	filter := map[string]interface{}{` + "`id`" + `: req.IdArr}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Delete` + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	_, err = service.` + tpl.LogicStructName + `().Delete(ctx, filter)
	return
}
`
	}

	if option.IsList && tpl.Handle.Pid.Pid != `` {
		tplController += `
// 列表（树状）
func (controllerThis *` + tpl.TableCaseCamel + `) Tree(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes, err error) {
	/**--------参数处理 开始--------**/
	filter := gconv.Map(req.Filter, gconv.MapOption{Deep: true, OmitEmpty: true})
	if filter == nil {
		filter = map[string]interface{}{}
	}

	allowField := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()`
		if len(controller.tree) > 0 {
			tplController += `
	allowField = append(allowField` + gstr.Join(append([]string{``}, controller.tree...), `, `) + `)`
		}
		/* if len(controller.diff) > 0 {
			tplController += `
		allowField = gset.NewStrSetFrom(allowField).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controller.diff, `, `) + `})).Slice() //移除敏感字段`
		} */
		tplController += `
	field := allowField
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(allowField)).Slice()
		if len(field) == 0 {
			field = allowField
		}
	}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Look` + "`" + `)
	if !isAuth {
		field = []string{` + gstr.Join(controller.noAuth, `, `) + `}
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	field = append(field, ` + "`tree`" + `)

	list, err :=dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field).HookSelect().ListPri()
	if err != nil {
		return
	}
	tree := utils.Tree(list.List(), 0, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `)

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes{}
	gconv.Structs(tree, &res.Tree)
	return
}
`
	}

	saveFile := gfile.SelfDir() + `/internal/controller/` + option.SceneCode + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	gfile.PutContents(saveFile, tplController)
	utils.GoFileFmt(saveFile)
}

func getControllerIdAndLabel(tpl myGenTpl) (controller myGenController) {
	if len(tpl.Handle.Id.List) > 1 || tpl.Handle.Id.List[0].FieldRaw != `id` {
		controller.list = []string{"`id`"}
		controller.info = []string{"`id`"}
		controller.tree = []string{"`id`"}
	}
	controller.noAuth = []string{"`id`"}

	if len(tpl.Handle.LabelList) > 0 {
		controller.list = append(controller.list, "`label`")
		controller.info = append(controller.info, "`label`")
		controller.tree = append(controller.tree, "`label`")
		if tpl.Handle.Pid.Pid != `` {
			controller.list = append(controller.list, "`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`")
			// controller.info = append(controller.info, "`p"+gstr.CaseCamel(tpl.Handle.LabelList[0])+"`")
		}
		controller.noAuth = append(controller.noAuth, "`label`")
		if len(tpl.Handle.Id.List) == 1 && tpl.Handle.Id.List[0].FieldRaw != `id` {
			controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel)
		}
		controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0]))
		/* for _, v := range tpl.Handle.LabelList {
			controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(v))
		} */
	}
	return
}

func getControllerFieldList(tpl myGenTpl, fieldArrOfIgnore ...string) (controller myGenController) {
	for _, v := range tpl.FieldList {
		if garray.NewStrArrayFrom(fieldArrOfIgnore).Contains(v.FieldRaw) {
			continue
		}

		/*--------根据字段主键类型处理 开始--------*/
		switch v.FieldTypePrimary {
		case TypePrimary: // 独立主键
		case TypePrimaryAutoInc: // 独立主键（自增）
			continue
		case TypePrimaryMany: // 联合主键
		case TypePrimaryManyAutoInc: // 联合主键（自增）
			continue
		}
		/*--------根据字段主键类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
		case TypeNameUpdated: // 更新时间字段
		case TypeNameCreated: // 创建时间字段
		case TypeNamePid: // pid；	类型：int等类型；
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			// controller.diff = append(controller.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			// controller.diff = append(controller.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
				daoPath := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
				importDaoStr := `dao` + relIdObj.tpl.ModuleDirCaseCamel + ` "api/internal/dao/` + relIdObj.tpl.ModuleDirCaseKebab + `"`
				if !garray.NewStrArrayFrom(controller.importDao).Contains(importDaoStr) {
					controller.importDao = append(controller.importDao, importDaoStr)
				}
				fieldTmp := daoPath + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])
				if relIdObj.Suffix != `` {
					fieldTmp += "+`" + relIdObj.Suffix + "`"
				}
				controller.list = append(controller.list, fieldTmp)
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/
	}

	for _, v := range tpl.Handle.ExtendTableOneList {
		controller.Merge(getControllerFieldList(v.tpl, v.FieldArrOfIgnore...))
	}
	return
}
