<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        uploadType: 'uploadOfLocal',
        uploadOfLocalUrl: '',
        uploadOfLocalSignKey: '',
        uploadOfLocalFileSaveDir: '',
        uploadOfLocalFileUrlPrefix: '',
        uploadOfAliyunOssHost: '',
        uploadOfAliyunOssBucket: '',
        uploadOfAliyunOssAccessKeyId: '',
        uploadOfAliyunOssAccessKeySecret: '',
        uploadOfAliyunOssCallbackUrl: '',
        uploadOfAliyunOssEndpoint: '',
        uploadOfAliyunOssRoleArn: ''
    } as { [propName: string]: any },
    rules: {
        uploadType: [{ type: 'enum', enum: [`uploadOfLocal`, `uploadOfAliyunOss`], trigger: 'change', message: t('validation.select') }],
        uploadOfLocalUrl: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        uploadOfLocalSignKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        uploadOfLocalFileSaveDir: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        uploadOfLocalFileUrlPrefix: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        uploadOfAliyunOssHost: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        uploadOfAliyunOssBucket: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        uploadOfAliyunOssAccessKeyId: [{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }],
        uploadOfAliyunOssAccessKeySecret: [{ pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }],
        uploadOfAliyunOssCallbackUrl: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        uploadOfAliyunOssEndpoint: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        uploadOfAliyunOssRoleArn: [{ type: 'string', trigger: 'blur', message: t('validation.input') }]
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config
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
    }
})

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => (saveForm.ref = el)" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('platform.config.name.uploadType')" prop="uploadType">
            <ElRadioGroup v-model="saveForm.data.uploadType">
                <ElRadio v-for="(item, index) in tm('platform.config.status.uploadType') as any" :key="index" :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.uploadType == 'uploadOfLocal'">
            <ElFormItem :label="t('platform.config.name.uploadOfLocalUrl')" prop="uploadOfLocalUrl">
                <ElInput v-model="saveForm.data.uploadOfLocalUrl" :placeholder="t('platform.config.name.uploadOfLocalUrl')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfLocalSignKey')" prop="uploadOfLocalSignKey">
                <ElInput v-model="saveForm.data.uploadOfLocalSignKey" :placeholder="t('platform.config.name.uploadOfLocalSignKey')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfLocalFileSaveDir')" prop="uploadOfLocalFileSaveDir">
                <ElInput v-model="saveForm.data.uploadOfLocalFileSaveDir" :placeholder="t('platform.config.name.uploadOfLocalFileSaveDir')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert :title="t('platform.config.tip.uploadOfLocalFileSaveDir')" type="info" :show-icon="true" :closable="false" />
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfLocalFileUrlPrefix')" prop="uploadOfLocalFileUrlPrefix">
                <ElInput v-model="saveForm.data.uploadOfLocalFileUrlPrefix" :placeholder="t('platform.config.name.uploadOfLocalFileUrlPrefix')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert :title="t('platform.config.tip.uploadOfLocalFileUrlPrefix')" type="info" :show-icon="true" :closable="false" />
                </label>
            </ElFormItem>
        </template>

        <template v-if="saveForm.data.uploadType == 'uploadOfAliyunOss'">
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssHost')" prop="uploadOfAliyunOssHost">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssHost" :placeholder="t('platform.config.name.uploadOfAliyunOssHost')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert :title="t('platform.config.tip.uploadOfAliyunOssHost')" type="info" :show-icon="true" :closable="false" />
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssBucket')" prop="uploadOfAliyunOssBucket">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssBucket" :placeholder="t('platform.config.name.uploadOfAliyunOssBucket')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssAccessKeyId')" prop="uploadOfAliyunOssAccessKeyId">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssAccessKeyId" :placeholder="t('platform.config.name.uploadOfAliyunOssAccessKeyId')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssAccessKeySecret')" prop="uploadOfAliyunOssAccessKeySecret">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssAccessKeySecret" :placeholder="t('platform.config.name.uploadOfAliyunOssAccessKeySecret')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssCallbackUrl')" prop="uploadOfAliyunOssCallbackUrl">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssCallbackUrl" :placeholder="t('platform.config.name.uploadOfAliyunOssCallbackUrl')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert :title="t('platform.config.tip.uploadOfAliyunOssCallbackUrl')" type="info" :show-icon="true" :closable="false" />
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssEndpoint')" prop="uploadOfAliyunOssEndpoint">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssEndpoint" :placeholder="t('platform.config.name.uploadOfAliyunOssEndpoint')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.tip.uploadOfAliyunOssEndpoint')"></span>
                        </template>
                    </ElAlert>
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.uploadOfAliyunOssRoleArn')" prop="uploadOfAliyunOssRoleArn">
                <ElInput v-model="saveForm.data.uploadOfAliyunOssRoleArn" :placeholder="t('platform.config.name.uploadOfAliyunOssRoleArn')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert :title="t('platform.config.tip.uploadOfAliyunOssRoleArn')" type="info" :show-icon="true" :closable="false" />
                </label>
            </ElFormItem>
        </template>

        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <AutoiconEpCircleCheck />{{ t('common.save') }} </ElButton>
            <ElButton type="info" @click="saveForm.reset"> <AutoiconEpCircleClose />{{ t('common.reset') }} </ElButton>
        </ElFormItem>
    </ElForm>
</template>
