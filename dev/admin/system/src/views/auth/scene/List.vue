<script setup lang="ts">
const { t } = useI18n()

const saveFormVisible = inject('saveFormVisible')

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
        width: 120,
    },
    {
        dataKey: 'sceneCode',
        title: '场景标识',
        key: 'sceneCode',
        width: 120,
    },
    {
        dataKey: 'sceneConfig',
        title: '场景配置',
        key: 'sceneConfig',
        width: 200,
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
                })
            ]
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
    },
    {
        dataKey: 'action',
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
    data: []
})

const pagination = reactive({
    total: 1,
    page: 1,
    limit: 10
})

const handleEdit = (id: number) => {

}
const handleDelete = (id: number) => {

}

const queryFormData = inject('queryFormData') as { [propName: string]: any }
const getList = async (resetPage: boolean = false) => {
    if (resetPage) {
        pagination.page = 1
    }
    let param = {
        field: [],
        where: removeObjectNullValue(queryFormData),
        order: {},
        page: pagination.page,
        limit: pagination.limit
    }
    const res = await request('auth.scene.list', param)
    table.data = res.data.list
}
getList()

defineExpose({
    getList
})
</script>

<template>
    <ElAutoResizer>
        <template #default="{ height, width }">
            <ElRow class="main-table-tool">
                <ElCol :span="16">
                    <ElSpace :size="10" style="height: 100%; margin-left: 10px;">
                        <ElButton type="primary" @click="saveFormVisible = true">
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

            <ElTableV2 class="main-table" :columns="table.columns" :data="table.data" :width="width"
                :height="height - 40" :footer-height="40" :fixed="true">
                <template #footer>
                    <ElPagination v-model:total="pagination.total" v-model:currentPage="pagination.page"
                        v-model:page-size="pagination.limit" @size-change="getList" @current-change="getList"
                        :page-sizes="[10, 20, 50, 100, 200, 500, 1000]" layout="total, sizes, prev, pager, next, jumper"
                        :background="true" />
                </template>
            </ElTableV2>
        </template>
    </ElAutoResizer>
</template>

<style scoped>
.main-table-tool {
    height: 40px;
    background-color: var(--el-bg-color);
    border-bottom: 2px dashed var(--el-border-color);
    border-top-right-radius: 8px;
    border-top-left-radius: 8px;
}

.main-table :deep(.el-table-v2__main) {
    position: static;
}

.main-table :deep(.el-table-v2__footer) {
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
}
</style>