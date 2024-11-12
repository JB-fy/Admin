<script setup lang="tsx">
import type { Action, MessageBoxState } from 'element-plus'
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const table = reactive({
    columns: [
        {
            dataKey: 'id',
            title: t('common.name.id'),
            key: 'id',
            align: 'center',
            width: 200,
            fixed: 'left',
            sortable: true,
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
            },
        },
        {
            dataKey: 'name_type',
            title: t('app.app.name.name_type'),
            key: 'name_type',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('app.app.status.name_type') as { value: any; label: string }[]
                let index = obj.findIndex((item) => {
                    return item.value == props.rowData.name_type
                })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            },
        },
        {
            dataKey: 'app_type',
            title: t('app.app.name.app_type'),
            key: 'app_type',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let obj = tm('app.app.status.app_type') as { value: any; label: string }[]
                let index = obj.findIndex((item) => {
                    return item.value == props.rowData.app_type
                })
                return <el-tag type={tagType[index % tagType.length]}>{obj[index]?.label}</el-tag>
            },
        },
        {
            dataKey: 'package_name',
            title: t('app.app.name.package_name'),
            key: 'package_name',
            align: 'center',
            width: 150,
            cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [<div class="el-table-v2__cell-text">{props.rowData.package_name}</div>]
                }
                if (!props.rowData?.editPackageName?.isEdit) {
                    return [
                        <div class="el-table-v2__cell-text inline-edit" onClick={() => (props.rowData.editPackageName = { isEdit: true, oldValue: props.rowData.package_name })}>
                            {props.rowData.package_name}
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
                        v-model={props.rowData.package_name}
                        placeholder={t('app.app.name.package_name')}
                        maxlength={60}
                        show-word-limit={true}
                        onBlur={() => {
                            props.rowData.editPackageName.isEdit = false
                            if (props.rowData.package_name == props.rowData.editPackageName.oldValue) {
                                return
                            }
                            if (!props.rowData.package_name) {
                                props.rowData.package_name = props.rowData.editPackageName.oldValue
                                return
                            }
                            handleUpdate(props.rowData.id, { package_name: props.rowData.package_name }).catch(() => (props.rowData.package_name = props.rowData.editPackageName.oldValue))
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
            },
        },
        {
            dataKey: 'package_file',
            title: t('app.app.name.package_file'),
            key: 'package_file',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                if (!props.rowData.package_file) {
                    return
                }
                const fileList = [props.rowData.package_file]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {fileList.map((item) => {
                                return <my-upload v-model={item} size="small" disabled={true} /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            },
        },
        {
            dataKey: 'ver_no',
            title: t('app.app.name.ver_no'),
            key: 'ver_no',
            align: 'center',
            width: 150,
            sortable: true,
            cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [<div class="el-table-v2__cell-text">{props.rowData.ver_no}</div>]
                }
                if (!props.rowData?.editVerNo?.isEdit) {
                    return [
                        <div class="el-table-v2__cell-text inline-edit" onClick={() => (props.rowData.editVerNo = { isEdit: true, oldValue: props.rowData.ver_no })}>
                            {props.rowData.ver_no}
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
                        v-model={props.rowData.ver_no}
                        placeholder={t('app.app.name.ver_no')}
                        min={0}
                        max={4294967295}
                        precision={0}
                        controls={false}
                        onBlur={() => {
                            props.rowData.editVerNo.isEdit = false
                            if (props.rowData.ver_no == props.rowData.editVerNo.oldValue) {
                                return
                            }
                            if (!(props.rowData.ver_no || props.rowData.ver_no === 0)) {
                                props.rowData.ver_no = props.rowData.editVerNo.oldValue
                                return
                            }
                            handleUpdate(props.rowData.id, { ver_no: props.rowData.ver_no }).catch(() => (props.rowData.ver_no = props.rowData.editVerNo.oldValue))
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
            },
        },
        {
            dataKey: 'ver_name',
            title: t('app.app.name.ver_name'),
            key: 'ver_name',
            align: 'center',
            width: 150,
            cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [<div class="el-table-v2__cell-text">{props.rowData.ver_name}</div>]
                }
                if (!props.rowData?.editVerName?.isEdit) {
                    return [
                        <div class="el-table-v2__cell-text inline-edit" onClick={() => (props.rowData.editVerName = { isEdit: true, oldValue: props.rowData.ver_name })}>
                            {props.rowData.ver_name}
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
                        v-model={props.rowData.ver_name}
                        placeholder={t('app.app.name.ver_name')}
                        maxlength={30}
                        show-word-limit={true}
                        onBlur={() => {
                            props.rowData.editVerName.isEdit = false
                            if (props.rowData.ver_name == props.rowData.editVerName.oldValue) {
                                return
                            }
                            if (!props.rowData.ver_name) {
                                props.rowData.ver_name = props.rowData.editVerName.oldValue
                                return
                            }
                            handleUpdate(props.rowData.id, { ver_name: props.rowData.ver_name }).catch(() => (props.rowData.ver_name = props.rowData.editVerName.oldValue))
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
            },
        },
        {
            dataKey: 'ver_intro',
            title: t('app.app.name.ver_intro'),
            key: 'ver_intro',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'extra_config',
            title: t('app.app.name.extra_config'),
            key: 'extra_config',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'remark',
            title: t('app.app.name.remark'),
            key: 'remark',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'is_force_prev',
            title: t('app.app.name.is_force_prev'),
            key: 'is_force_prev',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.is_force_prev}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        disabled={!authAction.isUpdate}
                        onChange={(val: number) => handleUpdate(props.rowData.id, { is_force_prev: val }).then(() => (props.rowData.is_force_prev = val))}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                    />,
                ]
            },
        },
        {
            dataKey: 'is_stop',
            title: t('app.app.name.is_stop'),
            key: 'is_stop',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.is_stop}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        disabled={!authAction.isUpdate}
                        onChange={(val: number) => handleUpdate(props.rowData.id, { is_stop: val }).then(() => (props.rowData.is_stop = val))}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                    />,
                ]
            },
        },
        {
            dataKey: 'updated_at',
            title: t('common.name.updatedAt'),
            key: 'updated_at',
            align: 'center',
            width: 150,
            sortable: true,
        },
        {
            dataKey: 'created_at',
            title: t('common.name.createdAt'),
            key: 'created_at',
            align: 'center',
            width: 150,
            sortable: true,
        },
        {
            title: t('common.name.action'),
            key: 'action',
            align: 'center',
            width: 80 * ((authAction.isCreate ? 1 : 0) + (authAction.isUpdate ? 1 : 0) + (authAction.isDelete ? 1 : 0)),
            fixed: 'right',
            hidden: !(authAction.isCreate || authAction.isUpdate || authAction.isDelete),
            cellRenderer: (props: any): any => {
                let vNode: any = []
                if (authAction.isUpdate) {
                    vNode.push(
                        <el-button type="primary" size="small" onClick={() => handleEditCopy(props.rowData.id)}>
                            <autoicon-ep-edit />
                            {t('common.edit')}
                        </el-button>
                    )
                }
                if (authAction.isDelete) {
                    vNode.push(
                        <el-button type="danger" size="small" onClick={() => handleDelete(props.rowData.id)}>
                            <autoicon-ep-delete />
                            {t('common.delete')}
                        </el-button>
                    )
                }
                if (authAction.isCreate) {
                    vNode.push(
                        <el-button type="warning" size="small" onClick={() => handleEditCopy(props.rowData.id, 'copy')}>
                            <autoicon-ep-document-copy />
                            {t('common.copy')}
                        </el-button>
                    )
                }
                return vNode
            },
        },
    ] as any,
    data: [],
    loading: false,
    sort: { key: 'id', order: 'desc' } as any,
    handleSort: (sort: any) => {
        table.sort.key = sort.key
        table.sort.order = sort.order
        getList()
    },
})

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
//新增
const handleAdd = () => {
    saveCommon.data = {}
    saveCommon.title = t('common.add')
    saveCommon.visible = true
}
//批量删除
const handleBatchDelete = () => {
    const idArr: number[] = []
    table.data.forEach((item: any) => item.checked && idArr.push(item.id))
    idArr.length == 0 ? ElMessage.error(t('common.tip.selectDelete')) : handleDelete(idArr)
}
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/app/app/info', { id: id }).then((res) => {
        saveCommon.data = { ...res.data.info }
        switch (type) {
            case 'edit':
                saveCommon.title = t('common.edit')
                break
            case 'copy':
                delete saveCommon.data.id
                saveCommon.title = t('common.copy')
                break
        }
        saveCommon.visible = true
    })
}
//删除
const handleDelete = (id: number | number[]) => {
    ElMessageBox.confirm('', {
        type: 'warning',
        title: t('common.tip.configDelete'),
        center: true,
        showClose: false,
        beforeClose: (action: Action, instance: MessageBoxState, done: Function) => {
            switch (action) {
                case 'confirm':
                    instance.confirmButtonLoading = true
                    request(t('config.VITE_HTTP_API_PREFIX') + '/app/app/del', { [Array.isArray(id) ? 'id_arr' : 'id']: id }, true)
                        .then(() => {
                            getList()
                            done()
                        })
                        .finally(() => (instance.confirmButtonLoading = false))
                    break
                default:
                    done()
                    break
            }
        },
    })
}
//更新
const handleUpdate = async (id: number | number[], param: { [propName: string]: any }) => {
    param[Array.isArray(id) ? 'id_arr' : 'id'] = id
    await request(t('config.VITE_HTTP_API_PREFIX') + '/app/app/update', param, true)
}

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
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/app/app/list', param)
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
            <el-space :size="10" style="height: 100%; margin-left: 10px">
                <el-button v-if="authAction.isCreate" type="primary" @click="handleAdd"><autoicon-ep-edit-pen />{{ t('common.add') }}</el-button>
                <el-button v-if="authAction.isDelete" type="danger" @click="handleBatchDelete"><autoicon-ep-delete-filled />{{ t('common.batchDelete') }}</el-button>
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%; margin-right: 10px">
                <my-export-button i18nPrefix="app.app" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/app/app/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
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
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="50">
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
