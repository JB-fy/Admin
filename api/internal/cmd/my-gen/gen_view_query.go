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

// 视图模板Query生成
func genViewQuery(option myGenOption, tpl myGenTpl) {
	viewQuery := getViewQueryFieldList(tpl)

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
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">
        <el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="1" :controls="false" />
        </el-form-item>` + gstr.Join(append([]string{``}, viewQuery.form...), `
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

func getViewQueryFieldList(tpl myGenTpl) (viewQuery myGenViewQuery) {

	for _, v := range tpl.FieldList {
		viewQueryField := myGenViewQueryField{}
		viewQueryField.formProp.Method = ReturnType
		viewQueryField.formProp.DataType = v.FieldRaw

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt: // `int等类型`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :controls="false" />`
		case TypeIntU: // `int等类型（unsigned）`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :controls="false" />`
		case TypeFloat: // `float等类型`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />`
		case TypeFloatU: // `float等类型（unsigned）`
			// viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="0" :precision="` + v.FieldLimitFloat[1] + `" :controls="false" />`
		case TypeVarchar: // `varchar类型`
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" maxlength="` + v.FieldLimitStr + `" :clearable="true" />`
		case TypeChar: // `char类型`
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" minlength="` + v.FieldLimitStr + `" maxlength="` + v.FieldLimitStr + `" :clearable="true" />`
		case TypeText: // `text类型`
		case TypeJson: // `json类型`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />`
		case TypeDate: // `date类型`
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="date" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 150px" />`
		default:
			viewQueryField.form.Method = ReturnType
			viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :clearable="true" />`
		}
		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

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
		case TypeNamePri: // 主键
		case TypeNamePriAutoInc: // 主键（自增）
			continue
		case TypeNamePid: // pid；	类型：int等类型；
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree' }" :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :props="{ checkStrictly: true, emitPath: false }" />`
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :min="1" :controls="false" />`
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
			viewQueryField.form.DataTypeName = `<my-select v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />`
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.Pid.Pid != `` {
				viewQueryField.form.DataTypeName = `<my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree' }" :props="{ emitPath: false }" />`
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			continue
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-select-v2 v-model="queryCommon.data.` + v.FieldRaw + `" :options="tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + v.FieldRaw + `')" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :clearable="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			viewQueryField.form.Method = ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-select-v2 v-model="queryCommon.data.` + v.FieldRaw + `" :options="tm('common.status.whether')" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" :clearable="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
			if v.FieldType != TypeDate {
				viewQueryField.form.Method = ReturnTypeName
				viewQueryField.form.DataTypeName = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :default-time="new Date(2000, 0, 1, 23, 59, 59)" />`
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

		if viewQueryField.dataInit.getData() != `` {
			viewQuery.dataInit = append(viewQuery.dataInit, viewQueryField.dataInit.getData())
		}
		if viewQueryField.form.getData() != `` {
			viewQuery.form = append(viewQuery.form, `<el-form-item prop="`+viewQueryField.formProp.getData()+`">
            `+viewQueryField.form.getData()+`
        </el-form-item>`)
		}
	}

	// 做一次去重
	viewQuery.dataInit = garray.NewStrArrayFrom(viewQuery.dataInit).Unique().Slice()
	viewQuery.form = garray.NewStrArrayFrom(viewQuery.form).Unique().Slice()
	return
}
