<script setup lang="ts">
const { t } = useI18n()

const table = reactive({
    columns: [{
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
                h('div', {
                    class: 'id-checkbox',
                    onClick: (event: any) => {
                        event.stopPropagation();    //阻止冒泡
                    },
                }, {
                    default: () => [
                        h(ElCheckbox as any, {
                            'model-value': table.data.length ? allChecked : false,
                            indeterminate: someChecked && !allChecked,
                            onChange: (val: boolean) => {
                                table.data.forEach((item: any) => {
                                    item.checked = val
                                })
                            }
                        })
                    ]
                }),
                h('div', {}, {
                    default: () => t('common.name.id')
                })
            ]
        },
        cellRenderer: (props: any): any => {
            return [
                h(ElCheckbox as any, {
                    class: 'id-checkbox',
                    'model-value': props.rowData.checked,
                    onChange: (val: boolean) => {
                        props.rowData.checked = val
                    }
                }),
                h('div', {}, {
                    default: () => props.rowData.id
                })
            ]
        },
    },
    {
        dataKey: 'account',
        title: t('platform.admin.name.account'),
        key: 'account',
        align: 'center',
        width: 150,
    },
    {
        dataKey: 'phone',
        title: t('platform.admin.name.phone'),
        key: 'phone',
        align: 'center',
        width: 150,
    },
    {
        dataKey: 'nickname',
        title: t('platform.admin.name.nickname'),
        key: 'nickname',
        align: 'center',
        width: 150,
    },
    {
        dataKey: 'avatar',
        title: t('platform.admin.name.avatar'),
        key: 'avatar',
        align: 'center',
        width: 100,
        cellRenderer: (props: any): any => {
            if (!props.rowData.avatar) {
                return
            }
            //const imageList= JSON.parse(props.rowData.avatar)
            const imageList = [props.rowData.avatar]
            return [
                h(ElScrollbar, {
                    'wrap-style': 'display: flex; align-items: center;',
                    'view-style': 'margin: auto;',
                }, {
                    default: () => {
                        const content = imageList.map((item) => {
                            return h(ElImage as any, {
                                'style': 'width: 45px;',    //不想显示滚动条，需设置table属性row-height增加行高
                                'src': item,
                                'lazy': true,
                                'hide-on-click-modal': true,
                                'preview-teleported': true,
                                'preview-src-list': imageList
                            })
                        })
                        return content
                    }
                })
            ]
        },
    },
    {
        dataKey: 'isStop',
        title: t('platform.admin.name.isStop'),
        key: 'isStop',
        align: 'center',
        width: 100,
        cellRenderer: (props: any): any => {
            return [
                h(ElSwitch as any, {
                    'model-value': props.rowData.isStop,
                    'active-value': 1,
                    'inactive-value': 0,
                    'inline-prompt': true,
                    'active-text': t('common.yes'),
                    'inactive-text': t('common.no'),
                    style: '--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)',
                    onChange: (val: number) => {
                        handleUpdate({
                            idArr: [props.rowData.id],
                            isStop: val
                        }).then((res) => {
                            props.rowData.isStop = val
                        }).catch((error) => { })
                    }
                })
            ]
        }
    },
    {
        dataKey: 'updatedAt',
        title: t('common.name.updatedAt'),
        key: 'updatedAt',
        align: 'center',
        width: 150,
        sortable: true,
    },
    {
        dataKey: 'createdAt',
        title: t('common.name.createdAt'),
        key: 'createdAt',
        align: 'center',
        width: 150,
        sortable: true
    },
    {
        title: t('common.name.action'),
        key: 'action',
        align: 'center',
        width: 250,
        fixed: 'right',
        cellRenderer: (props: any): any => {
            return [
                h(ElButton, {
                    type: 'primary',
                    size: 'small',
                    onClick: () => handleEditCopy(props.rowData.id)
                }, {
                    default: () => [h(AutoiconEpEdit), t('common.edit')]
                }),
                h(ElButton, {
                    type: 'danger',
                    size: 'small',
                    onClick: () => handleDelete([props.rowData.id])
                }, {
                    default: () => [h(AutoiconEpDelete), t('common.delete')]
                }),
                h(ElButton, {
                    type: 'warning',
                    size: 'small',
                    onClick: () => handleEditCopy(props.rowData.id, 'copy')
                }, {
                    default: () => [h(AutoiconEpDocumentCopy), t('common.copy')]
                }),
            ]
        },
    }] as any,
    data: [],
    loading: false,
    sort: { key: 'id', order: 'desc' } as any,
    handleSort: (sort: any) => {
        table.sort.key = sort.key
        table.sort.order = sort.order
        getList()
    },
})

