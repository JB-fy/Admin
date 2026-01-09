package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"api/internal/utils"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
)

type myGenController struct {
	importDao []string
	common    []string
	list      []string
	info      []string
	tree      []string
	noAuth    []string
	diff      []string // 可以不要。数据返回时，会根据API文件中的结构体做过滤
	update    []string // 参数结构体数组[]struct(用*[]struct会导致校验规则失效)时，此时传空数组无法删除扩展表，需增加判断参数是nil还是[]
}

func (controllerThis *myGenController) Merge(controllerOther myGenController) {
	controllerThis.importDao = append(controllerThis.importDao, controllerOther.importDao...)
	controllerThis.common = append(controllerThis.common, controllerOther.common...)
	controllerThis.list = append(controllerThis.list, controllerOther.list...)
	controllerThis.info = append(controllerThis.info, controllerOther.info...)
	controllerThis.tree = append(controllerThis.tree, controllerOther.tree...)
	controllerThis.noAuth = append(controllerThis.noAuth, controllerOther.noAuth...)
	controllerThis.diff = append(controllerThis.diff, controllerOther.diff...)
	controllerThis.update = append(controllerThis.update, controllerOther.update...)
}

func (controllerThis *myGenController) Unique() {
	controllerThis.importDao = garray.NewStrArrayFrom(controllerThis.importDao).Unique().Slice()
	// controllerThis.common = garray.NewStrArrayFrom(controllerThis.common).Unique().Slice()
	// controllerThis.list = garray.NewStrArrayFrom(controllerThis.list).Unique().Slice()
	// controllerThis.info = garray.NewStrArrayFrom(controllerThis.info).Unique().Slice()
	// controllerThis.tree = garray.NewStrArrayFrom(controllerThis.tree).Unique().Slice()
	// controllerThis.noAuth = garray.NewStrArrayFrom(controllerThis.noAuth).Unique().Slice()
	// controllerThis.diff = garray.NewStrArrayFrom(controllerThis.diff).Unique().Slice()
}

