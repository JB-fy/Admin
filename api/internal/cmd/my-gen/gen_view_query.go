package my_gen

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewQuery struct {
	dataInit []string
	form     []string
}

type myGenViewQueryField struct {
	dataInit myGenDataStrHandler
	formProp myGenDataStrHandler
	form     myGenDataStrHandler
}

func (viewQueryThis *myGenViewQuery) Add(viewQueryField myGenViewQueryField) {
	if viewQueryField.dataInit.getData() != `` {
		viewQueryThis.dataInit = append(viewQueryThis.dataInit, viewQueryField.dataInit.getData())
	}
	if viewQueryField.form.getData() != `` {
		viewQueryThis.form = append(viewQueryThis.form, `<el-form-item prop="`+viewQueryField.formProp.getData()+`">
            `+viewQueryField.form.getData()+`
        </el-form-item>`)
	}
}

func (viewQueryThis *myGenViewQuery) Merge(viewQueryOther myGenViewQuery) {
	viewQueryThis.dataInit = append(viewQueryThis.dataInit, viewQueryOther.dataInit...)
	viewQueryThis.form = append(viewQueryThis.form, viewQueryOther.form...)
}

func (viewQueryThis *myGenViewQuery) Unique() {
	// viewQueryThis.dataInit = garray.NewStrArrayFrom(viewQueryThis.dataInit).Unique().Slice()
	// viewQueryThis.form = garray.NewStrArrayFrom(viewQueryThis.form).Unique().Slice()
}

// 视图模板Query生成
func genViewQuery(option myGenOption, tpl myGenTpl) {
	viewQuery := getViewQueryIdAndLabel(tpl)
	viewQuery.Merge(getViewQueryFieldList(tpl, tpl.I18nPath, tpl.FieldArr...))
	for _, v := range tpl.Handle.ExtendTableOneList {
		viewQuery.Merge(getViewQueryExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		viewQuery.Merge(getViewQueryExtendMiddleOne(v))
	}
	for _, v := range tpl.FieldArrAfter {
		viewQuery.Merge(getViewQueryFieldList(tpl, tpl.I18nPath, v))
	}
	viewQuery.Unique()

	tplView := `<script setup lang="tsx">
import dayjs from 'dayjs'

const { t, tm } = useI18n()

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
queryCommon.data = {
    ...queryCommon.data,` + gstr.Join(append([]string{``}, viewQuery.dataInit...), `
    `) + `
}
const listCommon = inject('listCommon') as { ref: any }
const queryForm = reactive({
    ref: null as any,
    loading: false,
    submit: () => {
        queryForm.loading = true
        listCommon.ref.getList(true).finally(() => {
            queryForm.loading = false
        })
    },
    reset: () => {
        queryForm.ref.resetFields()
        //queryForm.submit()
    },
})
</script>

<template>
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">` + gstr.Join(append([]string{``}, viewQuery.form...), `
        `) + `
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"> <autoicon-ep-search />{{ t('common.query') }} </el-button>
            <el-button type="info" @click="queryForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Query.vue`
	gfile.PutContents(saveFile, tplView)
}

func getViewQueryIdAndLabel(tpl myGenTpl) (viewQuery myGenViewQuery) {
	if len(tpl.Handle.Id.List) == 1 {
		switch tpl.Handle.Id.List[0].FieldType {
		case TypeInt:
			viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :controls="false" />
        </el-form-item>`)
		case TypeIntU:
			viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </el-form-item>`)
		default:
			viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :maxlength="`+tpl.Handle.Id.List[0].FieldLimitStr+`" :clearable="true" />
        </el-form-item>`)
		}
	} else {
		viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :clearable="true" />
        </el-form-item>`)
	}
	return
}

