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
                        <el-checkbox model-value={table.data.length ? allChecked : false} indeterminate={someChecked && !allChecked} onChange={(val: boolean) => table.data.forEach((item: any) => (item.checked = val))} />
                    </div>,
                    <div>{t('common.name.id')}</div>,
                ]
            },
            cellRenderer: (props: any): any => {
                return [<el-checkbox class="id-checkbox" model-value={props.rowData.checked ? true : false} onChange={(val: boolean) => (props.rowData.checked = val)} />, <div>{props.rowData.id}</div>]
            },
        },
        {
            dataKey: 'app_name',
            title: t('app.pkg.name.app_id'),
            key: 'app_id',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'pkg_type',
            title: t('app.pkg.name.pkg_type'),
            key: 'pkg_type',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let tagType = tm('config.const.tagType') as string[]
                let statusList = tm('app.pkg.status.pkg_type') as { value: any; label: string }[]
                let statusIndex = statusList.findIndex((item) => item.value == props.rowData.pkg_type)
                return <el-tag type={tagType[statusIndex % tagType.length]}>{statusList[statusIndex]?.label}</el-tag>
            },
        },
        {
            dataKey: 'pkg_name',
            title: t('app.pkg.name.pkg_name'),
            key: 'pkg_name',
            align: 'center',
            width: 150,
            cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [
                        <el-text line-clamp="2" title={props.rowData.pkg_name}>
                            {props.rowData.pkg_name}
                        </el-text>,
                    ]
                }
                if (!props.rowData?.editPkgName?.isEdit) {
                    return [
                        <el-text class="inline-edit" type="primary" line-clamp="2" title={props.rowData.pkg_name} onClick={() => (props.rowData.editPkgName = { isEdit: true, oldValue: props.rowData.pkg_name })}>
                            {props.rowData.pkg_name}
                        </el-text>,
                    ]
                }
                let currentRef: any
                return [
                    <el-input
                        ref={(el: any) => (el?.focus(), (currentRef = el))}
                        v-model={props.rowData.pkg_name}
                        placeholder={t('app.pkg.name.pkg_name')}
                        maxlength={60}
                        show-word-limit={true}
                        onBlur={() => {
                            props.rowData.editPkgName.isEdit = false
                            if (props.rowData.pkg_name == props.rowData.editPkgName.oldValue) {
                                return
                            }
                            if (!props.rowData.pkg_name) {
                                props.rowData.pkg_name = props.rowData.editPkgName.oldValue
                                return
                            }
                            handleUpdate(props.rowData.id, { pkg_name: props.rowData.pkg_name }).catch(() => (props.rowData.pkg_name = props.rowData.editPkgName.oldValue))
                        }}
                        onKeydown={(event: any) => [13].includes(event.keyCode) && currentRef?.blur()} //13：Enter键 27：Esc键 32：空格键
                    />,
                ]
            },
        },
        {
            dataKey: 'pkg_file',
            title: t('app.pkg.name.pkg_file'),
            key: 'pkg_file',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                if (!props.rowData.pkg_file) {
                    return
                }
                const fileList = [props.rowData.pkg_file]
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
            title: t('app.pkg.name.ver_no'),
            key: 'ver_no',
            align: 'center',
            width: 150,
            sortable: true,
            cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [
                        <el-text line-clamp="2" title={props.rowData.ver_no}>
                            {props.rowData.ver_no}
                        </el-text>,
                    ]
                }
                if (!props.rowData?.editVerNo?.isEdit) {
                    return [
                        <el-text class="inline-edit" type="primary" line-clamp="2" title={props.rowData.ver_no} onClick={() => (props.rowData.editVerNo = { isEdit: true, oldValue: props.rowData.ver_no })}>
                            {props.rowData.ver_no}
                        </el-text>,
                    ]
                }
                let currentRef: any
                return [
                    <el-input-number
                        ref={(el: any) => (el?.focus(), (currentRef = el))}
                        v-model={props.rowData.ver_no}
                        placeholder={t('app.pkg.name.ver_no')}
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
                        onKeydown={(event: any) => [13].includes(event.keyCode) && currentRef?.blur()} //13：Enter键 27：Esc键 32：空格键
                    />,
                ]
            },
        },
        {
            dataKey: 'ver_name',
            title: t('app.pkg.name.ver_name'),
            key: 'ver_name',
            align: 'center',
            width: 150,
            cellRenderer: (props: any): any => {
                if (!authAction.isUpdate) {
                    return [
                        <el-text line-clamp="2" title={props.rowData.ver_name}>
                            {props.rowData.ver_name}
                        </el-text>,
                    ]
                }
                if (!props.rowData?.editVerName?.isEdit) {
                    return [
                        <el-text class="inline-edit" type="primary" line-clamp="2" title={props.rowData.ver_name} onClick={() => (props.rowData.editVerName = { isEdit: true, oldValue: props.rowData.ver_name })}>
                            {props.rowData.ver_name}
                        </el-text>,
                    ]
                }
                let currentRef: any
                return [
                    <el-input
                        ref={(el: any) => (el?.focus(), (currentRef = el))}
                        v-model={props.rowData.ver_name}
                        placeholder={t('app.pkg.name.ver_name')}
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
                        onKeydown={(event: any) => [13].includes(event.keyCode) && currentRef?.blur()} //13：Enter键 27：Esc键 32：空格键
                    />,
                ]
            },
        },
        {
            dataKey: 'ver_intro',
            title: t('app.pkg.name.ver_intro'),
            key: 'ver_intro',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'extra_config',
            title: t('app.pkg.name.extra_config'),
            key: 'extra_config',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'remark',
            title: t('app.pkg.name.remark'),
            key: 'remark',
            align: 'center',
            width: 200,
            hidden: true,
        },
        {
            dataKey: 'is_force_prev',
            title: t('app.pkg.name.is_force_prev'),
            key: 'is_force_prev',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let statusList = tm('common.status.whether') as { value: any; label: string }[]
                return [
                    <el-switch
                        model-value={props.rowData.is_force_prev}
                        active-value={statusList[1].value}
                        inactive-value={statusList[0].value}
                        active-text={statusList[1].label}
                        inactive-text={statusList[0].label}
                        inline-prompt={true}
                        disabled={!authAction.isUpdate}
                        onChange={(val: any) => handleUpdate(props.rowData.id, { is_force_prev: val }).then(() => (props.rowData.is_force_prev = val))}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                    />,
                ]
            },
        },
        {
            dataKey: 'is_stop',
            title: t('app.pkg.name.is_stop'),
            key: 'is_stop',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                let statusList = tm('common.status.whether') as { value: any; label: string }[]
                return [
                    <el-switch
                        model-value={props.rowData.is_stop}
                        active-value={statusList[1].value}
                        inactive-value={statusList[0].value}
                        active-text={statusList[1].label}
                        inactive-text={statusList[0].label}
                        inline-prompt={true}
                        disabled={!authAction.isUpdate}
                        onChange={(val: any) => handleUpdate(props.rowData.id, { is_stop: val }).then(() => (props.rowData.is_stop = val))}
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
    sort: { key: 'created_at', order: 'desc' } as any,
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
    request(t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/info', { id: id }).then((res) => {
        saveCommon.data = { ...res.data.info }
        saveCommon.title = t('common.' + type)
        if (type == 'copy') {
            delete saveCommon.data.id
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
            if (action == 'confirm') {
                instance.confirmButtonLoading = true
                request(t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/del', { [Array.isArray(id) ? 'id_arr' : 'id']: id }, true)
                    .then(() => (table.data = table.data.filter((rowData: any) => (Array.isArray(id) ? !id.includes(rowData.id) : rowData.id != id))) /* getList() */, done())
                    .finally(() => (instance.confirmButtonLoading = false))
            } else {
                done()
            }
        },
    })
}
//更新
const handleUpdate = async (id: number | number[], param: { [propName: string]: any }) => {
    param[Array.isArray(id) ? 'id_arr' : 'id'] = id
    await request(t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/update', param, true)
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
        filter: removeEmptyOfObj(queryCommon.data),
        sort: table.sort.key + ' ' + table.sort.order,
        page: pagination.page,
        limit: pagination.size,
    }
    table.loading = true
    try {
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/list', param)
        table.data = res.data.list?.length ? res.data.list : []
        pagination.total = res.data.count
    } finally {
        table.loading = false
    }
}
getList()

//暴露组件接口给父组件
defineExpose({ getList })
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
                <my-export-button i18nPrefix="app.pkg" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/app/pkg/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
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
