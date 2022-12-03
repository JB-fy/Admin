<script setup lang="ts">
const { t } = useI18n()

const saveVisible = inject('saveVisible')

const table = reactive({
    columns: [{
        dataKey: 'sceneId',
        title: 'ID',
        key: 'id',
        width: 120,
        align: 'left',
        fixed: 'left',
        sortable: true
    },
    {
        dataKey: 'sceneName',
        title: '场景名称',
        key: 'sceneName',
        align: 'center',
        width: 120,
    },
    {
        dataKey: 'sceneCode',
        title: '场景标识',
        key: 'sceneCode',
        align: 'center',
        width: 120,
    },
    {
        dataKey: 'sceneConfig',
        title: '场景配置',
        key: 'sceneConfig',
        width: 200,
        align: 'center',
        hidden: true
    },
    {
        dataKey: 'isStop',
        title: '停用',
        key: 'isStop',
        align: 'center',
        width: 120,
        cellRenderer: (data: any) => {
            return [
                h(ElButton, {
                    type: data.rowData.isStop ? 'danger' : 'primary',
                    size: 'small',
                    /* plain: true,
                    circle: true, */
                    round: true
                }, {
                    default: () => data.rowData.isStop ? '是' : '否'
                }),
            ] as any
        }
    },
    {
        dataKey: 'updateTime',
        title: '更新时间',
        key: 'updateTime',
        align: 'center',
        width: 150,
    },
    {
        dataKey: 'createTime',
        title: '创建时间',
        key: 'createTime',
        align: 'center',
        width: 150,
        sortable: true
    },
    {
        title: '操作',
        key: 'action',
        align: 'center',
        fixed: 'right',
        width: 200,
        cellRenderer: (data: any) => {
            return [
                h(ElButton, {
                    type: 'primary',
                    size: 'small',
                    onClick: () => handleEdit(data.rowData.sceneId)
                }, {
                    default: () => [h(AutoiconEpEdit), t('common.edit')]
                }),
                h(ElButton, {
                    type: 'danger',
                    size: 'small',
                    onClick: () => handleDelete(data.rowData.sceneId)
                }, {
                    default: () => [h(AutoiconEpDelete), t('common.delete')]
                })
            ]
        },
    }],
    data: [],
    loading: false,
    order: {
        key: 'id',
        order: 'desc',
    },
    handleOrder: (order: any) => {
        table.order = order
        table.data = table.data.reverse()
        getList()
    },
})

const pagination = reactive({
    total: 0,
    page: 1,
    limit: 10,
    sizeChange: (val: number) => {
        getList()
    },
    pageChange: (val: number) => {
        getList()
    }
})

const handleEdit = (id: number) => {

}
const handleDelete = (id: number) => {

}

const queryData = inject('queryData') as { [propName: string]: any }
const getList = async (resetPage: boolean = false) => {
    if (resetPage) {
        pagination.page = 1
    }
    const param = {
        field: [],
        where: removeEmptyOfObj(queryData),
        order: {
            [table.order.key]: table.order.order
        },
        page: pagination.page,
        limit: pagination.limit
    }
    table.loading = true
    try {
        const res = await request('auth.scene.list', param)
        table.data = res.data.list
        pagination.total = res.data.count
    } catch (error) {
    }
    table.loading = false
}
getList()

defineExpose({
    getList
})
</script>

<template>
    <ElRow class="main-table-tool">
        <ElCol :span="16">
            <ElSpace :size="10" style="height: 100%; margin-left: 10px;">
                <ElButton type="primary" @click="saveVisible = true">
                    <AutoiconEpEditPen />{{ t('common.add') }}
                </ElButton>
                <ElButton type="danger">
                    <AutoiconEpDelete />{{ t('common.delete') }}
                </ElButton>
            </ElSpace>
        </ElCol>
        <ElCol :span="8" style="text-align: right;">
            <ElSpace :size="10" style="height: 100%;">
                <ElDropdown max-height="300" :hide-on-click="false">
                    <ElButton type="info" :circle="true">
                        <AutoiconEpView />
                    </ElButton>
                    <template #dropdown>
                        <ElDropdownMenu>
                            <ElDropdownItem v-for="(item, key) in table.columns" :key="key">
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

    <ElMain style="padding: 0;">
        <ElAutoResizer>
            <template #default="{ height, width }">
                <ElTableV2 class="main-table" :columns="table.columns" :data="table.data" :sort-by="table.order"
                    @column-sort="table.handleOrder" :width="width" :height="height" :fixed="true">
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
                v-model:page-size="pagination.limit" @size-change="pagination.sizeChange"
                @current-change="pagination.pageChange" :page-sizes="[10, 20, 50, 100, 200, 500, 1000]"
                layout="total, sizes, prev, pager, next, jumper" :background="true" />
        </ElCol>
    </ElRow>

</template>

<style scoped>
.main-table-tool {
    height: 40px;
    background-color: var(--el-bg-color);
    border-bottom: 2px dashed var(--el-border-color);
    border-top-right-radius: 8px;
    border-top-left-radius: 8px;
}

/* .main-table :deep(.el-table-v2__main) {
    position: static;
} */

.main-table :deep(.el-table-v2__overlay) {
    z-index: 10;
    background-color: var(--el-mask-color);
    display: flex;
    align-items: center;
    justify-content: center;
}

.main-table-pagination {
    height: 40px;
    background-color: var(--el-bg-color);
    border-top: 2px dashed var(--el-border-color);
    border-bottom-right-radius: 8px;
    border-bottom-left-radius: 8px;
}

.main-table-pagination :deep(.el-pagination) {
    float: right;
    margin-right: 5px;
}

/* .main-table :deep(.el-table-v2__footer) {
    background-color: var(--el-bg-color);
    border-top: 2px dashed var(--el-border-color);
    border-bottom-right-radius: 8px;
    border-bottom-left-radius: 8px;
    padding-top: 1px;
    position: static;
}

.main-table :deep(.el-table-v2__footer .el-pagination) {
    float: right;
    margin-right: 5px;
} */
</style>