package my_gen

import (
	"api/internal/cmd/my-gen/internal"
	"slices"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewQuery struct {
	isI18nTm bool
	dataInit []string
	form     []string
}

type myGenViewQueryField struct {
	isI18nTm bool
	dataInit internal.MyGenDataStrHandler
	formProp internal.MyGenDataStrHandler
	form     internal.MyGenDataStrHandler
}

func (viewQueryThis *myGenViewQuery) Add(viewQueryField myGenViewQueryField) {
	if viewQueryField.form.GetData() == `` {
		return
	}
	if viewQueryField.isI18nTm {
		viewQueryThis.isI18nTm = true
	}
	if viewQueryField.dataInit.GetData() != `` {
		viewQueryThis.dataInit = append(viewQueryThis.dataInit, viewQueryField.dataInit.GetData())
	}
	viewQueryThis.form = append(viewQueryThis.form, `<el-form-item prop="`+viewQueryField.formProp.GetData()+`">
            `+viewQueryField.form.GetData()+`
        </el-form-item>`)
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
func genViewQuery(option myGenOption, tpl *myGenTpl) {
	viewQuery := getViewQueryIdAndLabel(tpl)
	for _, v := range tpl.FieldListOfDefault {
		viewQuery.Add(getViewQueryField(tpl, v, tpl.I18nPath, v.FieldRaw))
	}
	for _, v := range tpl.FieldListOfAfter1 {
		viewQuery.Add(getViewQueryField(tpl, v, tpl.I18nPath, v.FieldRaw))
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		viewQuery.Merge(getViewQueryExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		viewQuery.Merge(getViewQueryExtendMiddleOne(v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		viewQuery.Merge(getViewQueryExtendMiddleMany(v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		viewQuery.Merge(getViewQueryExtendMiddleMany(v))
	}
	for _, v := range tpl.FieldListOfAfter2 {
		viewQuery.Add(getViewQueryField(tpl, v, tpl.I18nPath, v.FieldRaw))
	}
	viewQuery.Unique()

	tplView := `<script setup lang="tsx">
import dayjs from 'dayjs'

const { t`
	if viewQuery.isI18nTm {
		tplView += `, tm`
	}
	tplView += ` } = useI18n()

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
        listCommon.ref.getList(true).finally(() => (queryForm.loading = false))
    },
    reset: () => queryForm.ref.resetFields(),
})
</script>

<template>
    <el-form class="query-form" :ref="(el: any) => queryForm.ref = el" :model="queryCommon.data" :inline="true" @keyup.enter="queryForm.submit">` + gstr.Join(append([]string{``}, viewQuery.form...), `
        `) + `
        <el-form-item>
            <el-button type="primary" @click="queryForm.submit" :loading="queryForm.loading"><autoicon-ep-search />{{ t('common.query') }}</el-button>
            <el-button type="info" @click="queryForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneId + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/Query.vue`
	gfile.PutContents(saveFile, tplView)
}

func getViewQueryIdAndLabel(tpl *myGenTpl) (viewQuery myGenViewQuery) {
	if len(tpl.Handle.Id.List) == 1 {
		switch tpl.Handle.Id.List[0].FieldType {
		case internal.TypeInt, internal.TypeIntU:
			viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input-number v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :min="`+tpl.Handle.Id.List[0].FieldLimitInt.Min+`" :max="`+tpl.Handle.Id.List[0].FieldLimitInt.Max+`" :precision="0" :controls="false" />
        </el-form-item>`)
		default:
			viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input v-model="queryCommon.data.id" :placeholder="t('common.name.id')" maxlength="`+tpl.Handle.Id.List[0].FieldLimitStr+`" :clearable="true" />
        </el-form-item>`)
		}
	} else {
		viewQuery.form = append(viewQuery.form, `<el-form-item prop="id">
            <el-input v-model="queryCommon.data.id" :placeholder="t('common.name.id')" :clearable="true" />
        </el-form-item>`)
	}

	if len(tpl.Handle.Label.List) == 1 && (!tpl.Handle.Label.IsDefault || slices.Contains([]internal.MyGenFieldType{internal.TypeVarchar, internal.TypeChar}, tpl.Handle.Label.List[0].FieldType)) {
		viewQuery.form = append(viewQuery.form, `<el-form-item prop="label">
			<el-input v-model="queryCommon.data.label" :placeholder="t('`+tpl.I18nPath+`.name.`+tpl.Handle.Label.List[0].FieldRaw+`')" maxlength="`+tpl.Handle.Label.List[0].FieldLimitStr+`" :clearable="true" />
		</el-form-item>`)
	} else if len(tpl.Handle.Label.List) > 1 {
		viewQuery.form = append(viewQuery.form, `<el-form-item prop="label">
			<el-input v-model="queryCommon.data.label" :placeholder="t('common.name.label')" maxlength="30" :clearable="true" />
		</el-form-item>`)
	}
	return
}

func getViewQueryField(tpl *myGenTpl, v myGenField, i18nPath string, i18nFieldPath string) (viewQueryField myGenViewQueryField) {
	viewQueryField.formProp.Method = internal.ReturnType
	viewQueryField.formProp.DataType = v.FieldRaw

	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt, internal.TypeIntU: // `int等类型`	// `int等类型（unsigned）`
		// viewQueryField.form.Method = internal.ReturnType
		viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :min="` + v.FieldLimitInt.Min + `" :max="` + v.FieldLimitInt.Max + `" :precision="0" :controls="false" />`
	case internal.TypeFloat, internal.TypeFloatU: // `float等类型`	// `float等类型（unsigned）`
		attrOfAdd := ``
		if v.FieldLimitFloat.Min != `` {
			attrOfAdd += ` :min="` + v.FieldLimitFloat.Min + `"`
		}
		if v.FieldLimitFloat.Max != `` {
			attrOfAdd += ` :max="` + v.FieldLimitFloat.Max + `"`
		}
		// viewQueryField.form.Method = internal.ReturnType
		viewQueryField.form.DataType = `<el-input-number v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')"` + attrOfAdd + ` :precision="` + gconv.String(v.FieldLimitFloat.Precision) + `" :controls="false" />`
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型`	// `char类型`
		if (v.IsUnique || gconv.Uint(v.FieldLimitStr) <= internal.ConfigMaxLenOfStrFilter) && !(len(tpl.Handle.Label.List) == 1 && tpl.Handle.Label.List[0].FieldRaw == v.FieldRaw) {
			attrOfAdd := ``
			if v.FieldType == internal.TypeChar /* && v.FieldTypeName != internal.TypeNameNameSuffix */ {
				attrOfAdd = ` minlength="` + v.FieldLimitStr + `"`
			}
			viewQueryField.form.Method = internal.ReturnType
			viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')"` + attrOfAdd + ` maxlength="` + v.FieldLimitStr + `" :clearable="true" />`
		}
	case internal.TypeText: // `text类型`
	case internal.TypeJson: // `json类型`
	case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		// viewQueryField.form.Method = internal.ReturnType
		viewQueryField.form.DataType = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />`
	case internal.TypeDate: // `date类型`
		viewQueryField.form.Method = internal.ReturnType
		viewQueryField.form.DataType = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="date" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />`
	case internal.TypeTime: // `time类型`
		// viewQueryField.form.Method = internal.ReturnType
		viewQueryField.form.DataType = `<el-time-picker v-model="queryCommon.data.` + v.FieldRaw + `" placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="HH:mm:ss" value-format="HH:mm:ss" />`
	default:
		viewQueryField.form.Method = internal.ReturnType
		viewQueryField.form.DataType = `<el-input v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="true" />`
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
		return myGenViewQueryField{}
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		return myGenViewQueryField{}
	case internal.TypePrimaryMany: // 联合主键
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
		viewQueryField.form.Method = internal.ReturnType
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
		return myGenViewQueryField{}
	case internal.TypeNameUpdated: // 更新时间字段
		return myGenViewQueryField{}
	case internal.TypeNameCreated: // 创建时间字段
		viewQueryField.dataInit.Method = internal.ReturnTypeName
		viewQueryField.dataInit.DataTypeName = internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`) + `: undefined, //[new Date().setHours(0, 0, 0), new Date().setHours(23, 59, 59)]
    ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range_start`) + `: computed(() => (queryCommon.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`) + `?.length ? dayjs(queryCommon.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`) + `[0]).format('YYYY-MM-DD HH:mm:ss') : undefined)),
    ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range_end`) + `: computed(() => (queryCommon.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`) + `?.length ? dayjs(queryCommon.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`) + `[1]).format('YYYY-MM-DD HH:mm:ss') : undefined)),`

		viewQueryField.formProp.Method = internal.ReturnTypeName
		viewQueryField.formProp.DataTypeName = internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`)
		viewQueryField.form.Method = internal.ReturnTypeName
		viewQueryField.form.DataTypeName = `<el-date-picker v-model="queryCommon.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `time_range`) + `" type="datetimerange" range-separator="-" :default-time="[new Date(2000, 0, 1, 0, 0, 0), new Date(2000, 0, 1, 23, 59, 59)]" :start-placeholder="t('common.name.timeRangeStart')" :end-placeholder="t('common.name.timeRangeEnd')" />`
	case internal.TypeNamePid: // pid，且与主键类型相同时（才）有效；	类型：int等类型或varchar或char；
		viewQueryField.isI18nTm = true
		viewQueryField.form.Method = internal.ReturnTypeName
		options := `tm('common.status.pid')`
		if !slices.Contains([]internal.MyGenFieldType{internal.TypeInt, internal.TypeIntU}, v.FieldType) {
			options = `tm('common.status.pidStr')`
		}
		viewQueryField.form.DataTypeName = `<my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/tree'` + tpl.Handle.Pid.Tpl.PidDefValOfView + ` }" :options="` + options + `" :props="{ checkStrictly: true, emitPath: false }" />`
	case internal.TypeNameIdPath, internal.TypeNameNamePath: // id_path|idPath，且pid同时存在时（才）有效；	类型：varchar或text；	// name_path|namePath，且pid，id_path|idPath同时存在时（才）有效；	类型：varchar或text；
		return myGenViewQueryField{}
	case internal.TypeNameLevel: // level，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
		viewQueryField.form.Method = internal.ReturnType
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		return myGenViewQueryField{}
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenViewQueryField{}
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		return myGenViewQueryField{}
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型或varchar或char；
		viewQueryField.form.Method = internal.ReturnTypeName
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl != nil {
			apiUrl := relIdObj.tpl.ModuleDirCaseKebab + `/` + relIdObj.tpl.TableCaseKebab
			if relIdObj.tpl.Handle.Pid.Pid != `` {
				viewQueryField.form.DataTypeName = `<my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree'` + tpl.Handle.Pid.Tpl.PidDefValOfView + ` }" :props="{ emitPath: false }" />`
			} else {
				viewQueryField.form.DataTypeName = `<my-select v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" />`
			}
		} else {
			apiUrl := tpl.ModuleDirCaseKebab + `/` + gstr.CaseKebab(gstr.SubStr(v.FieldCaseCamelRemove, 0, -2))
			viewQueryField.form.DataTypeName = `<!-- 可选择组件<my-select>或<my-cascader>使用，但需手动确认关联表，并修改接口路径 -->
            ` + viewQueryField.form.DataType + `
            <!-- <my-select v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/list' }" /> -->
            <!-- <my-cascader v-model="queryCommon.data.` + v.FieldRaw + `" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + apiUrl + `/tree'` + tpl.Handle.Pid.Tpl.PidDefValOfView + ` }" :props="{ emitPath: false }" /> -->`
		}
	case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		viewQueryField.isI18nTm = true
		viewQueryField.form.Method = internal.ReturnTypeName
		viewQueryField.form.DataTypeName = `<el-select-v2 v-model="queryCommon.data.` + v.FieldRaw + `" :options="tm('` + i18nPath + `.status.` + i18nFieldPath + `')" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
	case internal.TypeNameIsPrefix, internal.TypeNameIsLeaf: // is_前缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）	// is_leaf|isLeaf，且pid，id_path|idPath同时存在时（才）有效；	类型：int等类型；
		viewQueryField.isI18nTm = true
		viewQueryField.form.Method = internal.ReturnTypeName
		tmKey := i18nPath + `.status.` + i18nFieldPath
		if v.StatusWhetherI18n != `` {
			tmKey = v.StatusWhetherI18n
		}
		viewQueryField.form.DataTypeName = `<el-select-v2 v-model="queryCommon.data.` + v.FieldRaw + `" :options="tm('` + tmKey + `')" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" :clearable="true" style="width: ` + gconv.String(100+(v.FieldShowLenMax-3)*14) + `px" />`
	case internal.TypeNameSortSuffix: // sort,num,number,weight等后缀；	类型：int等类型；
		return myGenViewQueryField{}
	case internal.TypeNameNoSuffix: // no,level,rank等后缀；	类型：int等类型；
		viewQueryField.form.Method = internal.ReturnType
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
		viewQueryField.form.Method = internal.ReturnType
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
		switch v.FieldType {
		case internal.TypeDatetime, internal.TypeTimestamp:
			viewQueryField.form.Method = internal.ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-date-picker v-model="queryCommon.data.` + v.FieldRaw + `" type="datetime" :placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" :default-time="new Date(2000, 0, 1, 23, 59, 59)" />`
		case internal.TypeDate:
		case internal.TypeTime:
			viewQueryField.form.Method = internal.ReturnTypeName
			viewQueryField.form.DataTypeName = `<el-time-picker v-model="queryCommon.data.` + v.FieldRaw + `" placeholder="t('` + i18nPath + `.name.` + i18nFieldPath + `')" format="HH:mm:ss" value-format="HH:mm:ss" :default-value="new Date(2000, 0, 1, 23, 59, 59)" />`
		}
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		return myGenViewQueryField{}
	case internal.TypeNameImageSuffix, internal.TypeNameVideoSuffix, internal.TypeNameAudioSuffix, internal.TypeNameFileSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：varchar或json或text	// video,video_list,videoList,video_arr,videoArr等后缀；	类型：varchar或json或text	// audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：varchar或json或text	// file,file_list,fileList,file_arr,fileArr等后缀；	类型：varchar或json或text
		return myGenViewQueryField{}
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：varchar或json或text；
		return myGenViewQueryField{}
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getViewQueryExtendMiddleOne(tplEM handleExtendMiddle) (viewQuery myGenViewQuery) {
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, v.FieldRaw))
		}
	case internal.TableTypeMiddleOne:
		for _, v := range tplEM.FieldListOfIdSuffix {
			viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, v.FieldRaw))
		}
		for _, v := range tplEM.FieldListOfOther {
			viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, v.FieldRaw))
		}
	}
	return
}

func getViewQueryExtendMiddleMany(tplEM handleExtendMiddle) (viewQuery myGenViewQuery) {
	if len(tplEM.FieldList) == 1 {
		for _, v := range tplEM.FieldList {
			viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, tplEM.FieldVar))
		}
	} else {
		switch tplEM.TableType {
		case internal.TableTypeExtendMany:
			for _, v := range tplEM.FieldList {
				viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw))
			}
		case internal.TableTypeMiddleMany:
			for _, v := range tplEM.FieldListOfIdSuffix {
				viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw))
			}
			for _, v := range tplEM.FieldListOfOther {
				viewQuery.Add(getViewQueryField(tplEM.tpl, v, tplEM.tplOfTop.I18nPath, tplEM.FieldVar+`.`+v.FieldRaw))
			}
		}
	}
	return
}
