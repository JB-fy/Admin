<script setup lang="tsx">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        pushType: 'pushOfTx',
        pushOfTxHost: '',
        pushOfTxAndroidAccessID: '',
        pushOfTxAndroidSecretKey: '',
        pushOfTxIosAccessID: '',
        pushOfTxIosSecretKey: '',
        pushOfTxMacOSAccessID: '',
        pushOfTxMacOSSecretKey: ''
    } as { [propName: string]: any },
    rules: {
        pushType: [{ type: 'enum', enum: [`pushOfTx`], trigger: 'change', message: t('validation.select') }],
        pushOfTxHost: [{ type: 'url', trigger: 'blur', message: t('validation.url') }],
        pushOfTxAndroidAccessID: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxAndroidSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxIosAccessID: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxIosSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxMacOSAccessID: [{ type: 'string', trigger: 'blur', message: t('validation.input') }],
        pushOfTxMacOSSecretKey: [{ type: 'string', trigger: 'blur', message: t('validation.input') }]
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/config/get', param)
            saveForm.data = {
                ...saveForm.data,
                ...res.data.config
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
    }
})

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => (saveForm.ref = el)" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('platform.config.name.pushType')" prop="pushType">
            <ElRadioGroup v-model="saveForm.data.pushType">
                <ElRadio v-for="(item, index) in tm('platform.config.status.pushType') as any" :key="index" :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.pushType == 'pushOfTx'">
            <ElFormItem :label="t('platform.config.name.pushOfTxHost')" prop="pushOfTxHost">
                <ElInput v-model="saveForm.data.pushOfTxHost" :placeholder="t('platform.config.name.pushOfTxHost')" :clearable="true" style="max-width: 500px" />
                <label>
                    <ElAlert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.tip.pushOfTxHost')"></span>
                        </template>
                    </ElAlert>
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.pushOfTxAndroidAccessID')" prop="pushOfTxAndroidAccessID">
                <ElInput v-model="saveForm.data.pushOfTxAndroidAccessID" :placeholder="t('platform.config.name.pushOfTxAndroidAccessID')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.pushOfTxAndroidSecretKey')" prop="pushOfTxAndroidSecretKey">
                <ElInput v-model="saveForm.data.pushOfTxAndroidSecretKey" :placeholder="t('platform.config.name.pushOfTxAndroidSecretKey')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.pushOfTxIosAccessID')" prop="pushOfTxIosAccessID">
                <ElInput v-model="saveForm.data.pushOfTxIosAccessID" :placeholder="t('platform.config.name.pushOfTxIosAccessID')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.pushOfTxIosSecretKey')" prop="pushOfTxIosSecretKey">
                <ElInput v-model="saveForm.data.pushOfTxIosSecretKey" :placeholder="t('platform.config.name.pushOfTxIosSecretKey')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.pushOfTxMacOSAccessID')" prop="pushOfTxMacOSAccessID">
                <ElInput v-model="saveForm.data.pushOfTxMacOSAccessID" :placeholder="t('platform.config.name.pushOfTxMacOSAccessID')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.pushOfTxMacOSSecretKey')" prop="pushOfTxMacOSSecretKey">
                <ElInput v-model="saveForm.data.pushOfTxMacOSSecretKey" :placeholder="t('platform.config.name.pushOfTxMacOSSecretKey')" :clearable="true" />
            </ElFormItem>
        </template>

        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <AutoiconEpCircleCheck />{{ t('common.save') }} </ElButton>
            <ElButton type="info" @click="saveForm.reset"> <AutoiconEpCircleClose />{{ t('common.reset') }} </ElButton>
        </ElFormItem>
    </ElForm>
</template>
