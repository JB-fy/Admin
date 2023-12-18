<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        aliPayAppId: '',
        aliPaySignType: 'RSA2',
        aliPayPrivateKey: '',
        aliPayPublicKey: '',

        wxPayAppId: '',
        wxPayMchid: '',
        wxPaySerialNo: '',
        wxPayApiV3Key: '',
        wxPayPrivateKey: '',
    } as { [propName: string]: any },
    rules: {
        aliPayAppId: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        aliPaySignType: [
            { type: 'enum', enum: (tm('platform.config.status.aliPaySignType') as any).map((item: any) => item.value), trigger: 'change', message: t('validation.select') },
        ],
        aliPayPrivateKey: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        aliPayPublicKey: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],

        wxPayAppId: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        wxPayMchid: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        wxPaySerialNo: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        wxPayApiV3Key: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        wxPayPrivateKey: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
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
            <ElTabPane :label="t('platform.config.label.aliPay')" :lazy="true">
                <ElFormItem :label="t('platform.config.name.aliPayAppId')" prop="aliPayAppId">
                    <ElInput v-model="saveForm.data.aliPayAppId" :placeholder="t('platform.config.name.aliPayAppId')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.aliPaySignType')" prop="aliPaySignType">
                    <ElRadioGroup v-model="saveForm.data.aliPaySignType">
                        <ElRadio v-for="(item, index) in (tm('platform.config.status.aliPaySignType') as any)" :key="index"
                            :label="item.value">
                            {{ item.label }}
                        </ElRadio>
                    </ElRadioGroup>
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.aliPayPrivateKey')" prop="aliPayPrivateKey">
                    <ElInput v-model="saveForm.data.aliPayPrivateKey"
                        :placeholder="t('platform.config.name.aliPayPrivateKey')" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.aliPayPublicKey')" prop="aliPayPublicKey">
                    <ElInput v-model="saveForm.data.aliPayPublicKey"
                        :placeholder="t('platform.config.name.aliPayPublicKey')" :clearable="true" />
                </ElFormItem>
            </ElTabPane>

            <ElTabPane :label="t('platform.config.label.wxPay')" :lazy="true">
                <ElFormItem :label="t('platform.config.name.wxPayAppId')" prop="wxPayAppId">
                    <ElInput v-model="saveForm.data.wxPayAppId" :placeholder="t('platform.config.name.wxPayAppId')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.wxPayMchid')" prop="wxPayMchid">
                    <ElInput v-model="saveForm.data.wxPayMchid" :placeholder="t('platform.config.name.wxPayMchid')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.wxPaySerialNo')" prop="wxPaySerialNo">
                    <ElInput v-model="saveForm.data.wxPaySerialNo" :placeholder="t('platform.config.name.wxPaySerialNo')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.wxPayApiV3Key')" prop="wxPayApiV3Key">
                    <ElInput v-model="saveForm.data.wxPayApiV3Key" :placeholder="t('platform.config.name.wxPayApiV3Key')"
                        :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('platform.config.name.wxPayPrivateKey')" prop="wxPayPrivateKey">
                    <ElInput v-model="saveForm.data.wxPayPrivateKey"
                        :placeholder="t('platform.config.name.wxPayPrivateKey')" :clearable="true" />
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