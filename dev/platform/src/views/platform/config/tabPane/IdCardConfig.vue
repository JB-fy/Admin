<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置项key，用于向服务器获取对应的配置项value
        idCardType: 'aliyunIdCard',
        aliyunIdCardHost: '',
        aliyunIdCardPath: '',
        aliyunIdCardAppcode: '',
    } as { [propName: string]: any },
    rules: {
        idCardType: [
            { type: 'enum', enum: [`aliyunIdCard`], trigger: 'change', message: t('validation.select') }
        ],
        aliyunIdCardHost: [
            { type: 'url', trigger: 'blur', message: t('validation.url') }
        ],
        aliyunIdCardPath: [
            { type: 'string', trigger: 'blur' }
        ],
        aliyunIdCardAppcode: [
            { type: 'string', trigger: 'blur' }
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
        <ElFormItem :label="t('platform.config.name.idCardType')" prop="idCardType">
            <ElRadioGroup v-model="saveForm.data.idCardType">
                <ElRadio v-for="(item, index) in (tm('platform.config.status.idCardType') as any)" :key="index"
                    :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.idCardType == 'aliyunIdCard'">
            <ElFormItem :label="t('platform.config.name.aliyunIdCardHost')" prop="aliyunIdCardHost">
                <ElInput v-model="saveForm.data.aliyunIdCardHost" :placeholder="t('platform.config.name.aliyunIdCardHost')"
                    :clearable="true" style="max-width: 500px;" />
                <label>
                    <ElAlert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.tip.aliyunIdCardHost')"></span>
                        </template>
                    </ElAlert>
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunIdCardPath')" prop="aliyunIdCardPath">
                <ElInput v-model="saveForm.data.aliyunIdCardPath" :placeholder="t('platform.config.name.aliyunIdCardPath')"
                    :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunIdCardAppcode')" prop="aliyunIdCardAppcode">
                <ElInput v-model="saveForm.data.aliyunIdCardAppcode"
                    :placeholder="t('platform.config.name.aliyunIdCardAppcode')" :clearable="true" />
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
    </ElForm>
</template>