// controller生成
func genController(option myGenOption, tpl *myGenTpl) {
	controller := myGenController{}
	controller.importDao = append(controller.importDao, `dao`+tpl.ModuleDirCaseCamel+` "api/internal/dao/`+tpl.ModuleDirCaseKebab+`"`)
	if option.LoginIdStr != `` {
		loginIdStrArr := strings.Split(option.LoginIdStr, `.`)
		if len(loginIdStrArr) == 4 && strings.Index(loginIdStrArr[0], `dao`) == 0 {
			moduleDirCaseKebab := gstr.CaseKebab(strings.Replace(loginIdStrArr[0], `dao`, ``, 1))
			daoDirFormat := gfile.SelfDir() + `/internal/dao/%s`
			// daoFileFormat := daoDirFormat + `/` + gstr.CaseKebab(loginIdStrArr[1]) + `.go`
			if gfile.IsDir(fmt.Sprintf(daoDirFormat, moduleDirCaseKebab)) /* && gfile.IsFile(fmt.Sprintf(daoFileFormat, loginIdStrArr[0])) */ {
				controller.importDao = append(controller.importDao, loginIdStrArr[0]+` "api/internal/dao/`+moduleDirCaseKebab+`"`)
			} else if moduleDirCaseKebabArr := strings.Split(moduleDirCaseKebab, `_`); len(moduleDirCaseKebabArr) > 1 { //非default分组
				moduleDirCaseKebab = moduleDirCaseKebabArr[0] + `/` + strings.Join(moduleDirCaseKebabArr[1:], `_`)
				if gfile.IsDir(fmt.Sprintf(daoDirFormat, moduleDirCaseKebab)) /* && gfile.IsFile(fmt.Sprintf(daoFileFormat, moduleDirCaseKebab)) */ {
					controller.importDao = append(controller.importDao, loginIdStrArr[0]+` "api/internal/dao/`+moduleDirCaseKebab+`"`)
				}
			}
		}
	}
	if len(tpl.Handle.Id.List) > 1 || tpl.Handle.Id.List[0].FieldRaw != `id` {
		controller.common = append(controller.common, "`id`")
	}
	controller.common = append(controller.common, "`label`")
	controller.noAuth = append(controller.noAuth, "`id`", "`label`")
	/* if len(tpl.Handle.Id.List) == 1 && tpl.Handle.Id.List[0].FieldRaw != `id` {
		controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+tpl.Handle.Id.List[0].FieldCaseCamel)
	}
	controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.Label.List[0])) */
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
		defaultFieldObj.part2 = append(defaultFieldObj.part2, `appendFieldOfList := []string{`+gstr.Join(controller.list, `, `)+`}`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `defaultFieldOfList: slices.Clone(append(field, appendFieldOfList...)),`)
		defaultFieldObj.part4 = append(defaultFieldObj.part4, `appendFieldOfList`)
	}
	if option.IsInfo {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `defaultFieldOfInfo []string`)
		defaultFieldObj.part2 = append(defaultFieldObj.part2, `appendFieldOfInfo := []string{`+gstr.Join(controller.info, `, `)+`}`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `defaultFieldOfInfo: slices.Clone(append(field, appendFieldOfInfo...)),`)
		defaultFieldObj.part4 = append(defaultFieldObj.part4, `appendFieldOfInfo`)
	}
	if option.IsList && tpl.Handle.Pid.Pid != `` {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `defaultFieldOfTree []string`)
		defaultFieldObj.part2 = append(defaultFieldObj.part2, `appendFieldOfTree := []string{`+gstr.Join(controller.tree, `, `)+`}`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `defaultFieldOfTree: slices.Clone(append(field, appendFieldOfTree...)),`)
		defaultFieldObj.part4 = append(defaultFieldObj.part4, `appendFieldOfTree`)
	}
	if option.IsList || option.IsInfo {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `allowField         []string`)
		if len(controller.diff) > 0 {
			defaultFieldObj.part2 = append([]string{`field = gset.NewStrSetFrom(field).Diff(gset.NewStrSetFrom([]string{` + gstr.Join(controller.diff, `, `) + `})).Slice() //移除敏感字段`}, defaultFieldObj.part2...)
		}
		defaultFieldObj.part2 = append([]string{`field := slices.Clone(append(dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.ColumnArr(), ` + gstr.Join(controller.common, `, `) + `))`}, defaultFieldObj.part2...)
		part3Str := `allowField:         slices.Clone(append(field, `
		if len(defaultFieldObj.part4) == 1 {
			part3Str += defaultFieldObj.part4[0]
		} else {
			part3Str += `gset.NewStrSetFrom(slices.Concat(` + gstr.Join(defaultFieldObj.part4, `, `) + `)).Slice()`
		}
		part3Str += `...)),`
		defaultFieldObj.part3 = append(defaultFieldObj.part3, part3Str)
	}
	if option.IsList && option.IsAuthAction {
		defaultFieldObj.part1 = append(defaultFieldObj.part1, `noAuthField        []string`)
		defaultFieldObj.part3 = append(defaultFieldObj.part3, `noAuthField:        []string{`+gstr.Join(controller.noAuth, `, `)+`},`)
	}

	loginFilterStr := ``
	loginDataStr := ``
	if option.LoginRelId != `` {
		loginFilterStr = `

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	filter[dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(option.LoginRelId) + `] = loginInfo[` + option.LoginIdStr + `]`
		loginDataStr = `

	loginInfo := get_or_set_ctx.GetCtxLoginInfo(ctx)
	data[dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(option.LoginRelId) + `] = loginInfo[` + option.LoginIdStr + `]`
	}
	if option.FilterIsStop {
		loginFilterStr += `
	filter[dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(`is_stop`) + `] = 0`
	}

	tplController := `package ` + tpl.GetModuleName(`controller`) + `

