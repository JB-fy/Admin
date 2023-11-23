<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        pushType: 'txTpns',
        txTpnsHost: '',
        txTpnsAccessIDOfAndroid: '',
        txTpnsSecretKeyOfAndroid: '',
        txTpnsAccessIDOfIos: '',
        txTpnsSecretKeyOfIos: '',
        txTpnsAccessIDOfMacOS: '',
        txTpnsSecretKeyOfMacOS: '',
    } as { [propName: string]: any },
    rules: {
        pushType: [
            { type: 'enum', enum: [`txTpns`], trigger: 'change', message: t('validation.select') },
        ],
        txTpnsHost: [
            { type: 'url', trigger: 'blur', message: t('validation.url') },
        ],
        txTpnsAccessIDOfAndroid: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        txTpnsSecretKeyOfAndroid: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        txTpnsAccessIDOfIos: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        txTpnsSecretKeyOfIos: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        txTpnsAccessIDOfMacOS: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        txTpnsSecretKeyOfMacOS: [
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
    }
})

saveForm.initData()
</script>

<template>
    <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto"
        :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('platform.config.name.pushType')" prop="pushType">
            <ElRadioGroup v-model="saveForm.data.pushType">
                <ElRadio v-for="(item, index) in (tm('platform.config.status.pushType') as any)" :key="index"
                    :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.pushType == 'txTpns'">
            <ElFormItem :label="t('platform.config.name.txTpnsHost')" prop="txTpnsHost">
                <ElInput v-model="saveForm.data.txTpnsHost" :placeholder="t('platform.config.name.txTpnsHost')"
                    :clearable="true" style="max-width: 500px;" />
                <label>
                    <ElAlert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.tip.txTpnsHost')"></span>
                        </template>
                    </ElAlert>
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.txTpnsAccessIDOfAndroid')" prop="txTpnsAccessIDOfAndroid">
                <ElInput v-model="saveForm.data.txTpnsAccessIDOfAndroid"
                    :placeholder="t('platform.config.name.txTpnsAccessIDOfAndroid')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.txTpnsSecretKeyOfAndroid')" prop="txTpnsSecretKeyOfAndroid">
                <ElInput v-model="saveForm.data.txTpnsSecretKeyOfAndroid"
                    :placeholder="t('platform.config.name.txTpnsSecretKeyOfAndroid')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.txTpnsAccessIDOfIos')" prop="txTpnsAccessIDOfIos">
                <ElInput v-model="saveForm.data.txTpnsAccessIDOfIos"
                    :placeholder="t('platform.config.name.txTpnsAccessIDOfIos')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.txTpnsSecretKeyOfIos')" prop="txTpnsSecretKeyOfIos">
                <ElInput v-model="saveForm.data.txTpnsSecretKeyOfIos"
                    :placeholder="t('platform.config.name.txTpnsSecretKeyOfIos')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.txTpnsAccessIDOfMacOS')" prop="txTpnsAccessIDOfMacOS">
                <ElInput v-model="saveForm.data.txTpnsAccessIDOfMacOS"
                    :placeholder="t('platform.config.name.txTpnsAccessIDOfMacOS')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.txTpnsSecretKeyOfMacOS')" prop="txTpnsSecretKeyOfMacOS">
                <ElInput v-model="saveForm.data.txTpnsSecretKeyOfMacOS"
                    :placeholder="t('platform.config.name.txTpnsSecretKeyOfMacOS')" :clearable="true" />
            </ElFormItem>
        </template>

        <ElFormItem>
            <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                <AutoiconEpCircleCheck />{{ t('common.save') }}
            </ElButton>
            <ElButton type="info" @click="saveForm.reset">
                <AutoiconEpCircleClose />{{ t('common.reset') }}
            </ElButton>
        </ElFormItem>
</ElForm></template>