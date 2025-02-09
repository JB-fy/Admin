<script setup lang="tsx">
const { t } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        oneClickOfWx: {},
        oneClickOfYidun: {},
    } as { [propName: string]: any },
    rules: {
        'oneClickOfWx.host': [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        'oneClickOfWx.appId': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'oneClickOfWx.secret': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'oneClickOfYidun.secretId': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'oneClickOfYidun.secretKey': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'oneClickOfYidun.businessId': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
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
        <el-tabs tab-position="left">
            <el-tab-pane :label="t('platform.config.plugin.label.oneClickOfWx')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.oneClickOfWx.host')" prop="oneClickOfWx.host">
                    <el-input v-model="saveForm.data.oneClickOfWx.host" :placeholder="t('platform.config.plugin.name.oneClickOfWx.host')" :clearable="true" style="max-width: 500px" />
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.plugin.tip.wx.host')"></span>
                        </template>
                    </el-alert>
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.oneClickOfWx.appId')" prop="oneClickOfWx.appId">
                    <el-input v-model="saveForm.data.oneClickOfWx.appId" :placeholder="t('platform.config.plugin.name.oneClickOfWx.appId')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.oneClickOfWx.secret')" prop="oneClickOfWx.secret">
                    <el-input v-model="saveForm.data.oneClickOfWx.secret" :placeholder="t('platform.config.plugin.name.oneClickOfWx.secret')" :clearable="true" />
                </el-form-item>
            </el-tab-pane>

            <el-tab-pane :label="t('platform.config.plugin.label.oneClickOfYidun')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.oneClickOfYidun.secretId')" prop="oneClickOfYidun.secretId">
                    <el-input v-model="saveForm.data.oneClickOfYidun.secretId" :placeholder="t('platform.config.plugin.name.oneClickOfYidun.secretId')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.oneClickOfYidun.secretKey')" prop="oneClickOfYidun.secretKey">
                    <el-input v-model="saveForm.data.oneClickOfYidun.secretKey" :placeholder="t('platform.config.plugin.name.oneClickOfYidun.secretKey')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.oneClickOfYidun.businessId')" prop="oneClickOfYidun.businessId">
                    <el-input v-model="saveForm.data.oneClickOfYidun.businessId" :placeholder="t('platform.config.plugin.name.oneClickOfYidun.businessId')" :clearable="true" />
                </el-form-item>
            </el-tab-pane>
        </el-tabs>

        <el-form-item>
            <el-button v-if="authAction.isOneClickSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
