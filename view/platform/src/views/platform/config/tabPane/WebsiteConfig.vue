<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        hotSearch: [],
        userAgreement: '',
        privacyAgreement: '',
    } as { [propName: string]: any },
    rules: {
        hotSearch: [
            // { type: 'array', trigger: 'change', message: t('validation.required') },
            { type: 'array', max: 10, trigger: 'change', message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } },
        ],
        userAgreement: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        privacyAgreement: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config,
            }
        } catch (error) {}
    },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data, false)
            try {
                await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/save', param, true)
            } catch (error) {}
            saveForm.loading = false
        })
    },
    reset: () => {
        saveForm.ref.resetFields()
        saveForm.initData()
    },
})

const hotSearchHandle = reactive({
    ref: null as any,
    visible: false,
    value: undefined,
    tagType: tm('config.const.tagType') as any[],
    visibleChange: () => {
        hotSearchHandle.visible = true
        nextTick(() => {
            hotSearchHandle.ref?.focus()
        })
    },
    addValue: () => {
        if (hotSearchHandle.value) {
            saveForm.data.hotSearch.push(hotSearchHandle.value)
        }
        hotSearchHandle.visible = false
        hotSearchHandle.value = undefined
    },
    delValue: (item: any) => {
        saveForm.data.hotSearch.splice(saveForm.data.hotSearch.indexOf(item), 1)
    },
})

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => (saveForm.ref = el)" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('platform.config.name.hotSearch')" prop="hotSearch">
            <ElTag v-for="(item, index) in saveForm.data.hotSearch" :type="hotSearchHandle.tagType[index % hotSearchHandle.tagType.length]" @close="hotSearchHandle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px">
                {{ item }}
            </ElTag>
            <template v-if="saveForm.data.hotSearch.length < 10">
                <ElInput
                    v-if="hotSearchHandle.visible"
                    :ref="(el: any) => (hotSearchHandle.ref = el)"
                    v-model="hotSearchHandle.value"
                    :placeholder="t('platform.config.name.hotSearch')"
                    @keyup.enter="hotSearchHandle.addValue"
                    @blur="hotSearchHandle.addValue"
                    size="small"
                    style="width: 100px"
                />
                <ElButton v-else type="primary" size="small" @click="hotSearchHandle.visibleChange"> <AutoiconEpPlus />{{ t('common.add') }} </ElButton>
            </template>
        </ElFormItem>
        <ElFormItem :label="t('platform.config.name.userAgreement')" prop="userAgreement">
            <MyEditor v-model="saveForm.data.userAgreement" />
        </ElFormItem>
        <ElFormItem :label="t('platform.config.name.privacyAgreement')" prop="privacyAgreement">
            <MyEditor v-model="saveForm.data.privacyAgreement" />
        </ElFormItem>
        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <AutoiconEpCircleCheck />{{ t('common.save') }} </ElButton>
            <ElButton type="info" @click="saveForm.reset"> <AutoiconEpCircleClose />{{ t('common.reset') }} </ElButton>
        </ElFormItem>
    </ElForm>
</template>
