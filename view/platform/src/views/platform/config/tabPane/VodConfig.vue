<script setup lang="ts">
const { t, tm } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置Key，用于向服务器获取对应的配置值
        vodType: 'aliyunVod',
        aliyunVodAccessKeyId: '',
        aliyunVodAccessKeySecret: '',
        aliyunVodEndpoint: '',
        aliyunVodRoleArn: '',
    } as { [propName: string]: any },
    rules: {
        vodType: [
            { type: 'enum', enum: [`aliyunVod`], trigger: 'change', message: t('validation.select') },
        ],
        aliyunVodAccessKeyId: [
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
        ],
        aliyunVodAccessKeySecret: [
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') },
        ],
        aliyunVodEndpoint: [
            { type: 'string', trigger: 'blur', message: t('validation.input') },
        ],
        aliyunVodRoleArn: [
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
        <ElFormItem :label="t('platform.config.name.vodType')" prop="vodType">
            <ElRadioGroup v-model="saveForm.data.vodType">
                <ElRadio v-for="(item, index) in (tm('platform.config.status.vodType') as any)" :key="index"
                    :label="item.value">
                    {{ item.label }}
                </ElRadio>
            </ElRadioGroup>
        </ElFormItem>

        <template v-if="saveForm.data.vodType == 'aliyunVod'">
            <ElFormItem :label="t('platform.config.name.aliyunVodAccessKeyId')" prop="aliyunVodAccessKeyId">
                <ElInput v-model="saveForm.data.aliyunVodAccessKeyId"
                    :placeholder="t('platform.config.name.aliyunVodAccessKeyId')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunVodAccessKeySecret')" prop="aliyunVodAccessKeySecret">
                <ElInput v-model="saveForm.data.aliyunVodAccessKeySecret"
                    :placeholder="t('platform.config.name.aliyunVodAccessKeySecret')" :clearable="true" />
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunVodEndpoint')" prop="aliyunVodEndpoint">
                <ElInput v-model="saveForm.data.aliyunVodEndpoint"
                    :placeholder="t('platform.config.name.aliyunVodEndpoint')" :clearable="true"
                    style="max-width: 500px;" />
                <label>
                    <ElAlert type="info" :show-icon="true" :closable="false">
                        <template #title>
                            <span v-html="t('platform.config.tip.aliyunVodEndpoint')"></span>
                        </template>
                    </ElAlert>
                </label>
            </ElFormItem>
            <ElFormItem :label="t('platform.config.name.aliyunVodRoleArn')" prop="aliyunVodRoleArn">
                <ElInput v-model="saveForm.data.aliyunVodRoleArn" :placeholder="t('platform.config.name.aliyunVodRoleArn')"
                    :clearable="true" style="max-width: 500px;" />
                <label>
                    <ElAlert :title="t('platform.config.tip.aliyunVodRoleArn')" type="info" :show-icon="true"
                        :closable="false" />
                </label>
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