<script setup lang="tsx">
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
        uploadOfAliyunOssRoleArn: '',
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
        uploadOfAliyunOssRoleArn: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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
            const param = removeEmptyOfObj(saveForm.data)
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
    <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <el-form-item :label="t('platform.config.plugin.name.uploadType')" prop="uploadType">
            <el-radio-group v-model="saveForm.data.uploadType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.uploadType') as any" :key="index" :label="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.uploadType == 'uploadOfLocal'">
            <el-form-item :label="t('platform.config.plugin.name.uploadOfLocalUrl')" prop="uploadOfLocalUrl">
                <el-input v-model="saveForm.data.uploadOfLocalUrl" :placeholder="t('platform.config.plugin.name.uploadOfLocalUrl')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfLocalSignKey')" prop="uploadOfLocalSignKey">
                <el-input v-model="saveForm.data.uploadOfLocalSignKey" :placeholder="t('platform.config.plugin.name.uploadOfLocalSignKey')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfLocalFileSaveDir')" prop="uploadOfLocalFileSaveDir">
                <el-input v-model="saveForm.data.uploadOfLocalFileSaveDir" :placeholder="t('platform.config.plugin.name.uploadOfLocalFileSaveDir')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert :title="t('platform.config.plugin.tip.uploadOfLocalFileSaveDir')" type="info" :show-icon="true" :closable="false" />
                </label>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfLocalFileUrlPrefix')" prop="uploadOfLocalFileUrlPrefix">
                <el-input v-model="saveForm.data.uploadOfLocalFileUrlPrefix" :placeholder="t('platform.config.plugin.name.uploadOfLocalFileUrlPrefix')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert :title="t('platform.config.plugin.tip.uploadOfLocalFileUrlPrefix')" type="info" :show-icon="true" :closable="false" />
                </label>
            </el-form-item>
        </template>

        <template v-if="saveForm.data.uploadType == 'uploadOfAliyunOss'">
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssHost')" prop="uploadOfAliyunOssHost">
                <el-input v-model="saveForm.data.uploadOfAliyunOssHost" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssHost')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert :title="t('platform.config.plugin.tip.uploadOfAliyunOssHost')" type="info" :show-icon="true" :closable="false" />
                </label>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssBucket')" prop="uploadOfAliyunOssBucket">
                <el-input v-model="saveForm.data.uploadOfAliyunOssBucket" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssBucket')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssAccessKeyId')" prop="uploadOfAliyunOssAccessKeyId">
                <el-input v-model="saveForm.data.uploadOfAliyunOssAccessKeyId" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssAccessKeyId')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssAccessKeySecret')" prop="uploadOfAliyunOssAccessKeySecret">
                <el-input v-model="saveForm.data.uploadOfAliyunOssAccessKeySecret" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssAccessKeySecret')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssCallbackUrl')" prop="uploadOfAliyunOssCallbackUrl">
                <el-input v-model="saveForm.data.uploadOfAliyunOssCallbackUrl" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssCallbackUrl')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert :title="t('platform.config.plugin.tip.uploadOfAliyunOssCallbackUrl')" type="info" :show-icon="true" :closable="false" />
                </label>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssEndpoint')" prop="uploadOfAliyunOssEndpoint">
                <el-input v-model="saveForm.data.uploadOfAliyunOssEndpoint" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssEndpoint')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.plugin.tip.uploadOfAliyunOssEndpoint')"></span>
                        </template>
                    </el-alert>
                </label>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.uploadOfAliyunOssRoleArn')" prop="uploadOfAliyunOssRoleArn">
                <el-input v-model="saveForm.data.uploadOfAliyunOssRoleArn" :placeholder="t('platform.config.plugin.name.uploadOfAliyunOssRoleArn')" :clearable="true" style="max-width: 500px" />
                <label>
                    <el-alert :title="t('platform.config.plugin.tip.uploadOfAliyunOssRoleArn')" type="info" :show-icon="true" :closable="false" />
                </label>
            </el-form-item>
        </template>

        <el-form-item>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
