<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        emailType: 'emailOfCommon',
        emailCode: {},
        emailOfCommon: {},
    } as { [propName: string]: any },
    rules: {
        'emailCode.subject': [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'emailCode.template': [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        emailType: [
            { required: true, message: t('validation.required') },
            { type: 'enum', trigger: 'change', enum: [`emailOfCommon`], message: t('validation.select') },
        ],
        'emailOfCommon.smtpHost': [
            { required: computed((): boolean => (saveForm.data.emailType == `emailOfCommon` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'emailOfCommon.smtpPort': [
            { required: computed((): boolean => (saveForm.data.emailType == `emailOfCommon` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        'emailOfCommon.fromEmail': [
            { required: computed((): boolean => (saveForm.data.emailType == `emailOfCommon` ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', message: t('validation.input') },
            { type: 'email', trigger: 'blur', message: t('validation.email') },
        ],
        'emailOfCommon.password': [
            { required: computed((): boolean => (saveForm.data.emailType == `emailOfCommon` ? true : false)), message: t('validation.required') },
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
        <el-form-item :label="t('platform.config.plugin.name.emailCode.subject')" prop="emailCode.subject">
            <el-input v-model="saveForm.data.emailCode.subject" :placeholder="t('platform.config.plugin.name.emailCode.subject')" :clearable="true" />
        </el-form-item>
        <el-form-item :label="t('platform.config.plugin.name.emailCode.template')" prop="emailCode.template">
            <el-alert :title="t('platform.config.plugin.tip.emailCode.template')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
            <el-input v-model="saveForm.data.emailCode.template" type="textarea" :autosize="{ minRows: 3 }" />
        </el-form-item>
        <el-form-item :label="t('platform.config.plugin.name.emailType')" prop="emailType">
            <el-radio-group v-model="saveForm.data.emailType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.emailType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.emailType == 'emailOfCommon'">
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommon.smtpHost')" prop="emailOfCommon.smtpHost">
                <el-input v-model="saveForm.data.emailOfCommon.smtpHost" :placeholder="t('platform.config.plugin.name.emailOfCommon.smtpHost')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommon.smtpPort')" prop="emailOfCommon.smtpPort">
                <el-input v-model="saveForm.data.emailOfCommon.smtpPort" :placeholder="t('platform.config.plugin.name.emailOfCommon.smtpPort')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommon.fromEmail')" prop="emailOfCommon.fromEmail">
                <el-input v-model="saveForm.data.emailOfCommon.fromEmail" :placeholder="t('platform.config.plugin.name.emailOfCommon.fromEmail')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommon.password')" prop="emailOfCommon.password">
                <el-input v-model="saveForm.data.emailOfCommon.password" :placeholder="t('platform.config.plugin.name.emailOfCommon.password')" :clearable="true" style="max-width: 500px" />
                <el-alert :title="t('platform.config.plugin.tip.emailOfCommon.password')" type="info" :show-icon="true" :closable="false" />
            </el-form-item>
        </template>

        <el-form-item>
            <el-button v-if="authAction.isSmsSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
