<!-------- 使用示例 开始-------->
<!-- <my-cascader v-model="saveForm.data.menuId" :placeholder="t('common.name.rel.menuId')"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { field: ['id', 'menuName'] } }"
    :props="{ emitPath: false, value: 'id', label: 'menuName' }" />

<my-cascader v-model="saveForm.data.menuIdArr"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId } } }" :isPanel="true"
    :props="{ multiple: true }" />

<my-cascader v-model="saveForm.data.pid"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { sceneId: saveForm.data.sceneId } } }"
    :props="{ checkStrictly: true, emitPath: false }" />
<my-cascader v-model="saveForm.data.pid"
    :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/list', param: { filter: { sceneId: saveForm.data.sceneId } } }"
    :props="{ checkStrictly: true, emitPath: false, lazy: true }" />

<my-cascader v-model="queryCommon.data.pid" :placeholder="t('auth.menu.name.pid')"
    :defaultOptions="[{ id: 0, label: t('common.name.allTopLevel') }]" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree' }"
    :props="{ checkStrictly: true, emitPath: false }" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
const slots = useSlots()
const props = defineProps({
    modelValue: {
        type: [String, Number, Array],
    },
    defaultOptions: {
        //选项初始默认值。格式：[{ [cascader.props.value]: any, [cascader.props.label]: any },...]
        type: Array,
        default: () => [],
    },
    /**
     * 接口。格式：{ code: string, param: object, transform: function, pidField: string }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ filter: { [propName: string]: any }, field: string[], sort: string, page: number, limit: number }。其中field内第0，1字段默认用于cascader.props的value，label属性，cascader.api的transform属性，使用时请注意。或直接在props.props中设置对应参数
     *      transform：非必须。接口返回数据转换方法
     *      pidField：非必须。动态加载时用于获取子级，接口参数filter中使用的字段名
     */
    api: {
        type: Object,
        required: true,
    },
    isPanel: {
        //是否为面板
        type: Boolean,
        default: false,
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
    disabled: {
        type: Boolean,
        default: false,
    },
    collapseTags: {
        type: Boolean,
        default: true,
    },
    collapseTagsTooltip: {
        type: Boolean,
        default: true,
    },
    separator: {
        type: String,
        default: '/',
    },
    props: {
        type: Object,
        default: () => {},
    },
})

const emits = defineEmits(['update:modelValue', 'change'])
const cascader = reactive({
    ref: null as any,
    value: computed({
        get: (): any => {
            return props.modelValue
        },
        set: (val) => {
            emits('update:modelValue', val)
            emits('change')
        },
    }),
    options: [...props.defaultOptions] as any,
    props: {
        expandTrigger: 'hover' as any, //子级展开方式。click或hover
        checkStrictly: false,
        lazy: false, //不建议使用动态加载模式，使用体验很差
        lazyLoad: (node: any, resolve: any) => {
            if (node.level == 0) {
                cascader.api.param.filter[cascader.api.pidField] = 0
            } else {
                cascader.api.param.filter[cascader.api.pidField] = node.data.id
            }
            cascader.api.getOptions().then((options) => {
                if (options?.length === 0) {
                    node.data.leaf = true
                }
                resolve(options)
            })
            delete cascader.api.param.filter[cascader.api.pidField]
        },
        value: props.api?.param?.field?.[0] ?? 'id',
        label: props.api?.param?.field?.[1] ?? 'label',
        ...props.props,
    },
    initOptions: () => {
        cascader.api.addOptions()
    },
    resetOptions: () => {
        cascader.options = [...props.defaultOptions] as any
        cascader.api.param.page = 1
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
                      if (cascader.props.lazy) {
                          if (!cascader.props.checkStrictly) {
                              //这种情况暂时可以用非动态全部加载解决。等确实需要使用时在考虑修改。
                              //动态加载，且当checkStrictly为false时，leaf字段必须有，否则选中后值为null
                              /* const options: any = []
                        res.data.list.forEach((item: any) => {
                            options.push({
                                [cascader.props.value]: item[cascader.api.param.field[0]],
                                [cascader.props.label]: item[cascader.api.param.field[1]],
                                //[cascader.props.leaf]: false    //后端接口还得返回一个是否有子集的字段，暂时不考虑
                            })
                        })
                        return options */
                          }
                          return res.data.list
                      }
                      return res.data.tree
                  }
        }),
        pidField: computed((): string => {
            return props.api.pidField ?? 'pid'
        }),
        getOptions: async () => {
            if (cascader.api.loading) {
                return
            }
            cascader.api.loading = true
            let options = []
            try {
                const res = await request(props.api.code, cascader.api.param)
                options = cascader.api.transform(res)
            } finally {
                cascader.api.loading = false
            }
            return options
        },
        addOptions: () => {
            cascader.api.getOptions().then((options) => {
                if (options?.length) {
                    cascader.options = cascader.options.concat(options ?? [])
                }
            })
        },
    },
    visibleChange: (val: boolean) => {
        if (val) {
            //每次打开都重新加载
            if (cascader.props.lazy) {
                //重新触发下动态加载事件。
                /* cascader.props.lazy = false
                cascader.props.lazy = true */
            } else {
                cascader.resetOptions()
                cascader.api.addOptions()
            }
        }
    },
})
//组件创建时，如有初始值，需初始化options
if (props.isPanel || (!cascader.props.lazy && ((Array.isArray(props.modelValue) && props.modelValue.length) || props.modelValue))) {
    cascader.initOptions()
}

//当外部环境filter变化时，重置options
watch(
    () => props.api?.param?.filter,
    (newVal: any, oldVal: any) => {
        if (JSON.stringify(newVal) !== JSON.stringify(oldVal)) {
            cascader.resetOptions()
            cascader.api.addOptions()
        }
    }
)

//暴露组件属性给父组件
defineExpose({
    options: computed(() => {
        return cascader.options
    }),
})
</script>

<template>
    <el-cascader-panel v-if="props.isPanel" :ref="(el: any) => cascader.ref = el" v-model="cascader.value" :options="cascader.options" :props="cascader.props">
        <template v-if="slots.default" #default="{ node, data }">
            <slot name="default" :node="node" :data="data"></slot>
        </template>
    </el-cascader-panel>
    <el-cascader
        v-else-if="cascader.props.lazy"
        :ref="(el: any) => cascader.ref = el"
        v-model="cascader.value"
        :placeholder="placeholder"
        :clearable="clearable"
        :props="cascader.props"
        @visible-change="cascader.visibleChange"
        :disabled="disabled"
        :collapse-tags="collapseTags"
        :collapse-tags-tooltip="collapseTagsTooltip"
        :separator="separator"
    >
        <template v-if="slots.default" #default="{ node, data }">
            <slot name="default" :node="node" :data="data"></slot>
        </template>
        <template v-if="slots.empty" #empty>
            <slot name="empty"></slot>
        </template>
    </el-cascader>
    <el-cascader
        v-else
        :ref="(el: any) => cascader.ref = el"
        v-model="cascader.value"
        :placeholder="placeholder"
        :clearable="clearable"
        :options="cascader.options"
        :props="cascader.props"
        :filterable="filterable"
        @visible-change="cascader.visibleChange"
        :disabled="disabled"
        :collapse-tags="collapseTags"
        :collapse-tags-tooltip="collapseTagsTooltip"
        :separator="separator"
    >
        <template v-if="slots.default" #default="{ node, data }">
            <slot name="default" :node="node" :data="data"></slot>
        </template>
        <template v-if="slots.empty" #empty>
            <slot name="empty"></slot>
        </template>
    </el-cascader>
</template>
