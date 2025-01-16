<!-------- 使用示例 开始-------->
<!-- <my-transfer v-model="saveForm.data.scene_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" />

<my-transfer v-model="saveForm.data.scene_id_arr" :defaultOptions="tm('common.status.xxxx')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list', param: { field: ['id', 'scene_name'] } }" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
const slots = useSlots()
const props = defineProps({
    modelValue: {
        type: Array,
    },
    defaultOptions: {
        //选项初始默认值。格式：{ [transfer.props.key]: any, [transfer.props.label]: any, [propName: string]: any }[]
        type: Array,
        default: () => [],
    },
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
    placeholder: {
        type: String,
    },
    filterable: {
        type: Boolean,
        default: true,
    },
    props: {
        type: Object,
        default: () => {},
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const transfer = reactive({
    ref: null as any,
    value: computed({
        get: (): any => {
            return props.modelValue
        },
        set: (val) => {
            emits('update:modelValue', val)
            emits(
                'change',
                transfer.options.filter((item: any) => val.includes(item[transfer.props.key]))
            )
        },
    }),
    options: [...props.defaultOptions] as any,
    props: {
        key: 'value',
        label: 'label',
        ...props.props,
    },
    initOptions: () => {
        transfer.api.addOptions()
    },
    api: {
        loading: false,
        param: computed((): { filter: { [propName: string]: any }; field: string[]; sort: string; page: number; limit: number } => {
            return {
                filter: {} as { [propName: string]: any },
                field: ['id', 'label'],
                sort: 'id desc',
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
        addOptions: () => {
            transfer.api.getOptions().then((options) => {
                transfer.options = [...props.defaultOptions, ...(options ?? [])]
            })
        },
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
    options: computed(() => {
        return transfer.options
    }),
})
</script>

<template>
    <el-transfer :ref="(el: any) => transfer.ref = el" v-model="transfer.value" :data="transfer.options" :filterable="filterable" :filter-placeholder="placeholder" :props="transfer.props">
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
