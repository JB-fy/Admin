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
    apiCode: {  //格式：接口标识
        type: String,
        required: true,
    },
    apiParam: { //接口函数所需参数。格式：{ field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number }。其中field内第0个和第1个字段默认用于select.api的dataToOptions，selectedField，searchField三个属性。使用时请注意，否则需要设置props的apiDataToOptions，apiSelectedField，apiSearchField三个参数
        type: Object,
        required: true,
    },
    apiDataToOptions: { //接口返回数据转换方法。返回值格式：[{ value: string|number, label: string },...]
        type: Function
    },
    apiSelectedField: { //当组件初始化，modelValue有初始值时，接口参数where中使用的字段名。默认为props.apiParam.field[0]
        type: String
    },
    apiSearchField: {   //当用户输入关键字做查询时，接口参数where中使用的字段名。默认为props.apiParam.field[1]
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
    /* multiple: {
      type: Boolean,
      default: false
    }, */
})

const emits = defineEmits(['update:modelValue'])
const select = reactive({
    ref: null as any,
    value: computed({
        get: () => {
            return props.modelValue
        },
        set: (val) => {
            emits('update:modelValue', val)
        }
    }),
    options: [...props.defaultOptions] as { value: string | number, label: string }[],
    loading: computed((): boolean => {
        //ElSelectV2的loading属性建议在远程数据全部加载时使用，其他情况下都为false。
        //例如：分页加载时使用会导致因出现加载中元素节点而导致滚动条节点丢失再出现。虽然可根据这个重新处理滚动事件，但视觉效果也不好
        if (select.api.param.page == 1 && select.api.param.limit == 0) {
            return select.api.loading
        }
        return false
    }),
    api: {
        isEnd: false,
        loading: false,
        param: {
            field: [],
            where: {} as { [propName: string]: any },
            order: { id: 'desc' },
            page: 1,
            limit: 10,
            ...props.apiParam
        },
        dataToOptions: props.apiDataToOptions ? props.apiDataToOptions : (res: any) => {
            const options: { value: any, label: any }[] = []
            res.data.list.forEach((item: any) => {
                options.push({
                    value: item[select.api.param.field[0]],
                    label: item[select.api.param.field[1]]
                })
            })
            return options
        },
        selectedField: props.apiSelectedField ?? props.apiParam.field[0],
        searchField: props.apiSearchField ?? props.apiParam.field[1],
        addOptions: () => {
            if (select.api.loading) {
                return
            }
            if (select.api.isEnd) {
                return
            }
            select.api.loading = true
            request(props.apiCode, select.api.param).then((res) => {
                const options = select.api.dataToOptions(res)
                select.options = select.options.concat(options ?? [])
                if (select.api.param.limit === 0 || options.length < select.api.param.limit) {
                    select.api.isEnd = true
                }
            }).catch(() => {
            }).finally(() => {
                select.api.loading = false
            })
        },
    },
    resetOptions: () => {
        select.options = [...props.defaultOptions] as any
        select.api.param.page = 1
        select.api.isEnd = false
    },
    visibleChange: (val: boolean) => {
        //if (val && select.options.length == props.defaultOptions.length) {    //只在首次打开加载。但用户切换页面做数据变动，再返回时，需要刷新页面清理缓存才能获取最新数据
        if (val) {  //每次打开都加载
            delete select.api.param.where[select.api.searchField]
            select.resetOptions()
            select.api.addOptions()
        }
    },
    remoteMethod: (keyword: string) => {
        if (keyword) {
            select.api.param.where[select.api.searchField] = keyword
        } else {
            delete select.api.param.where[select.api.searchField]
        }
        select.resetOptions()
        select.api.addOptions()
    }
})
if (props.modelValue && select.options.findIndex((item) => {
    return item.value == props.modelValue
}) === -1) {
    select.resetOptions()
    select.api.param.where[select.api.selectedField] = props.modelValue
    select.api.addOptions()
    delete select.api.param.where[select.api.selectedField]
}

//滚动方法。需要写外面，否则无法通过移除事件removeEventListener移除
const scrollFunc = (event: any) => {
    if (event.target.scrollTop > 0 && event.target.scrollHeight - event.target.scrollTop <= event.target.clientHeight) {
        select.api.param.page++
        select.api.addOptions()
    }
}
/* //分页加载要使用动态设置select.loading时，使用这个方式设置滚动事件
watch(() => select.loading, (newVal: any, oldVal: any) => {
    if (select.loading === false) { */
watch(() => select.options, (newVal: any, oldVal: any) => {
    if (select.options.length) {
        nextTick(() => {
            /* const dropId = el.querySelector('.el-tooltip__trigger').getAttribute('aria-describedby')
            if (!dropId) {
                return
            }
            const scrollDom = document.getElementById(dropId).querySelector('.el-select-dropdown__list') */
            const scrollDom = select.ref.popperRef.querySelector('.el-select-dropdown__list')
            if (scrollDom) {
                scrollDom.removeEventListener('scroll', scrollFunc)
                scrollDom.addEventListener('scroll', scrollFunc)
            }
        })
    }
})
</script>

<template>
    <ElSelectV2 :ref="(el: any) => { select.ref = el }" v-model="select.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :options="select.options" :clearable="clearable"
        :filterable="filterable" @visible-change="select.visibleChange" :remote="true"
        :remote-method="select.remoteMethod" :loading="select.loading" :validate-event="false" />
</template>