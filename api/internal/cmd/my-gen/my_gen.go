package my_gen

import (
	daoAuth "api/internal/dao/auth"
	"context"
	"fmt"

	"github.com/fatih/color"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGen struct {
	ctx       context.Context
	sceneInfo gdb.Record //场景信息
	option    myGenOption
	tpl       myGenTpl
}

// 命令参数解析后的数据
type myGenOption struct {
	SceneCode          string     `json:"sceneCode"`          //场景标识，必须在数据库表auth_scene已存在。示例：platform
	DbGroup            string     `json:"dbGroup"`            //db分组。示例：default
	DbTable            string     `json:"dbTable"`            //db表。示例：auth_test
	RemovePrefixCommon string     `json:"removePrefixCommon"` //要删除的共有前缀，没有可为空。removePrefixCommon + removePrefixAlone必须和hack/config.yaml内removePrefix保持一致
	RemovePrefixAlone  string     `json:"removePrefixAlone"`  //要删除的独有前缀。removePrefixCommon + removePrefixAlone必须和hack/config.yaml内removePrefix保持一致，示例：auth_
	CommonName         string     `json:"commonName"`         //公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：用户，权限管理/测试
	IsList             bool       `json:"isList" `            //是否生成列表接口(0和no为false，1和yes为true)
	IsCount            bool       `json:"isCount" `           //列表接口是否返回总数
	IsInfo             bool       `json:"isInfo" `            //是否生成详情接口
	IsCreate           bool       `json:"isCreate"`           //是否生成创建接口
	IsUpdate           bool       `json:"isUpdate"`           //是否生成更新接口
	IsDelete           bool       `json:"isDelete"`           //是否生成删除接口
	IsApi              bool       `json:"isApi"`              //是否生成后端接口文件
	IsAuthAction       bool       `json:"isAuthAction"`       //是否判断操作权限，如是，则同时会生成操作权限
	IsView             bool       `json:"isView"`             //是否生成前端视图文件
	SceneInfo          gdb.Record //场景信息
}

func NewMyGen(ctx context.Context, parser *gcmd.Parser) (myGenObj myGen) {
	option := myGenOption{}
	defer func() {
		myGenObj.option = option
		myGenObj.sceneInfo = option.SceneInfo
		myGenObj.tpl = createTpl(ctx, option.DbGroup, option.DbTable, option.RemovePrefixCommon, option.RemovePrefixAlone)
	}()

	optionMap := parser.GetOptAll()
	gconv.Struct(optionMap, &option)

	// 命令执行前提示搭配Git使用
	gcmd.Scan(
		color.HiYellowString(`重要提示：强烈建议搭配Git使用，防止代码覆盖风险。`)+"\n",
		color.HiYellowString(`    Git库已创建或忽略风险，请按`)+color.HiGreenString(`[Enter]`)+color.HiYellowString(`继续执行`)+"\n",
		color.HiYellowString(`    Git库未创建，请按`)+color.HiRedString(`[Ctrl + C]`)+color.HiYellowString(`中断执行`)+"\n",
	)

	// 场景标识
	if option.SceneCode == `` {
		option.SceneCode = gcmd.Scan(color.BlueString(`> 请输入场景标识：`))
	}
	for {
		if option.SceneCode != `` {
			option.SceneInfo, _ = daoAuth.Scene.CtxDaoModel(ctx).Filter(daoAuth.Scene.Columns().SceneCode, option.SceneCode).One()
			if !option.SceneInfo.IsEmpty() {
				break
			}
		}
		option.SceneCode = gcmd.Scan(color.RedString(`    场景标识不存在，请重新输入：`))
	}
	// db分组
	if option.DbGroup == `` {
		option.DbGroup = gcmd.Scan(color.BlueString(`> 请输入db分组，默认(default)：`))
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	for {
		/* err := g.Try(ctx, func(ctx context.Context) {
			dbLink = gconv.String(gconv.SliceMap(g.Cfg().MustGet(ctx, `database`).MapStrAny()[option.DbGroup])[0][`link`])
		})
		if err == nil {
			break
		} */
		err := g.Try(ctx, func(ctx context.Context) {
			g.DB(option.DbGroup)
		})
		if err == nil {
			break
		}
		option.DbGroup = gcmd.Scan(color.RedString(`    db分组不存在，请重新输入，默认(default)：`))
		if option.DbGroup == `` {
			option.DbGroup = `default`
		}
	}
	// db表
	tableArr, _ := g.DB(option.DbGroup).Tables(ctx)
	if option.DbTable == `` {
		option.DbTable = gcmd.Scan(color.BlueString(`> 请输入db表：`))
	}
	for {
		if option.DbTable != `` && garray.NewStrArrayFrom(tableArr).Contains(option.DbTable) {
			break
		}
		option.DbTable = gcmd.Scan(color.RedString(`    db表不存在，请重新输入：`))
	}
	// 要删除的共有前缀
	if _, ok := optionMap[`removePrefixCommon`]; !ok {
		option.RemovePrefixCommon = gcmd.Scan(color.BlueString(`> 请输入要删除的共有前缀，默认(空)：`))
	}
	for {
		if option.RemovePrefixCommon == `` || gstr.Pos(option.DbTable, option.RemovePrefixCommon) == 0 {
			break
		}
		option.RemovePrefixCommon = gcmd.Scan(color.RedString(`    要删除的共有前缀不存在，请重新输入，默认(空)：`))
	}
	// 要删除的独有前缀
	if _, ok := optionMap[`removePrefixAlone`]; !ok {
		option.RemovePrefixAlone = gcmd.Scan(color.BlueString(`> 请输入要删除的独有前缀，默认(空)：`))
	}
	for {
		if option.RemovePrefixAlone == `` || gstr.Pos(option.DbTable, option.RemovePrefixCommon+option.RemovePrefixAlone) == 0 {
			break
		}
		option.RemovePrefixAlone = gcmd.Scan(color.RedString(`    要删除的独有前缀不存在，请重新输入，默认(空)：`))
	}
	// 公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用。示例：场景
	for {
		if option.CommonName != `` {
			break
		}
		option.CommonName = gcmd.Scan(color.BlueString(`> 请输入公共名称，将同时在swagger文档Tag标签，权限菜单和权限操作中使用：`))
	}
noAllRestart:
	// 是否生成列表接口
	isList, ok := optionMap[`isList`]
	if !ok {
		isList = gcmd.Scan(color.BlueString(`> 是否生成列表接口，默认(yes)：`))
	}
isListEnd:
	for {
		switch isList {
		case ``, `1`, `yes`:
			option.IsList = true
			break isListEnd
		case `0`, `no`:
			option.IsList = false
			break isListEnd
		default:
			isList = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成列表接口，默认(yes)：`))
		}
	}
	// 列表接口是否返回总数
	isCount, ok := optionMap[`isCount`]
	if !ok {
		isCount = gcmd.Scan(color.BlueString(`> 列表接口是否返回总数，默认(yes)：`))
	}
isCountEnd:
	for {
		switch isCount {
		case ``, `1`, `yes`:
			option.IsCount = true
			break isCountEnd
		case `0`, `no`:
			option.IsCount = false
			break isCountEnd
		default:
			isCount = gcmd.Scan(color.RedString(`    输入错误，请重新输入，列表接口是否返回总数，默认(yes)：`))
		}
	}
	// 是否生成详情接口
	isInfo, ok := optionMap[`isInfo`]
	if !ok {
		isInfo = gcmd.Scan(color.BlueString(`> 是否生成详情接口，默认(yes)：`))
	}
isInfoEnd:
	for {
		switch isInfo {
		case ``, `1`, `yes`:
			option.IsInfo = true
			break isInfoEnd
		case `0`, `no`:
			option.IsInfo = false
			break isInfoEnd
		default:
			isInfo = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成详情接口，默认(yes)：`))
		}
	}
	// 是否生成创建接口
	isCreate, ok := optionMap[`isCreate`]
	if !ok {
		isCreate = gcmd.Scan(color.BlueString(`> 是否生成创建接口，默认(yes)：`))
	}
isCreateEnd:
	for {
		switch isCreate {
		case ``, `1`, `yes`:
			option.IsCreate = true
			break isCreateEnd
		case `0`, `no`:
			option.IsCreate = false
			break isCreateEnd
		default:
			isCreate = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成创建接口，默认(yes)：`))
		}
	}
	// 是否生成更新接口
	isUpdate, ok := optionMap[`isUpdate`]
	if !ok {
		isUpdate = gcmd.Scan(color.BlueString(`> 是否生成更新接口，默认(yes)：`))
	}
isUpdateEnd:
	for {
		switch isUpdate {
		case ``, `1`, `yes`:
			option.IsUpdate = true
			break isUpdateEnd
		case `0`, `no`:
			option.IsUpdate = false
			break isUpdateEnd
		default:
			isUpdate = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成更新接口，默认(yes)：`))
		}
	}
	// 是否生成删除接口
	isDelete, ok := optionMap[`isDelete`]
	if !ok {
		isDelete = gcmd.Scan(color.BlueString(`> 是否生成删除接口，默认(yes)：`))
	}
