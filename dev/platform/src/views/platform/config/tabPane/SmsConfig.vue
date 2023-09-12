<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置项key，用于向服务器获取对应的配置项value
        smsType: 'aliyunSms',
        aliyunSmsAccessKeyId: '',
        aliyunSmsAccessKeySecret: '',
        aliyunSmsSignName: '',
        aliyunSmsTemplateCode: '',
    } as { [propName: string]: any },
    rules: {
        smsType: [
            { type: 'enum', enum: [`aliyunSms`], trigger: 'change', message: t('validation.select') }
        ],
        aliyunSmsAccessKeyId: [
            //{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        aliyunSmsAccessKeySecret: [
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        aliyunSmsSignName: [
            { type: 'string', trigger: 'blur' }
        ],
        aliyunSmsTemplateCode: [
            { type: 'string', trigger: 'blur' }
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
        <ElFormItem :label="t('platform.config.name.smsType')" prop="smsType">
            <ElRadioGroup v-model="saveForm.data.smsType">
                <ElRadio v-for="(item, index) in (tm('platform.config.status.smsType') as any)" :key="index"
                    :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.smsType == 'aliyunSms'">
            <ElFormItem :label="t('platform.config.name.aliyunSmsAccessKeyId')" prop="aliyunSmsAccessKeyId">
                <ElInput v-model="saveForm.data.aliyunSmsAccessKeyId"
                    :placeholder="t('platform.config.name.aliyunSmsAccessKeyId')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunSmsAccessKeySecret')" prop="aliyunSmsAccessKeySecret">
                <ElInput v-model="saveForm.data.aliyunSmsAccessKeySecret"
                    :placeholder="t('platform.config.name.aliyunSmsAccessKeySecret')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunSmsSignName')" prop="aliyunSmsSignName">
                <ElInput v-model="saveForm.data.aliyunSmsSignName"
                    :placeholder="t('platform.config.name.aliyunSmsSignName')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunSmsTemplateCode')" prop="aliyunSmsTemplateCode">
                <ElInput v-model="saveForm.data.aliyunSmsTemplateCode"
                    :placeholder="t('platform.config.name.aliyunSmsTemplateCode')" :clearable="true" />
            </ElFormItem>
        </template>

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