func getViewQueryFieldList(tpl myGenTpl, i18nPath string, fieldArr ...string) (viewQuery myGenViewQuery) {
	for _, v := range tpl.FieldList {
		if len(fieldArr) > 0 && !garray.NewStrArrayFrom(fieldArr).Contains(v.FieldRaw) {
			continue
		}

		viewQueryField := myGenViewQueryField{}
		viewQueryField.formProp.Method = ReturnType
		viewQueryField.formProp.DataType = v.FieldRaw

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :controls="false" />`
		case TypeIntU: // `int等类型（unsigned）`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :min="0" :controls="false" />`
		case TypeFloat: // `float等类型`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />`
		case TypeFloatU: // `float等类型（unsigned）`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />`
		case TypeVarchar: // `varchar类型`
			if gconv.Uint(v.FieldLimitStr) <= configMaxLenOfStrFilter {
				viewQueryField.form.Method = ReturnType
				viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" maxlength="` + v.FieldLimitStr + `" :clearable="true" />`
			}
		case TypeChar: // `char类型`
			if gconv.Uint(v.FieldLimitStr) <= configMaxLenOfStrFilter {
				viewQueryField.form.Method = ReturnType
				viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :clearable="true" />`
			}
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />`
		case TypeDate: // `date类型`
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="date" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />`
		default:
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :clearable="true" />`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

		/*--------根据字段主键类型处理 开始--------*/
		switch v.FieldTypePrimary {
		case TypePrimary: // 独立主键
			if v.FieldRaw == `id` {
				continue
			}
		case TypePrimaryAutoInc: // 独立主键（自增）
			continue
		case TypePrimaryMany: // 联合主键
		case TypePrimaryManyAutoInc: // 联合主键（自增）
		}
		/*--------根据字段主键类型处理 结束--------*/

		/*--------根据字段命名类型处理 开始--------*/
		switch v.FieldTypeName {
		case TypeNameDeleted: // 软删除字段
			continue
		case TypeNameUpdated: // 更新时间字段
			continue
		case TypeNameCreated: // 创建时间字段
			viewQueryField.dataInit.Method = ReturnTypeName
			viewQueryField.dataInit.DataTypeName = `timeRange: (() => {
        return undefined
        /* const date = new Date()
        return [
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 0, 0, 0),
            new Date(date.getFullYear(), date.getMonth(), date.getDate(), 23, 59, 59),
        ] */
    })(),
    timeRangeStart: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[0]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),
    timeRangeEnd: computed(() => {
        if (queryCommon.data.timeRange?.length) {
            return dayjs(queryCommon.data.timeRange[1]).format('YYYY-MM-DD HH:mm:ss')
        }
        return ''
    }),`

			viewQueryField.formProp.Method = ReturnTypeName
			viewQueryField.formProp.DataTypeName = `timeRange`
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-date-picker v-model="queryCommon.data.timeRange" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />`
		case TypeNamePid: // pid；	类型：int等类型；
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />`
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :min="1" :controls="false" />`
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			continue
		case TypeNamePasswordSuffix: // password,passwd后缀；		类型：char(32)；
			continue
		case TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
			continue
		case TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		case TypeNameCodeSuffix: // code后缀；	类型：varchar；
		case TypeNameAccountSuffix: // account后缀；	类型：varchar；
		case TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
		case TypeNameEmailSuffix: // email后缀；	类型：varchar；
		case TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
		case TypeNameIpSuffix: // IP后缀；	类型：varchar；
		case TypeNameIdSuffix: // id后缀；	类型：int等类型；
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` {
				relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
				apiUrl = relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			}

			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<my-select v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />`
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewQueryField.form.DataTypeName = `<my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />`
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-select-v2 v-model="queryCommon.data.` + v.FieldRaw + `" :options="tm('` + i18nPath + `.status.` + v.FieldRaw + `')" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :clearable="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-select-v2 v-model="queryCommon.data.` + v.FieldRaw + `" :options="tm('common.status.whether')" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" :clearable="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewQueryField.form.Method = ReturnTypeName
				viewQueryField.form.DataTypeName = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :default-time="new Date(2000, 0, 1, 23, 59, 59)" />`
			}
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
			continue
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			continue
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			continue
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			continue
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewQuery.Add(viewQueryField)
	}
	return
}
func getViewQueryExtendMiddleOne(tplEM handleExtendMiddle) (viewQuery myGenViewQuery) {
	switch tplEM.TableType {
	case TableTypeExtendOne:
		viewQuery.Merge(getViewQueryFieldList(tplEM.tpl, tplEM.tplOfTop.I18nPath, tplEM.FieldArr...))
	case TableTypeMiddleOne:
		viewQuery.Merge(getViewQueryFieldList(tplEM.tpl, tplEM.tplOfTop.I18nPath, tplEM.FieldArrOfIdSuffix...))
		viewQuery.Merge(getViewQueryFieldList(tplEM.tpl, tplEM.tplOfTop.I18nPath, tplEM.FieldArrOfOther...))
	}
	return
}
