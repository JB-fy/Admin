<script setup lang="ts">
const { t } = useI18n()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        nickname: '',
    } as { [propName: string]: any },
    rules: {
        nickname: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ]
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = {
                ...removeEmptyOfObj(saveForm.data, false)
            }
            try {
                await request('platform/config/save', param, true)
            } catch (error) { }
            saveForm.loading = false
        })
    },
    reset: () => {
        saveForm.ref.resetFields()
    }
})
</script>

<template>
    <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules" label-width="auto"
        :status-icon="true" :scroll-to-error="false">
        <ElFormItem :label="t('common.name.nickname')" prop="nickname">
            <ElInput v-model="saveForm.data.nickname" :placeholder="t('common.name.nickname')" minlength="1"
                maxlength="30" :show-word-limit="true" :clearable="true" />
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