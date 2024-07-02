<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        smsType: 'smsOfAliyun',
        smsOfAliyunAccessKeyId: '',
        smsOfAliyunAccessKeySecret: '',
        smsOfAliyunEndpoint: '',
        smsOfAliyunSignName: '',
        smsOfAliyunTemplateCode: '',
    } as { [propName: string]: any },
    rules: {
        smsType: [{ type: 'enum', trigger: 'change', enum: [`smsOfAliyun`], message: t('validation.select') }],
        smsOfAliyunAccessKeyId: [{ type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') }],
        smsOfAliyunAccessKeySecret: [{ type: 'string', trigger: 'blur', pattern: /^[\p{L}\p{N}_-]+$/u, message: t('validation.alpha_dash') }],
        smsOfAliyunEndpoint: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        smsOfAliyunSignName: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        smsOfAliyunTemplateCode: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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
            } catch (error) {
                /* eslint-disable-next-line no-empty */
            }
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
        <el-form-item :label="t('platform.config.plugin.name.smsType')" prop="smsType">
            <el-radio-group v-model="saveForm.data.smsType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.smsType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.smsType == 'smsOfAliyun'">
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyunAccessKeyId')" prop="smsOfAliyunAccessKeyId">
                <el-input v-model="saveForm.data.smsOfAliyunAccessKeyId" :placeholder="t('platform.config.plugin.name.smsOfAliyunAccessKeyId')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyunAccessKeySecret')" prop="smsOfAliyunAccessKeySecret">
                <el-input v-model="saveForm.data.smsOfAliyunAccessKeySecret" :placeholder="t('platform.config.plugin.name.smsOfAliyunAccessKeySecret')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyunEndpoint')" prop="smsOfAliyunEndpoint">
                <el-input v-model="saveForm.data.smsOfAliyunEndpoint" :placeholder="t('platform.config.plugin.name.smsOfAliyunEndpoint')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyunSignName')" prop="smsOfAliyunSignName">
                <el-input v-model="saveForm.data.smsOfAliyunSignName" :placeholder="t('platform.config.plugin.name.smsOfAliyunSignName')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyunTemplateCode')" prop="smsOfAliyunTemplateCode">
                <el-input v-model="saveForm.data.smsOfAliyunTemplateCode" :placeholder="t('platform.config.plugin.name.smsOfAliyunTemplateCode')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isSmsSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
