<!-------- 使用示例 开始-------->
<!-- <my-transfer v-model="saveForm.data.scene_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" />

<my-transfer v-model="saveForm.data.scene_id_arr" :options="tm('common.status.xxxx')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list', param: { field: ['id', 'scene_name'] } }" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
defineOptions({ inheritAttrs: false })
const attrs = useAttrs()
const slots = useSlots()
const model = defineModel()
const emits = defineEmits(['change'])
const props = defineProps({
    /**
     * 接口。格式：{ code: string, param: object, transform: function }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ filter: { [propName: string]: any }, field: string[], sort: string, page: number, limit: number }。其中field内第0，1字段默认用于transfer.props的key，label属性，transfer.api的transform属性，使用时请注意。或直接在props.props中设置对应参数
     *      transform：非必须。接口返回数据转换方法
     */
    api: {
        type: Object,
        required: true,
    },
})

const transfer = reactive({
    ref: null as any,
    options: [...((attrs.options as any[]) ?? [])] as any,
    props: {
        key: 'value',
        label: 'label',
        ...(attrs.props ?? {}),
    },
    initOptions: () => transfer.api.addOptions(),
    api: {
        loading: false,
        param: computed((): { filter: { [propName: string]: any }; field: string[]; sort: string; page: number; limit: number } => {
            return {
                filter: {} as { [propName: string]: any },
                field: ['id', 'label'],
                page: 1,
                limit: 0,
                ...(props.api?.param ?? {}),
            }
        }),
        transform: computed(() => {
            return props.api.transform
                ? props.api.transform
                : (res: any) => {
                      const options: { key: any; label: any }[] = []
                      res.data.list.forEach((item: any) => {
                          options.push({
                              ...item,
                              [transfer.props.key]: item[transfer.api.param.field[0]],
                              [transfer.props.label]: item[transfer.api.param.field[1]],
                          })
                      })
                      return options
                  }
        }),
        getOptions: async () => {
            if (transfer.api.loading) {
                return
            }
            transfer.api.loading = true
            let options = []
            try {
                const res = await request(props.api.code, transfer.api.param)
                options = transfer.api.transform(res)
            } finally {
                transfer.api.loading = false
            }
            return options
        },
        addOptions: () => transfer.api.getOptions().then((options) => (transfer.options = [...((attrs.options as any[]) ?? []), ...(options ?? [])])),
    },
})
//组件创建时，初始化options
transfer.initOptions()

//当外部环境filter变化时，重置options
watch(
    () => props.api?.param?.filter,
    (newVal: any, oldVal: any) => {
        if (JSON.stringify(newVal) !== JSON.stringify(oldVal)) {
            transfer.api.param.page = 1
            transfer.api.addOptions()
        }
    }
)

//暴露组件属性给父组件
defineExpose({
    options: computed(() => transfer.options),
})
</script>

<template>
    <el-transfer
        :ref="(el: any) => transfer.ref = el"
        v-model="(model as any)"
        :filterable="true"
        v-bind="$attrs"
        :data="transfer.options"
        :props="transfer.props"
        @change="(value: any, direction: any, movedKeys: any) => emits('change', value, direction, movedKeys, transfer.options.filter((item: any) => value.includes(item[transfer.props.key])))"
    >
        <template v-if="slots.default" #default="{ option }">
            <slot name="default" :option="option"></slot>
        </template>
        <template v-if="slots.leftFooter" #left-footer>
            <slot name="left-footer"></slot>
        </template>
        <template v-if="slots.rightFooter" #right-footer>
            <slot name="right-footer"></slot>
        </template>
    </el-transfer>
</template>

<style scoped>
:deep(.el-transfer-panel) {
    width: auto;
    min-width: var(--el-transfer-panel-width);
}
</style>
