package my_gen

import (
	"api/internal/cmd/my-gen/internal"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewList struct {
	rowHeight uint
	isI18nTm  bool
	idType    string
	columns   []string
}

type myGenViewListField struct {
	rowHeight    uint
	isI18nTm     bool
	dataKey      internal.MyGenDataStrHandler
	title        internal.MyGenDataStrHandler
	key          internal.MyGenDataStrHandler
	align        internal.MyGenDataStrHandler
	width        internal.MyGenDataStrHandler
	sortable     internal.MyGenDataStrHandler
	hidden       internal.MyGenDataStrHandler
	cellRenderer internal.MyGenDataStrHandler
}

func (viewListThis *myGenViewList) Add(viewListField myGenViewListField) {
	if viewListField.dataKey.GetData() == `` {
		return
	}
	if viewListThis.rowHeight < viewListField.rowHeight {
		viewListThis.rowHeight = viewListField.rowHeight
	}
	if viewListField.isI18nTm {
		viewListThis.isI18nTm = true
	}
	columnAttrStr := []string{
		`dataKey: ` + viewListField.dataKey.GetData() + `,`,
		`title: ` + viewListField.title.GetData() + `,`,
		`key: ` + viewListField.key.GetData() + `,`,
		`align: ` + viewListField.align.GetData() + `,`,
		`width: ` + viewListField.width.GetData() + `,`,
	}
	if viewListField.sortable.GetData() != `` {
		columnAttrStr = append(columnAttrStr, `sortable: `+viewListField.sortable.GetData()+`,`)
	}
	if viewListField.hidden.GetData() != `` {
		columnAttrStr = append(columnAttrStr, `hidden: `+viewListField.hidden.GetData()+`,`)
	}
	if viewListField.cellRenderer.GetData() != `` {
		columnAttrStr = append(columnAttrStr, `cellRenderer: `+viewListField.cellRenderer.GetData()+`,`)
	}
	viewListThis.columns = append(viewListThis.columns, `{`+gstr.Join(append([]string{``}, columnAttrStr...), `
            `)+`
        },`)
}

func (viewListThis *myGenViewList) Merge(viewListOther myGenViewList) {
	if viewListThis.rowHeight < viewListOther.rowHeight {
		viewListThis.rowHeight = viewListOther.rowHeight
	}
	viewListThis.columns = append(viewListThis.columns, viewListOther.columns...)
}

func (viewListThis *myGenViewList) Unique() {
	// viewListThis.columns = garray.NewStrArrayFrom(viewListThis.columns).Unique().Slice()
}

// 视图模板List生成
func genViewList(option myGenOption, tpl myGenTpl) {
	viewList := myGenViewList{
		rowHeight: 50,
		idType:    `number`,
	}
	if len(tpl.Handle.Id.List) > 1 || !garray.NewIntArrayFrom([]int{internal.TypeInt, internal.TypeIntU}).Contains(tpl.Handle.Id.List[0].FieldType) {
		viewList.idType = `string`
	}
	for _, v := range tpl.FieldListOfDefault {
		viewList.Add(getViewListField(option, tpl, v, tpl.I18nPath))
	}
	for _, v := range tpl.FieldListOfAfter1 {
		viewList.Add(getViewListField(option, tpl, v, tpl.I18nPath))
	}
	for _, v := range tpl.Handle.ExtendTableOneList {
		viewList.Merge(getViewListExtendMiddleOne(option, v))
	}
	for _, v := range tpl.Handle.MiddleTableOneList {
		viewList.Merge(getViewListExtendMiddleOne(option, v))
	}
	for _, v := range tpl.Handle.ExtendTableManyList {
		viewList.Merge(getViewListExtendMiddleMany(option, v))
	}
	for _, v := range tpl.Handle.MiddleTableManyList {
		viewList.Merge(getViewListExtendMiddleMany(option, v))
	}
	for _, v := range tpl.FieldListOfAfter2 {
		viewList.Add(getViewListField(option, tpl, v, tpl.I18nPath))
	}
	viewList.Unique()

	tplView := `<script setup lang="tsx">
const { t`
	if viewList.isI18nTm {
		tplView += `, tm`
	}
	tplView += ` } = useI18n()`
	if option.IsCreate || option.IsUpdate || option.IsDelete {
		tplView += `

const authAction = inject('authAction') as { [propName: string]: boolean }`
	}
	tplView += `

const table = reactive({
    columns: [
        {
            dataKey: 'id',
            title: t('common.name.id'),
            key: 'id',
            align: 'center',
            width: 200,
            fixed: 'left',
            sortable: true,`
	if option.IsUpdate || option.IsDelete {
		tplView += `
            headerCellRenderer: () => {
                const allChecked = table.data.every((item: any) => item.checked)
                const someChecked = table.data.some((item: any) => item.checked)
                return [
                    //阻止冒泡
                    <div class="id-checkbox" onClick={(event: any) => event.stopPropagation()}>
                        <el-checkbox
                            model-value={table.data.length ? allChecked : false}
                            indeterminate={someChecked && !allChecked}
                            onChange={(val: boolean) => {
                                table.data.forEach((item: any) => {
                                    item.checked = val
                                })
                            }}
                        />
                    </div>,
                    <div>{t('common.name.id')}</div>,
                ]
            },
            cellRenderer: (props: any): any => {
                return [<el-checkbox class="id-checkbox" model-value={props.rowData.checked} onChange={(val: boolean) => (props.rowData.checked = val)} />, <div>{props.rowData.id}</div>]
            },`
	}
	tplView += `
        },` + gstr.Join(append([]string{``}, viewList.columns...), `
        `)
	if option.IsCreate || option.IsUpdate || option.IsDelete {
		tplView += `
        {
            title: t('common.name.action'),
            key: 'action',
            align: 'center',
            width: 80 * ((authAction.isCreate ? 1 : 0) + (authAction.isUpdate ? 1 : 0) + (authAction.isDelete ? 1 : 0)),
            fixed: 'right',
            hidden: !(authAction.isCreate || authAction.isUpdate || authAction.isDelete),
            cellRenderer: (props: any): any => {
                let vNode: any = []`
		if option.IsUpdate {
			tplView += `
                if (authAction.isUpdate) {
                    vNode.push(
                        <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                            <autoicon-ep-edit />
                            {t('common.edit')}
                        </el-button>
                    )
                }`
		}
		if option.IsDelete {
			tplView += `
                if (authAction.isDelete) {
                    vNode.push(
                        <el-button type="danger" size="small" onClick={() => handleDelete(props.rowData.id)}>
                            <autoicon-ep-delete />
                            {t('common.delete')}
                        </el-button>
                    )
                }`
		}
		if option.IsCreate {
			tplView += `
                if (authAction.isCreate) {
                    vNode.push(
                        <el-button type="warning" size="small" onClick={() => handleEditCopy(props.rowData.id, 'copy')}>
                            <autoicon-ep-document-copy />
                            {t('common.copy')}
                        </el-button>
                    )
                }`
		}
		tplView += `
                return vNode
            },
        },`
	}
	tplView += `
    ] as any,
    data: [],
    loading: false,
    sort: { key: 'id', order: 'desc' } as any,
    handleSort: (sort: any) => {
        table.sort.key = sort.key
        table.sort.order = sort.order
        getList()
    },
})`
	if option.IsCreate || option.IsUpdate {
		tplView += `

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }`
	}
	if option.IsCreate {
		tplView += `
//新增
const handleAdd = () => {
    saveCommon.data = {}
    saveCommon.title = t('common.add')
    saveCommon.visible = true
}`
	}
	if option.IsDelete {
		tplView += `
//批量删除
const handleBatchDelete = () => {
    const idArr: ` + viewList.idType + `[] = []
    table.data.forEach((item: any) => item.checked && idArr.push(item.id))
    idArr.length == 0 ? ElMessage.error(t('common.tip.selectDelete')) : handleDelete(idArr)
}`
	}
	if option.IsCreate || option.IsUpdate {
		tplView += `
//编辑|复制
const handleEditCopy = (id: ` + viewList.idType + `, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/info', { id: id }).then((res) => {
        saveCommon.data = { ...res.data.info }
        switch (type) {
            case 'edit':
                saveCommon.data.` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + ` = [saveCommon.data.id]
                delete saveCommon.data.id
                saveCommon.title = t('common.edit')
                break
            case 'copy':
                delete saveCommon.data.id
                saveCommon.title = t('common.copy')
                break
        }
        saveCommon.visible = true
    })
}`
	}
	if option.IsDelete {
		tplView += `
//删除
const handleDelete = (idArr: ` + viewList.idType + `[]) => {
    ElMessageBox.confirm('', {
        type: 'warning',
        title: t('common.tip.configDelete'),
        center: true,
        showClose: false,
    }).then(() => request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/del', { ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `: idArr }, true).then(() => getList()))
}`
	}
	if option.IsUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `: ` + viewList.idType + `[]; [propName: string]: any }) => {
    await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/update', param, true)
}`
	}
	tplView += `

