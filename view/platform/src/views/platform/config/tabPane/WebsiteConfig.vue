<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        userAgreement: '',
        privacyAgreement: '',
    } as { [propName: string]: any },
    rules: {
        userAgreement: [
            //{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { type: 'string', trigger: 'blur' },
        ],
        privacyAgreement: [
            { type: 'string', trigger: 'blur' },
        ],
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config
            }
        } catch (error) { }
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
            } catch (error) { }
            saveForm.loading = false
        })
    },
    reset: () => {
        saveForm.ref.resetFields()
        saveForm.initData()
    }
})

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto"
        :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('platform.config.name.userAgreement')" prop="userAgreement">
            <MyEditor v-model="saveForm.data.userAgreement" />
        </ElFormItem>
        <ElFormItem :label="t('platform.config.name.privacyAgreement')" prop="privacyAgreement">
            <MyEditor v-model="saveForm.data.privacyAgreement" />
        </ElFormItem>
        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                <AutoiconEpCircleCheck />{{ t('common.save') }}
            </ElButton>
            <ElButton type="info" @click="saveForm.reset">
                <AutoiconEpCircleClose />{{ t('common.reset') }}
            </ElButton>
        </ElFormItem>
    </ElForm>
</template>