<!-------- 使用示例 开始-------->
<!-- <my-cascader v-model="saveForm.data.menu_id" :placeholder="t('common.name.rel.menu_id')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { field: ['id', 'menu_name'] } }" :props="{ emitPath: false, value: 'id', label: 'menu_name' }" />

<my-cascader v-model="saveForm.data.menu_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { scene_id: saveForm.data.scene_id } } }" :isPanel="true" :props="{ multiple: true }" />

<my-cascader v-model="saveForm.data.pid" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree', param: { filter: { scene_id: saveForm.data.scene_id } } }" :props="{ checkStrictly: true, emitPath: false }" />
<my-cascader v-model="saveForm.data.pid" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/list', param: { filter: { scene_id: saveForm.data.scene_id } } }" :props="{ checkStrictly: true, emitPath: false, lazy: true }" />

<my-cascader v-model="queryCommon.data.pid" :placeholder="t('auth.menu.name.pid')" :defaultOptions="tm('common.status.pid')" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/menu/tree' }" :props="{ checkStrictly: true, emitPath: false }" /> -->
<!-------- 使用示例 结束-------->
<script setup lang="tsx">
const slots = useSlots()
const props = defineProps({
    modelValue: {
        type: [String, Number, Array],
    },
    defaultOptions: {
        //选项初始默认值。格式：{ [cascader.props.value]: any, [cascader.props.label]: any, [cascader.props.leaf]: boolean, [cascader.props.children]: { [propName: string]: any }[], [propName: string]: any }[]
        type: Array,
        default: () => [],
    },
    /**
     * 接口。格式：{ code: string, param: object, transform: function, pidField: string, pidIsStr: true }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ filter: { [propName: string]: any }, field: string[], sort: string, page: number, limit: number }。其中field内第0，1字段默认用于cascader.props的value，label属性，cascader.api的transform属性，使用时请注意。或直接在props.props中设置对应参数
     *      transform：非必须。接口返回数据转换方法
     *      pidField：非必须。动态加载时用于获取子级，接口参数filter中使用的字段名
     *      pidIsStr：非必须。pid是否字符串类型
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
        lazy: false,
        lazyLoad: (node: any, resolve: any) => {
            if (node.level == 0) {
                cascader.api.param.filter[cascader.api.pidField] = props.api.pidIsStr ? '' : 0
            } else {
                cascader.api.param.filter[cascader.api.pidField] = node.data[cascader.props.value]
            }
            cascader.api.getOptions().then((options) => {
                if (!options?.length) {
                    node.data.leaf = true
                }
                if (node.level == 0) {
                    options = [...props.defaultOptions, ...options]
                }
                resolve(options)
            })
            delete cascader.api.param.filter[cascader.api.pidField]
        },
        value: 'value',
        label: 'label',
        leaf: 'leaf',
        children: 'children',
        ...props.props,
    },
    initOptions: () => {
        cascader.api.addOptions()
    },
    api: {
        loading: false,
        param: computed((): { filter: { [propName: string]: any }; field: string[]; sort: string; page: number; limit: number } => {
            const param = {
                filter: {} as { [propName: string]: any },
                field: ['id', 'label'],
                page: 1,
                limit: 0,
                ...(props.api?.param ?? {}),
            }
            if (cascader.props.lazy /* && !cascader.props.checkStrictly */) {
                // 当checkStrictly=true时，可在cascader.props.lazyLoad中动态改变leaf=true
                // 当checkStrictly=false时，可在cascader.props.lazyLoad中动态改变leaf=true。但选项选中后值为null，故服务器必须返回是否有子级is_has_child字段，用于直接确定leaf
                // 无子级设置leaf=true
                param.field.push('is_has_child')
            }
            return param
        }),
        transform: computed(() => {
            return props.api.transform
                ? props.api.transform
                : (res: any) => {
                      const handle = (tree: { [propName: string]: any }[]) => {
                          const treeTmp: { [propName: string]: any }[] = []
                          tree.forEach((item, index) => {
                              treeTmp[index] = {
                                  ...item,
                                  [cascader.props.value]: item[cascader.api.param.field[0]],
                                  [cascader.props.label]: item[cascader.api.param.field[1]],
                              }
                              if ('is_has_child' in item) {
                                  treeTmp[index][cascader.props.leaf] = item.is_has_child === 0 ? true : false
                              }
                              if (item.children?.length) {
                                  treeTmp[index][cascader.props.children] = handle(item.children)
                              }
                          })
                          return treeTmp
                      }
                      if (!cascader.props.lazy) {
                          return handle(res.data.tree ?? [])
                      }
                      return handle(res.data.list ?? [])
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
                cascader.options = [...props.defaultOptions, ...(options ?? [])]
            })
        },
    },
    visibleChange: (val: boolean) => {
        if (val) {
            //每次打开都重新加载
            if (cascader.props.lazy) {
                //重新触发下动态加载事件
                cascader.props.lazy = false
                cascader.props.lazy = true
            } else {
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
