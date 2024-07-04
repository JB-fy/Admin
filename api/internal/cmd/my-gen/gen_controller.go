package my_gen

import (
	"api/internal/cmd/my-gen/internal"
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
	diff      []string // 可以不要。数据返回时，会根据API文件中的结构体做过滤
}

func (controllerThis *myGenController) Merge(controllerOther myGenController) {
	controllerThis.importDao = append(controllerThis.importDao, controllerOther.importDao...)
	controllerThis.list = append(controllerThis.list, controllerOther.list...)
	controllerThis.info = append(controllerThis.info, controllerOther.info...)
	controllerThis.tree = append(controllerThis.tree, controllerOther.tree...)
	controllerThis.noAuth = append(controllerThis.noAuth, controllerOther.noAuth...)
	controllerThis.diff = append(controllerThis.diff, controllerOther.diff...)
}

func (controllerThis *myGenController) Unique() {
	controllerThis.importDao = garray.NewStrArrayFrom(controllerThis.importDao).Unique().Slice()
	// controllerThis.list = garray.NewStrArrayFrom(controllerThis.list).Unique().Slice()
	// controllerThis.info = garray.NewStrArrayFrom(controllerThis.info).Unique().Slice()
	// controllerThis.tree = garray.NewStrArrayFrom(controllerThis.tree).Unique().Slice()
	// controllerThis.noAuth = garray.NewStrArrayFrom(controllerThis.noAuth).Unique().Slice()
	// controllerThis.diff = garray.NewStrArrayFrom(controllerThis.diff).Unique().Slice()
}

// controller生成
func genController(option myGenOption, tpl myGenTpl) {
	controller := getControllerIdAndLabel(tpl)
	for _, v := range tpl.FieldList {
		controller.Merge(getControllerField(tpl, v))
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		controller.Merge(getControllerExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		controller.Merge(getControllerExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		controller.Merge(getControllerExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		controller.Merge(getControllerExtendMiddleMany(v))
	}
	controller.Unique()

	type defaultField struct {
		part1 []string
		part2 []string
		part3 []string
		part4 []string
	}
	defaultFieldObj := defaultField{}
	if option.IsList {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `defaultFieldOfList []string`)
		defaultFieldObj.part2 = append(defaultFieldObj.part2, `defaultFieldOfList := []string{`+gstr.Join(controller.list, `, `)+`}`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `defaultFieldOfList: append(field, defaultFieldOfList...),`)
		defaultFieldObj.part4 = append(defaultFieldObj.part4, `defaultFieldOfList`)
	}
	if option.IsInfo {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `defaultFieldOfInfo []string`)
		defaultFieldObj.part2 = append(defaultFieldObj.part2, `defaultFieldOfInfo := []string{`+gstr.Join(controller.info, `, `)+`}`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `defaultFieldOfInfo: append(field, defaultFieldOfInfo...),`)
		defaultFieldObj.part4 = append(defaultFieldObj.part4, `defaultFieldOfInfo`)
	}
	if option.IsList && tpl.Handle.Pid.Pid != `` {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `defaultFieldOfTree []string`)
		defaultFieldObj.part2 = append(defaultFieldObj.part2, `defaultFieldOfTree := []string{`+gstr.Join(controller.tree, `, `)+`}`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `defaultFieldOfTree: append(field, defaultFieldOfTree...),`)
		defaultFieldObj.part4 = append(defaultFieldObj.part4, `defaultFieldOfTree`)
	}
	if option.IsList || option.IsInfo {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `allowField         []string`)
		if len(controller.diff) > 0 {
			defaultFieldObj.part2 = append([]string{`field = gset.NewStrSetFrom(field).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controller.diff, `, `) + `})).Slice() //移除敏感字段`}, defaultFieldObj.part2...)
		}
		defaultFieldObj.part2 = append([]string{`field := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr().Slice()`}, defaultFieldObj.part2...)
		part3Str := `allowField:         append(field, `
		if len(defaultFieldObj.part4) == 1 {
			part3Str += defaultFieldObj.part4[0]
		} else {
			for k, v := range defaultFieldObj.part4 {
				if k == 0 {
					part3Str += `gset.NewStrSetFrom(` + v + `)`
				} else {
					part3Str += `.Merge(gset.NewStrSetFrom(` + v + `))`
				}
			}
			part3Str += `.Slice()`
		}
		part3Str += `...),`
		defaultFieldObj.part3 = append(defaultFieldObj.part3, part3Str)
	}
	if option.IsList && option.IsAuthAction {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `noAuthField        []string`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `noAuthField:        []string{`+gstr.Join(controller.noAuth, `, `)+`},`)
	}

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
`
	if option.IsList || option.IsInfo {
		tplController += `
type ` + tpl.TableCaseCamel + ` struct {
	` + gstr.Join(defaultFieldObj.part1, `
	`) + `
}

func New` + tpl.TableCaseCamel + `() *` + tpl.TableCaseCamel + ` {
	` + gstr.Join(defaultFieldObj.part2, `
	`) + `
	return &` + tpl.TableCaseCamel + `{
		` + gstr.Join(defaultFieldObj.part3, `
		`) + `
	}
}
`
	} else {
		tplController += `
type ` + tpl.TableCaseCamel + ` struct{}

func New` + tpl.TableCaseCamel + `() *` + tpl.TableCaseCamel + ` {
	return &` + tpl.TableCaseCamel + `{}
}
`
	}
	if option.IsList {
		tplController += `