import (
	"api/api"
	api` + tpl.ModuleDirCaseCamel + ` "api/api/` + option.SceneId + `/` + tpl.ModuleDirCaseKebab + `"` + gstr.Join(append([]string{``}, controller.importDao...), `
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
	}` + loginFilterStr + `
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
	filter := map[string]any{` + "`id`" + `: req.Id}` + loginFilterStr + `
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
	data := gconv.Map(req.` + tpl.TableCaseCamel + `CreateData, gconv.MapOption{Deep: true, OmitEmpty: true})` + loginDataStr + `
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
	filter := gconv.Map(req.` + tpl.TableCaseCamel + `UpdateDeleteFilter, gconv.MapOption{Deep: true, OmitEmpty: true})
	data := gconv.Map(req.` + tpl.TableCaseCamel + `UpdateData, gconv.MapOption{Deep: true, OmitEmpty: true})
	` + gstr.Join(append(controller.update, ``), `
	`) + `if len(data) == 0 {
		err = utils.NewErrorCode(ctx, 89999999, ` + "``" + `)
		return
	}` + loginFilterStr + `
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
	filter := gconv.Map(req.` + tpl.TableCaseCamel + `UpdateDeleteFilter, gconv.MapOption{Deep: true, OmitEmpty: true})` + loginFilterStr + `
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
	field = append(field, ` + "`tree`" + `)` + loginFilterStr + `
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
	tree := utils.Tree(list.List(), ` + tpl.Handle.Pid.Tpl.PidDefVal + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + tpl.Handle.Id.List[0].FieldCaseCamel + `, dao` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `.Columns().` + gstr.CaseCamel(tpl.Handle.Pid.Pid) + `)

	res = &api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `TreeRes{Tree: []api` + tpl.ModuleDirCaseCamel + `.` + tpl.TableCaseCamel + `Info{}}
	gconv.Structs(tree, &res.Tree)
	return
}
`
	}

	saveFile := gfile.SelfDir() + `/internal/controller/` + option.SceneId + `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseSnake + `.go`
	utils.FilePutFormat(saveFile, []byte(tplController)...)
}

func getControllerField(tpl *myGenTpl, v myGenField) (controller myGenController) {
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
	case internal.TypeNamePid: // pid，且与主键类型相同时（才）有效；	类型：int等类型或varchar或char；
		controller.list = append(controller.list, "`"+internal.GetStrByFieldStyle(tpl.FieldStyle, tpl.Handle.Label.List[0].FieldRaw, `p`)+"`")
		if tpl.Handle.Pid.IsLeaf == `` {
			controller.list = append(controller.list, "`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_leaf`)+"`")
			controller.noAuth = append(controller.noAuth, "`"+internal.GetStrByFieldStyle(tpl.FieldStyle, `is_leaf`)+"`")
		} else {
			// controller.list = append(controller.list, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf))
			controller.noAuth = append(controller.noAuth, `dao`+tpl.ModuleDirCaseCamel+`.`+tpl.TableCaseCamel+`.Columns().`+gstr.CaseCamel(tpl.Handle.Pid.IsLeaf))
		}
	case internal.TypeNameIdPath, internal.TypeNameNamePath: // id_path|idPath，且pid同时存在时（才）有效；	类型：varchar或text；	// name_path|namePath，且pid，id_path|idPath同时存在时（才）有效；	类型：varchar或text；
	case internal.TypeNameLevel, internal.TypeNameIsLeaf: // level，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；	// is_leaf|isLeaf，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
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
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl != nil && !relIdObj.IsRedundName {
			controller.importDao = append(controller.importDao, `dao`+relIdObj.tpl.ModuleDirCaseCamel+` "api/internal/dao/`+relIdObj.tpl.ModuleDirCaseKebab+`"`)
			daoPathRel := `dao` + relIdObj.tpl.ModuleDirCaseCamel + `.` + relIdObj.tpl.TableCaseCamel
			fieldTmp := daoPathRel + `.Columns().` + relIdObj.tpl.Handle.Label.List[0].FieldCaseCamel
			if relIdObj.Suffix != `` {
				fieldTmp += "+`" + relIdObj.Suffix + "`"
			}
			controller.list = append(controller.list, fieldTmp)
		}
	case internal.TypeNameStatusSuffix, internal.TypeNameIsPrefix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）	// is_前缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
	case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
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
				controller.common = append(controller.common, field)
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
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
			isShow = false
		case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
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
	} else {
		controller.update = append(controller.update, `if len(req.`+gstr.CaseCamel(tplEM.FieldVar)+`) == 0 && req.`+gstr.CaseCamel(tplEM.FieldVar)+` != nil {
		data[`+"`"+tplEM.FieldVar+"`"+`] = g.List{}
	}`)
	}
	return
}
