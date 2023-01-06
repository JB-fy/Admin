<script setup lang="ts">
const { t } = useI18n()

const table = reactive({
    columns: [{
        dataKey: 'id',
        title: t('common.name.id'),
        key: 'id',
        width: 150,
        align: 'center',
        fixed: 'left',
        sortable: true,
    },
    {
        dataKey: 'requestUrl',
        title: t('common.name.log.request.requestUrl'),
        key: 'requestUrl',
        width: 300,
        align: 'center',
    },
    {
        dataKey: 'requestHeader',
        title: t('common.name.log.request.requestHeader'),
        key: 'requestHeader',
        width: 200,
        align: 'center',
    },
    {
        dataKey: 'requestData',
        title: t('common.name.log.request.requestData'),
        key: 'requestData',
        width: 200,
        align: 'center',
    },
    {
        dataKey: 'responseBody',
        title: t('common.name.log.request.responseBody'),
        key: 'responseBody',
        width: 200,
        align: 'center',
    },
    {
        dataKey: 'runTime',
        title: t('common.name.log.request.runTime'),
        key: 'runTime',
        align: 'center',
        width: 150,
        sortable: true,
    },
    {
        dataKey: 'updateTime',
        title: t('common.name.updateTime'),
        key: 'updateTime',
        align: 'center',
        width: 150,
        sortable: true,
    },
    {
        dataKey: 'createTime',
        title: t('common.name.createTime'),
        key: 'createTime',
        align: 'center',
        width: 150,
        sortable: true
    }] as any,
    data: [],
    loading: false,
    order: { key: 'id', order: 'desc' } as any,
    handleOrder: (order: any) => {
        table.order = order
        table.data = table.data.reverse()
        getList()
    },
})
//导出
const exportButton = reactive({
    loading: false,
    click: () => {

        ElMessageBox.confirm('', {
            type: 'warning',
            title: t('common.tip.configExport'),
            center: true,
            showClose: false,
        }).then(() => {
            exportButton.loading = true
            /* import('@/vendor/Export2Excel').then(excel => {
                const tHeader = [
                    this.name.requestUrl,
                    this.name.requestData,
                    this.name.requestHeaders,
                    this.name.responseData,
                    this.name.runTime,
                    this.name.addTime
                ]
                const filterVal = ['requestUrl', 'requestData', 'requestHeaders', 'responseData', 'runTime', 'addTime']
                const data = (() => {
                    return this.table.list.map(v => filterVal.map(j => {
                        switch (j) {
                            default:
                                return v[j];
                        }
                    }))
                })()
                excel.export_json_to_excel({
                    header: tHeader,
                    data,
                    filename: 'logRequest'
                })
            }) */
            //exportButton.loading = false
        }).catch(() => { })
    }
})

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
        where: removeEmptyOfObj(queryCommon.data),
        order: { [table.order.key]: table.order.order },
        page: pagination.page,
        limit: pagination.size
    }
    table.loading = true
    try {
        const res = await request('log/request/list', param)
        table.data = res.data.list
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

            </ElSpace>
        </ElCol>
        <ElCol :span="8" style="text-align: right;">
            <ElSpace :size="10" style="height: 100%;">
                <ElButton type="primary" :round="true" @click="exportButton.click" :loading="exportButton.loading">
                    <AutoiconEpDownload />{{ t('common.export') }}
                </ElButton>
                <ElDropdown max-height="300" :hide-on-click="false">
                    <ElButton type="info" :circle="true">
                        <AutoiconEpHide />
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
                    @column-sort="table.handleOrder" :width="width" :height="height" :fixed="true" :row-height="50">
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