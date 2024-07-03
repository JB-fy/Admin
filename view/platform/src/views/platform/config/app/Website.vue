<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        hotSearch: [],
        userAgreement: '',
        privacyAgreement: '',
    } as { [propName: string]: any },
    rules: {
        hotSearch: [{ type: 'array', trigger: 'change', max: 10, message: t('validation.max.array', { max: 10 }), defaultField: { type: 'string', message: t('validation.input') } }],
        userAgreement: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        privacyAgreement: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
        const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
        saveForm.data = {
            ...saveForm.data,
            ...res.data.config,
        }
    },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            try {
                await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/save', param, true)
            } finally {
                saveForm.loading = false
            }
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
    <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <el-form-item :label="t('platform.config.platform.name.hotSearch')" prop="hotSearch">
            <el-tag v-for="(item, index) in saveForm.data.hotSearch" :type="hotSearchHandle.tagType[index % hotSearchHandle.tagType.length]" @close="hotSearchHandle.delValue(item)" :key="index" :closable="true" style="margin-right: 10px">
                {{ item }}
            </el-tag>
            <template v-if="saveForm.data.hotSearch.length < 10">
                <el-input
                    v-if="hotSearchHandle.visible"
                    :ref="(el: any) => hotSearchHandle.ref = el"
                    v-model="hotSearchHandle.value"
                    :placeholder="t('platform.config.platform.name.hotSearch')"
                    @keyup.enter="hotSearchHandle.addValue"
                    @blur="hotSearchHandle.addValue"
                    size="small"
                    style="width: 100px"
                />
                <el-button v-else type="primary" size="small" @click="hotSearchHandle.visibleChange"> <autoicon-ep-plus />{{ t('common.add') }}</el-button>
            </template>
        </el-form-item>
        <el-form-item :label="t('platform.config.platform.name.userAgreement')" prop="userAgreement">
            <my-editor v-model="saveForm.data.userAgreement" />
        </el-form-item>
        <el-form-item :label="t('platform.config.platform.name.privacyAgreement')" prop="privacyAgreement">
            <my-editor v-model="saveForm.data.privacyAgreement" />
        </el-form-item>
        <el-form-item>
            <el-button v-if="authAction.isWebsiteSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
