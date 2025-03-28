<!-------- 使用示例 开始-------->
<!-- <my-export-button i18nPrefix="auth.test" :headerList="table.columns"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/test/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order } }" />

<my-export-button fileName="文件名.xlsx" i18nPrefix="auth.test" :headerList="table.columns"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/test/list', param: { filter: queryCommon.data, sort: table.sort.key + ' ' + table.sort.order }, limit: 0 }" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
const { t, tm } = useI18n()

const props = defineProps({
    /**
     * 接口。格式：{ code: string, param: object, transform: function }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ filter: { [propName: string]: any }, field: string[], sort: string, page: number, limit: number }。
     *      transform：非必须。接口返回数据转换方法。当有字段转化时必须
     */
    api: {
        type: Object,
        required: true,
    },
    fileName: {
        type: String,
    },
    i18nPrefix: {
        //i18n包t, tm两个方法参数的前缀。示例：'auth.test'，则对应t('auth.test.name.xxxx')或tm('auth.test.status.xxxx')
        type: String,
        required: true,
    },
    headerList: {
        //表头。格式：{ [propName: string]: string }。也可直接传table.columns，即各个页面table组件的列定义
        type: [Object, Array],
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
                if (typeof item === 'string') {
                    headerListTmp[item] = tm(props.i18nPrefix + '.name.' + item)
                } else if (item.dataKey) {
                    headerListTmp[item.dataKey] = item.title
                }
                return headerListTmp
            }, {})
        }
        return props.headerList
    }),
    api: {
        param: computed((): { filter: { [propName: string]: any }; field: string[]; sort: string; page: number; limit: number } => {
            const param = {
                filter: {} as { [propName: string]: any },
                field: [],
                page: 1,
                limit: useSettingStore().exportButton.limit,
                ...(props.api?.param ?? {}),
            }
            param.filter = removeEmptyOfObj(param.filter)
            return param
        }),
        transform: computed(() => {
            return props.api.transform
                ? props.api.transform
                : (res: any) => {
                      return res.data.list.map((item: any) => {
                          for (const key in item) {
                              if (
                                  key
                                      .replace(/([A-Z])/g, '_$1')
                                      .toLowerCase()
                                      .startsWith('is_')
                              ) {
                                  item[key] = (tm('common.status.whether') as any).find((item1: any) => {
                                      return item1.value == item[key]
                                  }).label
                              } else {
                                  // let statusArr = tm(props.i18nPrefix + '.status.' + key) as { value: number | string, label: string }[]
                                  let statusArr = tm(props.i18nPrefix + '.status.' + key)
                                  if (statusArr.length) {
                                      item[key] = (statusArr as any).find((item1: any) => {
                                          return item1.value == item[key]
                                      }).label
                                  }
                              }
                          }
                          return item
                      })
                  }
        }),
        getData: async () => {
            let data = []
            const res = await request(props.api.code, exportButton.api.param)
            data = exportButton.api.transform(res)
            return data
        },
    },
    dataHandle: (data: any, headerList: { [propName: string]: string }) => {
        return data.map((item: any) => {
            const tmp: { [propName: string]: any } = {}
            for (const key in headerList) {
                key in item && (tmp[headerList[key]] = item[key])
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
            while (exportButton.loading) {
                let data = await exportButton.api.getData()
                if (data.length > 0) {
                    data = exportButton.dataHandle(data, exportButton.headerList)
                    exportExcel([{ data: data }], exportButton.fileName)
                    if (!(exportButton.api.param.limit === 0 || data.length < exportButton.api.param.limit)) {
                        exportButton.api.param.page++
                        continue
                    }
                }
                exportButton.loading = false
            }
        })
    },
})
</script>

<template>
    <el-button type="primary" :round="true" @click="exportButton.click" :loading="exportButton.loading"><autoicon-ep-download />{{ t('common.export') }}</el-button>
</template>
