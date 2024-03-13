package my_gen

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewSave struct {
	importModule   []string
	dataInitBefore []string
	dataInitAfter  []string
	rule           []string
	form           []string
	formHandle     []string
	paramHandle    []string
}

type myGenViewSaveField struct {
	importModule   []string
	dataInitBefore myGenDataStrHandler
	dataInitAfter  myGenDataStrHandler
	rule           myGenDataSliceHandler
	isRequired     bool
	form           myGenDataStrHandler
	formHandle     myGenDataStrHandler
	paramHandle    myGenDataStrHandler
}

// 视图模板Query生成
func genViewSave(option myGenOption, tpl myGenTpl) {
	if !(option.IsCreate || option.IsUpdate) {
		return
	}

	viewSave := getViewSaveFieldList(tpl)

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
    } as any,
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
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">` + gstr.Join(append([]string{``}, viewSave.form...), `
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

func getViewSaveFieldList(tpl myGenTpl) (viewSave myGenViewSave) {

	for _, v := range tpl.FieldList {
		viewSaveField := myGenViewSaveField{}

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
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input-number v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
		case TypeIntU: // `int等类型（unsigned）`
			defaultVal := gconv.Uint(v.Default)
			if defaultVal != 0 {
				viewSaveField.dataInitBefore.Method = ReturnType
				viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'integer', min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input-number v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
		case TypeFloat: // `float等类型`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveField.dataInitBefore.Method = ReturnType
				viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'number'/* 'float' */, trigger: 'change', message: t('validation.input') },    // 类型float值为0时验证不能通过`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input-number v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
		case TypeFloatU: // `float等类型（unsigned）`
			defaultVal := gconv.Float64(v.Default)
			if defaultVal != 0 {
				viewSaveField.dataInitBefore.Method = ReturnType
				viewSaveField.dataInitBefore.DataType = gconv.String(defaultVal)
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'number'/* 'float' */, min: 0, trigger: 'change', message: t('validation.min.number', { min: 0 }) },    // 类型float值为0时验证不能通过`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input-number v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" :value-on-clear="` + gconv.String(defaultVal) + `" />`
		case TypeVarchar: // `varchar类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveField.isRequired = true
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', max: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.max.string', { max: `+v.FieldLimitStr+` }) },`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
			if v.IndexRaw == `UNI` {
				viewSaveField.form.DataType = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>`
			}
		case TypeChar: // `char类型`
			if v.IndexRaw == `UNI` && !v.IsNull {
				viewSaveField.isRequired = true
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', len: `+v.FieldLimitStr+`, trigger: 'blur', message: t('validation.size.string', { size: `+v.FieldLimitStr+` }) },`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" />`
			if v.IndexRaw == `UNI` {
				viewSaveField.form.DataType = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <label>
                        <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                    </label>`
			}
		case TypeText: // `text类型`
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'blur', message: t('validation.input') },`)

			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<my-editor v-model="saveForm.data.` + v.FieldRaw + `" />`
		case TypeJson: // `json类型`
			if !v.IsNull {
				viewSaveField.isRequired = true
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{
                type: 'object',
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
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-alert :title="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + v.FieldRaw + `')" type="info" :show-icon="true" :closable="false" />
                    <el-input v-model="saveForm.data.` + v.FieldRaw + `" type="textarea" :autosize="{ minRows: 3 }" />`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveField.isRequired = true
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-date-picker v-model="saveForm.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />`
		case TypeDate: // `date类型`
			if !v.IsNull && gconv.String(v.Default) == `` {
				viewSaveField.isRequired = true
			}
			viewSaveField.rule.Method = ReturnType
			viewSaveField.rule.DataType = append(viewSaveField.rule.DataType, `{ type: 'string', trigger: 'change', message: t('validation.select') },`)
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-date-picker v-model="saveForm.data.` + v.FieldRaw + `" type="date" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" />`
		default:
			viewSaveField.form.Method = ReturnType
			viewSaveField.form.DataType = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :clearable="true" />`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

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
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', min: 0, trigger: 'change', message: t('validation.select') },`)
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<my-cascader v-model="saveForm.data.` + v.FieldRaw + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree', param: { filter: { excIdArr: saveForm.data.idArr } } }" :props="{ checkStrictly: true, emitPath: false }" />`
			viewSaveField.paramHandle.Method = ReturnTypeName
			viewSaveField.paramHandle.DataTypeName = `param.` + v.FieldRaw + ` === undefined ? param.` + v.FieldRaw + ` = 0 : null`
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			viewSaveField.importModule = append(viewSaveField.importModule, `import md5 from 'js-md5'`)
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'string', required: computed((): boolean => { return saveForm.data.idArr?.length ? false : true; }), min: 6, max: 20, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 20 }) },`)
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <label v-if="saveForm.data.idArr?.length">
                        <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>`
			viewSaveField.paramHandle.Method = ReturnTypeName
			viewSaveField.paramHandle.DataTypeName = `param.` + v.FieldRaw + ` ? param.` + v.FieldRaw + ` = md5(param.` + v.FieldRaw + `) : delete param.` + v.FieldRaw
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
			if len(tpl.Handle.LabelList) > 0 && gstr.CaseCamel(tpl.Handle.LabelList[0]) == v.FieldCaseCamel {
				viewSaveField.isRequired = true
			}
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
			viewSaveField.rule.Method = ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ pattern: /^[\p{L}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },`)
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
			viewSaveField.rule.Method = ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ pattern: /^[\p{L}][\p{L}\p{N}_]+$/u, trigger: 'blur', message: t('validation.account') },`)
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
			viewSaveField.rule.Method = ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') },`)
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
			viewSaveField.rule.Method = ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'email', trigger: 'blur', message: t('validation.email') },`)
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
			viewSaveField.rule.Method = ReturnUnion
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'url', trigger: 'blur', message: t('validation.url') },`)
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName,
				`// { required: true, message: t('validation.required') },`,
				`{ type: 'integer', min: 1, trigger: 'change', message: t('validation.select') },`,
			)
			viewSaveField.form.Method = ReturnTypeName
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewSaveField.form.DataTypeName = `<my-cascader v-model="saveForm.data.` + v.FieldRaw + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />`
			} else {
				viewSaveField.form.DataTypeName = `<my-select v-model="saveForm.data.` + v.FieldRaw + `" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />`
				viewSaveField.dataInitAfter.Method = ReturnTypeName
				viewSaveField.dataInitAfter.DataTypeName = `saveCommon.data.` + v.FieldRaw + ` ? saveCommon.data.` + v.FieldRaw + ` : undefined`
			}
			viewSaveField.paramHandle.Method = ReturnTypeName
			viewSaveField.paramHandle.DataTypeName = `param.` + v.FieldRaw + ` === undefined ? param.` + v.FieldRaw + ` = 0 : null`
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'integer', min: 0, max: 100, trigger: 'change', message: t('validation.between.number', { min: 0, max: 100 }) },`)
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<el-input-number v-model="saveForm.data.` + v.FieldRaw + `" :precision="0" :min="0" :max="100" :step="1" :step-strictly="true" controls-position="right" :value-on-clear="` + gconv.String(gconv.Int(v.Default)) + `" />
                    <label>
                        <el-alert :title="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + v.FieldRaw + `')" type="info" :show-icon="true" :closable="false" />
                    </label>`
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			defaultVal := gconv.String(v.Default)
			if defaultVal == `` {
				defaultVal = v.StatusList[0][0]
			}
			viewSaveField.dataInitBefore.Method = ReturnTypeName
			viewSaveField.dataInitBefore.DataTypeName = defaultVal
			if garray.NewFrom([]interface{}{TypeVarchar, TypeChar}).Contains(v.FieldType) {
				viewSaveField.dataInitBefore.DataTypeName = `'` + defaultVal + `'`
			}
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'enum', enum: (tm('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.status.`+v.FieldRaw+`') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },`)
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<el-radio-group v-model="saveForm.data.` + v.FieldRaw + `">
                        <el-radio v-for="(item, index) in (tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + v.FieldRaw + `') as any)" :key="index" :label="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>`
			if len(v.StatusList) > 5 { //超过5个状态用select组件，小于5个用radio组件
				viewSaveField.form.DataTypeName = `<el-select-v2 v-model="saveForm.data.` + v.FieldRaw + `" :options="tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + v.FieldRaw + `')" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :clearable="false" />`
			}
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'enum', enum: (tm('common.status.whether') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },`)
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<el-switch v-model="saveForm.data.` + v.FieldRaw + `" :active-value="1" :inactive-value="0" :inline-prompt="true" :active-text="t('common.yes')" :inactive-text="t('common.no')" style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);" />`
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewSaveField.form.Method = ReturnTypeName
				viewSaveField.form.DataTypeName = `<el-date-picker v-model="saveForm.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :default-time="new Date(2000, 0, 1, 23, 59, 59)" />`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			if v.FieldType == TypeVarchar {
				viewSaveField.form.Method = ReturnTypeName
				viewSaveField.form.DataTypeName = `<el-input v-model="saveForm.data.` + v.FieldRaw + `" type="textarea" :autosize="{ minRows: 3 }" maxlength="` + v.FieldLimitStr + `" :show-word-limit="true" />`
			}
		case TypeNameImageSuffix, TypeNameVideoSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if v.FieldType == TypeVarchar {
				viewSaveField.rule.Method = ReturnUnion
				viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName, `{ type: 'url', trigger: 'change', message: t('validation.upload') },`)
			} else {
				viewSaveField.rule.Method = ReturnTypeName
				viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName,
					`{ type: 'array', trigger: 'change', message: t('validation.upload'), defaultField: { type: 'url', message: t('validation.url') } },`,
					`// { type: 'array', max: 10, trigger: 'change', message: t('validation.max.upload', { max: 10 }), defaultField: { type: 'url', message: t('validation.url') } },`,
				)
				if !v.IsNull {
					viewSaveField.isRequired = true
				}
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
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<my-upload v-model="saveForm.data.` + v.FieldRaw + `"` + attrOfAdd + ` />`
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			if !v.IsNull {
				viewSaveField.isRequired = true
			}
			viewSaveField.dataInitBefore.Method = ReturnTypeName
			viewSaveField.dataInitBefore.DataTypeName = `[]`
			viewSaveField.rule.Method = ReturnTypeName
			viewSaveField.rule.DataTypeName = append(viewSaveField.rule.DataTypeName,
				`{ type: 'array', trigger: 'change', message: t('validation.required') },`,
				`// { type: 'array', max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },`,
			)
			viewSaveField.form.Method = ReturnTypeName
			viewSaveField.form.DataTypeName = `<el-tag v-for="(item, index) in saveForm.data.` + v.FieldRaw + `" :type="` + v.FieldRaw + `Handle.tagType[index % ` + v.FieldRaw + `Handle.tagType.length]" @close="` + v.FieldRaw + `Handle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px;">
                        {{ item }}
                    </el-tag>
                    <!-- <el-input-number v-if="` + v.FieldRaw + `Handle.visible" :ref="(el: any) => ` + v.FieldRaw + `Handle.ref = el" v-model="` + v.FieldRaw + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" @keyup.enter="` + v.FieldRaw + `Handle.addValue" @blur="` + v.FieldRaw + `Handle.addValue" size="small" style="width: 100px;" :controls="false" /> -->
                    <el-input v-if="` + v.FieldRaw + `Handle.visible" :ref="(el: any) => ` + v.FieldRaw + `Handle.ref = el" v-model="` + v.FieldRaw + `Handle.value" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" @keyup.enter="` + v.FieldRaw + `Handle.addValue" @blur="` + v.FieldRaw + `Handle.addValue" size="small" style="width: 100px;" />
                    <el-button v-else type="primary" size="small" @click="` + v.FieldRaw + `Handle.visibleChange">
                        <autoicon-ep-plus />{{ t('common.add') }}
                    </el-button>`
			viewSaveField.formHandle.Method = ReturnTypeName
			viewSaveField.formHandle.DataTypeName = `const ` + v.FieldRaw + `Handle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as string[],
    visibleChange: () => {
        ` + v.FieldRaw + `Handle.visible = true
        nextTick(() => {
            ` + v.FieldRaw + `Handle.ref?.focus()
        })
    },
    addValue: () => {
        if (` + v.FieldRaw + `Handle.value) {
            saveForm.data.` + v.FieldRaw + `.push(` + v.FieldRaw + `Handle.value)
        }
        ` + v.FieldRaw + `Handle.visible = false
        ` + v.FieldRaw + `Handle.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.` + v.FieldRaw + `.splice(saveForm.data.` + v.FieldRaw + `.indexOf(item), 1)
    },
})`
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewSave.importModule = append(viewSave.importModule, viewSaveField.importModule...)
		if viewSaveField.dataInitBefore.getData() != `` {
			viewSave.dataInitBefore = append(viewSave.dataInitBefore, v.FieldRaw+`: `+viewSaveField.dataInitBefore.getData()+`,`)
		}
		if viewSaveField.dataInitAfter.getData() != `` {
			viewSave.dataInitAfter = append(viewSave.dataInitAfter, v.FieldRaw+`: `+viewSaveField.dataInitAfter.getData()+`,`)
		}
		rule := viewSaveField.rule.getData()
		if viewSaveField.isRequired {
			rule = append([]string{`{ required: true, message: t('validation.required') },`}, rule...)
		}
		if len(rule) > 0 {
			viewSave.rule = append(viewSave.rule, v.FieldRaw+`: [`+gstr.Join(append([]string{``}, rule...), `
			`)+`
        ],`)
		} else {
			viewSave.rule = append(viewSave.rule, v.FieldRaw+`: [],`)
		}
		if viewSaveField.form.getData() != `` {
			viewSave.form = append(viewSave.form, `<el-form-item :label="t('`+tpl.ModuleDirCaseKebabReplace+`.`+tpl.TableCaseKebab+`.name.`+v.FieldRaw+`')" prop="`+v.FieldRaw+`">
                    `+viewSaveField.form.getData()+`
                </el-form-item>`)
		}
		if viewSaveField.formHandle.getData() != `` {
			viewSave.formHandle = append(viewSave.formHandle, viewSaveField.formHandle.getData())
		}
		if viewSaveField.paramHandle.getData() != `` {
			viewSave.paramHandle = append(viewSave.paramHandle, viewSaveField.paramHandle.getData())
		}
	}

	// 做一次去重
	viewSave.importModule = garray.NewStrArrayFrom(viewSave.importModule).Unique().Slice()
	viewSave.dataInitBefore = garray.NewStrArrayFrom(viewSave.dataInitBefore).Unique().Slice()
	viewSave.dataInitAfter = garray.NewStrArrayFrom(viewSave.dataInitAfter).Unique().Slice()
	viewSave.rule = garray.NewStrArrayFrom(viewSave.rule).Unique().Slice()
	viewSave.form = garray.NewStrArrayFrom(viewSave.form).Unique().Slice()
	viewSave.formHandle = garray.NewStrArrayFrom(viewSave.formHandle).Unique().Slice()
	viewSave.paramHandle = garray.NewStrArrayFrom(viewSave.paramHandle).Unique().Slice()
	return
}
