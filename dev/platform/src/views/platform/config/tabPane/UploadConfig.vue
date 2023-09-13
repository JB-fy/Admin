<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置项key，用于向服务器获取对应的配置项value
        uploadType: 'local',
        localUploadUrl: '',
        localUploadSignKey: '',
        localUploadFileSaveDir: '',
        localUploadFileUrlPrefix: '',
        aliyunOssHost: '',
        aliyunOssBucket: '',
        aliyunOssAccessKeyId: '',
        aliyunOssAccessKeySecret: '',
        aliyunOssRoleArn: '',
        aliyunOssCallbackUrl: '',
    } as { [propName: string]: any },
    rules: {
        uploadType: [
            { type: 'enum', enum: [`local`, `aliyunOss`], trigger: 'change', message: t('validation.select') }
        ],
        localUploadUrl: [
            { type: 'url', trigger: 'blur', message: t('validation.url') }
        ],
        localUploadSignKey: [
            { type: 'string', trigger: 'blur' }
        ],
        localUploadFileSaveDir: [
            { type: 'string', trigger: 'blur' }
        ],
        localUploadFileUrlPrefix: [
            { type: 'url', trigger: 'blur', message: t('validation.url') }
        ],
        aliyunOssHost: [
            { type: 'url', trigger: 'blur', message: t('validation.url') }
        ],
        aliyunOssBucket: [
            { type: 'string', trigger: 'blur' }
        ],
        aliyunOssAccessKeyId: [
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        aliyunOssAccessKeySecret: [
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        aliyunOssRoleArn: [
            { type: 'string', trigger: 'blur' }
        ],
        aliyunOssCallbackUrl: [
            { type: 'url', trigger: 'blur', message: t('validation.url') }
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
        <ElFormItem :label="t('platform.config.name.uploadType')" prop="uploadType">
            <ElRadioGroup v-model="saveForm.data.uploadType">
                <ElRadio v-for="(item, index) in (tm('platform.config.status.uploadType') as any)" :key="index"
                    :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.uploadType == 'local'">
            <ElFormItem :label="t('platform.config.name.localUploadUrl')" prop="localUploadUrl">
                <ElInput v-model="saveForm.data.localUploadUrl" :placeholder="t('platform.config.name.localUploadUrl')"
                    :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.localUploadSignKey')" prop="localUploadSignKey">
                <ElInput v-model="saveForm.data.localUploadSignKey"
                    :placeholder="t('platform.config.name.localUploadSignKey')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.localUploadFileSaveDir')" prop="localUploadFileSaveDir">
                <ElInput v-model="saveForm.data.localUploadFileSaveDir"
                    :placeholder="t('platform.config.name.localUploadFileSaveDir')" :clearable="true"
                    style="max-width: 500px;" />
                <label>
                    <ElAlert :title="t('platform.config.tip.localUploadFileSaveDir')" type="info" :show-icon="true"
                        :closable="false" />
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.localUploadFileUrlPrefix')" prop="localUploadFileUrlPrefix">
                <ElInput v-model="saveForm.data.localUploadFileUrlPrefix"
                    :placeholder="t('platform.config.name.localUploadFileUrlPrefix')" :clearable="true"
                    style="max-width: 500px;" />
                <label>
                    <ElAlert :title="t('platform.config.tip.localUploadFileUrlPrefix')" type="info" :show-icon="true"
                        :closable="false" />
                </label>
            </ElFormItem>
        </template>

        <template v-if="saveForm.data.uploadType == 'aliyunOss'">
            <ElFormItem :label="t('platform.config.name.aliyunOssHost')" prop="aliyunOssHost">
                <ElInput v-model="saveForm.data.aliyunOssHost" :placeholder="t('platform.config.name.aliyunOssHost')"
                    :clearable="true" style="max-width: 500px;" />
                <label>
                    <ElAlert :title="t('platform.config.tip.aliyunOssHost')" type="info" :show-icon="true"
                        :closable="false" />
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunOssBucket')" prop="aliyunOssBucket">
                <ElInput v-model="saveForm.data.aliyunOssBucket" :placeholder="t('platform.config.name.aliyunOssBucket')"
                    :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunOssAccessKeyId')" prop="aliyunOssAccessKeyId">
                <ElInput v-model="saveForm.data.aliyunOssAccessKeyId"
                    :placeholder="t('platform.config.name.aliyunOssAccessKeyId')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunOssAccessKeySecret')" prop="aliyunOssAccessKeySecret">
                <ElInput v-model="saveForm.data.aliyunOssAccessKeySecret"
                    :placeholder="t('platform.config.name.aliyunOssAccessKeySecret')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunOssRoleArn')" prop="aliyunOssRoleArn">
                <ElInput v-model="saveForm.data.aliyunOssRoleArn" :placeholder="t('platform.config.name.aliyunOssRoleArn')"
                    :clearable="true" style="max-width: 500px;" />
                <label>
                    <ElAlert :title="t('platform.config.tip.aliyunOssRoleArn')" type="info" :show-icon="true"
                        :closable="false" />
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunOssCallbackUrl')" prop="aliyunOssCallbackUrl">
                <ElInput v-model="saveForm.data.aliyunOssCallbackUrl"
                    :placeholder="t('platform.config.name.aliyunOssCallbackUrl')" :clearable="true"
                    style="max-width: 500px;" />
                <label>
                    <ElAlert :title="t('platform.config.tip.aliyunOssCallbackUrl')" type="info" :show-icon="true"
                        :closable="false" />
                </label>
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