isDeleteEnd:
	for {
		switch isDelete {
		case ``, `1`, `yes`:
			option.IsDelete = true
			break isDeleteEnd
		case `0`, `no`:
			option.IsDelete = false
			break isDeleteEnd
		default:
			isDelete = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成删除接口，默认(yes)：`))
		}
	}
	if !(option.IsList || option.IsInfo || option.IsCreate || option.IsUpdate || option.IsDelete) {
		fmt.Println(`请重新选择生成哪些接口，不能全是no！`)
		goto noAllRestart
	}
	// 是否生成后端接口文件
	isApi, ok := optionMap[`isApi`]
	if !ok {
		isApi = gcmd.Scan(color.BlueString(`> 是否生成后端接口文件，默认(yes)：`))
	}
isApiEnd:
	for {
		switch isApi {
		case ``, `1`, `yes`:
			option.IsApi = true
			break isApiEnd
		case `0`, `no`:
			option.IsApi = false
			break isApiEnd
		default:
			isApi = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成后端接口文件，默认(yes)：`))
		}
	}
	if option.IsApi {
		// 是否判断操作权限，如是，则同时会生成操作权限
		isAuthAction, ok := optionMap[`isAuthAction`]
		if !ok {
			isAuthAction = gcmd.Scan(color.BlueString(`> 是否判断操作权限，如是，则同时会生成操作权限，默认(yes)：`))
		}
	isAuthActionEnd:
		for {
			switch isAuthAction {
			case ``, `1`, `yes`:
				option.IsAuthAction = true
				break isAuthActionEnd
			case `0`, `no`:
				option.IsAuthAction = false
				break isAuthActionEnd
			default:
				isAuthAction = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否判断操作权限，如是，则同时会生成操作权限，默认(yes)：`))
			}
		}
	}
	// 是否生成前端视图文件
	isView, ok := optionMap[`isView`]
	if !ok {
		isView = gcmd.Scan(color.BlueString(`> 是否生成前端视图文件，默认(yes)：`))
	}
isViewEnd:
	for {
		switch isView {
		case ``, `1`, `yes`:
			option.IsView = true
			break isViewEnd
		case `0`, `no`:
			option.IsView = false
			break isViewEnd
		default:
			isView = gcmd.Scan(color.RedString(`    输入错误，请重新输入，是否生成前端视图文件，默认(yes)：`))
		}
	}

	return
}

// 生成代码
func (myGenThis *myGen) Handle() {
	genDao(myGenThis.tpl)   // dao模板生成
	genLogic(myGenThis.tpl) // logic模板生成

	if myGenThis.option.IsApi {
		genApi(myGenThis.option, myGenThis.tpl)                   // api模板生成
		genController(myGenThis.option, myGenThis.tpl)            // controller模板生成
		genAction(myGenThis.ctx, myGenThis.option, myGenThis.tpl) // 操作权限生成
		genRouter(myGenThis.option, myGenThis.tpl)                // 后端路由生成
	}

	if myGenThis.option.IsView {
		genViewIndex(myGenThis.option, myGenThis.tpl) // 视图模板Index生成
		genViewList(myGenThis.option, myGenThis.tpl)  // 视图模板List生成
		genViewQuery(myGenThis.option, myGenThis.tpl) // 视图模板Query生成
		myGenThis.genViewSave()                       // 视图模板Save生成
		genViewI18n(myGenThis.option, myGenThis.tpl)  // 视图模板I18n生成
		myGenThis.genViewRouter()                     // 前端路由生成

		command(`前端代码格式化`, false, gfile.SelfDir()+`/../view/`+myGenThis.option.SceneCode, `npm`, `run`, `format`) // 前端代码格式化
	}
}

// 视图模板Save生成
func (myGenThis *myGen) genViewSave() {
	tpl := myGenThis.tpl

	if !(myGenThis.option.IsCreate || myGenThis.option.IsUpdate) {
		return
	}

	type viewSave struct {
		importModule   []string
		paramHandle    []string
		dataInitBefore []string
		dataInitAfter  []string
		rule           []string
		form           []string
		formHandle     []string
	}
	viewSaveObj := viewSave{
		importModule:   []string{},
		paramHandle:    []string{},
		dataInitBefore: []string{},
		dataInitAfter:  []string{},
		rule:           []string{},
		form:           []string{},
		formHandle:     []string{},
	}

	type viewSaveItem struct {
		rule         []string
		required     string
		elDatePicker struct {
			defaultTime string
		}
	}
	for _, v := range tpl.FieldList {
		viewSaveItemObj := viewSaveItem{
			rule: []string{},
		}

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			continue
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` === undefined ? param.`+v.FieldRaw+` = 0 : null`)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', min: 0, trigger: 'change', message: t('validation.select') },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-cascader v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+tpl.ModuleDirCaseKebab+`/`+tpl.TableCaseKebab+`/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />
                </el-form-item>`)
			continue
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			if !garray.NewStrArrayFrom(viewSaveObj.importModule).Contains(`import md5 from 'js-md5'`) {
				viewSaveObj.importModule = append(viewSaveObj.importModule, `import md5 from 'js-md5'`)
			}
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` ? param.`+v.FieldRaw+` = md5(param.`+v.FieldRaw+`) : delete param.`+v.FieldRaw)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 6, max: 20, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 20 }) },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <label v-if="saveForm.data.idArr?.length">
                        <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				viewSaveItemObj.required = ` required: true,`
			}
			// 去掉该验证规则。有时会用到特殊符号
			// viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^(?!\d*$)[\p{L}\p{M}\p{N}_]+$/u, trigger: 'blur', message: t('validation.account') },`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'email', trigger: 'blur', message: t('validation.email') },`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'url', trigger: 'blur', message: t('validation.url') },`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			viewSaveObj.paramHandle = append(viewSaveObj.paramHandle, `param.`+v.FieldRaw+` === undefined ? param.`+v.FieldRaw+` = 0 : null`)
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', /* required: true, */ min: 1, trigger: 'change', message: t('validation.select') },
        ],`)
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-cascader v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/tree' }" :props="{ emitPath: false }" />
                </el-form-item>`)
			} else {
				viewSaveObj.dataInitAfter = append(viewSaveObj.dataInitAfter, v.FieldRaw+`: saveCommon.data.`+v.FieldRaw+` ? saveCommon.data.`+v.FieldRaw+` : undefined,`)
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-select v-model="saveForm.data.`+v.FieldRaw+`" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/`+apiUrl+`/list' }" />
                </el-form-item>`)
			}
			continue
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="`+gconv.String(defaultVal)+`" />
                    <label>
                        <el-alert :title="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.tip.`+v.FieldRaw+`')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			defaultVal := gconv.String(v.Default)
			if defaultVal == `` {
				defaultVal = v.StatusList[0][0]
			}
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: '`+defaultVal+`',`)
			} else {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+defaultVal+`,`)
			}

			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'enum', enum: (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`)

			//超过5个状态用select组件，小于5个用radio组件
			if len(v.StatusList) > 5 {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-select-v2 v-model="saveForm.data.`+v.FieldRaw+`" :options="tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`')" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="false" />
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-radio-group v-model="saveForm.data.`+v.FieldRaw+`">
                        <el-radio v-for="(item, index) in (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>`)
			}
			continue
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-switch v-model="saveForm.data.`+v.FieldRaw+`" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />
                </el-form-item>`)
			continue
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewSaveItemObj.elDatePicker.defaultTime = ` :default-time="new Date(2000, 0, 1, 23, 59, 59)"`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			if v.FieldType == TypeVarchar {
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },
        ],`)
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" type="textarea" :autosize="{ minRows: 3 }" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" />
                </el-form-item>`)
				continue
			}
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			multipleStr := ``
			if v.FieldType == TypeVarchar {
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],`)
			} else {
				multipleStr = ` :multiple="true"`
				requiredStr := ``
				if !v.IsNull {
					requiredStr = ` required: true,`
				}
				viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'array',`+requiredStr+` trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },
            // { type: 'array',`+requiredStr+` max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },
        ],`)
			}

			acceptStr := ` accept="image/*"`
			if v.FieldTypeName == TypeNameVideoSuffix {
				acceptStr = ` accept="video/*" :isImage="false"`
			}
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-upload v-model="saveForm.data.`+v.FieldRaw+`"`+multipleStr+acceptStr+` />
                </el-form-item>`)
			continue
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: [],`)
			requiredStr := ``
			if !v.IsNull {
				requiredStr = ` required: true,`
			}
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [
            { type: 'array',`+requiredStr+` trigger: 'change', message: t('validation.required') },
            // { type: 'array',`+requiredStr+` max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },
        ],`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-tag v-for="(item, index) in saveForm.data.`+v.FieldRaw+`" :type="`+v.FieldRaw+`Handle.tagType[index % `+v.FieldRaw+`Handle.tagType.length]" @close="`+v.FieldRaw+`Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="`+v.FieldRaw+`Handle.visible" :ref="(el: any) => `+v.FieldRaw+`Handle.ref = el" v-model="`+v.FieldRaw+`Handle.value" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" @keyup.enter="`+v.FieldRaw+`Handle.addValue" @blur="`+v.FieldRaw+`Handle.addValue" size="small" style="width: 100px;" :controls="false" /> -->
                    <el-input v-if="`+v.FieldRaw+`Handle.visible" :ref="(el: any) => `+v.FieldRaw+`Handle.ref = el" v-model="`+v.FieldRaw+`Handle.value" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" @keyup.enter="`+v.FieldRaw+`Handle.addValue" @blur="`+v.FieldRaw+`Handle.addValue" size="small" style="width: 100px;" />
                    <el-button v-else type="primary" size="small" @click="`+v.FieldRaw+`Handle.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>
                </el-form-item>`)
			viewSaveObj.formHandle = append(viewSaveObj.formHandle, `const `+v.FieldRaw+`Handle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        `+v.FieldRaw+`Handle.visible = true
        nextTick(() => {
            `+v.FieldRaw+`Handle.ref?.focus()
        })
    },
    addValue: () => {
        if (`+v.FieldRaw+`Handle.value) {
            saveForm.data.`+v.FieldRaw+`.push(`+v.FieldRaw+`Handle.value)
        }
        `+v.FieldRaw+`Handle.visible = false
        `+v.FieldRaw+`Handle.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.`+v.FieldRaw+`.splice(saveForm.data.`+v.FieldRaw+`.indexOf(item), 1)
    },
})`)
			continue
		}
		/*--------根据字段命名类型处理 结束--------*/

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			defaultVal := gconv.Int(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'integer', trigger: 'change', message: t('validation.input') },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeIntU: // `int等类型（unsigned）`
			defaultVal := gconv.Uint(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="0" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeFloat: // `float等类型`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },    // 类型float值为0时验证不能通过`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :precision="`+v.FieldLimitFloat[1]+`" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeFloatU: // `float等类型（unsigned）`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveObj.dataInitBefore = append(viewSaveObj.dataInitBefore, v.FieldRaw+`: `+gconv.String(defaultVal)+`,`)
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{ type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },    // 类型float值为0时验证不能通过`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input-number v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :min="0" :precision="`+v.FieldLimitFloat[1]+`" :controls="false" :value-on-clear="`+gconv.String(defaultVal)+`" />
                </el-form-item>`)
		case TypeVarchar: // `varchar类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` max: ` + v.FieldLimitStr + `, trigger: 'blur', message: t('validation.max.string', { max: ` + v.FieldLimitStr + ` }) },`}, viewSaveItemObj.rule...)

			if v.IndexRaw == `UNI` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" />
                </el-form-item>`)
			}
		case TypeChar: // `char类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` len: ` + v.FieldLimitStr + `, trigger: 'blur', message: t('validation.size.string', { size: ` + v.FieldLimitStr + ` }) },`}, viewSaveItemObj.rule...)

			if v.IndexRaw == `UNI` {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" />
                </el-form-item>`)
			} else {
				viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" minlength="`+v.FieldLimitStr+`" maxlength="`+v.FieldLimitStr+`" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>`)
			}
		case TypeText: // `text类型`
			viewSaveItemObj.rule = append([]string{`{ type: 'string', trigger: 'blur', message: t('validation.input') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <my-editor v-model="saveForm.data.`+v.FieldRaw+`" />
                </el-form-item>`)
		case TypeJson: // `json类型`
			if !v.IsNull {
				viewSaveItemObj.required = `
                required: true,`
			}
			viewSaveItemObj.rule = append(viewSaveItemObj.rule, `{
                type: 'object',`+viewSaveItemObj.required+`
                /* fields: {
                    xxxx: { type: 'string', required: true, message: 'xxxx' + t('validation.required') },
                    xxxx: { type: 'integer', required: true, min: 1, message: 'xxxx' + t('validation.min.number', { min: 1 }) },
                }, */
                transform(value: any) {
                    if (value === '' || value === null || value === undefined) {
                        return undefined
                    }
                    try {
                        return JSON.parse(value)
                    } catch (e) {
                        return value
                    }
                },
                trigger: 'blur',
                message: t('validation.json'),
            },`)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-alert :title="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.tip.`+v.FieldRaw+`')" type="info" :show-icon="true" :closable="false" />
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" type="textarea" :autosize="{ minRows: 3 }" />
                </el-form-item>`)
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` trigger: 'change', message: t('validation.select') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-date-picker v-model="saveForm.data.`+v.FieldRaw+`" type="datetime" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss"`+viewSaveItemObj.elDatePicker.defaultTime+` />
                </el-form-item>`)
		case TypeDate: // `date类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveItemObj.required = ` required: true,`
			}
			viewSaveItemObj.rule = append([]string{`{ type: 'string',` + viewSaveItemObj.required + ` trigger: 'change', message: t('validation.select') },`}, viewSaveItemObj.rule...)
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-date-picker v-model="saveForm.data.`+v.FieldRaw+`" type="date" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />
                </el-form-item>`)
		default:
			viewSaveObj.form = append(viewSaveObj.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    <el-input v-model="saveForm.data.`+v.FieldRaw+`" :placeholder="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" :clearable="true" />
                </el-form-item>`)
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		if len(viewSaveItemObj.rule) > 0 {
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [`+gstr.Join(append([]string{``}, viewSaveItemObj.rule...), `
			`)+`
        ],`)
		} else {
			viewSaveObj.rule = append(viewSaveObj.rule, v.FieldRaw+`: [],`)
		}
	}

	tplView := `<script setup lang="tsx">` + gstr.Join(append([]string{``}, viewSaveObj.importModule...), `
`) + `
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {` + gstr.Join(append([]string{``}, viewSaveObj.dataInitBefore...), `
        `) + `
        ...saveCommon.data,` + gstr.Join(append([]string{``}, viewSaveObj.dataInitAfter...), `
        `) + `
    } as { [propName: string]: any },
    rules: {` + gstr.Join(append([]string{``}, viewSaveObj.rule...), `
        `) + `
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + gstr.Join(append([]string{``}, viewSaveObj.paramHandle...), `
            `) + `
            try {
                if (param?.idArr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } catch (error) {}
            saveForm.loading = false
        })
    },
})

const saveDrawer = reactive({
    ref: null as any,
    size: useSettingStore().saveDrawer.size,
    beforeClose: (done: Function) => {
        if (useSettingStore().saveDrawer.isTipClose) {
            ElMessageBox.confirm('', {
                type: 'info',
                title: t('common.tip.configExit'),
                center: true,
                showClose: false,
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})` + gstr.Join(append([]string{``}, viewSaveObj.formHandle...), `

`) + `
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + gstr.Join(append([]string{``}, viewSaveObj.form...), `
                `) + `
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Save.vue`
	gfile.PutContents(saveFile, tplView)
}

// 前端路由生成
func (myGenThis *myGen) genViewRouter() {
	tpl := myGenThis.tpl

	saveFile := gfile.SelfDir() + `/../view/` + myGenThis.option.SceneCode + `/src/router/index.ts`
	tplViewRouter := gfile.GetContents(saveFile)

	path := `/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab
	replaceStr := `{
                path: '` + path + `',
                component: async () => {
                    const component = await import('@/views` + path + `/Index.vue')
                    component.default.name = '` + path + `'
                    return component
                },
                meta: { isAuth: true, keepAlive: true, componentName: '` + path + `' }
            },`

	if gstr.Pos(tplViewRouter, `'`+path+`'`) == -1 { //路由不存在时新增
		tplViewRouter = gstr.Replace(tplViewRouter, `/*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, replaceStr+`
            /*--------前端路由自动代码生成锚点（不允许修改和删除，否则将不能自动生成路由）--------*/`, 1)
	} else { //路由已存在则替换
		tplViewRouter, _ = gregex.ReplaceString(`\{
                path: '`+path+`',[\s\S]*'`+path+`' \}
            \},`, replaceStr, tplViewRouter)
	}
	gfile.PutContents(saveFile, tplViewRouter)

	myGenThis.genMenu(myGenThis.sceneInfo[daoAuth.Scene.Columns().SceneId].Uint(), path, myGenThis.option.CommonName, tpl.TableCaseCamel) // 数据库权限菜单处理
}

// 自动生成菜单
func (myGenThis *myGen) genMenu(sceneId uint, menuUrl string, menuName string, menuNameOfEn string) {
	ctx := myGenThis.ctx

	menuNameArr := gstr.Split(menuName, `/`)

	var pid int64 = 0
	for _, v := range menuNameArr[:len(menuNameArr)-1] {
		pidVar, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{
			daoAuth.Menu.Columns().SceneId:  sceneId,
			daoAuth.Menu.Columns().MenuName: v,
		}).Value(daoAuth.Menu.PrimaryKey())
		if pidVar.Uint() == 0 {
			pid, _ = daoAuth.Menu.CtxDaoModel(ctx).HookInsert(g.Map{
				daoAuth.Menu.Columns().SceneId:   sceneId,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().MenuName:  v,
				daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
				daoAuth.Menu.Columns().MenuUrl:   ``,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "", "zh-cn": "` + v + `"}}}`,
			}).InsertAndGetId()
		} else {
			pid = pidVar.Int64()
		}
	}

	menuName = menuNameArr[len(menuNameArr)-1]
	idVar, _ := daoAuth.Menu.CtxDaoModel(ctx).Filters(g.Map{
		daoAuth.Menu.Columns().SceneId: sceneId,
		daoAuth.Menu.Columns().MenuUrl: menuUrl,
	}).Value(daoAuth.Menu.PrimaryKey())
	id := idVar.Uint()
	if id == 0 {
		daoAuth.Menu.CtxDaoModel(ctx).HookInsert(g.Map{
			daoAuth.Menu.Columns().SceneId:   sceneId,
			daoAuth.Menu.Columns().Pid:       pid,
			daoAuth.Menu.Columns().MenuName:  menuName,
			daoAuth.Menu.Columns().MenuIcon:  `autoicon-ep-link`,
			daoAuth.Menu.Columns().MenuUrl:   menuUrl,
			daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
		}).Insert()
	} else {
		daoAuth.Menu.CtxDaoModel(ctx).Filter(daoAuth.Menu.PrimaryKey(), id).
			SetIdArr().
			HookUpdate(g.Map{
				daoAuth.Menu.Columns().MenuName:  menuName,
				daoAuth.Menu.Columns().Pid:       pid,
				daoAuth.Menu.Columns().ExtraData: `{"i18n": {"title": {"en": "` + menuNameOfEn + `", "zh-cn": "` + menuName + `"}}}`,
			}).Update()
	}
}
