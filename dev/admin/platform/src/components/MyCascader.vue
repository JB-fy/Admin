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
    apiCode: {  //接口标识。参考common/utils/common.js文件内request方法的参数说明
        type: String,
        required: true,
    },
    apiParam: { //接口函数所需参数。格式：{ field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number }。其中field内第0个和第1个字段默认用于cascader.api的dataToOptions，selectedField，searchField三个属性。使用时请注意，否则需要设置props的apiDataToOptions，apiSelectedField，apiSearchField三个参数
        type: Object,
        required: true,
    },
    apiDataToOptions: { //接口返回数据转换方法。返回值格式：[{ value: string|number, label: string },...]
        type: Function
    },
    apiSelectedField: { //当组件初始化，modelValue有初始值时，接口参数where中使用的字段名。默认：props.apiParam.field[0]
        type: String
    },
    apiSearchField: {   //当用户输入关键字做查询时，接口参数where中使用的字段名。默认：props.apiParam.field[1]
        type: String
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
    lazy: {
        type: Boolean,
        default: true
    },
    disabled: {
        type: Boolean,
        default: false
    },
    multiple: {
        type: Boolean,
        default: false
    },
    collapseTags: {
        type: Boolean,
        default: true
    },
    collapseTagsTooltip: {
        type: Boolean,
        default: true
    },
    multipleLimit: {
        type: Number,
        default: 0
    },
})

const emits = defineEmits(['update:modelValue'])
const cascader = reactive({
    ref: null as any,
    value: computed({
        get: () => {
            return props.modelValue
        },
        set: (val) => {
            console.log(val)
            emits('update:modelValue', val)
        }
    }),
    options: [...props.defaultOptions] as { value: string | number, label: string }[],
    initOptions: () => {
        cascader.api.param.where[cascader.api.selectedField] = props.modelValue
        cascader.api.addOptions()
        delete cascader.api.param.where[cascader.api.selectedField]
    },
    resetOptions: () => {
        cascader.options = [...props.defaultOptions] as any
        cascader.api.param.page = 1
        cascader.api.isEnd = false
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
        isEnd: false,
        loading: false,
        param: computed((): { field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number } => {
            return {
                field: [],
                where: {} as { [propName: string]: any },
                order: { id: 'desc' },
                page: 1,
                limit: 0,
                ...props.apiParam
            }
        }),
        dataToOptions: computed(() => {
            return props.apiDataToOptions ? props.apiDataToOptions : (res: any) => {
                const options: { value: any, label: any }[] = []
                res.data.list.forEach((item: any) => {
                    options.push({
                        value: item[cascader.api.param.field[0]],
                        label: item[cascader.api.param.field[1]]
                    })
                })
                return options
            }
        }),
        selectedField: computed((): string => {
            if (props.apiSelectedField) {
                return props.apiSelectedField
            }
            if (props.apiParam.field[0] == 'id') {
                return props.multiple ? 'idArr' : 'id'
            }
            return props.apiParam.field[0]
        }),
        searchField: computed((): string => {
            return props.apiSearchField ?? props.apiParam.field[1]
        }),
        addOptions: () => {
            if (cascader.api.loading) {
                return
            }
            if (cascader.api.isEnd) {
                return
            }
            cascader.api.loading = true
            request(props.apiCode, cascader.api.param).then((res) => {
                const options = cascader.api.dataToOptions(res)
                cascader.options = cascader.options.concat(options ?? [])
                if (cascader.api.param.limit === 0 || options.length < cascader.api.param.limit) {
                    cascader.api.isEnd = true
                }
            }).catch(() => {
            }).finally(() => {
                cascader.api.loading = false
            })
        },
    },
    visibleChange: (val: boolean) => {
        //if (val && cascader.options.length == props.defaultOptions.length) {    //只在首次打开加载。但用户切换页面做数据变动，再返回时，需要刷新页面清理缓存才能获取最新数据
        if (val) {  //每次打开都加载
            delete cascader.api.param.where[cascader.api.searchField]
            cascader.resetOptions()
            cascader.api.addOptions()
        }
    },
    lazyLoad: (node: any, resolve: any) => {
        console.log(node)
        const { level } = node
        const options = Array.from({ length: level + 1 }).map((item) => ({
            value: 2,
            label: 'Option',
            leaf: level >= 2,   //用于终止继续加载
        }))
        resolve(options)
        /* if (keyword) {
            cascader.api.param.where[cascader.api.searchField] = keyword
        } else {
            delete cascader.api.param.where[cascader.api.searchField]
        }
        cascader.resetOptions()
        cascader.api.addOptions() */
    }
})
//组件创建时，如有初始值，需初始化options。
if (props.filterable) {
    cascader.initOptions()
}
</script>

<template>
    <ElCascader v-if="filterable" :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :options="cascader.options" :clearable="clearable"
        :filterable="filterable" @visible-change="cascader.visibleChange"
        :props="{ multiple: multiple, emitPath: false }" :disabled="disabled" />
    <ElCascader v-else :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :clearable="clearable"
        @visible-change="cascader.visibleChange"
        :props="{ multiple: multiple, emitPath: false, lazy: lazy, lazyLoad: cascader.lazyLoad }"
        :disabled="disabled" />
</template>