<script setup lang="tsx">
const { t, tm } = useI18n()

const authAction = inject('authAction') as { [propName: string]: boolean }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置键，用于向服务器获取对应的配置值
        wxGzhHost: '',
        wxGzhAppId: '',
        wxGzhSecret: '',
        wxGzhToken: '',
        wxGzhEncodingAESKey: '',
    } as { [propName: string]: any },
    rules: {
        wxGzhHost: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        wxGzhAppId: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        wxGzhSecret: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        wxGzhToken: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        wxGzhEncodingAESKey: [{ type: 'string', trigger: 'blur', len: 43, message: t('validation.size.string', { size: 43 }) }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    initData: async () => {
        const param = { config_key_arr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config,
            }
        } catch (error) {
            /* eslint-disable-next-line no-empty */
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
        <el-tabs tab-position="left">
            <el-tab-pane :label="t('platform.config.plugin.label.wxGzh')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.wxGzhHost')" prop="wxGzhHost">
                    <el-input v-model="saveForm.data.wxGzhHost" :placeholder="t('platform.config.plugin.name.wxGzhHost')" :clearable="true" style="max-width: 500px" />
                    <el-alert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.plugin.tip.wxHost')"></span>
                        </template>
                    </el-alert>
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzhAppId')" prop="wxGzhAppId">
                    <el-input v-model="saveForm.data.wxGzhAppId" :placeholder="t('platform.config.plugin.name.wxGzhAppId')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzhSecret')" prop="wxGzhSecret">
                    <el-input v-model="saveForm.data.wxGzhSecret" :placeholder="t('platform.config.plugin.name.wxGzhSecret')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzhToken')" prop="wxGzhToken">
                    <el-input v-model="saveForm.data.wxGzhToken" :placeholder="t('platform.config.plugin.name.wxGzhToken')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.wxGzhEncodingAESKey')" prop="wxGzhEncodingAESKey">
                    <el-input v-model="saveForm.data.wxGzhEncodingAESKey" :placeholder="t('platform.config.plugin.name.wxGzhEncodingAESKey')" :clearable="true" />
                </el-form-item>
            </el-tab-pane>
        </el-tabs>

        <el-form-item>
            <el-button v-if="authAction.isWxSave" type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
