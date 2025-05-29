<script setup lang="tsx">
const { t } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        wx_gzh: {},
    } as { [propName: string]: any },
    rules: {
        'wx_gzh.host': [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        'wx_gzh.app_id': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'wx_gzh.secret': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'wx_gzh.token': [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        'wx_gzh.encoding_aes_key': [{ type: 'string', trigger: 'blur', len: 43, message: t('validation.size.string', { size: 43 }) }],
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
            <el-tab-pane :label="t('platform.config.plugin.label.wx_gzh')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.wx_gzh.host')" prop="wx_gzh.host">
                    <el-input v-model="saveForm.data.wx_gzh.host" :placeholder="t('platform.config.plugin.name.wx_gzh.host')" :clearable="true" style="max-width: 500px" />
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.plugin.tip.wx.host')"></span>
                        </template>
                    </el-alert>
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wx_gzh.app_id')" prop="wx_gzh.app_id">
                    <el-input v-model="saveForm.data.wx_gzh.app_id" :placeholder="t('platform.config.plugin.name.wx_gzh.app_id')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wx_gzh.secret')" prop="wx_gzh.secret">
                    <el-input v-model="saveForm.data.wx_gzh.secret" :placeholder="t('platform.config.plugin.name.wx_gzh.secret')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wx_gzh.token')" prop="wx_gzh.token">
                    <el-input v-model="saveForm.data.wx_gzh.token" :placeholder="t('platform.config.plugin.name.wx_gzh.token')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wx_gzh.encoding_aes_key')" prop="wx_gzh.encoding_aes_key">
                    <el-input v-model="saveForm.data.wx_gzh.encoding_aes_key" :placeholder="t('platform.config.plugin.name.wx_gzh.encoding_aes_key')" :clearable="true" />
                </el-form-item>
            </el-tab-pane>
        </el-tabs>

        <el-form-item>
            <el-button v-if="authAction.isWxSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"><autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
            <el-button type="info" @click="saveForm.reset"><autoicon-ep-circle-close />{{ t('common.reset') }}</el-button>
        </el-form-item>
    </el-form>
</template>
