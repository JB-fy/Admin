<!-------- 使用示例 开始-------->
<!-- <my-select v-model="saveForm.data.sceneId" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list' }" />

<my-select v-model="queryCommon.data.sceneId" :placeholder="t('auth.role.name.sceneId')"
    :defaultOptions="[{ value: 0, label: t('common.name.allTopLevel') }]"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/scene/list', param: { field: ['id', 'sceneName'] } }" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
const slots = useSlots()
const props = defineProps({
    modelValue: {
        type: [String, Number, Array],
    },
    defaultOptions: {
        //选项初始默认值。格式：[{ value: any, label: any },...]
        type: Array,
        default: () => [],
    },
    /**
     * 接口。格式：{ code: string, param: object, transform: function, selectedField: string, searchField: string }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ filter: { [propName: string]: any }, field: string[], sort: string, page: number, limit: number }。其中field内第0，1字段默认用于select.api的transform，selectedField，searchField属性，使用时请注意。或直接在props.api中设置对应参数
     *      transform：非必须。接口返回数据转换方法。返回值格式：[{ value: any, label: any },...]
     *      selectedField：非必须。当组件初始化，modelValue有初始值时，接口参数filter中使用的字段名。默认：props.api.param.field[0]
     *      searchField：非必须。当用户输入关键字做查询时，接口参数filter中使用的字段名。默认：props.api.param.field[1]
     */
    api: {
        type: Object,
        required: true,
    },
    placeholder: {
        type: String,
    },
    clearable: {
        type: Boolean,
        default: true,
    },
    filterable: {
        type: Boolean,
        default: true,
    },
    remote: {
        type: Boolean,
        default: true,
    },
    disabled: {
        type: Boolean,
        default: false,
    },
    multiple: {
        type: Boolean,
        default: false,
    },
    multipleLimit: {
        type: Number,
        default: 0,
    },
    collapseTags: {
        type: Boolean,
        default: true,
    },
    collapseTagsTooltip: {
        type: Boolean,
        default: true,
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const select = reactive({
    ref: null as any,
    value: computed({
        get: () => {
            return props.modelValue
        },
        set: (val) => {
            emits('update:modelValue', val)
            emits(
                'change',
                props.multiple
                    ? select.options.filter((item) => {
                          return (val as any).indexOf(item.value) !== -1
                      })
                    : select.options.find((item) => item.value == val)
            )
        },
    }),
    options: [...props.defaultOptions] as { value: any; label: any; [propName: string]: any }[],
    initOptions: () => {
        select.api.param.filter[select.api.selectedField] = props.modelValue
        select.api.addOptions()
        delete select.api.param.filter[select.api.selectedField]
    },
    resetOptions: () => {
        select.options = [...props.defaultOptions] as { value: any; label: any; [propName: string]: any }[]
        select.api.param.page = 1
        select.api.isEnd = false
    },
    loading: computed((): boolean => {
        //ElSelectV2的loading属性建议在远程数据全部加载时使用，其它情况下都为false。
        //例如：分页加载时使用会导致因出现加载中元素节点而导致滚动条节点丢失再出现。虽然可根据这个重新处理滚动事件，但视觉效果也不好
        if (select.api.param.page == 1 && select.api.param.limit == 0) {
            return select.api.loading
        }
        return false
    }),
    api: {
        isEnd: false,
        loading: false,
        param: computed((): { filter: { [propName: string]: any }; field: string[]; sort: string; page: number; limit: number } => {
            return {
                filter: {} as { [propName: string]: any },
                field: ['id', 'label'],
                sort: 'id desc',
                page: 1,
                limit: useSettingStore().scrollSize,
                ...(props.api?.param ?? {}),
            }
        }),
        transform: computed(() => {
            return props.api.transform
                ? props.api.transform
                : (res: any) => {
                      const options: { value: any; label: any }[] = []
                      res.data.list.forEach((item: any) => {
                          options.push({
                              ...item,
                              value: item[select.api.param.field[0]],
                              label: item[select.api.param.field[1]],
                          })
                      })
                      return options
                  }
        }),
        selectedField: computed((): string => {
            if (props.api.selectedField) {
                return props.api.selectedField
            }
            if (select.api.param.field[0] == 'id') {
                return props.multiple ? 'id_arr' : 'id'
            }
            return select.api.param.field[0]
        }),
        searchField: computed((): string => {
            return props.api.searchField ?? select.api.param.field[1]
        }),
        getOptions: async () => {
            if (select.api.loading) {
                return
            }
            if (select.api.isEnd) {
                return
            }
            select.api.loading = true
            let options = []
            try {
                const res = await request(props.api.code, select.api.param)
                options = select.api.transform(res)
                if (select.api.param.limit === 0 || options.length < select.api.param.limit) {
                    select.api.isEnd = true
                }
            } catch (error) {}
            select.api.loading = false
            return options
        },
        addOptions: () => {
            select.api.getOptions().then((options) => {
                if (options?.length) {
                    select.options = select.options.concat(options ?? [])
                }
            })
        },
    },
    visibleChange: (val: boolean) => {
        //if (val && select.options.length == props.defaultOptions.length) {    //只在首次打开加载。但用户切换页面做数据变动，再返回时，需要刷新页面清理缓存才能获取最新数据
        if (val) {
            //每次打开都重新加载
            delete select.api.param.filter[select.api.searchField]
            select.resetOptions()
            select.api.addOptions()
        }
    },
    remoteMethod: (label: string) => {
        if (label) {
            select.api.param.filter[select.api.searchField] = label
        } else {
            delete select.api.param.filter[select.api.searchField]
        }
        select.resetOptions()
        select.api.addOptions()
    },
})
//组件创建时，如有初始值，需初始化options
if ((Array.isArray(props.modelValue) && props.modelValue.length) || props.modelValue) {
    select.initOptions()
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
        if (props.modelValue.length && select.options.filter((item) => {
            //return (<string[] | number[]>props.modelValue).indexOf(item.value) !== -1
            return (<any>props.modelValue).indexOf(item.value) !== -1
        }).length !== props.modelValue.length) {
            select.resetOptions()
            select.initOptions()
        }
    } else if (props.modelValue && select.options.findIndex((item) => {
        return item.value == props.modelValue
    }) === -1) {
        select.resetOptions()
        select.initOptions()
    }
}) */

//滚动方法。需要写外面，否则无法通过removeEventListener移除事件
const scrollFunc = (event: any) => {
    if (event.target.scrollTop > 0 && event.target.scrollHeight - event.target.scrollTop <= event.target.clientHeight) {
        select.api.param.page++
        select.api.addOptions()
    }
}
/* //分页加载，用到动态设置select.loading时，用这个方式设置滚动事件
watch(() => select.loading, (newVal: any, oldVal: any) => {
    if (select.loading === false) { */
watch(
    () => select.options,
    () => {
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
    }
)

//暴露组件属性给父组件
defineExpose({
    options: computed(() => {
        return select.options
    }),
})
</script>

<template>
    <el-select-v2
        :ref="(el: any) => select.ref = el"
        v-model="select.value"
        :placeholder="placeholder"
        :options="select.options"
        :clearable="clearable"
        :filterable="filterable"
        @visible-change="select.visibleChange"
        :remote="remote"
        :remote-method="select.remoteMethod"
        :loading="select.loading"
        :disabled="disabled"
        :multiple="multiple"
        :multiple-limit="multipleLimit"
        :collapse-tags="collapseTags"
        :collapse-tags-tooltip="collapseTagsTooltip"
    >
        <template v-if="slots.default" #default="{ item }">
            <slot name="default" :item="item"></slot>
        </template>
        <template v-if="slots.empty" #empty>
            <slot name="empty"></slot>
        </template>
        <template v-if="slots.prefix" #prefix>
            <slot name="prefix"></slot>
        </template>
        <template v-if="slots.tag" #tag>
            <slot name="tag"></slot>
        </template>
    </el-select-v2>
</template>

<style scoped>
.el-select.el-select--default {
    /* 
        width设置原因：（外部可设置同属性覆盖）
            1、multiple设置为true时，显示时宽度很小（11px）
            2、更新到element-2.5.0版本以上后，el-select-v2组件在el-form组件内使用时，当el-form组件设置:inline="true"时，显示时宽度很小（11px），当:inline="false"时，显示时宽度很大（100%）
    */
    width: 214px;
}
</style>
