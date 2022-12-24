<script setup lang="ts">
const { t } = useI18n()

const props = defineProps({
    modelValue: {
        type: [String, Number, Array],
        //required: true,
    },
    defaultOptions: {   //选项初始默认值。格式：[{ value: string | number, label: string },...]
        type: Array,
        default: []
    },
    /**
     * 接口。格式：{ code: string, param: object, dataToOptions: function }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number }。其中field内第0个和第1个字段默认用于select.api的dataToOptions，selectedField，searchField三个属性。使用时请注意，否则需要设置props.api中对应的三个参数
     *      dataToOptions：非必须。接口返回数据转换方法。返回值格式：[{ value: string|number, label: string },...]
     */
    api: {
        type: Object,
        required: true,
    },
    placeholder: {
        type: String,
        //default: t('common.tip.pleaseSelect') //defineProps会被提取到setup外执行，故这里t函数是不存在的
        //default: i18n.global.t('common.tip.pleaseSelect') //动态切换时不会改变，需直接写在html中（框架语言切换默认会做页面刷新）
    },
    clearable: {
        type: Boolean,
        default: true
    },
    filterable: {
        type: Boolean,
        default: true
    },
    disabled: {
        type: Boolean,
        default: false
    },
    props: {
        type: Object,
        default: {}
    },
    collapseTags: {
        type: Boolean,
        default: true
    },
    collapseTagsTooltip: {
        type: Boolean,
        default: true
    },
})

const emits = defineEmits(['update:modelValue'])
const cascader = reactive({
    ref: null as any,
    value: computed({
        get: (): any => {
            return props.modelValue
        },
        set: (val) => {
            console.log(val)
            emits('update:modelValue', val)
        }
    }),
    options: [...props.defaultOptions] as { value: string | number, label: string }[],
    props: {
        multiple: false,
        checkStrictly: false,
        emitPath: false,
        lazy: false,
        lazyLoad: (node: any, resolve: any) => {
            if (node.level == 0) {
                cascader.api.param.where['pid'] = 0
            } else {
                cascader.api.param.where['pid'] = node.data.id
            }
            cascader.api.getOptions().then((options) => {
                if (options.length === 0) {
                    node.data.leaf = true
                }
                //cascader.options = cascader.options.concat(options ?? [])
                resolve(options)
            }).catch((error) => { })
            delete cascader.api.param.where['pid']
        },
        value: 'id',
        label: 'menuName',
        //value: 'value',
        //label: 'label',
        //children: 'children',
        //disabled: 'disabled',
        //leaf: 'leaf', //动态加载时用于终止继续加载。当checkStrictly为false时，该字段必须有，否则选中后值为null
        ...props.props
    },
    initOptions: () => {
        cascader.api.addOptions()
    },
    resetOptions: () => {
        cascader.options = [...props.defaultOptions] as any
        cascader.api.param.page = 1
    },
    loading: computed((): boolean => {
        //ElSelectV2的loading属性建议在远程数据全部加载时使用，其他情况下都为false。
        //例如：分页加载时使用会导致因出现加载中元素节点而导致滚动条节点丢失再出现。虽然可根据这个重新处理滚动事件，但视觉效果也不好
        if (cascader.api.param.page == 1 && cascader.api.param.limit == 0) {
            return cascader.api.loading
        }
        return false
    }),
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
                return cascader.props.lazy ? res.data.list : res.data.tree
                /* const options: { value: any, label: any }[] = []
                if (cascader.props.lazy) {
                    res.data.list.forEach((item: any) => {
                        options.push({
                            value: item[cascader.api.param.field[0]],
                            label: item[cascader.api.param.field[1]]
                        })
                    })
                } else {
                    res.data.tree.forEach((item: any) => {
                        options.push({
                            value: item[cascader.api.param.field[0]],
                            label: item[cascader.api.param.field[1]]
                        })
                    })
                }
                return options */
            }
        }),
        getOptions: async () => {
            if (cascader.api.loading) {
                return
            }
            cascader.api.loading = true
            let options = []
            try {
                const res = await request(props.api.code, cascader.api.param)
                options = cascader.api.dataToOptions(res)
            } catch (error) { }
            cascader.api.loading = false
            return options
        },
        addOptions: () => {
            cascader.api.getOptions().then((options) => {
                if (options.length) {
                    cascader.options = cascader.options.concat(options ?? [])
                }
            }).catch((error) => { })
        },
    },
    visibleChange: (val: boolean) => {
        if (val) {  //每次打开都重新加载
            cascader.resetOptions()
            cascader.api.addOptions()
        }
    }
})
//组件创建时，如有初始值，需初始化options。
if (!cascader.props.lazy && ((Array.isArray(props.modelValue) && props.modelValue.length) || props.modelValue)) {
    cascader.initOptions()
}
</script>

<template>
    <ElCascader v-if="cascader.props.lazy" :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :clearable="clearable" :props="cascader.props"
        :disabled="disabled" />
    <ElCascader v-else :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :clearable="clearable" :options="cascader.options"
        :filterable="filterable" @visible-change="cascader.visibleChange" :props="cascader.props"
        :disabled="disabled" />
</template>