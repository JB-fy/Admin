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
    selectedField: {    //有初始值时，用于查询条件的字段
        type: String,
        default: 'id'
    },
    searchField: {  //远程查询时，用于查询条件的字段
        type: String,
        required: true
    },
    apiFunc: {   //接口函数。函数的返回值格式：[{ value: string | Number, label: string },...]
        type: Function,
        required: true
    },
    apiParam: { //接口函数所需参数
        type: Object,
        required: true,
    },
    placeholder: {
        type: String,
        //default: t('common.tip.pleaseSelect') //defineProps会被提取到setup外执行，故这里t函数是不存在的
        //default: i18n.global.t('common.tip.pleaseSelect') //切换时不会改变
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
    options: [...props.defaultOption] as { value: string | Number, label: string }[],
    loading: false,
    isEnd: false,
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
        if (select.loading) {
            return
        }
        if (select.isEnd) {
            return
        }
        select.loading = true
        props.apiFunc(select.param).then((options: []) => {
            if (select.param.limit === 0 || options.length < select.param.limit) {
                select.isEnd = true
            }
            select.options = select.options.concat(options ?? [])
        }).catch(() => {
        }).finally(() => {
            select.loading = false
        })
    },
    visibleChange: (val: boolean) => {
        //if (val && select.options.length == 0) {    //只在首次打开加载。但用户切换页面做数据变动，再返回时，需要刷新页面清理缓存才能获取最新数据
        if (val) {  //每次打开都加载
            delete select.param.where[props.searchField]
            select.resetOptions()
            select.param.page = 1
            select.isEnd = false
            select.setOptions()
        }
    },
    remoteMethod: (keyword: string) => {
        if (keyword) {
            select.param.where[props.searchField] = keyword
        } else {
            delete select.param.where[props.searchField]
        }
        select.resetOptions()
        select.param.page = 1
        select.isEnd = false
        select.setOptions()
    },
    change: (val: any) => {
        emits('update:modelValue', val)
        select.value = val
    },
})
if (props.modelValue) {
    select.param.where[props.selectedField] = props.modelValue
    select.setOptions()
    delete select.param.where[props.selectedField]
}
/* watch(() => select.value, (newValue: any, oldValue: any) => {
    console.log(newValue)
    console.log(oldValue)
    console.log(props.modelValue)
    console.log(props.searchField)
    if (newValue && !oldValue) {
        select.param.where[props.selectedField] = newValue
        select.setOptions()
    }
    delete select.param.where[props.selectedField]
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
            //scrollDom.removeEventListener('scroll', scrollFunc)
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
    <!-- <ElSelectV2 :ref="(el: any) => { select.ref = el }" v-model="select.value" :placeholder="placeholder"
        :options="select.options" :clearable="clearable" :filterable="filterable" @visible-change="select.visibleChange"
        :remote="true" :remote-method="select.remoteMethod" :loading="select.loading" v-scroll /> -->
</template>