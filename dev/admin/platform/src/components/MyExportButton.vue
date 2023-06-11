<script setup lang="ts">
const { t, tm } = useI18n()

const props = defineProps({
    fileName: {
        type: String
    },
    headerList: {   //表头。格式：{ [propName: string]: string }。也可直接传table.columns，即各个页面table组件的列定义
        type: [Object, Array],
        required: true
    },
    /**
     * 接口。格式：{ code: string, param: object, transform: function }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ filter: { [propName: string]: any }, field: string[], sort: { key: string, order: string }, page: number, limit: number }。
     *      transform：非必须。接口返回数据转换方法。当有字段转化时必须
     */
    api: {
        type: Object,
        required: true,
    },
})

//导出
const exportButton = reactive({
    loading: false,
    fileName: props.fileName ?? useLanguageStore().getPageTitle() + '.xlsx',
    headerList: computed(() => {
        if (Array.isArray(props.headerList)) {
            return props.headerList.reduce((headerListTmp: { [propName: string]: string }, item: any) => {
                item.dataKey ? headerListTmp[item.dataKey] = item.title : null
                return headerListTmp
            }, {})
        }
        return props.headerList
    }),
    api: {
        param: computed((): { filter: { [propName: string]: any }, field: string[], sort: { key: string, order: string }, page: number, limit: number } => {
            const param = {
                filter: {} as { [propName: string]: any },
                field: [],
                sort: { key: 'id', order: 'desc' },
                page: 1,
                limit: useSettingStore().exportButton.limit,
                ...(props.api?.param ?? {})
            }
            param.filter = removeEmptyOfObj(param.filter)
            return param
        }),
        transform: computed(() => {
            return props.api.transform ? props.api.transform : (res: any) => {
                return res.data.list.map((item: any) => {
                    for (const key in item) {
                        switch (key) {
                            case 'isStop':
                                item[key] = (<any>tm('common.status.whether')).find((item: any) => {
                                    return item.value == item[key]
                                }).label
                                break;
                        }
                    }
                    return item
                })
            }
        }),
        getData: async () => {
            let data = []
            try {
                const res = await request(props.api.code, exportButton.api.param)
                data = exportButton.api.transform(res)
            } catch (error) { }
            return data
        },
    },
    dataHandle: (data: any, headerList: { [propName: string]: string }) => {
        return data.map((item: any) => {
            const tmp: { [propName: string]: any } = {}
            for (const key in headerList) {
                item[key] ? tmp[headerList[key]] = item[key] : null
            }
            return tmp
        })
    },
    click: () => {
        ElMessageBox.confirm('', {
            type: 'warning',
            title: t('common.tip.configExport'),
            center: true,
            showClose: false,
        }).then(async () => {
            exportButton.loading = true
            while (true) {
                try {
                    let data = await exportButton.api.getData()
                    const length = data.length
                    if (data.length == 0) {
                        break
                    }
                    data = exportButton.dataHandle(data, exportButton.headerList)
                    exportExcel([{ data: data }], exportButton.fileName)
                    if (exportButton.api.param.limit === 0 || length < exportButton.api.param.limit) {
                        break
                    }
                    exportButton.api.param.page++
                } catch (error) {
                    break
                }
            }
            exportButton.loading = false
        }).catch(() => { })
    }
})

</script>

<template>
    <ElButton type="primary" :round="true" @click="exportButton.click" :loading="exportButton.loading">
        <AutoiconEpDownload />{{ t('common.export') }}
    </ElButton>

    <!-------- 使用示例 开始-------->
    <!-- <MyExportButton :headerList="table.columns"
        :api="{ code: 'log/request/list', param: { filter: queryCommon.data, sort: table.sort }" />

    <MyExportButton fileName="文件名.xlsx" :headerList="table.columns"
        :api="{ code: 'log/request/list', param: { filter: queryCommon.data, sort: table.sort }, limit: 0 }" /> -->
    <!-------- 使用示例 结束-------->
</template>