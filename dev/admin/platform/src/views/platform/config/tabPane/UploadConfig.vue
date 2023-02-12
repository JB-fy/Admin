<script setup lang="ts">
const { t } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: { //此处必须列出全部需要设置的配置项key，用于向服务器获取对应的配置项value
        aliyunOssAccessId: '',
        aliyunOssAccessSecret: '',
        aliyunOssHost: '',
        aliyunOssBucket: '',
    } as { [propName: string]: any },
    rules: {
        aliyunOssAccessId: [
            //{ type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        aliyunOssAccessSecret: [
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        aliyunOssHost: [
            { type: 'url', trigger: 'blur', message: t('validation.url') }
        ],
        aliyunOssBucket: [
            { type: 'string', trigger: 'blur' }
        ],
    } as any,
    initData: async () => {
        const param = { configKeyArr: Object.keys(saveForm.data) }
        try {
            const res = await request('platform/config/get', param)
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
                await request('platform/config/save', param, true)
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
        <ElFormItem :label="t('view.platform.config.name.aliyunOssAccessId')" prop="aliyunOssAccessId">
            <!-- <ElInput v-model="saveForm.data.aliyunOssAccessId"
                :placeholder="t('view.platform.config.name.aliyunOssAccessId')" minlength="1" maxlength="30"
                :show-word-limit="true" :clearable="true" /> -->
            <ElInput v-model="saveForm.data.aliyunOssAccessId"
                :placeholder="t('view.platform.config.name.aliyunOssAccessId')" :clearable="true" />
        </ElFormItem>
        <ElFormItem :label="t('view.platform.config.name.aliyunOssAccessSecret')" prop="aliyunOssAccessSecret">
            <ElInput v-model="saveForm.data.aliyunOssAccessSecret"
                :placeholder="t('view.platform.config.name.aliyunOssAccessSecret')" :clearable="true" />
        </ElFormItem>
        <ElFormItem :label="t('view.platform.config.name.aliyunOssHost')" prop="aliyunOssHost">
            <ElInput v-model="saveForm.data.aliyunOssHost" :placeholder="t('view.platform.config.name.aliyunOssHost')"
                :clearable="true" />
        </ElFormItem>
        <ElFormItem :label="t('view.platform.config.name.aliyunOssBucket')" prop="aliyunOssBucket">
            <ElInput v-model="saveForm.data.aliyunOssBucket"
                :placeholder="t('view.platform.config.name.aliyunOssBucket')" :clearable="true" />
        </ElFormItem>
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