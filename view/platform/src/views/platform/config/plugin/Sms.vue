<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        smsType: 'smsOfAliyun',
        smsOfAliyun: {},
    } as { [propName: string]: any },
    rules: {
        smsType: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: [`smsOfAliyun`], message: t('validation.select') },
        ],
        'smsOfAliyun.accessKeyId': [
            { required: computed((): boolean => (saveForm.data.smsType == `smsOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'smsOfAliyun.accessKeySecret': [
            { required: computed((): boolean => (saveForm.data.smsType == `smsOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'smsOfAliyun.endpoint': [
            { required: computed((): boolean => (saveForm.data.smsType == `smsOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'smsOfAliyun.signName': [
            { required: computed((): boolean => (saveForm.data.smsType == `smsOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'smsOfAliyun.templateCode': [
            { required: computed((): boolean => (saveForm.data.smsType == `smsOfAliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
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
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyun.accessKeyId')" prop="smsOfAliyun.accessKeyId">
                <el-input v-model="saveForm.data.smsOfAliyun.accessKeyId" :placeholder="t('platform.config.plugin.name.smsOfAliyun.accessKeyId')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyun.accessKeySecret')" prop="smsOfAliyun.accessKeySecret">
                <el-input v-model="saveForm.data.smsOfAliyun.accessKeySecret" :placeholder="t('platform.config.plugin.name.smsOfAliyun.accessKeySecret')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyun.endpoint')" prop="smsOfAliyun.endpoint">
                <el-input v-model="saveForm.data.smsOfAliyun.endpoint" :placeholder="t('platform.config.plugin.name.smsOfAliyun.endpoint')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyun.signName')" prop="smsOfAliyun.signName">
                <el-input v-model="saveForm.data.smsOfAliyun.signName" :placeholder="t('platform.config.plugin.name.smsOfAliyun.signName')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.smsOfAliyun.templateCode')" prop="smsOfAliyun.templateCode">
                <el-input v-model="saveForm.data.smsOfAliyun.templateCode" :placeholder="t('platform.config.plugin.name.smsOfAliyun.templateCode')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isSmsSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
