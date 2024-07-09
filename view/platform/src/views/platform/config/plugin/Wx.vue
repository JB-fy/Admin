<script setup lang="tsx">
const { t } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        wxGzh: {},
    } as { [propName: string]: any },
    rules: {
        'wxGzh.host': [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        'wxGzh.appId': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'wxGzh.secret': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'wxGzh.token': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'wxGzh.encodingAESKey': [{ type: 'string', trigger: 'blur', len: 43, message: t('validation.size.string', { size: 43 }) }],
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
            <el-tab-pane :label="t('platform.config.plugin.label.wxGzh')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.wxGzh.host')" prop="wxGzh.host">
                    <el-input v-model="saveForm.data.wxGzh.host" :placeholder="t('platform.config.plugin.name.wxGzh.host')" :clearable="true" style="max-width: 500px" />
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.plugin.tip.wxHost')"></span>
                        </template>
                    </el-alert>
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzh.appId')" prop="wxGzh.appId">
                    <el-input v-model="saveForm.data.wxGzh.appId" :placeholder="t('platform.config.plugin.name.wxGzh.appId')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzh.secret')" prop="wxGzh.secret">
                    <el-input v-model="saveForm.data.wxGzh.secret" :placeholder="t('platform.config.plugin.name.wxGzh.secret')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzh.token')" prop="wxGzh.token">
                    <el-input v-model="saveForm.data.wxGzh.token" :placeholder="t('platform.config.plugin.name.wxGzh.token')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzh.encodingAESKey')" prop="wxGzh.encodingAESKey">
                    <el-input v-model="saveForm.data.wxGzh.encodingAESKey" :placeholder="t('platform.config.plugin.name.wxGzh.encodingAESKey')" :clearable="true" />
                </el-form-item>
            </el-tab-pane>
        </el-tabs>

        <el-form-item>
            <el-button v-if="authAction.isWxSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
