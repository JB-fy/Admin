<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        emailType: 'emailOfCommon',
        emailOfCommonSmtpHost: '',
        emailOfCommonSmtpPort: '',
        emailOfCommonFromEmail: '',
        emailOfCommonPassword: '',
        emailCodeSubject: '',
        emailCodeTemplate: '',
    } as { [propName: string]: any },
    rules: {
        emailType: [{ type: 'enum', trigger: 'change', enum: [`emailOfCommon`], message: t('validation.select') }],
        emailOfCommonSmtpHost: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        emailOfCommonSmtpPort: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        emailOfCommonFromEmail: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
            { type: 'email', trigger: 'blur', message: t('validation.email') },
        ],
        emailOfCommonPassword: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        emailCodeSubject: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        emailCodeTemplate: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
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
        <el-form-item :label="t('platform.config.plugin.name.emailType')" prop="emailType">
            <el-radio-group v-model="saveForm.data.emailType">
                <el-radio v-for="(item, index) in tm('platform.config.plugin.status.emailType') as any" :key="index" :value="item.value">
                    {{ item.label }}
                </el-radio>
            </el-radio-group>
        </el-form-item>

        <template v-if="saveForm.data.emailType == 'emailOfCommon'">
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommonSmtpHost')" prop="emailOfCommonSmtpHost">
                <el-input v-model="saveForm.data.emailOfCommonSmtpHost" :placeholder="t('platform.config.plugin.name.emailOfCommonSmtpHost')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommonSmtpPort')" prop="emailOfCommonSmtpPort">
                <el-input v-model="saveForm.data.emailOfCommonSmtpPort" :placeholder="t('platform.config.plugin.name.emailOfCommonSmtpPort')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommonFromEmail')" prop="emailOfCommonFromEmail">
                <el-input v-model="saveForm.data.emailOfCommonFromEmail" :placeholder="t('platform.config.plugin.name.emailOfCommonFromEmail')" :clearable="true" />
            </el-form-item>
            <el-form-item :label="t('platform.config.plugin.name.emailOfCommonPassword')" prop="emailOfCommonPassword">
                <el-input v-model="saveForm.data.emailOfCommonPassword" :placeholder="t('platform.config.plugin.name.emailOfCommonPassword')" :clearable="true" style="max-width: 500px" />
                <el-alert :title="t('platform.config.plugin.tip.emailOfCommonPassword')" type="info" :show-icon="true" :closable="false" />
            </el-form-item>
        </template>

        <el-form-item :label="t('platform.config.plugin.name.emailCodeSubject')" prop="emailCodeSubject">
            <el-input v-model="saveForm.data.emailCodeSubject" :placeholder="t('platform.config.plugin.name.emailCodeSubject')" :clearable="true" />
        </el-form-item>
        <el-form-item :label="t('platform.config.plugin.name.emailCodeTemplate')" prop="emailCodeTemplate">
            <el-alert :title="t('platform.config.plugin.tip.emailCodeTemplate')" type="info" :show-icon="true" :closable="false" style="width: 100%" />
                    <el-input v-model="saveForm.data.emailCodeTemplate" type="textarea" :autosize="{ minRows: 3 }" />
        </el-form-item>
        <el-form-item>
            <el-button v-if="authAction.isSmsSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
