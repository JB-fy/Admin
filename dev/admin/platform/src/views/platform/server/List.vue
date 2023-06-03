<script setup lang="ts">
const { t } = useI18n()

const table = reactive({
    columns: [{
        dataKey: 'id',
        title: t('common.name.id'),
        key: 'id',
        width: 200,
        align: 'center',
        fixed: 'left',
        sortable: true,
    },
    {
        dataKey: 'networkIp',
        title: t('common.name.platform.server.networkIp'),
        key: 'networkIp',
        width: 200,
        align: 'center',
    },
    {
        dataKey: 'localIp',
        title: t('common.name.platform.server.localIp'),
        key: 'localIp',
        width: 200,
        align: 'center',
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
        getList()
    },
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
        filter: removeEmptyOfObj(queryCommon.data),
        order: { [table.order.key]: table.order.order },
        page: pagination.page,
        limit: pagination.size
    }
    table.loading = true
    try {
        const res = await request('/platform/server/list', param)
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

    <ElMain>
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