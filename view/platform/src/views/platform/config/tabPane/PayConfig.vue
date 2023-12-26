<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        payOfAliAppId: '',
        payOfAliSignType: 'RSA2',
        payOfAliPrivateKey: '',
        payOfAliPublicKey: '',
        payOfAliNotifyUrl: '',
        payOfAliReturnUrl: '',

        payOfWxAppId: '',
        payOfWxMchid: '',
        payOfWxSerialNo: '',
        payOfWxApiV3Key: '',
        payOfWxPrivateKey: '',
        payOfWxNotifyUrl: '',
    } as { [propName: string]: any },
    rules: {
        payOfAliAppId: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfAliSignType: [
            { type: 'enum', enum: (tm('platform.config.status.payOfAliSignType') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],
        payOfAliPrivateKey: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfAliPublicKey: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfAliNotifyUrl: [
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        payOfAliReturnUrl: [
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],

        payOfWxAppId: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfWxMchid: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfWxSerialNo: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfWxApiV3Key: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfWxPrivateKey: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        payOfWxNotifyUrl: [
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config
            }
        } catch (error) { }
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
            } catch (error) { }
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
    <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto"
        :status-icon="true" :scroll-to-error="false">
        <ElTabs tab-position="left">
            <ElTabPane :label="t('platform.config.label.payOfAli')" :lazy="true">
                <ElFormItem :label="t('platform.config.name.payOfAliAppId')" prop="payOfAliAppId">
                    <ElInput v-model="saveForm.data.payOfAliAppId" :placeholder="t('platform.config.name.payOfAliAppId')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfAliSignType')" prop="payOfAliSignType">
                    <ElRadioGroup v-model="saveForm.data.payOfAliSignType">
                        <ElRadio v-for="(item, index) in (tm('platform.config.status.payOfAliSignType') as any)"
                            :key="index" :label="item.value">
                            {{ item.label }}
                        </ElRadio>
                    </ElRadioGroup>
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfAliPrivateKey')" prop="payOfAliPrivateKey">
					<ElAlert :title="t('platform.config.tip.payOfAliPrivateKey')" type="info" :show-icon="true" :closable="false" />
                    <ElInput v-model="saveForm.data.payOfAliPrivateKey" type="textarea" :autosize="{ minRows: 5 }" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfAliPublicKey')" prop="payOfAliPublicKey">
                    <ElInput v-model="saveForm.data.payOfAliPublicKey" type="textarea" :autosize="{ minRows: 5 }" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfAliNotifyUrl')" prop="payOfAliNotifyUrl">
                    <ElInput v-model="saveForm.data.payOfAliNotifyUrl"
                        :placeholder="t('platform.config.name.payOfAliNotifyUrl')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfAliReturnUrl')" prop="payOfAliReturnUrl">
                    <ElInput v-model="saveForm.data.payOfAliReturnUrl"
                        :placeholder="t('platform.config.name.payOfAliReturnUrl')" :clearable="true" />
                </ElFormItem>
            </ElTabPane>

            <ElTabPane :label="t('platform.config.label.payOfWx')" :lazy="true">
                <ElFormItem :label="t('platform.config.name.payOfWxAppId')" prop="payOfWxAppId">
                    <ElInput v-model="saveForm.data.payOfWxAppId" :placeholder="t('platform.config.name.payOfWxAppId')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfWxMchid')" prop="payOfWxMchid">
                    <ElInput v-model="saveForm.data.payOfWxMchid" :placeholder="t('platform.config.name.payOfWxMchid')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfWxSerialNo')" prop="payOfWxSerialNo">
                    <ElInput v-model="saveForm.data.payOfWxSerialNo"
                        :placeholder="t('platform.config.name.payOfWxSerialNo')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfWxApiV3Key')" prop="payOfWxApiV3Key">
                    <ElInput v-model="saveForm.data.payOfWxApiV3Key"
                        :placeholder="t('platform.config.name.payOfWxApiV3Key')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfWxPrivateKey')" prop="payOfWxPrivateKey">
                    <ElInput v-model="saveForm.data.payOfWxPrivateKey" type="textarea" :autosize="{ minRows: 5 }" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.payOfWxNotifyUrl')" prop="payOfWxNotifyUrl">
                    <ElInput v-model="saveForm.data.payOfWxNotifyUrl"
                        :placeholder="t('platform.config.name.payOfWxNotifyUrl')" :clearable="true" />
                </ElFormItem>
            </ElTabPane>
        </ElTabs>

        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                <AutoiconEpCircleCheck />{{ t('common.save') }}
            </ElButton>
            <ElButton type="info" @click="saveForm.reset">
                <AutoiconEpCircleClose />{{ t('common.reset') }}
            </ElButton>
        </ElFormItem>
    </ElForm>
</template>