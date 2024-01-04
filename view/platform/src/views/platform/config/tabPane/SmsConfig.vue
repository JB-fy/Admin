<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        smsType: 'smsOfAliyun',
        smsOfAliyunAccessKeyId: '',
        smsOfAliyunAccessKeySecret: '',
        smsOfAliyunEndpoint: '',
        smsOfAliyunSignName: '',
        smsOfAliyunTemplateCode: '',
    } as { [propName: string]: any },
    rules: {
        smsType: [{ type: 'enum', enum: [`smsOfAliyun`], trigger: 'change', message: t('validation.select') }],
        smsOfAliyunAccessKeyId: [{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }],
        smsOfAliyunAccessKeySecret: [{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }],
        smsOfAliyunEndpoint: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        smsOfAliyunSignName: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        smsOfAliyunTemplateCode: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => (saveForm.ref = el)" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('platform.config.name.smsType')" prop="smsType">
            <ElRadioGroup v-model="saveForm.data.smsType">
                <ElRadio v-for="(item, index) in tm('platform.config.status.smsType') as any" :key="index" :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.smsType == 'smsOfAliyun'">
            <ElFormItem :label="t('platform.config.name.smsOfAliyunAccessKeyId')" prop="smsOfAliyunAccessKeyId">
                <ElInput v-model="saveForm.data.smsOfAliyunAccessKeyId" :placeholder="t('platform.config.name.smsOfAliyunAccessKeyId')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.smsOfAliyunAccessKeySecret')" prop="smsOfAliyunAccessKeySecret">
                <ElInput v-model="saveForm.data.smsOfAliyunAccessKeySecret" :placeholder="t('platform.config.name.smsOfAliyunAccessKeySecret')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.smsOfAliyunEndpoint')" prop="smsOfAliyunEndpoint">
                <ElInput v-model="saveForm.data.smsOfAliyunEndpoint" :placeholder="t('platform.config.name.smsOfAliyunEndpoint')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.smsOfAliyunSignName')" prop="smsOfAliyunSignName">
                <ElInput v-model="saveForm.data.smsOfAliyunSignName" :placeholder="t('platform.config.name.smsOfAliyunSignName')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.smsOfAliyunTemplateCode')" prop="smsOfAliyunTemplateCode">
                <ElInput v-model="saveForm.data.smsOfAliyunTemplateCode" :placeholder="t('platform.config.name.smsOfAliyunTemplateCode')" :clearable="true" />
            </ElFormItem>
        </template>

        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <AutoiconEpCircleCheck />{{ t('common.save') }} </ElButton>
            <ElButton type="info" @click="saveForm.reset"> <AutoiconEpCircleClose />{{ t('common.reset') }} </ElButton>
        </ElFormItem>
    </ElForm>
</template>
