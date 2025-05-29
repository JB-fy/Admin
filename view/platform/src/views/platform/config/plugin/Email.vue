<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        email_type: 'email_of_common',
        email_code: {},
        email_of_common: {},
    } as { [propName: string]: any },
    rules: {
        'email_code.subject': [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'email_code.template': [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        email_type: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: (tm('platform.config.plugin.status.email_type') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') },
        ],
        'email_of_common.smtp_host': [
            { required: computed((): boolean => (saveForm.data.email_type == `email_of_common` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'email_of_common.smtp_port': [
            { required: computed((): boolean => (saveForm.data.email_type == `email_of_common` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'email_of_common.from_email': [
            { required: computed((): boolean => (saveForm.data.email_type == `email_of_common` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
            { type: 'email', trigger: 'blur', message: t('validation.email') },
        ],
        'email_of_common.password': [
            { required: computed((): boolean => (saveForm.data.email_type == `email_of_common` ? true : false)), message: t('validation.required') },
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
        <el-form-item :label="t('platform.config.plugin.name.email_code.subject')" prop="email_code.subject">
            <el-input v-model="saveForm.data.email_code.subject" :placeholder="t('platform.config.plugin.name.email_code.subject')" :clearable="true" />
        </el-form-item>
        <el-form-item :label="t('platform.config.plugin.name.email_code.template')" prop="email_code.template">
            <el-alert :title="t('platform.config.plugin.tip.email_code.template')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
            <el-input v-model="saveForm.data.email_code.template" type="textarea" :autosize="{ minRows: 3 }" />
        </el-form-item>
        <el-form-item :label="t('platform.config.plugin.name.email_type')" prop="email_type">
            <el-radio-group v-model="saveForm.data.email_type">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.email_type') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.email_type == 'email_of_common'">
            <el-form-item :label="t('platform.config.plugin.name.email_of_common.smtp_host')" prop="email_of_common.smtp_host">
                <el-input v-model="saveForm.data.email_of_common.smtp_host" :placeholder="t('platform.config.plugin.name.email_of_common.smtp_host')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.email_of_common.smtp_port')" prop="email_of_common.smtp_port">
                <el-input v-model="saveForm.data.email_of_common.smtp_port" :placeholder="t('platform.config.plugin.name.email_of_common.smtp_port')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.email_of_common.from_email')" prop="email_of_common.from_email">
                <el-input v-model="saveForm.data.email_of_common.from_email" :placeholder="t('platform.config.plugin.name.email_of_common.from_email')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.email_of_common.password')" prop="email_of_common.password">
                <el-input v-model="saveForm.data.email_of_common.password" :placeholder="t('platform.config.plugin.name.email_of_common.password')" :clearable="true" style="max-width: 500px" />
                <el-alert :title="t('platform.config.plugin.tip.email_of_common.password')" type="info" :show-icon="true" :closable="false" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isSmsSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
