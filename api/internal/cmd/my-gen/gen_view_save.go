package my_gen

import (
	"api/internal/cmd/my-gen/internal"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewSave struct {
	isI18nTm       bool
	importModule   []string
	dataInitBefore []string
	dataInitAfter  []string
	rule           []string
	formItem       []string
	formContent    []string
	formHandle     []string
	paramHandle    []string
}

type myGenViewSaveField struct {
	isI18nTm       bool
	importModule   []string
	dataInitBefore internal.MyGenDataStrHandler
	dataInitAfter  internal.MyGenDataStrHandler
	isRequired     bool
	rule           internal.MyGenDataSliceHandler
	formContent    internal.MyGenDataStrHandler
	formHandle     internal.MyGenDataStrHandler
	paramHandle    internal.MyGenDataStrHandler
}

func (viewSaveThis *myGenViewSave) Add(viewSaveField myGenViewSaveField, field string, i18nPath string, i18nFieldPath string, tableType internal.MyGenTableType, fieldIf string) {
	if viewSaveField.formContent.GetData() == `` {
		return
	}
	viewSaveThis.importModule = append(viewSaveThis.importModule, viewSaveField.importModule...)
	if viewSaveField.isI18nTm {
		viewSaveThis.isI18nTm = true
	}
	if viewSaveField.dataInitBefore.GetData() != `` {
		viewSaveThis.dataInitBefore = append(viewSaveThis.dataInitBefore, field+`: `+viewSaveField.dataInitBefore.GetData()+`,`)
	}
	if viewSaveField.dataInitAfter.GetData() != `` {
		viewSaveThis.dataInitAfter = append(viewSaveThis.dataInitAfter, field+`: `+viewSaveField.dataInitAfter.GetData()+`,`)
	}
	rule := viewSaveField.rule.GetData()
	if viewSaveField.isRequired {
		if fieldIf != `` && tableType == internal.TableTypeMiddleOne {
			rule = append([]string{`{ required: computed((): boolean => ` + fieldIf + `), message: t('validation.required') },`}, rule...)
		} else {
			switch tableType {
			case internal.TableTypeExtendMany, internal.TableTypeMiddleMany:
				rule = append([]string{`{ required: true, message: t('` + i18nPath + `.name.` + i18nFieldPath + `.` + field + `') + t('validation.required') },`}, rule...)
			default:
				rule = append([]string{`{ required: true, message: t('validation.required') },`}, rule...)
			}
		}
	}
	if len(rule) > 0 {
		viewSaveThis.rule = append(viewSaveThis.rule, field+`: [`+gstr.Join(append([]string{``}, rule...), `
            `)+`
        ],`)
	} else {
		viewSaveThis.rule = append(viewSaveThis.rule, field+`: [],`)
	}
	if fieldIf == `` {
		viewSaveThis.formItem = append(viewSaveThis.formItem, `<el-form-item :label="t('`+i18nPath+`.name.`+i18nFieldPath+`')" prop="`+field+`">
                    {{formContent}}
                </el-form-item>`)
		viewSaveThis.formContent = append(viewSaveThis.formContent, viewSaveField.formContent.GetData())
	} else {
		viewSaveThis.formItem = append(viewSaveThis.formItem, `<el-form-item v-if="`+fieldIf+`" :label="t('`+i18nPath+`.name.`+i18nFieldPath+`')" prop="`+field+`">
					{{formContent}}
				</el-form-item>`)
		if tableType == internal.TableTypeMiddleMany {
			formContent := gstr.TrimStr(viewSaveField.formContent.GetData(), ` `)
			formContent = gstr.Replace(formContent, ` `, ` v-if="`+fieldIf+`"`, 1)
			/* switch gstr.Split(formContent, ` `)[0] {
			case `<el-input`:
				formContent = gstr.SubStr(formContent, 0, -2) + `style="width: 200px;" />`
			case `<el-input-number`:
				formContent = gstr.SubStr(formContent, 0, -2) + `style="width: 150px;" />`
			} */
			viewSaveThis.formContent = append(viewSaveThis.formContent, formContent)
		} else {
			viewSaveThis.formContent = append(viewSaveThis.formContent, viewSaveField.formContent.GetData())
		}
	}
	if viewSaveField.formHandle.GetData() != `` {
		viewSaveThis.formHandle = append(viewSaveThis.formHandle, viewSaveField.formHandle.GetData())
	}
	if viewSaveField.paramHandle.GetData() != `` {
		viewSaveThis.paramHandle = append(viewSaveThis.paramHandle, viewSaveField.paramHandle.GetData())
	}
}

func (viewSaveThis *myGenViewSave) Merge(viewSaveOther myGenViewSave) {
	viewSaveThis.importModule = append(viewSaveThis.importModule, viewSaveOther.importModule...)
	viewSaveThis.dataInitBefore = append(viewSaveThis.dataInitBefore, viewSaveOther.dataInitBefore...)
	viewSaveThis.dataInitAfter = append(viewSaveThis.dataInitAfter, viewSaveOther.dataInitAfter...)
	viewSaveThis.rule = append(viewSaveThis.rule, viewSaveOther.rule...)
	viewSaveThis.formItem = append(viewSaveThis.formItem, viewSaveOther.formItem...)
	viewSaveThis.formContent = append(viewSaveThis.formContent, viewSaveOther.formContent...)
	viewSaveThis.formHandle = append(viewSaveThis.formHandle, viewSaveOther.formHandle...)
	viewSaveThis.paramHandle = append(viewSaveThis.paramHandle, viewSaveOther.paramHandle...)
}

func (viewSaveThis *myGenViewSave) Unique() {
	viewSaveThis.importModule = garray.NewStrArrayFrom(viewSaveThis.importModule).Unique().Slice()
	// viewSaveThis.dataInitBefore = garray.NewStrArrayFrom(viewSaveThis.dataInitBefore).Unique().Slice()
	// viewSaveThis.dataInitAfter = garray.NewStrArrayFrom(viewSaveThis.dataInitAfter).Unique().Slice()
	// viewSaveThis.rule = garray.NewStrArrayFrom(viewSaveThis.rule).Unique().Slice()
	// viewSaveThis.formItem = garray.NewStrArrayFrom(viewSaveThis.formItem).Unique().Slice()
	// viewSaveThis.formContent = garray.NewStrArrayFrom(viewSaveThis.formContent).Unique().Slice()
	// viewSaveThis.formHandle = garray.NewStrArrayFrom(viewSaveThis.formHandle).Unique().Slice()
	// viewSaveThis.paramHandle = garray.NewStrArrayFrom(viewSaveThis.paramHandle).Unique().Slice()
}

func (viewSaveThis *myGenViewSave) CreateForm() (form []string) {
	for k, v := range viewSaveThis.formContent {
		form = append(form, gstr.Replace(viewSaveThis.formItem[k], `{{formContent}}`, v))
	}
	return
}

// 视图模板Save生成
func genViewSave(option myGenOption, tpl myGenTpl) {
	if !(option.IsCreate || option.IsUpdate) {
		return
	}
	viewSave := myGenViewSave{}
	for _, v := range tpl.FieldListOfDefault {
		viewSave.Add(getViewSaveField(tpl, v, v.FieldRaw, tpl.I18nPath, v.FieldRaw), v.FieldRaw, tpl.I18nPath, v.FieldRaw, internal.TableTypeDefault, ``)
	}
	for _, v := range tpl.FieldListOfAfter1 {
		viewSave.Add(getViewSaveField(tpl, v, v.FieldRaw, tpl.I18nPath, v.FieldRaw), v.FieldRaw, tpl.I18nPath, v.FieldRaw, internal.TableTypeDefault, ``)
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		viewSave.Merge(getViewSaveExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		viewSave.Merge(getViewSaveExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		viewSave.Merge(getViewSaveExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		viewSave.Merge(getViewSaveExtendMiddleMany(v))
	}
	for _, v := range tpl.FieldListOfAfter2 {
		viewSave.Add(getViewSaveField(tpl, v, v.FieldRaw, tpl.I18nPath, v.FieldRaw), v.FieldRaw, tpl.I18nPath, v.FieldRaw, internal.TableTypeDefault, ``)
	}
	viewSave.Unique()

	tplView := `<script setup lang="tsx">` + gstr.Join(append([]string{``}, viewSave.importModule...), `
`) + `
const { t`
	if viewSave.isI18nTm {
		tplView += `, tm`
	}
	tplView += ` } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {` + gstr.Join(append([]string{``}, viewSave.dataInitBefore...), `
        `) + `
        ...saveCommon.data,` + gstr.Join(append([]string{``}, viewSave.dataInitAfter...), `
        `) + `
    } as { [propName: string]: any },
    rules: {` + gstr.Join(append([]string{``}, viewSave.rule...), `
        `) + `
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + gstr.Join(append([]string{``}, viewSave.paramHandle...), `
            `) + `
            try {
                if (param?.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } finally {
                saveForm.loading = false
            }
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
            }).then(() => done())
        } else {
            done()
        }
    },
    buttonClose: () => saveDrawer.ref.handleClose(), //saveCommon.visible = false //不会触发beforeClose
})` + gstr.Join(append([]string{``}, viewSave.formHandle...), `