//分页
const settingStore = useSettingStore()
const pagination = reactive({
    total: 0,
    page: 1,
    size: settingStore.pagination.size,
    sizeList: settingStore.pagination.sizeList,
    layout: settingStore.pagination.layout,
    sizeChange: () => getList(),
    pageChange: () => getList(),
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
    resetPage && (pagination.page = 1)
    const param = {
        field: [],
        filter: removeEmptyOfObj(queryCommon.data, true, true),
        sort: table.sort.key + ' ' + table.sort.order,
        page: pagination.page,
        limit: pagination.size,
    }
    table.loading = true
    try {
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param)
        table.data = res.data.list?.length ? res.data.list : []
        pagination.total = res.data.count
    } finally {
        table.loading = false
    }
}
getList()

//暴露组件接口给父组件
defineExpose({
    getList,
})
</script>

<template>
    <el-row class="main-table-tool">
        <el-col :span="16">
            <el-space :size="10" style="height: 100%; margin-left: 10px">`
	if option.IsCreate {
		tplView += `
                <el-button v-if="authAction.isCreate" type="primary" @click="handleAdd"> <autoicon-ep-edit-pen />{{ t('common.add') }}</el-button>`
	}
	if option.IsDelete {
		tplView += `
                <el-button v-if="authAction.isDelete" type="danger" @click="handleBatchDelete"> <autoicon-ep-delete-filled />{{ t('common.batchDelete') }}</el-button>`
	}
	tplView += `
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%; margin-right: 10px">
                <my-export-button i18nPrefix="` + tpl.I18nPath + `" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
                <el-dropdown max-height="300" :hide-on-click="false">
                    <el-button type="info" :circle="true"><autoicon-ep-hide /></el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item v-for="(item, index) in table.columns" :key="index">
                                <el-checkbox v-model="item.hidden">{{ item.title }}</el-checkbox>
                            </el-dropdown-item>
                        </el-dropdown-menu>
                    </template>
                </el-dropdown>
            </el-space>
        </el-col>
    </el-row>

    <el-main>
        <el-auto-resizer>
            <template #default="{ height, width }">
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="` + gconv.String(viewList.rowHeight) + `">
                    <template v-if="table.loading" #overlay>
                        <el-icon class="is-loading" color="var(--el-color-primary)" :size="25"><autoicon-ep-loading /></el-icon>
                    </template>
                </el-table-v2>
            </template>
        </el-auto-resizer>
    </el-main>

    <el-row class="main-table-pagination">
        <el-col :span="24">
            <el-pagination
                :total="pagination.total"
                v-model:currentPage="pagination.page"
                v-model:page-size="pagination.size"
                @size-change="pagination.sizeChange"
                @current-change="pagination.pageChange"
                :page-sizes="pagination.sizeList"
                :layout="pagination.layout"
                :background="true"
            />
        </el-col>
    </el-row>