// 列表
func (controllerThis *` + tpl.TableCaseCamel + `) List(ctx context.Context, req *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListReq) (res *api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes, err error) {
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
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Read` + "`" + `)
	if !isAuth {
		field = controllerThis.noAuthField
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
	list, err := daoModelThis.Fields(field...).Order(req.Sort).Page(req.Page, req.Limit).ListPri()
	if err != nil {
		return
	}

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `ListRes{`
		if option.IsCount {
			tplController += `Count: count, `
		}
		tplController += `List: []api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `Info{}}
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
	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfInfo
	}
	filter := map[string]any{` + "`id`" + `: req.Id}
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	_, err = service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Read` + "`" + `)
	if err != nil {
		return
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	info, err := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field...).InfoPri()
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
	delete(data, ` + "`" + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + "`" + `)
	if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ` + "``" + `)
		return
	}
	filter := map[string]any{` + "`id`" + `: req.IdArr}
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
	filter := map[string]any{` + "`id`" + `: req.IdArr}
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
		filter = map[string]any{}
	}

	var field []string
	if len(req.Field) > 0 {
		field = gset.NewStrSetFrom(req.Field).Intersect(gset.NewStrSetFrom(controllerThis.allowField)).Slice()
	}
	if len(field) == 0 {
		field = controllerThis.defaultFieldOfTree
	}
	field = append(field, ` + "`tree`" + `)
	/**--------参数处理 结束--------**/
`
		if option.IsAuthAction {
			tplController += `
	/**--------权限验证 开始--------**/
	isAuth, _ := service.AuthAction().CheckAuth(ctx, ` + "`" + gstr.CaseCamelLower(tpl.LogicStructName) + `Read` + "`" + `)
	if !isAuth {
		field = controllerThis.noAuthField
	}
	/**--------权限验证 结束--------**/
`
		}
		tplController += `
	list, err := dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.CtxDaoModel(ctx).Filters(filter).Fields(field...).ListPri()
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
		controller.list = append(controller.list, "`id`")
		controller.info = append(controller.info, "`id`")
		controller.tree = append(controller.tree, "`id`")
	}
	controller.list = append(controller.list, "`label`")
	controller.info = append(controller.info, "`label`")
	controller.tree = append(controller.tree, "`label`")

	controller.noAuth = append(controller.noAuth, "`id`", "`label`")
	/* if len(tpl.Handle.Id.List) == 1 && tpl.Handle.Id.List[0].FieldRaw != `id` {
		controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel)
	}
	controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.LabelList[0])) */
	return
}

func getControllerField(tpl myGenTpl, v myGenField) (controller myGenController) {
	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		return
	case internal.TypePrimaryMany: // 联合主键
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
		return
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
	case internal.TypeNameUpdated: // 更新时间字段
	case internal.TypeNameCreated: // 创建时间字段
	case internal.TypeNamePid: // pid；	类型：int等类型；
		controller.list = append(controller.list,
			"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, tpl.Handle.LabelList[0], `p`)+"`",
			"`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_has_child`)+"`",
		)
		controller.noAuth = append(controller.noAuth, "`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_has_child`)+"`")
	case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
	case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		controller.diff = append(controller.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		controller.diff = append(controller.diff, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+v.FieldCaseCamel)
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
			controller.importDao = append(controller.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
			daoPathRel := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
			fieldTmp := daoPathRel + `.Columns().` + gstr.CaseCamel(relIdObj.tpl.Handle.LabelList[0])
			if relIdObj.Suffix != `` {
				fieldTmp += "+`" + relIdObj.Suffix + "`"
			}
			controller.list = append(controller.list, fieldTmp)
		}
	case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
	case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
	case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getControllerExtendMiddleOne(tplEM handleExtendMiddle) (controller myGenController) {
	tpl := tplEM.tpl
	controller.importDao = append(controller.importDao, `dao`+tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tpl.ModuleDirCaseKebab+`"`)
	daoPath := `dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel
	controllerAppend := func(field []myGenField) {
		for _, v := range field {
			switch v.FieldTypeName {
			case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
			case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			default:
				field := daoPath + `.Columns().` + v.FieldCaseCamel
				controller.list = append(controller.list, field)
				controller.info = append(controller.info, field)
				controller.tree = append(controller.tree, field)
			}
		}
	}
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		controllerAppend(tplEM.FieldList)
	case internal.TableTypeMiddleOne:
		controllerAppend(append(tplEM.FieldListOfIdSuffix, tplEM.FieldListOfOther...))
	}
	for _, v := range tplEM.FieldList {
		controllerField := getControllerField(tpl, v)
		controllerField.diff = []string{}
		controller.Merge(controllerField)
	}
	return
}

func getControllerExtendMiddleMany(tplEM handleExtendMiddle) (controller myGenController) {
	controller.info = append(controller.info, "`"+tplEM.FieldVar+"`")
	if len(tplEM.FieldList) == 1 {
		isShow := true
		v := tplEM.FieldList[0]
		switch v.FieldTypeName {
		case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			isShow = false
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
			isShow = false
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
			if v.FieldType != internal.TypeVarchar {
				isShow = false
			}
		}
		if isShow {
			controller.list = append(controller.list, "`"+tplEM.FieldVar+"`")
			controller.tree = append(controller.tree, "`"+tplEM.FieldVar+"`")
		}
	}
	return
}
