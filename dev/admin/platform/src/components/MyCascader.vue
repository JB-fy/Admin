<script setup lang="ts">
const props = defineProps({
    modelValue: {
        type: [String, Number, Array],
        //required: true,
    },
    defaultOptions: {   //选项初始默认值。格式：[{ [cascader.props.value]: string | number, [cascader.props.label]: string },...]
        type: Array,
        default: []
    },
    /**
     * 接口。格式：{ code: string, param: object, dataToOptions: function }
     *      code：必须。接口标识。参考common/utils/common.js文件内request方法的参数说明
     *      param：必须。接口函数所需参数。格式：{ field: string[], where: { [propName: string]: any }, order: { [propName: string]: any }, page: number, limit: number }。其中field内第0，1字段默认用于cascader.props的value，label属性，cascader.api的dataToOptions属性，使用时请注意。或直接在props.props中设置对应参数
     *      dataToOptions：非必须。接口返回数据转换方法
     *      pidField：非必须。动态加载时用于获取子级，接口参数where中使用的字段名
     */
    api: {
        type: Object,
        required: true,
    },
    isPanel: {  //是否为面板
        type: Boolean,
        default: false
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
    disabled: {
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
    separator: {
        type: String,
        default: '/'
    },
    props: {
        type: Object,
        default: {}
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
            emits('change')
            emits('update:modelValue', val)
        }
    }),
    options: [...props.defaultOptions] as any,
    props: {
        multiple: false,
        checkStrictly: true,
        emitPath: false,
        lazy: false,    //不建议使用动态加载模式，使用体验很差
        lazyLoad: (node: any, resolve: any) => {
            if (node.level == 0) {
                cascader.api.param.where[cascader.api.pidField] = 0
            } else {
                cascader.api.param.where[cascader.api.pidField] = node.data.id
            }
            cascader.api.getOptions().then((options) => {
                if (options.length === 0) {
                    node.data.leaf = true
                }
                resolve(options)
            }).catch((error) => { })
            delete cascader.api.param.where[cascader.api.pidField]
        },
        value: props.props.value ?? props.api.param.field[0] ?? 'value',
        label: props.props.label ?? props.api.param.field[1] ?? 'label',
        children: props.props.children ?? 'children',
        disabled: props.props.disabled ?? 'disabled',
        leaf: props.props.leaf ?? 'leaf',   //动态加载时用于终止继续加载。当checkStrictly为false时，该字段必须有，否则选中后值为null
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
            if (cascader.props.lazy) {
                //重新触发下动态加载事件。
                /* cascader.props.lazy = false
                cascader.props.lazy = true */
            } else {
                cascader.resetOptions()
                cascader.api.addOptions()
            }
        }
    }
})
//组件创建时，如有初始值，需初始化options
if (!cascader.props.lazy && ((Array.isArray(props.modelValue) && props.modelValue.length) || props.modelValue)) {
    cascader.initOptions()
}
console.log(props.props)
console.log(cascader.props)
</script>

<template>
    <ElCascaderPanel v-if="props.isPanel" :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value"
        :options="cascader.options" :props="cascader.props" />
    <ElCascader v-else-if="cascader.props.lazy" :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value"
        :placeholder="placeholder" :clearable="clearable" :props="cascader.props"
        @visible-change="cascader.visibleChange" :disabled="disabled" :collapse-tags="collapseTags"
        :collapse-tags-tooltip="collapseTagsTooltip" :separator="separator" />
    <ElCascader v-else :ref="(el: any) => { cascader.ref = el }" v-model="cascader.value" :placeholder="placeholder"
        :clearable="clearable" :options="cascader.options" :props="cascader.props" :filterable="filterable"
        @visible-change="cascader.visibleChange" :disabled="disabled" :collapse-tags="collapseTags"
        :collapse-tags-tooltip="collapseTagsTooltip" :separator="separator" />

    <!-------- 使用示例 开始-------->
   <!--  <MyCascader v-model="saveCommon.data.menuIdArr"
        :api="{ code: 'auth/menu/tree', param: { field: ['id', 'menuName'], where: { sceneId: saveCommon.data.sceneId } } }"
        :isPanel="true" :props="{ multiple: true }" /> -->

    <!-- <MyCascader v-model="saveCommon.data.pid"
        :api="{ code: 'auth/menu/tree', param: { field: ['id', 'menuName'], where: { sceneId: saveCommon.data.sceneId } } }" />
    <MyCascader v-model="saveCommon.data.pid"
        :api="{ code: 'auth/menu/list', param: { field: ['id', 'menuName'], where: { sceneId: saveCommon.data.sceneId } } }"
        :props="{ lazy: true }" />

    <MyCascader v-model="queryCommon.data.pid" :placeholder="t('common.name.rel.pid')"
        :defaultOptions="[{ id: 0, menuName: t('common.name.allTopLevel') }]"
        :api="{ code: 'auth/menu/tree', param: { field: ['id', 'menuName'] } }" /> -->
    <!-------- 使用示例 结束-------->
</template>