</template>
`

	saveFile := gfile.SelfDir() + `/../view/` + option.SceneCode + `/src/views/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/List.vue`
	gfile.PutContents(saveFile, tplView)
}

func getViewListField(option myGenOption, tpl myGenTpl, v myGenField, i18nPath string) (viewListField myGenViewListField) {
	viewListField.dataKey.Method = internal.ReturnType
	viewListField.dataKey.DataType = `'` + v.FieldRaw + `'`
	viewListField.title.Method = internal.ReturnType
	viewListField.title.DataType = `t('` + i18nPath + `.name.` + v.FieldRaw + `')`
	viewListField.key.Method = internal.ReturnType
	viewListField.key.DataType = `'` + v.FieldRaw + `'`
	viewListField.align.Method = internal.ReturnType
	viewListField.align.DataType = `'center'`
	viewListField.width.Method = internal.ReturnType
	viewListField.width.DataType = `150`

	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
	switch v.FieldType {
	case internal.TypeInt, internal.TypeIntU: // `int等类型` // `int等类型（unsigned）`
		switch v.FieldLimitInt.Size {
		case 2:
			viewListField.width.DataType = `100`
		case 8:
			viewListField.width.DataType = `200`
		}
	case internal.TypeFloat, internal.TypeFloatU: // `float等类型`  // `float等类型（unsigned）`
	case internal.TypeVarchar, internal.TypeChar: // `varchar类型` // `char类型`
		if gconv.Uint(v.FieldLimitStr) >= internal.ConfigMaxLenOfStrHiddle {
			viewListField.width.DataType = `200`
			viewListField.hidden.Method = internal.ReturnType
			viewListField.hidden.DataType = `true`
		}
	case internal.TypeText, internal.TypeJson: // `text类型` // `json类型`
		viewListField.width.DataType = `200`
		viewListField.hidden.Method = internal.ReturnType
		viewListField.hidden.DataType = `true`
	case internal.TypeDatetime, internal.TypeTimestamp: // `datetime类型`	// `timestamp类型`
		viewListField.sortable.Method = internal.ReturnType
		viewListField.sortable.DataType = `true`
	case internal.TypeDate: // `date类型`
		viewListField.width.DataType = `100`
		viewListField.sortable.Method = internal.ReturnType
		viewListField.sortable.DataType = `true`
	case internal.TypeTime: // `time类型`
		viewListField.sortable.Method = internal.ReturnType
		viewListField.sortable.DataType = `true`
	}
	/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 结束--------*/

	/*--------根据字段主键类型处理 开始--------*/
	switch v.FieldTypePrimary {
	case internal.TypePrimary: // 独立主键
		if v.FieldRaw == `id` {
			return myGenViewListField{}
		}
	case internal.TypePrimaryAutoInc: // 独立主键（自增）
		return myGenViewListField{}
	case internal.TypePrimaryMany: // 联合主键
	case internal.TypePrimaryManyAutoInc: // 联合主键（自增）
	}
	/*--------根据字段主键类型处理 结束--------*/

	/*--------根据字段命名类型处理 开始--------*/
	switch v.FieldTypeName {
	case internal.TypeNameDeleted: // 软删除字段
		return myGenViewListField{}
	case internal.TypeNameUpdated: // 更新时间字段
		viewListField.title.Method = internal.ReturnTypeName
		viewListField.title.DataTypeName = `t('common.name.updatedAt')`
	case internal.TypeNameCreated: // 创建时间字段
		viewListField.title.Method = internal.ReturnTypeName
		viewListField.title.DataTypeName = `t('common.name.createdAt')`
	case internal.TypeNamePid: // pid；	类型：int等类型；
		viewListField.dataKey.Method = internal.ReturnTypeName
		viewListField.dataKey.DataTypeName = `'` + internal.GetStrByFieldStyle(tpl.FieldStyle, tpl.Handle.LabelList[0], `p`) + `'`
	case internal.TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
		viewListField.sortable.Method = internal.ReturnTypeName
		viewListField.sortable.DataTypeName = `true`
	case internal.TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
		viewListField.hidden.Method = internal.ReturnTypeName
		viewListField.hidden.DataTypeName = `true`
	case internal.TypeNamePasswordSuffix: // password,passwd后缀；	类型：char(32)；
		return myGenViewListField{}
	case internal.TypeNameSaltSuffix: // salt后缀，且对应的password,passwd后缀存在时（才）有效；	类型：char；
		return myGenViewListField{}
	case internal.TypeNameNameSuffix: // name,title后缀；	类型：varchar；
		if option.IsUpdate {
			viewListField.cellRenderer.Method = internal.ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!authAction.isUpdate) {
                    return [<div class="el-table-v2__cell-text">{props.rowData.` + v.FieldRaw + `}</div>]
                }
                if (!props.rowData?.edit` + gstr.CaseCamel(v.FieldRaw) + `?.isEdit) {
                    return [
                        <div class="el-table-v2__cell-text inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = { isEdit: true, oldValue: props.rowData.` + v.FieldRaw + ` })}>
                            {props.rowData.` + v.FieldRaw + `}
                        </div>,
                    ]
                }
                let currentRef: any
                return [
                    <el-input
                        ref={(el: any) => {
                            el?.focus()
                            currentRef = el
                        }}
                        v-model={props.rowData.` + v.FieldRaw + `}
                        placeholder={t('` + i18nPath + `.name.` + v.FieldRaw + `')}
                        maxlength={` + v.FieldLimitStr + `}
                        show-word-limit={true}
                        onBlur={() => {
                            props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.isEdit = false
                            if (props.rowData.` + v.FieldRaw + ` == props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.oldValue) {
                                return
                            }
                            if (!props.rowData.` + v.FieldRaw + `) {
                                props.rowData.` + v.FieldRaw + ` = props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.oldValue
                                return
                            }
                            handleUpdate({ ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `: [props.rowData.id], ` + v.FieldRaw + `: props.rowData.` + v.FieldRaw + ` }).catch(() => (props.rowData.` + v.FieldRaw + ` = props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.oldValue))
                        }}
                        onKeydown={(event: any) => {
                            switch (event.keyCode) {
                                case 13: //13：Enter键 27：Esc键 32：空格键
                                    currentRef?.blur()
                                    break
                            }
                        }}
                    />,
                ]
            }`
		}
	case internal.TypeNameCodeSuffix: // code后缀；	类型：varchar；
	case internal.TypeNameAccountSuffix: // account后缀；	类型：varchar；
	case internal.TypeNamePhoneSuffix: // phone,mobile后缀；	类型：varchar；
	case internal.TypeNameEmailSuffix: // email后缀；	类型：varchar；
	case internal.TypeNameUrlSuffix: // url,link后缀；	类型：varchar；
	case internal.TypeNameIpSuffix: // IP后缀；	类型：varchar；
	case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
		viewListField.width.Method = internal.ReturnTypeName
		viewListField.width.DataTypeName = `100`
		cellRendererStr := `disabled={true}`
		if option.IsUpdate {
			cellRendererStr = `disabled={!authAction.isUpdate}
                        onChange={(val: string) => {
                            if (val != props.rowData.color) {
                                handleUpdate({  ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `: [props.rowData.id], ` + v.FieldRaw + `: val ? val : '' }).then(() => (props.rowData.` + v.FieldRaw + ` = val))
                            }
                        }}`
		}
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                return [
                    <el-color-picker
                        model-value={props.rowData.` + v.FieldRaw + `}
                        show-alpha={true}
                        ` + cellRendererStr + `
                    />,
                ]
            }`
	case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
		relIdObj := tpl.Handle.RelIdMap[v.FieldRaw]
		if relIdObj.tpl.Table != `` && !relIdObj.IsRedundName {
			viewListField.dataKey.Method = internal.ReturnTypeName
			viewListField.dataKey.DataTypeName = `'` + relIdObj.tpl.Handle.LabelList[0] + relIdObj.Suffix + `'`
		}
	case internal.TypeNameSortSuffix, internal.TypeNameNoSuffix: // sort,num,number,weight等后缀；	类型：int等类型；	// no,level,rank等后缀；	类型：int等类型；
		viewListField.sortable.Method = internal.ReturnTypeName
		viewListField.sortable.DataTypeName = `true`
		if option.IsUpdate {
			viewListField.cellRenderer.Method = internal.ReturnTypeName
			attrOfAdd := `placeholder={t('` + i18nPath + `.name.` + v.FieldRaw + `')}`
			if v.FieldTip != `` {
				attrOfAdd = `placeholder={t('` + i18nPath + `.tip.` + v.FieldRaw + `')}`
			}
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!authAction.isUpdate) {
                    return [<div class="el-table-v2__cell-text">{props.rowData.` + v.FieldRaw + `}</div>]
                }
                if (!props.rowData?.edit` + gstr.CaseCamel(v.FieldRaw) + `?.isEdit) {
                    return [
                        <div class="el-table-v2__cell-text inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = { isEdit: true, oldValue: props.rowData.` + v.FieldRaw + ` })}>
                            {props.rowData.` + v.FieldRaw + `}
                        </div>,
                    ]
                }
                let currentRef: any
                return [
                    <el-input-number
                        ref={(el: any) => {
                            el?.focus()
                            currentRef = el
                        }}
                        v-model={props.rowData.` + v.FieldRaw + `}
                        ` + attrOfAdd + `
                        precision={0}
                        min={` + v.FieldLimitInt.Min + `}
                        max={` + v.FieldLimitInt.Max + `}
                        controls={false}
                        onBlur={() => {
                            props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.isEdit = false
                            if (props.rowData.` + v.FieldRaw + ` == props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.oldValue) {
                                return
                            }
                            if (!(props.rowData.` + v.FieldRaw + ` || props.rowData.` + v.FieldRaw + ` === 0)) {
                                props.rowData.` + v.FieldRaw + ` = props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.oldValue
                                return
                            }
                            handleUpdate({ ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `: [props.rowData.id], ` + v.FieldRaw + `: props.rowData.` + v.FieldRaw + ` }).catch(() => (props.rowData.` + v.FieldRaw + ` = props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `.oldValue))
                        }}
                        onKeydown={(event: any) => {
                            switch (event.keyCode) {
                                case 13: //13：Enter键 27：Esc键 32：空格键
                                    currentRef?.blur()
                                    break
                            }
                        }}
                    />,
                ]
            }`
		}
	case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
		viewListField.isI18nTm = true
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('` + i18nPath + `.status.` + v.FieldRaw + `') as { value: any, label: string }[]
                let index = obj.findIndex((item) => { return item.value == props.rowData.` + v.FieldRaw + ` })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            }`
	case internal.TypeNameIsPrefix: // is_前缀；	类型：int等类型；注释：多状态之间用[\s,，.。;；]等字符分隔。示例（停用：0否 1是）
		viewListField.width.Method = internal.ReturnTypeName
		viewListField.width.DataTypeName = `100`
		cellRendererStr := `disabled={true}`
		if option.IsUpdate {
			cellRendererStr = `disabled={!authAction.isUpdate}
                        onChange={(val: number) => handleUpdate({ ` + internal.GetStrByFieldStyle(tpl.FieldStyle, `id_arr`) + `: [props.rowData.id], ` + v.FieldRaw + `: val }).then(() => (props.rowData.` + v.FieldRaw + ` = val))}`
		}
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.` + v.FieldRaw + `}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        ` + cellRendererStr + `
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                    />,
                ]
            }`
	case internal.TypeNameStartPrefix: // start_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameEndPrefix: // end_前缀；	类型：datetime或date或timestamp或time；
	case internal.TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
	case internal.TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
		viewListField.width.Method = internal.ReturnTypeName
		viewListField.width.DataTypeName = `100`
		viewListField.hidden.Method = internal.ReturnEmpty
		cellRendererStr := `
                const imageList = [props.rowData.` + v.FieldRaw + `]`
		if v.FieldType != internal.TypeVarchar {
			cellRendererStr = `
                let imageList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    imageList = props.rowData.` + v.FieldRaw + `
                } else {
                    imageList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
		}
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererStr + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {imageList.map((item) => {
                                return <el-image src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>
                ]
            }`
	case internal.TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text
		viewListField.rowHeight = 100
		viewListField.hidden.Method = internal.ReturnEmpty
		cellRendererStr := `
                const videoList = [props.rowData.` + v.FieldRaw + `]`
		if v.FieldType != internal.TypeVarchar {
			cellRendererStr = `
                let videoList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    videoList = props.rowData.` + v.FieldRaw + `
                } else {
                    videoList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
		}
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererStr + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {videoList.map((item) => {
                                return <video preload="none" controls={true} src={item} style="width: 100%;" /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
	case internal.TypeNameAudioSuffix: // audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text
		viewListField.hidden.Method = internal.ReturnEmpty
		cellRendererStr := `
                const audioList = [props.rowData.` + v.FieldRaw + `]`
		if v.FieldType != internal.TypeVarchar {
			cellRendererStr = `
                let audioList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    audioList = props.rowData.` + v.FieldRaw + `
                } else {
                    audioList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
		}
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererStr + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {audioList.map((item) => {
                                return <audio preload="none" controls={true} src={item} style="height: 30px;" /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
	case internal.TypeNameFileSuffix: // file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
	case internal.TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
		viewListField.isI18nTm = true
		viewListField.hidden.Method = internal.ReturnEmpty
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }
                let arrList: any[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    arrList = props.rowData.` + v.FieldRaw + `
                } else {
                    arrList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }
                let tagType = tm('config.const.tagType') as string[]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {arrList.map((item, index) => {
                                return <el-tag type={tagType[index % tagType.length]}>{item}</el-tag>
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
	}
	/*--------根据字段命名类型处理 结束--------*/
	return
}

func getViewListExtendMiddleOne(option myGenOption, tplEM handleExtendMiddle) (viewList myGenViewList) {
	switch tplEM.TableType {
	case internal.TableTypeExtendOne:
		for _, v := range tplEM.FieldList {
			viewList.Add(getViewListField(option, tplEM.tpl, v, tplEM.tplOfTop.I18nPath))
		}
	case internal.TableTypeMiddleOne:
		for _, v := range tplEM.FieldListOfIdSuffix {
			viewList.Add(getViewListField(option, tplEM.tpl, v, tplEM.tplOfTop.I18nPath))
		}
		for _, v := range tplEM.FieldListOfOther {
			viewList.Add(getViewListField(option, tplEM.tpl, v, tplEM.tplOfTop.I18nPath))
		}
	}
	return
}

func getViewListExtendMiddleMany(option myGenOption, tplEM handleExtendMiddle) (viewList myGenViewList) {
	if len(tplEM.FieldList) == 1 {
		v := tplEM.FieldList[0]

		isReturn := false
		viewListField := myGenViewListField{}
		viewListField.dataKey.Method = internal.ReturnType
		viewListField.dataKey.DataType = `'` + tplEM.FieldVar + `'`
		viewListField.title.Method = internal.ReturnType
		viewListField.title.DataType = `t('` + tplEM.tplOfTop.I18nPath + `.name.` + tplEM.FieldVar + `')`
		viewListField.key.Method = internal.ReturnType
		viewListField.key.DataType = `'` + tplEM.FieldVar + `'`
		viewListField.align.Method = internal.ReturnType
		viewListField.align.DataType = `'center'`
		viewListField.width.Method = internal.ReturnType
		viewListField.width.DataType = `150`
		/*--------部分命名类型直接处理后返回 开始--------*/
		switch v.FieldTypeName {
		case internal.TypeNameColorSuffix: // color后缀；	类型：varchar；
			isReturn = true
			viewListField.width.Method = internal.ReturnTypeName
			viewListField.width.DataTypeName = `100`
			viewListField.hidden.Method = internal.ReturnEmpty
			viewListField.cellRenderer.Method = internal.ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + tplEM.FieldVar + `) {
                    return
                }
                let arrList: string[] = props.rowData.` + tplEM.FieldVar + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {arrList.map((item, index) => {
                                return <el-color-picker model-value={item} show-alpha={true} disabled={true} />
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
		case internal.TypeNameStatusSuffix: // status,type,scene,method,pos,position,gender,currency等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，.。;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			return myGenViewList{}
		case internal.TypeNameIdSuffix: // id后缀；	类型：int等类型；
			return myGenViewList{}
		case internal.TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			if v.FieldType != internal.TypeVarchar {
				return myGenViewList{}
			}
			isReturn = true
			viewListField.width.Method = internal.ReturnTypeName
			viewListField.width.DataTypeName = `100`
			viewListField.hidden.Method = internal.ReturnEmpty
			viewListField.cellRenderer.Method = internal.ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + tplEM.FieldVar + `) {
                    return
                }
                let imageList: string[] = props.rowData.` + tplEM.FieldVar + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {imageList.map((item) => {
                                return <el-image src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>
                ]
            }`
		case internal.TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；	类型：单视频varchar，多视频json或text
			if v.FieldType != internal.TypeVarchar {
				return myGenViewList{}
			}
			isReturn = true
			viewListField.rowHeight = 100
			viewListField.hidden.Method = internal.ReturnEmpty
			viewListField.cellRenderer.Method = internal.ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + tplEM.FieldVar + `) {
                    return
                }
                let videoList: string[] = props.rowData.` + tplEM.FieldVar + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {videoList.map((item) => {
                                return <video preload="none" controls={true} src={item} style="width: 100%;" /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
		case internal.TypeNameAudioSuffix: // audio,audio_list,audioList,audio_arr,audioArr等后缀；	类型：单音频varchar，多音频json或text
			if v.FieldType != internal.TypeVarchar {
				return myGenViewList{}
			}
			isReturn = true
			viewListField.hidden.Method = internal.ReturnEmpty
			viewListField.cellRenderer.Method = internal.ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + tplEM.FieldVar + `) {
                    return
                }
                let videoList: string[] = props.rowData.` + tplEM.FieldVar + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {audioList.map((item) => {
                                return <audio preload="none" controls={true} src={item} style="height: 30px;" /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
			// case internal.TypeNameFileSuffix: // file,file_list,fileList,file_arr,fileArr等后缀；	类型：单文件varchar，多文件json或text
		}
		if isReturn {
			viewList.Add(viewListField)
			return
		}
		/*--------部分命名类型直接处理后返回 结束--------*/
		viewListField.isI18nTm = true
		viewListField.hidden.Method = internal.ReturnEmpty
		viewListField.cellRenderer.Method = internal.ReturnTypeName
		viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + tplEM.FieldVar + `) {
                    return
                }
                let arrList: any[] = props.rowData.` + tplEM.FieldVar + `
                let tagType = tm('config.const.tagType') as string[]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {arrList.map((item, index) => {
                                return <el-tag type={tagType[index % tagType.length]}>{item}</el-tag>
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            }`
		viewList.Add(viewListField)
	}
	return
}
