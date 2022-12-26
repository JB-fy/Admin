<script setup lang="ts">
const props = defineProps({
    modelValue: {
        type: Array,
        //required: true,
    },
    defaultOptions: {   //选项初始默认值。格式：[{ [transfer.props.key]: string | number, [transfer.props.label]: string },...]
        type: Array,
        default: []
    },
    /**
     * 接口。格式：{ code: string, param: object, dataToOptions: function }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number }。其中field内第0，1字段默认用于transfer.props的key，label属性，transfer.api的dataToOptions属性，使用时请注意。或直接在props.props中设置对应参数
     *      dataToOptions：非必须。接口返回数据转换方法
     */
    api: {
        type: Object,
        required: true,
    },
    placeholder: {
        type: String
    },
    filterable: {
        type: Boolean,
        default: true
    },
    props: {
        type: Object,
        default: {}
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
            emits('change')
            emits('update:modelValue', val)
        }
    }),
    options: [...props.defaultOptions] as any,
    props: {
        key: props.props.key ?? props.api.param.field[0] ?? 'key',
        label: props.props.label ?? props.api.param.field[1] ?? 'label',
        disabled: props.props.disabled ?? 'disabled',
        ...props.props,
    },
    initOptions: () => {
        transfer.api.addOptions()
    },
    resetOptions: () => {
        transfer.options = [...props.defaultOptions] as any
        transfer.api.param.page = 1
    },
    api: {
        loading: false,
        param: computed((): { field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number } => {
            return {
                field: [],
                where: {} as { [propName: string]: any },
                order: { id: 'desc' },
                page: 1,
                limit: 0,
                ...props.api.param
            }
        }),
        dataToOptions: computed(() => {
            return props.api.dataToOptions ? props.api.dataToOptions : (res: any) => {
                return res.data.list
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
                options = transfer.api.dataToOptions(res)
            } catch (error) { }
            transfer.api.loading = false
            return options
        },
        addOptions: () => {
            transfer.api.getOptions().then((options) => {
                if (options.length) {
                    transfer.options = transfer.options.concat(options ?? [])
                }
            }).catch((error) => { })
        },
    },
})
//组件创建时，初始化options
transfer.initOptions()
</script>

<template>
    <ElTransfer :ref="(el: any) => { transfer.ref = el }" v-model="transfer.value" :data="transfer.options"
        :filterable="filterable" :filter-placeholder="placeholder" :props="transfer.props" />

    <!-------- 使用示例 开始-------->
    <!-- <MyTransfer v-model="saveCommon.data.sceneIdArr"
        :api="{ code: 'auth/scene/list', param: { field: ['id', 'sceneName'] } }" /> -->
    <!-------- 使用示例 结束-------->
</template>