package my_gen

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewSave struct {
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
	importModule   []string
	dataInitBefore myGenDataStrHandler
	dataInitAfter  myGenDataStrHandler
	isRequired     bool
	rule           myGenDataSliceHandler
	formContent    myGenDataStrHandler
	formHandle     myGenDataStrHandler
	paramHandle    myGenDataStrHandler
}

func (viewSaveThis *myGenViewSave) Add(viewSaveField myGenViewSaveField, field string, i18nPath string, i18nFieldPath string, tableType myGenTableType, fieldIf string) {
	if viewSaveField.formContent.getData() == `` {
		return
	}
	viewSaveThis.importModule = append(viewSaveThis.importModule, viewSaveField.importModule...)
	if viewSaveField.dataInitBefore.getData() != `` {
		viewSaveThis.dataInitBefore = append(viewSaveThis.dataInitBefore, field+`: `+viewSaveField.dataInitBefore.getData()+`,`)
	}
	if viewSaveField.dataInitAfter.getData() != `` {
		viewSaveThis.dataInitAfter = append(viewSaveThis.dataInitAfter, field+`: `+viewSaveField.dataInitAfter.getData()+`,`)
	}
	rule := viewSaveField.rule.getData()
	if viewSaveField.isRequired {
		if fieldIf != `` && tableType == TableTypeMiddleOne {
			rule = append([]string{`{ required: computed((): boolean => ` + fieldIf + `), message: t('validation.required') },`}, rule...)
		} else {
			switch tableType {
			case TableTypeExtendMany, TableTypeMiddleMany:
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
		viewSaveThis.formContent = append(viewSaveThis.formContent, viewSaveField.formContent.getData())
	} else {
		viewSaveThis.formItem = append(viewSaveThis.formItem, `<el-form-item v-if="`+fieldIf+`" :label="t('`+i18nPath+`.name.`+i18nFieldPath+`')" prop="`+field+`">
					{{formContent}}
				</el-form-item>`)
		if tableType == TableTypeMiddleMany {
			formContent := gstr.TrimStr(viewSaveField.formContent.getData(), ` `)
			formContent = gstr.Replace(formContent, ` `, ` v-if="`+fieldIf+`"`, 1)
			/* switch gstr.Split(formContent, ` `)[0] {
			case `<el-input`:
				formContent = gstr.SubStr(formContent, 0, -2) + `style="width: 200px;" />`
			case `<el-input-number`:
				formContent = gstr.SubStr(formContent, 0, -2) + `style="width: 150px;" />`
			} */
			viewSaveThis.formContent = append(viewSaveThis.formContent, formContent)
		} else {
			viewSaveThis.formContent = append(viewSaveThis.formContent, viewSaveField.formContent.getData())
		}
	}
	if viewSaveField.formHandle.getData() != `` {
		viewSaveThis.formHandle = append(viewSaveThis.formHandle, viewSaveField.formHandle.getData())
	}
	if viewSaveField.paramHandle.getData() != `` {
		viewSaveThis.paramHandle = append(viewSaveThis.paramHandle, viewSaveField.paramHandle.getData())
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

// 视图模板Query生成
func genViewSave(option myGenOption, tpl myGenTpl) {
	if !(option.IsCreate || option.IsUpdate) {
		return
	}
	viewSave := myGenViewSave{}
	for _, v := range tpl.FieldListOfDefault {
		viewSave.Add(getViewSaveField(tpl, v, v.FieldRaw, tpl.I18nPath, v.FieldRaw), v.FieldRaw, tpl.I18nPath, v.FieldRaw, TableTypeDefault, ``)
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
	for _, v := range tpl.FieldListOfAfter {
		viewSave.Add(getViewSaveField(tpl, v, v.FieldRaw, tpl.I18nPath, v.FieldRaw), v.FieldRaw, tpl.I18nPath, v.FieldRaw, TableTypeDefault, ``)
	}
	viewSave.Unique()

	tplView := `<script setup lang="tsx">` + gstr.Join(append([]string{``}, viewSave.importModule...), `
`) + `
const { t, tm } = useI18n()

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
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)` + gstr.Join(append([]string{``}, viewSave.paramHandle...), `
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
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
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
	case TypeInt: // `int等类型`
		defaultVal := gconv.Int(v.Default)
		if defaultVal != 0 {
			viewSaveField.dataInitBefore.Method = ReturnType
			viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
		}
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'integer', trigger: 'change', message: t('validation.input') },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
	case TypeIntU: // `int等类型（unsigned）`
		defaultVal := gconv.Uint(v.Default)
		if defaultVal != 0 {
			viewSaveField.dataInitBefore.Method = ReturnType
			viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
		}
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'integer', trigger: 'change', min: 0, message: t('validation.min.number', { min: 0 }) },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :min="0" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
	case TypeFloat: // `float等类型`
		defaultVal := gconv.Float64(v.Default)
		if defaultVal != 0 {
			viewSaveField.dataInitBefore.Method = ReturnType
			viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
		}
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'number', trigger: 'change', message: t('validation.input') },    // type: 'float'在值为0时验证不能通过`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
	case TypeFloatU: // `float等类型（unsigned）`
		defaultVal := gconv.Float64(v.Default)
		if defaultVal != 0 {
			viewSaveField.dataInitBefore.Method = ReturnType
			viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
		}
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'number', trigger: 'change', min: 0, message: t('validation.min.number', { min: 0 }) },    // type: 'float'在值为0时验证不能通过`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
	case TypeVarchar: // `varchar类型`
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'blur', max: `+v.FieldLimitStr+`, message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
		if v.IsUnique {
			viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />`
		}
	case TypeChar: // `char类型`
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'blur', len: `+v.FieldLimitStr+`, message: t('validation.size.string', { size: `+v.FieldLimitStr+` }) },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
		if v.IsUnique {
			viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />`
		}
	case TypeText: // `text类型`
		if !v.IsNull {
			viewSaveField.isRequired = true
		}
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'blur', message: t('validation.input') },`)

		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<my-editor v-model="saveForm.data.` + dataFieldPath + `" />`
	case TypeJson: // `json类型`
		if !v.IsNull {
			viewSaveField.isRequired = true
		}
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{
                type: 'object',
                trigger: 'blur',
                message: t('validation.json'),
                /* fields: {
                    xxxx: [
						{ type: 'string', message: 'xxxx' + t('validation.input') },
						// { type: 'integer', min: 1, message: 'xxxx' + t('validation.min.number', { min: 1 }) },
					],
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
            },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" type="textarea" :autosize="{ minRows: 3 }" />`
		if v.FieldTip != `` {
			viewSaveField.formContent.DataType = `<el-alert :title="t('` + i18nPath + `.tip.` + i18nFieldPath + `')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    ` + viewSaveField.formContent.DataType
		}
	case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-date-picker v-model="saveForm.data.` + dataFieldPath + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />`
	case TypeDate: // `date类型`
		viewSaveField.rule.Method = ReturnType
		viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-date-picker v-model="saveForm.data.` + dataFieldPath + `" type="date" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />`
	default:
		viewSaveField.formContent.Method = ReturnType
		viewSaveField.formContent.DataType = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="true" />`
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case TypePrimary: // 独立主键
	case TypePrimaryAutoInc: // 独立主键（自增）
		return myGenViewSaveField{}
	case TypePrimaryMany: // 联合主键
	case TypePrimaryManyAutoInc: // 联合主键（自增）
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case TypeNameDeleted: // 软删除字段
		return myGenViewSaveField{}
	case TypeNameUpdated: // 更新时间字段
		return myGenViewSaveField{}
	case TypeNameCreated: // 创建时间字段
		return myGenViewSaveField{}
	case TypeNamePid: // pid；	类型：int等类型；
		viewSaveField.dataInitAfter.Method = ReturnTypeName
		viewSaveField.dataInitAfter.DataTypeName = `saveCommon.data.` + dataFieldPath + ` ? saveCommon.data.` + dataFieldPath + ` : undefined`
		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', trigger: 'change', min: 1, message: t('validation.select') },`)
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<my-cascader v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />`
		viewSaveField.paramHandle.Method = ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` === undefined ? param.` + dataFieldPath + ` = 0 : null`
	case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		return myGenViewSaveField{}
	case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		return myGenViewSaveField{}
	case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		viewSaveField.importModule = append(viewSaveField.importModule, `import md5 from 'js-md5'`)
		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName,
			`{ required: computed((): boolean => (saveForm.data.idArr?.length ? false : true)), message: t('validation.required') },`,
			`{ type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },`,
		)
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-input v-model="saveForm.data.` + dataFieldPath + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert v-if="saveForm.data.idArr?.length" :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />`
		viewSaveField.paramHandle.Method = ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` ? param.` + dataFieldPath + ` = md5(param.` + dataFieldPath + `) : delete param.` + dataFieldPath
	case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenViewSaveField{}
	case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		if gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
			viewSaveField.isRequired = true
		}
	case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		viewSaveField.rule.Method = ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },`)
	case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		viewSaveField.rule.Method = ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', trigger: 'blur', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },`)
	case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		viewSaveField.rule.Method = ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', trigger: 'blur', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },`)
	case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		viewSaveField.rule.Method = ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'email', trigger: 'blur', message: t('validation.email') },`)
	case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		viewSaveField.rule.Method = ReturnUnion
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'url', trigger: 'blur', message: t('validation.url') },`)
	case TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
		if relIdObj.tpl.Table != `` {
			apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
		}
		viewSaveField.dataInitAfter.Method = ReturnTypeName
		viewSaveField.dataInitAfter.DataTypeName = `saveCommon.data.` + dataFieldPath + ` ? saveCommon.data.` + dataFieldPath + ` : undefined`
		viewSaveField.rule.Method = ReturnTypeName
		if !viewSaveField.isRequired {
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `// { required: true, message: t('validation.required') },`)
		}
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', trigger: 'change', min: 1, message: t('validation.select') },`)
		viewSaveField.formContent.Method = ReturnTypeName
		if relIdObj.tpl.Handle.Pid.Pid != `` {
			viewSaveField.formContent.DataTypeName = `<my-cascader v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />`
		} else {
			viewSaveField.formContent.DataTypeName = `<my-select v-model="saveForm.data.` + dataFieldPath + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />`
		}
		viewSaveField.paramHandle.Method = ReturnTypeName
		viewSaveField.paramHandle.DataTypeName = `param.` + dataFieldPath + ` === undefined ? param.` + dataFieldPath + ` = 0 : null`
	case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', trigger: 'change', min: 0, max: 100, message: t('validation.between.number', { min: 0, max: 100 }) },`)
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-input-number v-model="saveForm.data.` + dataFieldPath + `" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="` + gconv.String(gconv.Int(v.Default)) + `" />
                    <el-alert :title="t('` + i18nPath + `.tip.` + i18nFieldPath + `')" type="info" :show-icon="true" :closable="false" />`
	case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		defaultVal := gconv.String(v.Default)
		if defaultVal == `` {
			defaultVal = v.StatusList[0][0]
		}
		viewSaveField.dataInitBefore.Method = ReturnTypeName
		viewSaveField.dataInitBefore.DataTypeName = defaultVal
		if garray.NewIntArrayFrom([]int{TypeVarchar, TypeChar}).Contains(v.FieldType) {
			viewSaveField.dataInitBefore.DataTypeName = `'` + defaultVal + `'`
		}
		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'enum', trigger: 'change', enum: (tm('`+i18nPath+`.status.`+i18nFieldPath+`') as any).map((item: any) => item.value), message: t('validation.select') },`)
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-radio-group v-model="saveForm.data.` + dataFieldPath + `">
                        <el-radio v-for="(item, index) in (tm('` + i18nPath + `.status.` + i18nFieldPath + `') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>`
		if len(v.StatusList) > 5 { //超过5个状态用select组件，小于5个用radio组件
			viewSaveField.formContent.DataTypeName = `<el-select-v2 v-model="saveForm.data.` + dataFieldPath + `" :options="tm('` + i18nPath + `.status.` + i18nFieldPath + `')" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="false" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		}
	case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') },`)
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-switch v-model="saveForm.data.` + dataFieldPath + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />`
	case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
	case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		if v.FieldType != TypeDate {
			viewSaveField.formContent.Method = ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-date-picker v-model="saveForm.data.` + dataFieldPath + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :default-time="new Date(2000, 0, 1, 23, 59, 59)" />`
		}
	case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		if v.FieldType == TypeVarchar {
			viewSaveField.formContent.Method = ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-input v-model="saveForm.data.` + dataFieldPath + `" type="textarea" :autosize="{ minRows: 3 }" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" />`
		}
	case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		if v.FieldType == TypeVarchar {
			viewSaveField.rule.Method = ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'url', trigger: 'change', message: t('validation.upload') },`)
		} else {
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },	// 限制数组数量时用：max: 10, message: t('validation.max.upload', { max: 10 })`)
		}
		attrOfAdd := ``
		if v.FieldType != TypeVarchar {
			attrOfAdd += ` :multiple="true"`
		}
		if v.FieldTypeName == TypeNameVideoSuffix {
			attrOfAdd += ` accept="video/*" :isImage="false"`
		} else {
			attrOfAdd += ` accept="image/*"`
		}
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<my-upload v-model="saveForm.data.` + dataFieldPath + `"` + attrOfAdd + ` />`
	case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'blur', message: t('validation.array'), defaultField: { type: 'string', message: t('validation.input') } },	// 限制数组数量时用：max: 10, message: t('validation.max.array', { max: 10 })`)
		fieldHandle := gstr.CaseCamelLower(gstr.Replace(dataFieldPath, `.`, `_`)) + `Handle`
		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-tag v-for="(item, index) in saveForm.data.` + dataFieldPath + `" :type="` + fieldHandle + `.tagType[index % ` + fieldHandle + `.tagType.length] as any" @close="` + fieldHandle + `.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="` + fieldHandle + `.visible" :ref="(el: any) => ` + fieldHandle + `.ref = el" v-model="` + fieldHandle + `.value" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" @keyup.enter="` + fieldHandle + `.addValue" @blur="` + fieldHandle + `.addValue" :controls="false" size="small" style="width: 150px;" /> -->
                    <el-input v-if="` + fieldHandle + `.visible" :ref="(el: any) => ` + fieldHandle + `.ref = el" v-model="` + fieldHandle + `.value" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" @keyup.enter="` + fieldHandle + `.addValue" @blur="` + fieldHandle + `.addValue" :show-word-limit="true" :clearable="true" size="small" style="width: 200px;" />
                    <el-button v-else type="primary" size="small" @click="` + fieldHandle + `.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>`
		viewSaveField.formHandle.Method = ReturnTypeName
		viewSaveField.formHandle.DataTypeName = `const ` + fieldHandle + ` = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        ` + fieldHandle + `.visible = true
        nextTick(() => {
            ` + fieldHandle + `.ref?.focus()
        })
    },
    addValue: () => {
        if (!(` + fieldHandle + `.value === undefined || ` + fieldHandle + `.value === '')) {
			Array.isArray(saveForm.data.` + dataFieldPath + `) ? null : (saveForm.data.` + dataFieldPath + ` = [])
            saveForm.data.` + dataFieldPath + `.push(` + fieldHandle + `.value)
        }
        ` + fieldHandle + `.visible = false
        ` + fieldHandle + `.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.` + dataFieldPath + `.splice(saveForm.data.` + dataFieldPath + `.indexOf(item), 1)
    },
})`
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getViewSaveExtendMiddleOne(tplEM handleExtendMiddle) (viewSave myGenViewSave) {
	switch tplEM.TableType {
	case TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			viewSave.Add(getViewSaveField(tplEM.tpl, v, v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw), v.FieldRaw, tplEM.tplOfTop.I18nPath, v.FieldRaw, tplEM.TableType, ``)
		}
	case TableTypeMiddleOne:
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
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			isReturn = true

			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'enum', enum: (tm('`+i18nPath+`.status.`+i18nFieldPath+`') as any).map((item: any) => item.value), message: t('validation.select') } },	// 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })`)

			viewSaveField.formContent.Method = ReturnTypeName
			viewSaveField.formContent.DataTypeName = `<el-select-v2 v-model="saveForm.data.` + tplEM.FieldVar + `" :options="tm('` + i18nPath + `.status.` + i18nFieldPath + `')" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :multiple="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			isReturn = true

			relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if relIdObj.tpl.Table != `` {
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			viewSaveField.formContent.Method = ReturnTypeName
			if relIdObj.tpl.Handle.Pid.Pid != `` {
				viewSaveField.rule.Method = ReturnTypeName
				viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.select')/* , defaultField: { type: 'array', defaultField: { type: 'integer', min: 1, message: t('validation.min.number', { min: 1 }) } } */ },`)

				viewSaveField.formContent.DataTypeName = `<my-cascader v-model="saveForm.data.` + tplEM.FieldVar + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :isPanel="true" :props="{ multiple: true }" />`

				viewSaveField.paramHandle.Method = ReturnTypeName
				viewSaveField.paramHandle.DataTypeName = `if (param.` + tplEM.FieldVar + ` === undefined) {
                param.` + tplEM.FieldVar + ` = []
            } else {
                let ` + tplEM.FieldVar + `: any = []
                param.` + tplEM.FieldVar + `.forEach((item: any) => {
                    ` + tplEM.FieldVar + ` = ` + tplEM.FieldVar + `.concat(item)
                })
                //param.` + tplEM.FieldVar + ` = [...new Set(` + tplEM.FieldVar + `)]
                param.` + tplEM.FieldVar + ` = ` + tplEM.FieldVar + `.filter((item: any, index: any) => {
                    return ` + tplEM.FieldVar + `.indexOf(item) === index
                })
            }`
			} else {
				viewSaveField.rule.Method = ReturnTypeName
				viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'integer', min: 1, message: t('validation.min.number', { min: 1 }) } },	// 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })`)

				viewSaveField.formContent.DataTypeName = `<!-- 建议：大表用<my-select>（滚动分页），小表用<my-transfer>（无分页） -->
					<my-select v-model="saveForm.data.` + tplEM.FieldVar + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" :multiple="true" />
                    <!-- <my-transfer v-model="saveForm.data.` + tplEM.FieldVar + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" /> -->`
			}
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if v.FieldType == TypeVarchar {
				isReturn = true

				viewSaveField.rule.Method = ReturnTypeName
				viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },	// 限制数组数量时用：max: 10, message: t('validation.max.upload', { max: 10 })`)
				attrOfAdd := ` :multiple="true"`
				if v.FieldTypeName == TypeNameVideoSuffix {
					attrOfAdd += ` accept="video/*" :isImage="false"`
				} else {
					attrOfAdd += ` accept="image/*"`
				}
				viewSaveField.formContent.Method = ReturnTypeName
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
		case TypeInt: // `int等类型`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'integer', message: t('validation.input') },`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input-number :controls="false" />`
		case TypeIntU: // `int等类型（unsigned）`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'integer', min: 0, message: t('validation.min.number', { min: 0 }) },`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input-number :min="0" :controls="false" />`
		case TypeFloat: // `float等类型`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'number', message: t('validation.input') },    // type: 'float'在值为0时验证不能通过`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input-number :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />`
		case TypeFloatU: // `float等类型（unsigned）`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'number', min: 0, message: t('validation.min.number', { min: 0 }) },    // type: 'float'在值为0时验证不能通过`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input-number :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />`
		case TypeVarchar: // `varchar类型`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'string', max: `+v.FieldLimitStr+`, message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
		case TypeChar: // `char类型`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'string', len: `+v.FieldLimitStr+`, message: t('validation.size.string', { size: `+v.FieldLimitStr+` }) },`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
		case TypeText, TypeJson: // `text类型` // `json类型`
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'string', message: t('validation.input') },`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataTypeName = `<el-input type="textarea" :autosize="{ minRows: 3 }" />`
		default:
			viewSaveFieldTmp.rule.Method = ReturnType
			viewSaveFieldTmp.rule.DataType = append(viewSaveFieldTmp.rule.DataType, `{ type: 'string', message: t('validation.input') },`)
			viewSaveFieldTmp.formContent.Method = ReturnType
			viewSaveFieldTmp.formContent.DataType = `<el-input :clearable="true" />`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted, TypeNameUpdated, TypeNameCreated: // 软删除字段 // 更新时间字段 // 创建时间字段
		case TypeNamePid: // pid；	类型：int等类型；
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'string', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') },`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'string', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'string', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'email', message: t('validation.email') },`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			viewSaveFieldTmp.rule.Method = ReturnUnion
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'url', message: t('validation.url') },`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewSaveFieldTmp.rule.Method = ReturnTypeName
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'integer', min: 0, max: 100, message: t('validation.between.number', { min: 0, max: 100 }) },`)
			viewSaveFieldTmp.formContent.Method = ReturnTypeName
			viewSaveFieldTmp.formContent.DataTypeName = `<el-input-number :precision="0" :min="0" :max="100" :step="1" /* :step-strictly="true" controls-position="right" */ />`
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewSaveFieldTmp.rule.Method = ReturnTypeName
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'enum', enum: (tm('`+i18nPath+`.status.`+i18nFieldPath+`') as any).map((item: any) => item.value), message: t('validation.select') },`)
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			viewSaveFieldTmp.rule.Method = ReturnTypeName
			viewSaveFieldTmp.rule.DataTypeName = append(viewSaveFieldTmp.rule.DataTypeName, `{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') },`)
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			if v.FieldType == TypeVarchar {
				viewSaveFieldTmp.formContent.Method = ReturnTypeName
				viewSaveFieldTmp.formContent.DataTypeName = `<el-input type="textarea" :autosize="{ minRows: 3 }" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" />`
			}
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewSaveField.rule.Method = ReturnTypeName
		viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'array', trigger: 'blur', message: t('validation.array'), defaultField: [`+gstr.Join(append([]string{``}, viewSaveFieldTmp.rule.getData()...), `
                `)+`] },	// 限制数组数量时用：max: 10, message: t('validation.max.array', { max: 10 })`)

		fieldHandle := gstr.CaseCamelLower(tplEM.FieldVar) + `Handle`
		formContent := gstr.TrimStr(viewSaveFieldTmp.formContent.getData(), ` `)
		formContent = gstr.Replace(formContent, ` `, ` v-if="`+fieldHandle+`.visible" :ref="(el: any) => `+fieldHandle+`.ref = el" v-model="`+fieldHandle+`.value" :placeholder="t('`+i18nPath+`.name.`+i18nFieldPath+`')" @keyup.enter="`+fieldHandle+`.addValue" @blur="`+fieldHandle+`.addValue"`, 1)
		switch gstr.Split(formContent, ` `)[0] {
		case `<el-input`:
			formContent = gstr.SubStr(formContent, 0, -2) + `size="small" style="width: 200px;" />`
		case `<el-input-number`:
			formContent = gstr.SubStr(formContent, 0, -2) + `size="small" style="width: 150px;" />`
		}

		viewSaveField.formContent.Method = ReturnTypeName
		viewSaveField.formContent.DataTypeName = `<el-tag v-for="(item, index) in saveForm.data.` + tplEM.FieldVar + `" :type="` + fieldHandle + `.tagType[index % ` + fieldHandle + `.tagType.length] as any" @close="` + fieldHandle + `.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
					` + formContent + `
                    <el-button v-else type="primary" size="small" @click="` + fieldHandle + `.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>`
		viewSaveField.formHandle.Method = ReturnTypeName
		viewSaveField.formHandle.DataTypeName = `const ` + fieldHandle + ` = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        ` + fieldHandle + `.visible = true
        nextTick(() => {
            ` + fieldHandle + `.ref?.focus()
        })
    },
    addValue: () => {
        if (!(` + fieldHandle + `.value === undefined || ` + fieldHandle + `.value === '')) {
			Array.isArray(saveForm.data.` + tplEM.FieldVar + `) ? null : (saveForm.data.` + tplEM.FieldVar + ` = [])
            saveForm.data.` + tplEM.FieldVar + `.push(` + fieldHandle + `.value)
        }
        ` + fieldHandle + `.visible = false
        ` + fieldHandle + `.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.` + tplEM.FieldVar + `.splice(saveForm.data.` + tplEM.FieldVar + `.indexOf(item), 1)
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
		case TableTypeExtendMany:
			for _, v := range tplEM.FieldList {
				viewSaveField := getViewSaveField(tplEM.tpl, v, tplEM.FieldVar+`[index].`+v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw)
				viewSaveField = viewSaveFieldAddMessagePrefix(viewSaveField, `t('`+tplEM.tplOfTop.I18nPath+`.name.`+tplEM.FieldVar+`.`+v.FieldRaw+`')`)
				viewSaveTmp.Add(viewSaveField, v.FieldRaw, tplEM.tplOfTop.I18nPath, tplEM.FieldVar, tplEM.TableType, ``)
			}
		case TableTypeMiddleMany:
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
		viewSave.formItem = append(viewSave.formItem, `<el-form-item :label="t('`+tplEM.tplOfTop.I18nPath+`.name.`+tplEM.FieldVar+`Label')" prop="`+tplEM.FieldVar+`" style="min-height: 60px">
                    <template #label>
                        <span style="text-align: right">
                            <div>{{ t('`+tplEM.tplOfTop.I18nPath+`.name.`+tplEM.FieldVar+`Label') }}</div>
                            <el-button type="primary" size="small" @click="() => saveForm.data.`+tplEM.FieldVar+`.push({})"> <autoicon-ep-plus />{{ t('common.add') }} </el-button>
                        </span>
                    </template>

                    <template v-for="(item, index) in saveForm.data.`+tplEM.FieldVar+`" :key="index">
                        <div style="width: 100%; margin: 2.5px 0; display: flex; align-items: center; gap: 10px">
                            <el-button type="danger" size="small" @click="() => saveForm.data.`+tplEM.FieldVar+`.splice(index, 1)"> <autoicon-ep-close />{{ t('common.delete') }} </el-button>
							{{formContent}}
                        </div>
                    </template>
                </el-form-item>`)
		viewSave.formContent = append(viewSave.formContent, gstr.Join(viewSaveTmp.formContent, `
                            `))
	}
	return
}
