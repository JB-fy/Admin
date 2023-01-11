<script setup lang="ts">
import md5 from 'js-md5'

const { t } = useI18n()
const adminStore = useAdminStore()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        nickname: adminStore.info.nickname,
        avatar: adminStore.info.avatar,
        password: '',
        checkPassword: '',
        oldPassword: '',
    } as { [propName: string]: any },
    rules: {
        nickname: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        avatar: [
            { type: 'string', min: 1, max: 120, trigger: 'change', message: t('validation.upload') }
        ],
        password: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        checkPassword: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            {
                required: computed((): boolean => { return saveForm.data.password ? true : false; }),
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password != saveForm.data.checkPassword) {
                        callback(new Error())
                    }
                    callback()
                }, trigger: 'blur', message: t('validation.checkPassword')
            }
        ],
        oldPassword: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            {
                required: computed((): boolean => { return saveForm.data.password ? true : false; }), trigger: 'blur', message: t('validation.oldPassword')
            }
        ],
    } as any,
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data, false)
            param.password ? param.password = md5(param.password) : delete param.password
            delete param.checkPassword
            param.oldPassword ? param.oldPassword = md5(param.oldPassword) : delete param.oldPassword
            try {
                await request('login/updateInfo', param, true)
                //成功则更新用户信息
                for (let k in param) {
                    switch (k) {
                        case 'nickname':
                        case 'avatar':
                            adminStore.info[k] = param[k]
                            break;
                    }
                }
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
    <ElContainer class="common-container">
        <ElMain>
            <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                label-width="auto" :status-icon="true" :scroll-to-error="false">
                <ElFormItem :label="t('common.name.nickname')" prop="nickname">
                    <ElInput v-model="saveForm.data.nickname" :placeholder="t('common.name.nickname')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.avatar')" prop="avatar">
                    <MyUpload v-model="saveForm.data.avatar" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.newPassword')" prop="password">
                    <ElInput v-model="saveForm.data.password" :placeholder="t('common.name.newPassword')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('common.name.checkPassword')" prop="checkPassword">
                    <ElInput v-model="saveForm.data.checkPassword" :placeholder="t('common.name.checkPassword')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('common.name.oldPassword')" prop="oldPassword">
                    <ElInput v-model="saveForm.data.oldPassword" :placeholder="t('common.name.oldPassword')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.updatePasswordRequired')" type="info" :show-icon="true"
                            :closable="false" />
                    </label>
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
        </ElMain>
    </ElContainer>
</template>