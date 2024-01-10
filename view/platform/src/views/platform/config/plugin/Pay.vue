<script setup lang="tsx">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        payOfAliAppId: '',
        payOfAliPrivateKey: '',
        payOfAliPublicKey: '',
        payOfAliNotifyUrl: '',
        payOfAliOpAppId: '',

        payOfWxAppId: '',
        payOfWxMchid: '',
        payOfWxSerialNo: '',
        payOfWxApiV3Key: '',
        payOfWxPrivateKey: '',
        payOfWxNotifyUrl: '',
    } as { [propName: string]: any },
    rules: {
        payOfAliAppId: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfAliPrivateKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfAliPublicKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfAliNotifyUrl: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        payOfAliOpAppId: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],

        payOfWxAppId: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfWxMchid: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfWxSerialNo: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfWxApiV3Key: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfWxPrivateKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        payOfWxNotifyUrl: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
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
            const param = removeEmptyOfObj(saveForm.data, false)
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
        <el-tabs tab-position="left">
            <el-tab-pane :label="t('platform.config.plugin.label.payOfAli')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.payOfAliAppId')" prop="payOfAliAppId">
                    <el-input v-model="saveForm.data.payOfAliAppId" :placeholder="t('platform.config.plugin.name.payOfAliAppId')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfAliPrivateKey')" prop="payOfAliPrivateKey">
                    <el-input v-model="saveForm.data.payOfAliPrivateKey" type="textarea" :autosize="{ minRows: 5 }" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfAliPublicKey')" prop="payOfAliPublicKey">
                    <el-input v-model="saveForm.data.payOfAliPublicKey" type="textarea" :autosize="{ minRows: 5 }" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfAliNotifyUrl')" prop="payOfAliNotifyUrl">
                    <el-input v-model="saveForm.data.payOfAliNotifyUrl" :placeholder="t('platform.config.plugin.name.payOfAliNotifyUrl')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfAliOpAppId')" prop="payOfAliOpAppId">
                    <el-input v-model="saveForm.data.payOfAliOpAppId" :placeholder="t('platform.config.plugin.name.payOfAliOpAppId')" :clearable="true" style="max-width: 500px" />
                    <label>
                        <el-alert :title="t('platform.config.plugin.tip.payOfAliOpAppId')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </el-form-item>
            </el-tab-pane>

            <el-tab-pane :label="t('platform.config.plugin.label.payOfWx')" :lazy="true">
                <el-form-item :label="t('platform.config.plugin.name.payOfWxAppId')" prop="payOfWxAppId">
                    <el-input v-model="saveForm.data.payOfWxAppId" :placeholder="t('platform.config.plugin.name.payOfWxAppId')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfWxMchid')" prop="payOfWxMchid">
                    <el-input v-model="saveForm.data.payOfWxMchid" :placeholder="t('platform.config.plugin.name.payOfWxMchid')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfWxSerialNo')" prop="payOfWxSerialNo">
                    <el-input v-model="saveForm.data.payOfWxSerialNo" :placeholder="t('platform.config.plugin.name.payOfWxSerialNo')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfWxApiV3Key')" prop="payOfWxApiV3Key">
                    <el-input v-model="saveForm.data.payOfWxApiV3Key" :placeholder="t('platform.config.plugin.name.payOfWxApiV3Key')" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfWxPrivateKey')" prop="payOfWxPrivateKey">
                    <el-input v-model="saveForm.data.payOfWxPrivateKey" type="textarea" :autosize="{ minRows: 5 }" />
                </el-form-item>
                <el-form-item :label="t('platform.config.plugin.name.payOfWxNotifyUrl')" prop="payOfWxNotifyUrl">
                    <el-input v-model="saveForm.data.payOfWxNotifyUrl" :placeholder="t('platform.config.plugin.name.payOfWxNotifyUrl')" :clearable="true" />
                </el-form-item>
            </el-tab-pane>
        </el-tabs>

        <el-form-item>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }} </el-button>
            <el-button type="info" @click="saveForm.reset"> <autoicon-ep-circle-close />{{ t('common.reset') }} </el-button>
        </el-form-item>
    </el-form>
</template>
