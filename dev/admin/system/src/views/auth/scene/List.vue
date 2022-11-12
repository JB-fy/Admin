<script setup lang="ts">
const { t } = useI18n()

const saveDrawerVisible = inject('saveDrawerVisible')

const generateColumns = (length = 10, prefix = 'column-', props?: any) =>
    Array.from({ length }).map((_, columnIndex) => ({
        ...props,
        key: `${prefix}${columnIndex}`,
        dataKey: `${prefix}${columnIndex}`,
        title: `Column ${columnIndex}`,
        width: 150,
    }))

const generateData = (
    columns: ReturnType<typeof generateColumns>,
    length = 200,
    prefix = 'row-'
) => Array.from({ length }).map((_, rowIndex) => {
    return columns.reduce(
        (rowData, column, columnIndex) => {
            rowData[column.dataKey] = `Row ${rowIndex} - Col ${columnIndex}`
            return rowData
        },
        {
            id: `${prefix}${rowIndex}`,
            parentId: null,
        }
    )
})

/* const columns = generateColumns(10)
const data = generateData(columns, 100) */

//const data = await request('auth.scene.list', { where: queryForm.data })
const columns = [
    {
        dataKey: 'id',
        title: 'ID',
        key: 'id',
        width: 150,
        hidden: false,
        align: 'left',
        fixed: 'left',
        sortable: true
    },
    {
        dataKey: 'sceneName',
        title: '场景名称',
        key: 'sceneName',
        width: 120,
        style: { width: 'auto' },
        minWidth: 120,
        maxWidth: 200,
    },
    {
        dataKey: 'sceneCode',
        title: '场景标识',
        key: 'sceneCode',
        width: 120,
        style: { width: 'auto' },
        minWidth: 120,
        maxWidth: 200,
    },
    {
        dataKey: 'action',
        title: '操作',
        key: 'action',
        fixed: 'right',
        width: 150,
    }
]
const data = [
    {
        parentId: null,
        id: '1',
        sceneName: '场景名称1',
        sceneCode: '场景标识1'
    },
    {
        parentId: null,
        id: '2',
        sceneName: '场景名称2',
        sceneCode: '场景标识2'
    }
]

const pagination = reactive({
    data: {
        currentPage: 1,
        pageSize: 10,
    },
    sizeChange: (val: number) => {
        console.log(`${val} items per page`)
    },
    currentChange: (val: number) => {
        console.log(`current page: ${val}`)
    }
})
</script>

<template>
    <ElAutoResizer>
        <template #default="{ height, width }">
            <ElRow class="main-table-tool">
                <ElCol :span="16">
                    <ElSpace :size="10" style="height: 100%; margin-left: 10px;">
                        <ElButton type="primary" @click="saveDrawerVisible = true">
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
                                    <ElDropdownItem>
                                        <ElCheckbox>
                                            字段1
                                        </ElCheckbox>
                                    </ElDropdownItem>
                                    <ElDropdownItem>
                                        <ElCheckbox>
                                            字段2
                                        </ElCheckbox>
                                    </ElDropdownItem>
                                </ElDropdownMenu>
                            </template>
                        </ElDropdown>
                    </ElSpace>
                </ElCol>
            </ElRow>

            <ElTableV2 class="main-table" :columns="columns" :data="data" :width="width" :height="0"
                :max-height="height - 40" :footer-height="40" :fixed="true">
                <template #footer>
                    <ElPagination v-model:currentPage="pagination.data.currentPage"
                        v-model:page-size="pagination.data.pageSize" :page-sizes="[10, 20, 50, 100, 200, 500, 1000]"
                        :background="true" layout="total, sizes, prev, pager, next, jumper" :total="400"
                        @size-change="pagination.sizeChange" @current-change="pagination.currentChange" />
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