const saveCommon = inject('saveCommon') as { visible: boolean, title: string, data: { [propName: string]: any } }
//新增
const handleAdd = () => {
    saveCommon.data = {}
    saveCommon.title = t('common.add')
    saveCommon.visible = true
}
//批量删除
const handleBatchDelete = () => {
    const idArr: number[] = [];
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
}
//编辑|复制
const handleEditCopy = (id: number, type: string = 'edit') => {
    request(t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/info', { id: id }).then((res) => {
        saveCommon.data = { ...res.data.info }
        switch (type) {
            case 'edit':
                saveCommon.data.idArr = [saveCommon.data.id]
                delete saveCommon.data.id
                saveCommon.title = t('common.edit')
                break;
            case 'copy':
                delete saveCommon.data.id
                saveCommon.title = t('common.copy')
                break;
        }
        saveCommon.visible = true
    }).catch(() => { })
}
//删除
const handleDelete = (idArr: number[]) => {
    ElMessageBox.confirm('', {
        type: 'warning',
        title: t('common.tip.configDelete'),
        center: true,
        showClose: false,
    }).then(() => {
        request(t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/del', { idArr: idArr }, true).then((res) => {
            getList()
        }).catch(() => { })
    }).catch(() => { })
}
//更新
const handleUpdate = async (param: { idArr: number[], [propName: string]: any }) => {
    await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/update', param, true)
}

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
    }
})

const queryCommon = inject('queryCommon') as { data: { [propName: string]: any } }
//列表
const getList = async (resetPage: boolean = false) => {
    if (resetPage) {
        pagination.page = 1
    }
    const param = {
        field: [],
        filter: removeEmptyOfObj(queryCommon.data),
        sort: table.sort.key + ' ' + table.sort.order,
        page: pagination.page,
        limit: pagination.size
    }
    table.loading = true
    try {
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/list', param)
        table.data = res.data.list?.length ? res.data.list : []
        pagination.total = res.data.count
    } catch (error) { }
    table.loading = false
}
getList()

//暴露组件接口给父组件
defineExpose({
    getList
})
</script>

<template>
    <ElRow class="main-table-tool">
        <ElCol :span="16">
            <ElSpace :size="10" style="height: 100%; margin-left: 10px;">
                <ElButton type="primary" @click="handleAdd">
                    <AutoiconEpEditPen />{{ t('common.add') }}
                </ElButton>
                <ElButton type="danger" @click="handleBatchDelete">
                    <AutoiconEpDeleteFilled />{{ t('common.batchDelete') }}
                </ElButton>
            </ElSpace>
        </ElCol>
        <ElCol :span="8" style="text-align: right;">
            <ElSpace :size="10" style="height: 100%;">
                <MyExportButton :headerList="table.columns"
                    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />
                <ElDropdown max-height="300" :hide-on-click="false">
                    <ElButton type="info" :circle="true">
                        <AutoiconEpHide />
                    </ElButton>
                    <template #dropdown>
                        <ElDropdownMenu>
                            <ElDropdownItem v-for="(item, index) in table.columns" :key="index">
                                <ElCheckbox v-model="item.hidden">
                                    {{ item.title }}
                                </ElCheckbox>
                            </ElDropdownItem>
                        </ElDropdownMenu>
                    </template>
                </ElDropdown>
            </ElSpace>
        </ElCol>
    </ElRow>

    <ElMain>
        <ElAutoResizer>
            <template #default="{ height, width }">
                <ElTableV2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.sort"
                    @column-sort="table.handleSort" :width="width" :height="height" :fixed="true" :row-height="50">
                    <template v-if="table.loading" #overlay>
                        <ElIcon class="is-loading" color="var(--el-color-primary)" :size="25">
                            <AutoiconEpLoading />
                        </ElIcon>
                    </template>
                </ElTableV2>
            </template>
        </ElAutoResizer>
    </ElMain>

    <ElRow class="main-table-pagination">
        <ElCol :span="24">
            <ElPagination :total="pagination.total" v-model:currentPage="pagination.page"
                v-model:page-size="pagination.size" @size-change="pagination.sizeChange"
                @current-change="pagination.pageChange" :page-sizes="pagination.sizeList" :layout="pagination.layout"
                :background="true" />
        </ElCol>
    </ElRow>
</template>