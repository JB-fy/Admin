<script setup lang="ts">
const props = defineProps({
    modelValue: {
        type: Array,
        //required: true,
    },
    defaultOptions: {   //选项初始默认值。格式：[{ value: string | number, label: string },...]
        type: Array,
        default: []
    },
    /**
     * 接口。格式：{ code: string, param: object, dataToOptions: function, selectedField: string, searchField: string }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number }。其中field内第0，1字段默认用于transfer.api的dataToOptions，selectedField，searchField属性，使用时请注意。或直接在props.api中设置对应参数
     *      dataToOptions：非必须。接口返回数据转换方法。返回值格式：[{ value: string|number, label: string },...]
     *      selectedField：非必须。当组件初始化，modelValue有初始值时，接口参数where中使用的字段名。默认：props.api.param.field[0]
     *      searchField：非必须。当用户输入关键字做查询时，接口参数where中使用的字段名。默认：props.api.param.field[1]
     */
    api: {
        type: Object,
        required: true,
    },
    placeholder: {
        type: String
    },
    clearable: {
        type: Boolean,
        default: true
    },
    filterable: {
        type: Boolean,
        default: true
    },
    remote: {
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
    options: [...props.defaultOptions] as { value: string | number, label: string }[],
    initOptions: () => {
        transfer.api.param.where[transfer.api.selectedField] = props.modelValue
        transfer.api.addOptions()
        delete transfer.api.param.where[transfer.api.selectedField]
    },
    resetOptions: () => {
        transfer.options = [...props.defaultOptions] as any
        transfer.api.param.page = 1
        transfer.api.isEnd = false
    },
    loading: computed((): boolean => {
        //ElSelectV2的loading属性建议在远程数据全部加载时使用，其他情况下都为false。
        //例如：分页加载时使用会导致因出现加载中元素节点而导致滚动条节点丢失再出现。虽然可根据这个重新处理滚动事件，但视觉效果也不好
        if (transfer.api.param.page == 1 && transfer.api.param.limit == 0) {
            return transfer.api.loading
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
                limit: useSettingStore().scrollSize,
                ...props.api.param
            }
        }),
        dataToOptions: computed(() => {
            return props.api.dataToOptions ? props.api.dataToOptions : (res: any) => {
                const options: { value: any, label: any }[] = []
                res.data.list.forEach((item: any) => {
                    options.push({
                        value: item[transfer.api.param.field[0]],
                        label: item[transfer.api.param.field[1]]
                    })
                })
                return options
            }
        }),
        selectedField: computed((): string => {
            if (props.api.selectedField) {
                return props.api.selectedField
            }
            if (props.api.param.field[0] == 'id') {
                return props.multiple ? 'idArr' : 'id'
            }
            return props.api.param.field[0]
        }),
        searchField: computed((): string => {
            return props.api.searchField ?? props.api.param.field[1]
        }),
        getOptions: async () => {
            if (transfer.api.loading) {
                return
            }
            if (transfer.api.isEnd) {
                return
            }
            transfer.api.loading = true
            let options = []
            try {
                const res = await request(props.api.code, transfer.api.param)
                options = transfer.api.dataToOptions(res)
                if (transfer.api.param.limit === 0 || options.length < transfer.api.param.limit) {
                    transfer.api.isEnd = true
                }
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
    visibleChange: (val: boolean) => {
        //if (val && transfer.options.length == props.defaultOptions.length) {    //只在首次打开加载。但用户切换页面做数据变动，再返回时，需要刷新页面清理缓存才能获取最新数据
        if (val) {  //每次打开都重新加载
            delete transfer.api.param.where[transfer.api.searchField]
            transfer.resetOptions()
            transfer.api.addOptions()
        }
    },
    remoteMethod: (keyword: string) => {
        if (keyword) {
            transfer.api.param.where[transfer.api.searchField] = keyword
        } else {
            delete transfer.api.param.where[transfer.api.searchField]
        }
        transfer.resetOptions()
        transfer.api.addOptions()
    }
})
//组件创建时，如有初始值，需初始化options
if ((Array.isArray(props.modelValue) && props.modelValue.length) || props.modelValue) {
    transfer.initOptions()
}
/**
 * 因上面的代码只在组件创建时初始化一次，所以当表的不同记录先后点击编辑按钮时，第二次编辑不会初始化options。
 *  解决方法
 *      1：在组件使用的地方用v-if来重新创建组件
 *          优点：适用于各种复杂情况
 *      2：参考下面的监听器代码
 *          优点：可减少对服务器的请求。当切换记录编辑时，如果两条记录数据是一样，不用重新请求接口初始化options
 *          缺点：必须设置validateEvent为false，否则当点击编辑，再点击新增，会直接提示错误信息
 */
/* watch(() => props.modelValue, (newVal: any, oldVal: any) => {
    if (Array.isArray(props.modelValue)) {
        if (props.modelValue.length && transfer.options.filter((item) => {
            //return (<string[] | number[]>props.modelValue).indexOf(item.value) !== -1
            return (<any>props.modelValue).indexOf(item.value) !== -1
        }).length !== props.modelValue.length) {
            transfer.resetOptions()
            transfer.initOptions()
        }
    } else if (props.modelValue && transfer.options.findIndex((item) => {
        return item.value == props.modelValue
    }) === -1) {
        transfer.resetOptions()
        transfer.initOptions()
    }
}) */

//滚动方法。需要写外面，否则无法通过removeEventListener移除事件
const scrollFunc = (event: any) => {
    if (event.target.scrollTop > 0 && event.target.scrollHeight - event.target.scrollTop <= event.target.clientHeight) {
        transfer.api.param.page++
        transfer.api.addOptions()
    }
}
/* //分页加载要使用动态设置transfer.loading时，使用这个方式设置滚动事件
watch(() => transfer.loading, (newVal: any, oldVal: any) => {
    if (transfer.loading === false) { */
watch(() => transfer.options, (newVal: any, oldVal: any) => {
    if (transfer.options.length) {
        nextTick(() => {
            /* const dropId = el.querySelector('.el-tooltip__trigger').getAttribute('aria-describedby')
            if (!dropId) {
                return
            }
            const scrollDom = document.getElementById(dropId).querySelector('.el-select-dropdown__list') */
            const scrollDom = transfer.ref.popperRef.querySelector('.el-select-dropdown__list')
            if (scrollDom) {
                scrollDom.removeEventListener('scroll', scrollFunc)
                scrollDom.addEventListener('scroll', scrollFunc)
            }
        })
    }
})
</script>

<template>
    <ElTransfer :ref="(el: any) => { transfer.ref = el }" v-model="transfer.value"
        :filter-placeholder="placeholder"
        :options="transfer.options" :clearable="clearable" :filterable="filterable"
        @visible-change="transfer.visibleChange" :remote="remote" :remote-method="transfer.remoteMethod"
        :loading="transfer.loading" :disabled="disabled" />

    <!-------- 使用示例 开始-------->
    <!-------- 使用示例 结束-------->
</template>