`) + `
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + gstr.Join(append([]string{``}, viewSave.CreateForm()...), `
                `) + `
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Save.vue`
	gfile.PutContents(saveFile, tplView)
}

func getViewSaveField(tpl myGenTpl, v myGenField, dataFieldPath string, i18nPath string, i18nFieldPath string) (viewSaveField myGenViewSaveField) {
	if !v.IsNull && (gvar.New(v.Default).IsNil() || v.IsUnique) {
		viewSaveField.isRequired = true
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt, internal.TypeIntU: // `int等类型`	// `int等类型（unsigned）`
		defaultVal := gconv.String(gconv.Int(v.Default))
		if defaultVal != `0` {
			viewSaveField.dataInitBefore.Method = internal.ReturnType
			viewSaveField.dataInitBefore.DataType = defaultVal
		}
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'integer', trigger: 'change', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.between.number', { min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+` }) },`)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :min="` + v.FieldLimitInt.Min + `" :max="` + v.FieldLimitInt.Max + `" :precision="0" :controls="false" :value-on-clear="` + defaultVal + `" />`
	case internal.TypeFloat, internal.TypeFloatU: // `float等类型`	 // `float等类型（unsigned）`
		defaultVal := gconv.Float64(v.Default)
		if defaultVal != 0 {
			viewSaveField.dataInitBefore.Method = internal.ReturnType
			viewSaveField.dataInitBefore.DataType = gconv.String(v.Default)
		}
		rule := `{ type: 'number', trigger: 'change', message: t('validation.input') },`
		attrOfAdd := ``
		if v.FieldLimitFloat.Min != `` && v.FieldLimitFloat.Max != `` {
			rule = `{ type: 'number', trigger: 'change', min: ` + v.FieldLimitFloat.Min + `, max: ` + v.FieldLimitFloat.Max + `, message: t('validation.between.number', { min: ` + v.FieldLimitFloat.Min + `, max: ` + v.FieldLimitFloat.Max + ` }) },`
			attrOfAdd = ` :min="` + v.FieldLimitFloat.Min + `" :max="` + v.FieldLimitFloat.Max + `"`
		} else if v.FieldLimitFloat.Min != `` {
			rule = `{ type: 'number', trigger: 'change', min: ` + v.FieldLimitFloat.Min + `, message: t('validation.min.number', { min: ` + v.FieldLimitFloat.Min + ` }) },`
			attrOfAdd = ` :min="` + v.FieldLimitFloat.Min + `"`
		} else if v.FieldLimitFloat.Max != `` {
			rule = `{ type: 'number', trigger: 'change', max: ` + v.FieldLimitFloat.Max + `, message: t('validation.max.number', { max: ` + v.FieldLimitFloat.Max + ` }) },`
			attrOfAdd = ` :max="` + v.FieldLimitFloat.Max + `"`
		}
		rule += `    // type: 'float'在值为0时验证不能通过`
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, rule)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')"` + attrOfAdd + ` :precision="` + gconv.String(v.FieldLimitFloat.Precision) + `" :controls="false" :value-on-clear="` + gconv.String(v.Default) + `" />`
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
		rule := `{ type: 'string', trigger: 'blur', max: ` + v.FieldLimitStr + `, message: t('validation.max.string', { max: ` + v.FieldLimitStr + ` }) },`
		attrOfAdd := ``
		if v.FieldType == internal.TypeChar {
			rule = `{ type: 'string', trigger: 'blur', len: ` + v.FieldLimitStr + `, message: t('validation.size.string', { size: ` + v.FieldLimitStr + ` }) },`
			attrOfAdd = ` minlength="` + v.FieldLimitStr + `"`
		}
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, rule)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')"` + attrOfAdd + ` maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
		if v.IsUnique {
			viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')"` + attrOfAdd + ` maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />`
		}
	case internal.TypeText: // `text类型`
		if !v.IsNull {
			viewSaveField.isRequired = true
		}
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'blur', message: t('validation.input') },`)

		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<my-editor v-model="saveForm.data.` + dataFieldPath + `" />`
	case internal.TypeJson: // `json类型`
		if !v.IsNull {
			viewSaveField.isRequired = true
		}
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                // fields: { xxxx: [{ required: true, message: 'xxxx' + t('validation.required') }] }, //内部添加规则时，不再需要设置trigger属性
                transform: (value: any) => {
                    if (!value) {
                        return undefined
                    }
                    try {
                        return JSON.parse(value)
                    } catch (error) {
                        return value
                    }
                },
            },`)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" type="textarea" :autosize="{ minRows: 3 }" />`
		if v.FieldTip != `` {
			viewSaveField.formContent.DataType = `<el-alert :title="t('` + i18nPath + `.tip.` + i18nFieldPath + `')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    ` + viewSaveField.formContent.DataType
		}
	case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-date-picker v-model="saveForm.data.` + dataFieldPath + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />`
	case internal.TypeDate: // `date类型`
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-date-picker v-model="saveForm.data.` + dataFieldPath + `" type="date" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />`
	case internal.TypeTime: // `time类型`
		viewSaveField.rule.Method = internal.ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-time-picker v-model="saveForm.data.` + dataFieldPath + `" placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="HH:mm:ss" value-format="HH:mm:ss" />`
	default:
		viewSaveField.formContent.Method = internal.ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="true" />`
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		return myGenViewSaveField{}
	case internal.TypePrimaryMany: // 联合主键
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
		return myGenViewSaveField{}
	case internal.TypeNameUpdated: // 更新时间字段
		return myGenViewSaveField{}
	case internal.TypeNameCreated: // 创建时间字段
		return myGenViewSaveField{}
	case internal.TypeNamePid: // pid；	类型：int等类型；
		viewSaveField.dataInitAfter.Method = internal.ReturnTypeName
		viewSaveField.dataInitAfter.DataTypeName = `saveCommon.data.` + dataFieldPath + ` ? saveCommon.data.` + dataFieldPath + ` : undefined`
		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', trigger: 'change', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<my-cascader v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree', param: { filter: { ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `exc_id_arr`) + `: saveForm.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + ` } } }" :props="{ checkStrictly: true, emitPath: false }" />`
		viewSaveField.paramHandle.Method = internal.ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` === undefined && (param.` + dataFieldPath + ` = 0)`
	case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		return myGenViewSaveField{}
	case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		return myGenViewSaveField{}
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		viewSaveField.importModule = append(viewSaveField.importModule, `import md5 from 'js-md5'`)
		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName,
			`{ required: computed((): boolean => (saveForm.data.`+internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`)+`?.length ? false : true)), message: t('validation.required') },`,
			`{ type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },`,
		)
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert v-if="saveForm.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `?.length" :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />`
		viewSaveField.paramHandle.Method = internal.ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` ? param.` + dataFieldPath + ` = md5(param.` + dataFieldPath + `) : delete param.` + dataFieldPath
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenViewSaveField{}
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		if gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
			viewSaveField.isRequired = true
		}
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
		viewSaveField.rule.Method = internal.ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },`)
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
		viewSaveField.rule.Method = internal.ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', trigger: 'blur', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },`)
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		viewSaveField.rule.Method = internal.ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', trigger: 'blur', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },`)
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
		viewSaveField.rule.Method = internal.ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'email', trigger: 'blur', message: t('validation.email') },`)
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		viewSaveField.rule.Method = internal.ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'url', trigger: 'blur', message: t('validation.url') },`)
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-color-picker v-model="saveForm.data.` + dataFieldPath + `" :show-alpha="true" />`
		viewSaveField.paramHandle.Method = internal.ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` === undefined && (param.` + dataFieldPath + ` = '')`
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		viewSaveField.dataInitAfter.Method = internal.ReturnTypeName
		viewSaveField.dataInitAfter.DataTypeName = `saveCommon.data.` + dataFieldPath + ` ? saveCommon.data.` + dataFieldPath + ` : undefined`
		viewSaveField.rule.Method = internal.ReturnTypeName
		if !viewSaveField.isRequired {
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `// { required: true, message: t('validation.required') },`)
		}
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', trigger: 'change', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnTypeName
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` {
			apiUrl := relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			if relIdObj.tpl.Handle.Pid.Pid != `` {
				viewSaveField.formContent.DataTypeName = `<my-cascader v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />`
			} else {
				viewSaveField.formContent.DataTypeName = `<my-select v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />`
			}
		} else {
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			viewSaveField.formContent.DataTypeName = `<!-- 可选择组件<my-select>或<my-cascader>使用，但需手动确认关联表，并修改接口路径 -->
                    ` + viewSaveField.formContent.DataType + `
                    <!-- <my-select v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" /> -->
                    <!-- <my-cascader v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" /> -->`
		}
		viewSaveField.paramHandle.Method = internal.ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` === undefined && (param.` + dataFieldPath + ` = 0)`
	case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', trigger: 'change', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.between.number', { min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+` }) },`)
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :min="` + v.FieldLimitInt.Min + `" :max="` + v.FieldLimitInt.Max + `" :precision="0" :value-on-clear="` + gconv.String(gconv.Int(v.Default)) + `" />`
		if v.FieldTip != `` {
			viewSaveField.formContent.DataTypeName += `
                    <el-alert :title="t('` + i18nPath + `.tip.` + i18nFieldPath + `')" type="info" :show-icon="true" :closable="false" />`
		}
	case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		viewSaveField.isI18nTm = true
		defaultVal := gconv.String(v.Default)
		if defaultVal == `` {
			defaultVal = v.StatusList[0][0]
		}
		viewSaveField.dataInitBefore.Method = internal.ReturnTypeName
		viewSaveField.dataInitBefore.DataTypeName = defaultVal
		if garray.NewIntArrayFrom([]int{internal.TypeVarchar, internal.TypeChar}).Contains(v.FieldType) {
			viewSaveField.dataInitBefore.DataTypeName = `'` + defaultVal + `'`
		}
		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'enum', trigger: 'change', enum: (tm('`+i18nPath+`.status.`+i18nFieldPath+`') as any).map((item: any) => item.value), message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-radio-group v-model="saveForm.data.` + dataFieldPath + `">
                        <el-radio v-for="(item, index) in (tm('` + i18nPath + `.status.` + i18nFieldPath + `') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>`
		if len(v.StatusList) > 5 { //超过5个状态用select组件，小于5个用radio组件
			viewSaveField.formContent.DataTypeName = `<el-select-v2 v-model="saveForm.data.` + dataFieldPath + `" :options="tm('` + i18nPath + `.status.` + i18nFieldPath + `')" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="false" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		}
	case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
		viewSaveField.isI18nTm = true
		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') },`)
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-switch v-model="saveForm.data.` + dataFieldPath + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />`
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		switch v.FieldType {
		case internal.TypeDatetime, internal.TypeTimestamp:
			viewSaveField.formContent.Method = internal.ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-date-picker v-model="saveForm.data.` + dataFieldPath + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :default-time="new Date(2000, 0, 1, 23, 59, 59)" />`
		case internal.TypeDate:
		case internal.TypeTime:
			viewSaveField.formContent.Method = internal.ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-time-picker v-model="saveForm.data.` + dataFieldPath + `" placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="HH:mm:ss" value-format="HH:mm:ss" :default-value="new Date(2000, 0, 1, 23, 59, 59)" />`
		}
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		if v.FieldType == internal.TypeVarchar {
			viewSaveField.formContent.Method = internal.ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-input v-model="saveForm.data.` + dataFieldPath + `" type="textarea" :autosize="{ minRows: 3 }" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" />`
		}
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
		if v.FieldType == internal.TypeVarchar {
			viewSaveField.rule.Method = internal.ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'url', trigger: 'change', message: t('validation.upload') },`)
		} else {
			viewSaveField.rule.Method = internal.ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },	// 限制数组数量时用：max: 10, message: t('validation.max.upload', { max: 10 })`)
		}
		attrOfAdd := ``
		if v.FieldType != internal.TypeVarchar {
			attrOfAdd += ` :multiple="true"`
		}
		switch v.FieldTypeName {
		case internal.TypeNameImageSuffix:
			attrOfAdd += ` accept="image/*"`
		case internal.TypeNameVideoSuffix:
			attrOfAdd += ` accept="video/*" show-type="video"`
		case internal.TypeNameAudioSuffix:
			attrOfAdd += ` accept="audio/*" show-type="audio"`
		case internal.TypeNameFileSuffix:
			attrOfAdd += ` show-type="text"`
		}
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<my-upload v-model="saveForm.data.` + dataFieldPath + `"` + attrOfAdd + ` />`
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		viewSaveField.isI18nTm = true
		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'blur', message: t('validation.array'), defaultField: { type: 'string', message: t('validation.input') } },	// 限制数组数量时用：max: 10, message: t('validation.max.array', { max: 10 })`)
		fieldHandle := gstr.CaseCamelLower(gstr.Replace(dataFieldPath, `.`, `_`)) + `Handle`
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<template v-for="(_, index) in saveForm.data.` + dataFieldPath + `" :key="index">
                    <el-tag type="info" :closable="true" @close="` + fieldHandle + `.del(index)" size="large" style="padding-left: 0; margin: 3px 10px 3px 0;">
                        <el-input :ref="(el: any) => ` + fieldHandle + `.ref[index] = el" v-model="saveForm.data.` + dataFieldPath + `[index]" @blur="` + fieldHandle + `.del(index, true)" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" style="width: 150px" />
                        <!-- <el-input-number :ref="(el: any) => ` + fieldHandle + `.ref[index] = el" v-model="saveForm.data.` + dataFieldPath + `[index]" @blur="` + fieldHandle + `.del(index, true)" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :controls="false" style="width: 150px" /> -->
                    </el-tag>
                </template>
                <el-button type="primary" @click="` + fieldHandle + `.add" style="margin: 3px 0"> <autoicon-ep-plus />{{ t('common.add') }} </el-button>`
		viewSaveField.formHandle.Method = internal.ReturnTypeName
		viewSaveField.formHandle.DataTypeName = `const ` + fieldHandle + ` = reactive({
    ref: [] as any[],
    add: () => {
        !Array.isArray(saveForm.data.` + dataFieldPath + `) && (saveForm.data.` + dataFieldPath + ` = [])
        saveForm.data.` + dataFieldPath + `.push(undefined)
        nextTick(() => ` + fieldHandle + `.ref[` + fieldHandle + `.ref.length - 1].focus())
    },
    del: (index: number, isBlur: boolean = false) => {
        if (isBlur && saveForm.data.` + dataFieldPath + `[index] !== undefined && saveForm.data.` + dataFieldPath + `[index] !== null && saveForm.data.` + dataFieldPath + `[index] !== '') {
            return
        }
        saveForm.data.` + dataFieldPath + `.splice(index, 1)
        ` + fieldHandle + `.ref.splice(index, 1)
    },
})`
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getViewSaveExtendMiddleOne(tplEM handleExtendMiddle) (viewSave myGenViewSave) {
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			viewSave.Add(getViewSaveField(tplEM.tpl, v, v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw), v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw, tplEM.TableType, ``)
		}
	case internal.TableTypeMiddleOne:
		for _, v := range tplEM.FieldListOfIdSuffix {
			viewSave.Add(getViewSaveField(tplEM.tpl, v, v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw), v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw, tplEM.TableType, ``)
		}
		if len(tplEM.FieldListOfOther) > 0 {
			fieldIfArr := []string{}
			for _, v := range tplEM.FieldListOfIdSuffix {
				fieldIfArr = append(fieldIfArr, `saveForm.data.`+v.FieldRaw)
			}
			fieldIf := gstr.Join(fieldIfArr, ` || `)
			for _, v := range tplEM.FieldListOfOther {
				viewSave.Add(getViewSaveField(tplEM.tpl, v, v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw), v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw, tplEM.TableType, fieldIf)
			}
		}
	}
	return
}

func getViewSaveExtendMiddleMany(tplEM handleExtendMiddle) (viewSave myGenViewSave) {
	if len(tplEM.FieldList) == 1 {
		v := tplEM.FieldList[0]

		tpl := tplEM.tpl
		i18nPath := tplEM.tplOfTop.I18nPath
		i18nFieldPath := tplEM.FieldVar

		viewSaveField := myGenViewSaveField{}
		/*--------部分命名类型直接处理后返回 开始--------*/
		isReturn := false
		switch v.FieldTypeName {
		case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
			isReturn = true

			viewSaveField.dataInitAfter.Method = internal.ReturnTypeName
			viewSaveField.dataInitAfter.DataTypeName = `saveCommon.data.` + tplEM.FieldVar + ` ? saveCommon.data.` + tplEM.FieldVar + ` : undefined`
			viewSaveField.rule.Method = internal.ReturnTypeName
			rule := []string{`{ required: true, message: t('validation.required') },`}
			if v.FieldType == internal.TypeVarchar {
				rule = append(rule, `{ type: 'string', max: `+v.FieldLimitStr+`, message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },`)
			} else {
				rule = append(rule, `{ type: 'string', len: `+v.FieldLimitStr+`, message: t('validation.size.string', { size: `+v.FieldLimitStr+` }) },`)
			}
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.array'), defaultField: [`+gstr.Join(append([]string{``}, rule...), `
					`)+`] },	// 限制数组数量时用：max: 10, message: t('validation.max.array', { max: 10 })`)
			viewSaveField.formContent.Method = internal.ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-space :size="10">
                        <el-color-picker v-for="(item, index) in saveForm.data.` + tplEM.FieldVar + `" :key="index" v-model="saveForm.data.` + tplEM.FieldVar + `[index]" :show-alpha="true" @change="(val) => (val ? null : saveForm.data.` + tplEM.FieldVar + `.splice(index, 1))" />
                        <el-button v-if="saveForm.data.` + tplEM.FieldVar + `.length == 0 || saveForm.data.` + tplEM.FieldVar + `[saveForm.data.` + tplEM.FieldVar + `.length - 1]" type="primary" size="small" @click="saveForm.data.` + tplEM.FieldVar + `.push('')">
                            <autoicon-ep-plus />{{ t('common.add') }}
                        </el-button>
                    </el-space>`
		case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			isReturn = true

			viewSaveField.isI18nTm = true
			viewSaveField.rule.Method = internal.ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'enum', enum: (tm('`+i18nPath+`.status.`+i18nFieldPath+`') as any).map((item: any) => item.value), message: t('validation.select') } },	// 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })`)

			viewSaveField.formContent.Method = internal.ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<!-- 根据个人喜好选择组件<el-transfer>或<el-select-v2> -->
                    <el-transfer v-model="saveForm.data.` + tplEM.FieldVar + `" :data="tm('` + i18nPath + `.status.` + i18nFieldPath + `')" :props="{ key: 'value', label: 'label' }" />
                    <!-- <el-select-v2 v-model="saveForm.data.` + tplEM.FieldVar + `" :options="tm('` + i18nPath + `.status.` + i18nFieldPath + `')" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :multiple="true" :collapse-tags="true" :collapse-tags-tooltip="true" style="width: ` + gconv.String(170+(v.FieldShowLenMax-3)*14) + `px" /> -->`
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			if relIdObj.tpl.Table != `` {
				isReturn = true
				apiUrl := relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
				viewSaveField.formContent.Method = internal.ReturnTypeName
				if relIdObj.tpl.Handle.Pid.Pid != `` {
					viewSaveField.rule.Method = internal.ReturnTypeName
					viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.select') },`)

					viewSaveField.formContent.DataTypeName = `<my-cascader v-model="saveForm.data.` + tplEM.FieldVar + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :isPanel="true" :props="{ multiple: true }" />`

					viewSaveField.paramHandle.Method = internal.ReturnTypeName
					viewSaveField.paramHandle.DataTypeName = `if (param.` + tplEM.FieldVar + ` === undefined) {
                param.` + tplEM.FieldVar + ` = []
            } else {
                let ` + gstr.CaseCamelLower(tplEM.FieldVar) + `: any = []
                param.` + tplEM.FieldVar + `.forEach((item: any) => {
                    ` + gstr.CaseCamelLower(tplEM.FieldVar) + ` = ` + gstr.CaseCamelLower(tplEM.FieldVar) + `.concat(item)
                })
                param.` + tplEM.FieldVar + ` = [...new Set(` + gstr.CaseCamelLower(tplEM.FieldVar) + `)]
            }`
				} else {
					viewSaveField.rule.Method = internal.ReturnTypeName
					viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'integer', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.select') } },	// 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })`)

					viewSaveField.formContent.DataTypeName = `<!-- 建议：大表用<my-select>（滚动分页），小表用<my-transfer>（无分页） -->
					<my-select v-model="saveForm.data.` + tplEM.FieldVar + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" :multiple="true" />
                    <!-- <my-transfer v-model="saveForm.data.` + tplEM.FieldVar + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" /> -->`
				}
			}
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
			if v.FieldType == internal.TypeVarchar {
				isReturn = true

				viewSaveField.rule.Method = internal.ReturnTypeName
				viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },	// 限制数组数量时用：max: 10, message: t('validation.max.upload', { max: 10 })`)
				attrOfAdd := ` :multiple="true"`
				switch v.FieldTypeName {
				case internal.TypeNameImageSuffix:
					attrOfAdd += ` accept="image/*"`
				case internal.TypeNameVideoSuffix:
					attrOfAdd += ` accept="video/*" show-type="video"`
				case internal.TypeNameAudioSuffix:
					attrOfAdd += ` accept="audio/*" show-type="audio"`
				case internal.TypeNameFileSuffix:
					attrOfAdd += ` show-type="text"`
				}
				viewSaveField.formContent.Method = internal.ReturnTypeName
				viewSaveField.formContent.DataTypeName = `<my-upload v-model="saveForm.data.` + tplEM.FieldVar + `"` + attrOfAdd + ` />`
			}
		}
		if isReturn {
			viewSave.Add(viewSaveField, tplEM.FieldVar, i18nPath, i18nFieldPath, tplEM.TableType, ``)
			return
		}
		/*--------部分命名类型直接处理后返回 结束--------*/

		viewSaveFieldTmp := myGenViewSaveField{}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
			viewSaveFieldTmp.rule.Method = internal.ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'integer', trigger: 'change', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.between.number', { min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+` }) },`)
			viewSaveFieldTmp.formContent.Method = internal.ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input-number :min="` + v.FieldLimitInt.Min + `" :max="` + v.FieldLimitInt.Max + `" :precision="0" :controls="false" />`
		case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
			rule := `{ type: 'number', message: t('validation.input') },`
			attrOfAdd := ``
			if v.FieldLimitFloat.Min != `` && v.FieldLimitFloat.Max != `` {
				rule = `{ type: 'number', min: ` + v.FieldLimitFloat.Min + `, max: ` + v.FieldLimitFloat.Max + `, message: t('validation.between.number', { min: ` + v.FieldLimitFloat.Min + `, max: ` + v.FieldLimitFloat.Max + ` }) },`
				attrOfAdd = ` :min="` + v.FieldLimitFloat.Min + `" :max="` + v.FieldLimitFloat.Max + `"`
			} else if v.FieldLimitFloat.Min != `` {
				rule = `{ type: 'number', min: ` + v.FieldLimitFloat.Min + `, message: t('validation.min.number', { min: ` + v.FieldLimitFloat.Min + ` }) },`
				attrOfAdd = ` :min="` + v.FieldLimitFloat.Min + `"`
			} else if v.FieldLimitFloat.Max != `` {
				rule = `{ type: 'number', max: ` + v.FieldLimitFloat.Max + `, message: t('validation.max.number', { max: ` + v.FieldLimitFloat.Max + ` }) },`
				attrOfAdd = ` :max="` + v.FieldLimitFloat.Max + `"`
			}
			rule += `    // type: 'float'在值为0时验证不能通过`
			viewSaveFieldTmp.rule.Method = internal.ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, rule)
			viewSaveFieldTmp.formContent.Method = internal.ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input-number` + attrOfAdd + ` :precision="` + gconv.String(v.FieldLimitFloat.Precision) + `" :controls="false" />`
		case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
			rule := `{ type: 'string', max: ` + v.FieldLimitStr + `, message: t('validation.max.string', { max: ` + v.FieldLimitStr + ` }) },`
			attrOfAdd := ``
			if v.FieldType == internal.TypeChar {
				rule = `{ type: 'string', len: ` + v.FieldLimitStr + `, message: t('validation.size.string', { size: ` + v.FieldLimitStr + ` }) },`
				attrOfAdd = ` minlength="` + v.FieldLimitStr + `"`
			}
			viewSaveFieldTmp.rule.Method = internal.ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, rule)
			viewSaveFieldTmp.formContent.Method = internal.ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input` + attrOfAdd + ` maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
		case internal.TypeText, internal.TypeJson: // `text类型` // `json类型`
			viewSaveFieldTmp.rule.Method = internal.ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'string', message: t('validation.input') },`)
			viewSaveFieldTmp.formContent.Method = internal.ReturnType
			viewSaveFieldTmp.formContent.DataTypeName = `<el-input type="textarea" :autosize="{ minRows: 3 }" />`
		default:
			viewSaveFieldTmp.rule.Method = internal.ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'string', message: t('validation.input') },`)
			viewSaveFieldTmp.formContent.Method = internal.ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input :clearable="true" />`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case internal.TypeNameDeleted, internal.TypeNameUpdated, internal.TypeNameCreated: // 软删除字段 // 更新时间字段 // 创建时间字段
		case internal.TypeNamePid: // pid；	类型：int等类型；
		case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = internal.ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'string', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },`)
		case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = internal.ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'string', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },`)
		case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = internal.ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'string', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },`)
		case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = internal.ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'email', message: t('validation.email') },`)
		case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = internal.ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'url', message: t('validation.url') },`)
		case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
			viewSaveFieldTmp.rule.Method = internal.ReturnTypeName
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'integer', min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+`, message: t('validation.between.number', { min: `+v.FieldLimitInt.Min+`, max: `+v.FieldLimitInt.Max+` }) },`)
			/* viewSaveFieldTmp.formContent.Method = internal.ReturnTypeName
			viewSaveFieldTmp.formContent.DataTypeName = `<el-input-number :min="` + v.FieldLimitInt.Min + `" :max="` + v.FieldLimitInt.Max + `" :precision="0" />` */
		case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewSaveFieldTmp.isI18nTm = true
			viewSaveFieldTmp.rule.Method = internal.ReturnTypeName
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'enum', enum: (tm('`+i18nPath+`.status.`+i18nFieldPath+`') as any).map((item: any) => item.value), message: t('validation.select') },`)
		case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
			viewSaveFieldTmp.isI18nTm = true
			viewSaveFieldTmp.rule.Method = internal.ReturnTypeName
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') },`)
		case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
		case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			if v.FieldType == internal.TypeVarchar {
				viewSaveFieldTmp.formContent.Method = internal.ReturnTypeName
				viewSaveFieldTmp.formContent.DataTypeName = `<el-input type="textarea" :autosize="{ minRows: 3 }" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" />`
			}
		case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
		case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewSaveField.rule.Method = internal.ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'blur', message: t('validation.array'), defaultField: [`+gstr.Join(append([]string{``}, viewSaveFieldTmp.rule.GetData()...), `
                `)+`] },	// 限制数组数量时用：max: 10, message: t('validation.max.array', { max: 10 })`)

		fieldHandle := gstr.CaseCamelLower(tplEM.FieldVar) + `Handle`
		formContent := gstr.TrimStr(viewSaveFieldTmp.formContent.GetData(), ` `)
		formContent = gstr.Replace(formContent, ` `, ` :ref="(el: any) => `+fieldHandle+`.ref[index] = el" v-model="saveForm.data.`+tplEM.FieldVar+`[index]" @blur="`+fieldHandle+`.del(index, true)" :placeholder="t('`+i18nPath+`.name.`+i18nFieldPath+`')"`, 1)
		switch gstr.Split(formContent, ` `)[0] {
		case `<el-input`, `<el-input-number`:
			formContent = gstr.SubStr(formContent, 0, -2) + `style="width: 150px;" />`
		}

		viewSaveField.isI18nTm = true
		viewSaveField.formContent.Method = internal.ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<template v-for="(_, index) in saveForm.data.` + tplEM.FieldVar + `" :key="index">
                    <el-tag type="info" :closable="true" @close="` + fieldHandle + `.del(index)" size="large" style="padding-left: 0; margin: 3px 10px 3px 0;">
                        ` + formContent + `
                    </el-tag>
                </template>
                <el-button type="primary" @click="` + fieldHandle + `.add" style="margin: 3px 0"> <autoicon-ep-plus />{{ t('common.add') }} </el-button>`
		viewSaveField.formHandle.Method = internal.ReturnTypeName
		viewSaveField.formHandle.DataTypeName = `const ` + fieldHandle + ` = reactive({
    ref: [] as any[],
    add: () => {
        !Array.isArray(saveForm.data.` + tplEM.FieldVar + `) && (saveForm.data.` + tplEM.FieldVar + ` = [])
        saveForm.data.` + tplEM.FieldVar + `.push(undefined)
        nextTick(() => ` + fieldHandle + `.ref[` + fieldHandle + `.ref.length - 1].focus())
    },
    del: (index: number, isBlur: boolean = false) => {
        if (isBlur && saveForm.data.` + tplEM.FieldVar + `[index] !== undefined && saveForm.data.` + tplEM.FieldVar + `[index] !== null && saveForm.data.` + tplEM.FieldVar + `[index] !== '') {
            return
        }
        saveForm.data.` + tplEM.FieldVar + `.splice(index, 1)
        ` + fieldHandle + `.ref.splice(index, 1)
    },
})`
		viewSave.Add(viewSaveField, tplEM.FieldVar, i18nPath, i18nFieldPath, tplEM.TableType, ``)
	} else {
		viewSaveFieldAddMessagePrefix := func(viewSaveField myGenViewSaveField, messagePrefix string) myGenViewSaveField {
			for k, v := range viewSaveField.rule.DataType {
				viewSaveField.rule.DataType[k] = gstr.Replace(v, `message: `, `message: `+messagePrefix+` + `, 1)
			}
			for k, v := range viewSaveField.rule.DataTypeName {
				viewSaveField.rule.DataTypeName[k] = gstr.Replace(v, `message: `, `message: `+messagePrefix+` + `, 1)
			}
			return viewSaveField
		}
		viewSaveTmp := myGenViewSave{}
		switch tplEM.TableType {
		case internal.TableTypeExtendMany:
			for _, v := range tplEM.FieldList {
				viewSaveField := getViewSaveField(tplEM.tpl, v, tplEM.FieldVar+`[index].`+v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw)
				viewSaveField = viewSaveFieldAddMessagePrefix(viewSaveField, `t('`+tplEM.tplOfTop.I18nPath+`.name.`+tplEM.FieldVar+`.`+v.FieldRaw+`')`)
				viewSaveTmp.Add(viewSaveField, v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar, tplEM.TableType, ``)
			}
		case internal.TableTypeMiddleMany:
			for _, v := range tplEM.FieldListOfIdSuffix {
				viewSaveField := getViewSaveField(tplEM.tpl, v, tplEM.FieldVar+`[index].`+v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw)
				viewSaveField = viewSaveFieldAddMessagePrefix(viewSaveField, `t('`+tplEM.tplOfTop.I18nPath+`.name.`+tplEM.FieldVar+`.`+v.FieldRaw+`')`)
				viewSaveTmp.Add(viewSaveField, v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar, tplEM.TableType, ``)
			}
			if len(tplEM.FieldListOfOther) > 0 {
				fieldIfArr := []string{}
				for _, v := range tplEM.FieldListOfIdSuffix {
					fieldIfArr = append(fieldIfArr, `saveForm.data.`+tplEM.FieldVar+`[index].`+v.FieldRaw)
				}
				fieldIf := gstr.Join(fieldIfArr, ` || `)
				for _, v := range tplEM.FieldListOfOther {
					viewSaveField := getViewSaveField(tplEM.tpl, v, tplEM.FieldVar+`[index].`+v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw)
					viewSaveField = viewSaveFieldAddMessagePrefix(viewSaveField, `t('`+tplEM.tplOfTop.I18nPath+`.name.`+tplEM.FieldVar+`.`+v.FieldRaw+`')`)
					viewSaveTmp.Add(viewSaveField, v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar, tplEM.TableType, fieldIf)
				}
			}
		}

		viewSave.importModule = append(viewSave.importModule, viewSaveTmp.importModule...)
		viewSave.dataInitAfter = append(viewSave.dataInitAfter, tplEM.FieldVar+`: saveCommon.data.`+tplEM.FieldVar+` ? saveCommon.data.`+tplEM.FieldVar+` : [],`)
		viewSave.rule = append(viewSave.rule, tplEM.FieldVar+`: [
            {
                type: 'array',
                trigger: 'change',
                message: t('validation.array'),
                defaultField: [
                    {
                        type: 'object',
                        fields: {`+gstr.Join(append([]string{``}, viewSaveTmp.rule...), `
                            `)+`
                        },
                    },
                ],
            },
        ],`)
		viewSave.formItem = append(viewSave.formItem, `<el-form-item :label="t('`+tplEM.tplOfTop.I18nPath+`.name.`+internal.GetStrByFieldStyle(tplEM.tplOfTop.FieldStyle, tplEM.FieldVar, ``, `label`)+`')" prop="`+tplEM.FieldVar+`" style="min-height: 60px">
                    <template #label>
                        <span style="text-align: right">
                            <div>{{ t('`+tplEM.tplOfTop.I18nPath+`.name.`+internal.GetStrByFieldStyle(tplEM.tplOfTop.FieldStyle, tplEM.FieldVar, ``, `label`)+`') }}</div>
                            <el-button type="primary" size="small" @click="() => saveForm.data.`+tplEM.FieldVar+`.push({})"><autoicon-ep-plus />{{ t('common.add') }}</el-button>
                        </span>
                    </template>

                    <template v-for="(item, index) in saveForm.data.`+tplEM.FieldVar+`" :key="index">
                        <div style="width: 100%; margin: 3px 0; display: flex; align-items: center; gap: 10px">
                            <el-button type="danger" size="small" @click="() => saveForm.data.`+tplEM.FieldVar+`.splice(index, 1)"><autoicon-ep-close />{{ t('common.delete') }}</el-button>
                            {{formContent}}
                        </div>
                    </template>
                </el-form-item>`)
		viewSave.formContent = append(viewSave.formContent, gstr.Join(viewSaveTmp.formContent, `
                            `))
	}
	return
}
