package my_gen

import (
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type myGenViewList struct {
	rowHeight uint
	idType    string
	columns   []string
}

type myGenViewListField struct {
	dataKey      myGenDataStrHandler
	title        myGenDataStrHandler
	key          myGenDataStrHandler
	align        myGenDataStrHandler
	width        myGenDataStrHandler
	sortable     myGenDataStrHandler
	hidden       myGenDataStrHandler
	cellRenderer myGenDataStrHandler
}

func (viewListThis *myGenViewList) Add(viewListField myGenViewListField) {
	columnAttrStr := []string{
		`dataKey: ` + viewListField.dataKey.getData() + `,`,
		`title: ` + viewListField.title.getData() + `,`,
		`key: ` + viewListField.key.getData() + `,`,
		`align: ` + viewListField.align.getData() + `,`,
		`width: ` + viewListField.width.getData() + `,`,
	}
	if viewListField.sortable.getData() != `` {
		columnAttrStr = append(columnAttrStr, `sortable: `+viewListField.sortable.getData()+`,`)
	}
	if viewListField.hidden.getData() != `` {
		columnAttrStr = append(columnAttrStr, `hidden: `+viewListField.hidden.getData()+`,`)
	}
	if viewListField.cellRenderer.getData() != `` {
		columnAttrStr = append(columnAttrStr, `cellRenderer: `+viewListField.cellRenderer.getData()+`,`)
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
	viewListThis.columns = garray.NewStrArrayFrom(viewListThis.columns).Unique().Slice()
}

// 视图模板List生成
func genViewList(option myGenOption, tpl myGenTpl) {
	viewList := myGenViewList{
		rowHeight: 50,
		idType:    `number`,
	}
	if len(tpl.Handle.Id.List) > 1 || !garray.NewIntArrayFrom([]int{TypeInt, TypeIntU}).Contains(tpl.Handle.Id.List[0].FieldType) {
		viewList.idType = `string`
	}
	viewList.Merge(getViewListFieldList(option, tpl))
	viewList.Unique()

	tplView := `<script setup lang="tsx">
const { t, tm } = useI18n()

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
            width: 250,
            fixed: 'right',
            cellRenderer: (props: any): any => {
                return [`
		if option.IsUpdate {
			tplView += `
                    <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                        <autoicon-ep-edit />
                        {t('common.edit')}
                    </el-button>,`
		}
		if option.IsDelete {
			tplView += `
                    <el-button type="danger" size="small" onClick={() => handleDelete(props.rowData.id)}>
                        <autoicon-ep-delete />
                        {t('common.delete')}
                    </el-button>,`
		}
		if option.IsCreate {
			tplView += `
                    <el-button type="warning" size="small" onClick={() => handleEditCopy(props.rowData.id, 'copy')}>
                        <autoicon-ep-document-copy />
                        {t('common.copy')}
                    </el-button>,`
		}
		tplView += `
                ]
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
    table.data.forEach((item: any) => {
        if (item.checked) {
            idArr.push(item.id)
        }
    })
    if (idArr.length) {
        handleDelete(idArr)
    } else {
        ElMessage.error(t('common.tip.selectDelete'))
    }
}`
	}
	if option.IsCreate || option.IsUpdate {
		tplView += `
//编辑|复制
const handleEditCopy = (id: ` + viewList.idType + `, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/info', { id: id })
        .then((res) => {
            saveCommon.data = { ...res.data.info }
            switch (type) {
                case 'edit':
                    saveCommon.data.idArr = [saveCommon.data.id]
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
        .catch(() => {})
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
    })
        .then(() => {
            request(t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/del', { idArr: idArr }, true)
                .then((res) => {
                    getList()
                })
                .catch(() => {})
        })
        .catch(() => {})
}`
	}
	if option.IsUpdate {
		tplView += `
//更新
const handleUpdate = async (param: { idArr: ` + viewList.idType + `[]; [propName: string]: any }) => {
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
    sizeChange: (val: number) => {
        getList()
    },
    pageChange: (val: number) => {
        getList()
    },
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
    if (resetPage) {
        pagination.page = 1
    }
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
    } catch (error) {}
    table.loading = false
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
                <el-button type="primary" @click="handleAdd"> <autoicon-ep-edit-pen />{{ t('common.add') }} </el-button>`
	}
	if option.IsDelete {
		tplView += `
                <el-button type="danger" @click="handleBatchDelete"> <autoicon-ep-delete-filled />{{ t('common.batchDelete') }} </el-button>`
	}
	tplView += `
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%">
                <my-export-button i18nPrefix="` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/` + tpl.ModuleDirCaseKebab + `/` + tpl.TableCaseKebab + `/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
                <el-dropdown max-height="300" :hide-on-click="false">
                    <el-button type="info" :circle="true">
                        <autoicon-ep-hide />
                    </el-button>
                    <template #dropdown>
                        <el-dropdown-menu>
                            <el-dropdown-item v-for="(item, index) in table.columns" :key="index">
                                <el-checkbox v-model="item.hidden">
                                    {{ item.title }}
                                </el-checkbox>
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
                        <el-icon class="is-loading" color="var(--el-color-primary)" :size="25">
                            <autoicon-ep-loading />
                        </el-icon>
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

func getViewListFieldList(option myGenOption, tpl myGenTpl) (viewList myGenViewList) {
	for _, v := range tpl.FieldList {
		viewListField := myGenViewListField{}
		viewListField.dataKey.Method = ReturnType
		viewListField.dataKey.DataType = `'` + v.FieldRaw + `'`
		viewListField.title.Method = ReturnType
		viewListField.title.DataType = `t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.name.` + v.FieldRaw + `')`
		viewListField.key.Method = ReturnType
		viewListField.key.DataType = `'` + v.FieldRaw + `'`
		viewListField.align.Method = ReturnType
		viewListField.align.DataType = `'center'`
		viewListField.width.Method = ReturnType
		viewListField.width.DataType = `150`

		/*--------根据字段数据类型处理（注意：这里的代码改动对字段命名类型处理有影响） 开始--------*/
		switch v.FieldType {
		case TypeInt, TypeIntU: // `int等类型` // `int等类型（unsigned）`
			if gstr.Pos(v.FieldTypeRaw, `tinyint`) != -1 || gstr.Pos(v.FieldTypeRaw, `smallint`) != -1 {
				viewListField.width.DataType = `100`
			}
		case TypeFloat: // `float等类型`
		case TypeFloatU: // `float等类型（unsigned）`
		case TypeVarchar, TypeChar: // `varchar类型` // `char类型`
			if gconv.Uint(v.FieldLimitStr) >= 120 {
				viewListField.width.DataType = `200`
				viewListField.hidden.Method = ReturnType
				viewListField.hidden.DataType = `true`
			}
		case TypeText, TypeJson: // `text类型` // `json类型`
			viewListField.width.DataType = `200`
			viewListField.hidden.Method = ReturnType
			viewListField.hidden.DataType = `true`
		case TypeTimestamp, TypeDatetime: // `timestamp类型` // `datetime类型`
			viewListField.sortable.Method = ReturnType
			viewListField.sortable.DataType = `true`
		case TypeDate: // `date类型`
			viewListField.width.DataType = `100`
			viewListField.sortable.Method = ReturnType
			viewListField.sortable.DataType = `true`
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
			viewListField.title.Method = ReturnTypeName
			viewListField.title.DataTypeName = `t('common.name.updatedAt')`
		case TypeNameCreated: // 创建时间字段
			viewListField.title.Method = ReturnTypeName
			viewListField.title.DataTypeName = `t('common.name.createdAt')`
		case TypeNamePid: // pid；	类型：int等类型；
			if len(tpl.Handle.LabelList) > 0 {
				viewListField.dataKey.Method = ReturnTypeName
				viewListField.dataKey.DataTypeName = `'p` + gstr.CaseCamel(tpl.Handle.LabelList[0]) + `'`
			}
		case TypeNameLevel: // level，且pid,level,idPath|id_path同时存在时（才）有效；	类型：int等类型；
			viewListField.sortable.Method = ReturnTypeName
			viewListField.sortable.DataTypeName = `true`
		case TypeNameIdPath: // idPath|id_path，且pid,level,idPath|id_path同时存在时（才）有效；	类型：varchar或text；
			viewListField.hidden.Method = ReturnTypeName
			viewListField.hidden.DataTypeName = `true`
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
			if tpl.Handle.RelIdMap[v.FieldRaw].tpl.Table != `` && !tpl.Handle.RelIdMap[v.FieldRaw].IsRedundName {
				viewListField.dataKey.Method = ReturnTypeName
				viewListField.dataKey.DataTypeName = `'` + tpl.Handle.RelIdMap[v.FieldRaw].tpl.Handle.LabelList[0] + tpl.Handle.RelIdMap[v.FieldRaw].Suffix + `'`
			}
		case TypeNameSortSuffix, TypeNameSort: // sort,weight等后缀；	类型：int等类型； // sort，且pid,level,idPath|id_path,sort同时存在时（才）有效；	类型：int等类型；
			viewListField.sortable.Method = ReturnTypeName
			viewListField.sortable.DataTypeName = `true`
			if option.IsUpdate {
				viewListField.cellRenderer.Method = ReturnTypeName
				viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + `) {
                    let currentRef: any
                    let currentVal = props.rowData.` + v.FieldRaw + `
                    return [
                        <el-input-number
                            ref={(el: any) => {
                                currentRef = el
                                el?.focus()
                            }}
                            model-value={currentVal}
                            placeholder={t('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.tip.` + v.FieldRaw + `')}
                            precision={0}
                            min={0}
                            max={100}
                            step={1}
                            step-strictly={true}
                            controls={false} //控制按钮会导致诸多问题。如：焦点丢失；` + v.FieldRaw + `是0或100时，只一个按钮可点击
                            controls-position="right"
                            onChange={(val: number) => (currentVal = val)}
                            onBlur={() => {
                                props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = false
                                if ((currentVal || currentVal === 0) && currentVal != props.rowData.` + v.FieldRaw + `) {
                                    handleUpdate({
                                        idArr: [props.rowData.id],
                                        ` + v.FieldRaw + `: currentVal,
                                    })
                                        .then((res) => {
                                            props.rowData.` + v.FieldRaw + ` = currentVal
                                        })
                                        .catch((error) => {})
                                }
                            }}
                            onKeydown={(event: any) => {
                                switch (event.keyCode) {
                                    // case 27:    //Esc键：Escape
                                    // case 32:    //空格键：" "
                                    case 13: //Enter键：Enter
                                        // props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = false    //也会触发onBlur事件
                                        currentRef?.blur()
                                        break
                                }
                            }}
                        />,
                    ]
                }
                return [
                    <div class="inline-edit" onClick={() => (props.rowData.edit` + gstr.CaseCamel(v.FieldRaw) + ` = true)}>
                        {props.rowData.` + v.FieldRaw + `}
                    </div>,
                ]
            }`
			}
		case TypeNameStatusSuffix: // status,type,method,pos,position,gender等后缀；	类型：int等类型或varchar或char；	注释：多状态之间用[\s,，;；]等字符分隔。示例（状态：0待处理 1已处理 2驳回 yes是 no否）
			viewListField.cellRenderer.Method = ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('` + tpl.ModuleDirCaseKebabReplace + `.` + tpl.TableCaseKebab + `.status.` + v.FieldRaw + `') as { value: any, label: string }[]
                let index = obj.findIndex((item) => { return item.value == props.rowData.` + v.FieldRaw + ` })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            }`
		case TypeNameIsPrefix: // is_前缀；		类型：int等类型；注释：多状态之间用[\s,，;；]等字符分隔。示例（停用：0否 1是）
			cellRendererStr := `disabled={true}`
			if option.IsUpdate {
				cellRendererStr = `onChange={(val: number) => {
                            handleUpdate({
                                idArr: [props.rowData.id],
                                ` + v.FieldRaw + `: val,
                            })
                                .then((res) => {
                                    props.rowData.` + v.FieldRaw + ` = val
                                })
                                .catch((error) => {})
                        }}`
			}
			viewListField.cellRenderer.Method = ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.` + v.FieldRaw + `}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                        ` + cellRendererStr + `
                    />,
                ]
            }`
		case TypeNameStartPrefix: // start_前缀；	类型：timestamp或datetime或date；
		case TypeNameEndPrefix: // end_前缀；	类型：timestamp或datetime或date；
		case TypeNameRemarkSuffix: // remark,desc,msg,message,intro,content后缀；	类型：varchar或text；前端对应组件：varchar文本输入框，text富文本编辑器
		case TypeNameImageSuffix: // icon,cover,avatar,img,img_list,imgList,img_arr,imgArr,image,image_list,imageList,image_arr,imageArr等后缀；	类型：单图片varchar，多图片json或text
			viewListField.width.Method = ReturnTypeName
			viewListField.width.DataTypeName = `100`
			viewListField.hidden.Method = ReturnEmpty
			cellRendererStr := `
                const imageList = [props.rowData.` + v.FieldRaw + `]`
			if v.FieldType != TypeVarchar {
				cellRendererStr = `
                let imageList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    imageList = props.rowData.` + v.FieldRaw + `
                } else {
                    imageList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
			}
			viewListField.cellRenderer.Method = ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererStr + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {imageList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <el-image style="width: 45px;" src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} />
                        })}
                    </el-scrollbar>
                ]
            }`
		case TypeNameVideoSuffix: // video,video_list,videoList,video_arr,videoArr等后缀；		类型：单视频varchar，多视频json或text
			if viewList.rowHeight < 100 {
				viewList.rowHeight = 100
			}
			viewListField.hidden.Method = ReturnEmpty
			cellRendererStr := `
                const videoList = [props.rowData.` + v.FieldRaw + `]`
			if v.FieldType != TypeVarchar {
				cellRendererStr = `
                let videoList: string[]
                if (Array.isArray(props.rowData.` + v.FieldRaw + `)) {
                    videoList = props.rowData.` + v.FieldRaw + `
                } else {
                    videoList = JSON.parse(props.rowData.` + v.FieldRaw + `)
                }`
			}
			viewListField.cellRenderer.Method = ReturnTypeName
			viewListField.cellRenderer.DataTypeName = `(props: any): any => {
                if (!props.rowData.` + v.FieldRaw + `) {
                    return
                }` + cellRendererStr + `
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        {videoList.map((item) => {
                            //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            return <video style="width: 120px; height: 80px;" preload="none" controls={true} src={item} />
                        })}
                    </el-scrollbar>,
                ]
            }`
		case TypeNameArrSuffix: // list,arr等后缀；	类型：json或text；
			viewListField.width.Method = ReturnTypeName
			viewListField.width.DataTypeName = `100`
			viewListField.hidden.Method = ReturnEmpty
			viewListField.cellRenderer.Method = ReturnTypeName
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
                        {arrList.map((item, index) => {
                            return [
                                <el-tag style="margin: auto 5px 5px auto;" type={tagType[index % tagType.length]}>
                                    {item}
                                </el-tag>,
                            ]
                        })}
                    </el-scrollbar>,
                ]
            }`
		}
		/*--------根据字段命名类型处理 结束--------*/

		viewList.Add(viewListField)
	}
	return
}
