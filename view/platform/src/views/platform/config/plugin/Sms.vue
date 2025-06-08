<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        sms_type: 'sms_of_aliyun',
        sms_of_aliyun: {},
    } as { [propName: string]: any },
    rules: {
        sms_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('platform.config.plugin.status.sms_type') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') },
        ],
        'sms_of_aliyun.access_key_id': [
            { required: computed((): boolean => (saveForm.data.sms_type == `sms_of_aliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'sms_of_aliyun.access_key_secret': [
            { required: computed((): boolean => (saveForm.data.sms_type == `sms_of_aliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'sms_of_aliyun.endpoint': [
            { required: computed((): boolean => (saveForm.data.sms_type == `sms_of_aliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'sms_of_aliyun.sign_name': [
            { required: computed((): boolean => (saveForm.data.sms_type == `sms_of_aliyun` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'sms_of_aliyun.template_code': [
            { required: computed((): boolean => (saveForm.data.sms_type == `sms_of_aliyun` ? true : false)), message: t('validation.required') },
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
        <el-form-item :label="t('platform.config.plugin.name.sms_type')" prop="sms_type">
            <el-radio-group v-model="saveForm.data.sms_type">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.sms_type') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.sms_type == 'sms_of_aliyun'">
            <el-form-item :label="t('platform.config.plugin.name.sms_of_aliyun.access_key_id')" prop="sms_of_aliyun.access_key_id">
                <el-input v-model="saveForm.data.sms_of_aliyun.access_key_id" :placeholder="t('platform.config.plugin.name.sms_of_aliyun.access_key_id')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.sms_of_aliyun.access_key_secret')" prop="sms_of_aliyun.access_key_secret">
                <el-input v-model="saveForm.data.sms_of_aliyun.access_key_secret" :placeholder="t('platform.config.plugin.name.sms_of_aliyun.access_key_secret')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.sms_of_aliyun.endpoint')" prop="sms_of_aliyun.endpoint">
                <el-input v-model="saveForm.data.sms_of_aliyun.endpoint" :placeholder="t('platform.config.plugin.name.sms_of_aliyun.endpoint')" :clearable="true" style="max-width: 500px" />
                <el-alert type="info" :show-icon="true" :closable="false">
                    <template #title>
                        <span v-html="t('platform.config.plugin.tip.sms_of_aliyun.endpoint')"></span>
                    </template>
                </el-alert>
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.sms_of_aliyun.sign_name')" prop="sms_of_aliyun.sign_name">
                <el-input v-model="saveForm.data.sms_of_aliyun.sign_name" :placeholder="t('platform.config.plugin.name.sms_of_aliyun.sign_name')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.sms_of_aliyun.template_code')" prop="sms_of_aliyun.template_code">
                <el-input v-model="saveForm.data.sms_of_aliyun.template_code" :placeholder="t('platform.config.plugin.name.sms_of_aliyun.template_code')" :clearable="true" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isSmsSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
