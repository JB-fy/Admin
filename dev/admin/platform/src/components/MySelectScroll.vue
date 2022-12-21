<script setup lang="ts">
const { t } = useI18n()

const props = defineProps({
    modelValue: {
        type: [String, Number, Array],
        //required: true,
    },
    defaultOption: {    //是否含有默认值。格式：[{ value: string | Number, label: string },...]
        type: Array,
        default: []
    },
    apiFunc: {   //接口函数。函数的返回值格式：[{ value: string|number, label: string },...]
        type: Function,
        required: true
    },
    apiParam: { //接口函数所需参数。格式：{ field: srting[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number,limit: number }
        type: Object,
        required: true,
    },
    apiSelectedField: {    //有初始值时，用于查询条件的字段
        type: String,
        //default: 'id' //默认为props.apiParam.field[0]
    },
    apiSearchField: {  //远程查询时，用于查询条件的字段
        type: String,
        //required: true,
        //default: 'id' //默认为props.apiParam.field[1]
    },
    placeholder: {
        type: String,
        //default: t('common.tip.pleaseSelect') //defineProps会被提取到setup外执行，故这里t函数是不存在的
        //default: i18n.global.t('common.tip.pleaseSelect') //动态切换时不会改变，需直接写在html中（系统默认页面刷新会切换）
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
    value: props.modelValue,
    /* value: computed({
        get: () => {
            return props.modelValue
        },
        set: (val) => {
            emits('update:modelValue', val)
        }
    }), */
    options: [...props.defaultOption] as { value: string | number, label: string }[],
    loading: computed((): boolean => {
        //ElSelectV2的loading属性建议在远程数据全部加载时启用
        if (select.param.page == 1 && select.param.limit == 0) {
            return select.apiLoading
        }
        return false
    }),
    isEnd: false,
    apiLoading: false,
    param: {
        field: [],
        where: {} as { [propName: string]: any },
        order: { id: 'desc' },
        page: 1,
        limit: 10,
        ...props.apiParam
    },
    resetOptions: () => {
        select.options = [...props.defaultOption] as any
    },
    setOptions: () => {
        if (select.apiLoading) {
            return
        }
        if (select.isEnd) {
            return
        }
        select.apiLoading = true
        props.apiFunc(select.param).then((options: []) => {
            if (select.param.limit === 0 || options.length < select.param.limit) {
                select.isEnd = true
            }
            select.options = select.options.concat(options ?? [])
        }).catch(() => {
        }).finally(() => {
            select.apiLoading = false
        })
    },
    visibleChange: (val: boolean) => {
        //if (val && select.options.length == 0) {    //只在首次打开加载。但用户切换页面做数据变动，再返回时，需要刷新页面清理缓存才能获取最新数据
        if (val) {  //每次打开都加载
            delete select.param.where[props.apiSearchField ?? props.apiParam.field[1]]
            select.resetOptions()
            select.param.page = 1
            select.isEnd = false
            select.setOptions()
        }
    },
    remoteMethod: (keyword: string) => {
        if (keyword) {
            select.param.where[props.apiSearchField ?? props.apiParam.field[1]] = keyword
        } else {
            delete select.param.where[props.apiSearchField ?? props.apiParam.field[1]]
        }
        select.resetOptions()
        select.param.page = 1
        select.isEnd = false
        select.setOptions()
    },
    change: (val: any) => {
        emits('update:modelValue', val)
    },
})
if (props.modelValue) {
    select.param.where[props.apiSelectedField ?? props.apiParam.field[0]] = props.modelValue
    select.setOptions()
    delete select.param.where[props.apiSelectedField ?? props.apiParam.field[0]]
}
watch(() => props.modelValue, (newVal: any, oldVal: any) => {
    console.log(newVal)
    select.value = newVal
})
/* watch(() => select.value, (newVal: any, oldVal: any) => {
    console.log(newVal)
    console.log(oldVal)
    console.log(props.modelValue)
    console.log(props.apiSearchField??props.apiParam.field[1])
    if (newVal && !oldVal) {
        select.param.where[props.apiSelectedField ?? props.apiParam.field[0]] = newVal
        select.setOptions()
    }
    delete select.param.where[props.apiSelectedField ?? props.apiParam.field[0]]
}) */

const vScroll = {
    updated: () => {
        /* const dropId = el.querySelector('.el-tooltip__trigger').getAttribute('aria-describedby')
        if (!dropId) {
            return
        }
        const scrollDom = document.getElementById(dropId).querySelector('.el-select-dropdown__list') */
        const scrollDom = select.ref.popperRef.querySelector('.el-select-dropdown__list')
        if (scrollDom) {
            //scrollDom.scrollBy(0, 400)
            const scrollFunc = () => {
                console.log(scrollDom.scrollTop)
                if (scrollDom.scrollHeight - scrollDom.scrollTop <= scrollDom.clientHeight) {
                    select.param.page++
                    select.setOptions()
                }
            }
            scrollDom.removeEventListener('scroll', scrollFunc)
            scrollDom.addEventListener('scroll', scrollFunc)
        }
    }
}
</script>

<template>
    <ElSelectV2 :ref="(el: any) => { select.ref = el }" v-model="select.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :options="select.options" :clearable="clearable"
        :filterable="filterable" @visible-change="select.visibleChange" :remote="true"
        :remote-method="select.remoteMethod" :loading="select.loading" v-scroll @change="select.change"
        :validate-event="false" />
    <!-- <ElSelectV2 :ref="(el: any) => { select.ref = el }" v-model="select.value"
        :placeholder="placeholder ?? t('common.tip.pleaseSelect')" :options="select.options" :clearable="clearable"
        :filterable="filterable" @visible-change="select.visibleChange" :remote="true"
        :remote-method="select.remoteMethod" :loading="select.loading" v-scroll @change="select.change"
        :validate-event="false" /> -->
</template>