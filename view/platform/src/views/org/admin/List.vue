<script setup lang="tsx">
const { t } = useI18n()

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
                                    if (item.is_super == 0) {
                                        return
                                    }
                                    item.checked = val
                                })
                            }}
                        />
                    </div>,
                    <div>{t('common.name.id')}</div>,
                ]
            },
            cellRenderer: (props: any): any => {
                return [<el-checkbox class="id-checkbox" model-value={props.rowData.checked} disabled={props.rowData.is_super == 0} onChange={(val: boolean) => (props.rowData.checked = val)} />, <div>{props.rowData.id}</div>]
            },
        },
        {
            dataKey: 'org_name',
            title: t('org.admin.name.org_id'),
            key: 'org_id',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'nickname',
            title: t('org.admin.name.nickname'),
            key: 'nickname',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'avatar',
            title: t('org.admin.name.avatar'),
            key: 'avatar',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                if (!props.rowData.avatar) {
                    return
                }
                const imageList = [props.rowData.avatar]
                return [
                    <el-scrollbar wrap-style="display: flex; align-items: center;" view-style="margin: auto;">
                        <el-space direction="vertical" style="margin: 5px 10px;">
                            {imageList.map((item) => {
                                return <el-image src={item} lazy={true} hide-on-click-modal={true} preview-teleported={true} preview-src-list={imageList} /> //修改宽高时，可同时修改table属性row-height增加行高，则不会显示滚动条
                            })}
                        </el-space>
                    </el-scrollbar>,
                ]
            },
        },
        {
            dataKey: 'phone',
            title: t('org.admin.name.phone'),
            key: 'phone',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'email',
            title: t('org.admin.name.email'),
            key: 'email',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'account',
            title: t('org.admin.name.account'),
            key: 'account',
            align: 'center',
            width: 150,
        },
        {
            dataKey: 'is_super',
            title: t('org.admin.name.is_super'),
            key: 'is_super',
            align: 'center',
            width: 100,
            cellRenderer: (props: any): any => {
                return [
                    <el-switch
                        model-value={props.rowData.is_super}
                        active-value={1}
                        inactive-value={0}
                        inline-prompt={true}
                        active-text={t('common.yes')}
                        inactive-text={t('common.no')}
                        disabled={true}
                        onChange={(val: number) => handleUpdate({ id_arr: [props.rowData.id], is_super: val }).then(() => (props.rowData.is_super = val))}
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success);"
                    />,
                ]
            },
        },
        {
            dataKey: 'is_stop',
            title: t('org.admin.name.is_stop'),
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
                        disabled={!authAction.isUpdate || props.rowData.is_super == 0}
                        onChange={(val: number) => handleUpdate({ id_arr: [props.rowData.id], is_stop: val }).then(() => (props.rowData.is_stop = val))}
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
                if (props.rowData.is_super == 0) {
                    return
                }
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
    if (idArr.length) {
        handleDelete(idArr)
    } else {
        ElMessage.error(t('common.tip.selectDelete'))
    }
}
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/org/admin/info', { id: id }).then((res) => {
        saveCommon.data = { ...res.data.info }
        switch (type) {
            case 'edit':
                saveCommon.data.id_arr = [saveCommon.data.id]
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
}
//删除
const handleDelete = (idArr: number[]) => {
    ElMessageBox.confirm('', {
        type: 'warning',
        title: t('common.tip.configDelete'),
        center: true,
        showClose: false,
    }).then(() => request(t('config.VITE_HTTP_API_PREFIX') + '/org/admin/del', { id_arr: idArr }, true).then(() => getList()))
}
//更新
const handleUpdate = async (param: { id_arr: number[]; [propName: string]: any }) => {
    await request(t('config.VITE_HTTP_API_PREFIX') + '/org/admin/update', param, true)
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
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/org/admin/list', param)
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
                <el-button v-if="authAction.isCreate" type="primary" @click="handleAdd"> <autoicon-ep-edit-pen />{{ t('common.add') }} </el-button>
                <el-button v-if="authAction.isDelete" type="danger" @click="handleBatchDelete"> <autoicon-ep-delete-filled />{{ t('common.batchDelete') }} </el-button>
            </el-space>
        </el-col>
        <el-col :span="8" style="text-align: right">
            <el-space :size="10" style="height: 100%">
                <my-export-button i18nPrefix="org.admin" :headerList="table.columns" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/org/admin/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
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
                <el-table-v2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort" @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="50">
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
