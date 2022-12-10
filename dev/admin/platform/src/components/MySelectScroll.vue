<script setup lang="ts">
const { t } = useI18n()

const props = defineProps({
    /* value: {
        type: [String, Number, Array],
        //required: true,
        default: undefined
    }, */
    selectedField: {    //有初始值时，用于查询条件的字段
        type: String,
        default: 'id'
    },
    searchField: {
        type: String,
        required: true
    },
    apiFunc: {   //接口函数
        type: Function,
        required: true
    },
    apiParam: { //接口函数所需参数
        type: Object,
        required: true,
    },
    placeholder: {
        type: String,
        default: '请选择'
        //default: t('view.auth.scene.sceneId')
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

const vScroll = {
    updated: () => {
        /* const dropId = el.querySelector('.el-tooltip__trigger').getAttribute('aria-describedby')
        if (!dropId) {
            return
        }
        const scrollDom = document.getElementById(dropId).querySelector('.el-select-dropdown__list') */
        const scrollDom = select.ref.popperRef.querySelector('.el-select-dropdown__list')
        console.log(scrollDom)
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
const select = reactive({
    ref: null as any,
    value: undefined,
    options: [],
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
            select.options = []
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
        select.options = []
        select.param.page = 1
        select.isEnd = false
        select.setOptions()
    }
})
watch(() => select.value, (newValue: any, oldValue: any) => {
    console.log(oldValue)
    if (newValue > 0 && !oldValue) {
        select.param.where[props.selectedField] = newValue
        select.setOptions()
    }
})
</script>

<template>
    <ElSelectV2 :ref="(el: any) => { select.ref = el }" v-model="select.value" :placeholder="placeholder"
        :options="select.options" :clearable="clearable" :filterable="filterable" @visible-change="select.visibleChange"
        :remote="true" :remote-method="select.remoteMethod" :loading="select.loading" v-scroll />
    <!-- <ElSelectV2 :ref="(el: any) => { select.ref = el }" v-model="select.value" :placeholder="placeholder"
        :options="select.options" :clearable="clearable" :filterable="filterable" @visible-change="select.visibleChange"
        :remote="true" :remote-method="select.remoteMethod" :loading="select.loading" v-scroll /> -->
</template>