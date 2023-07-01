<script setup lang="ts">
import md5 from 'js-md5'

const { t } = useI18n()
const adminStore = useAdminStore()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        account: '',
        phone: '',
        nickname: adminStore.info.nickname,
        avatar: adminStore.info.avatar,
        password: '',
        repeatPassword: '',
        checkPassword: '',
    } as { [propName: string]: any },
    rules: {
        account: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^(?!\d*$)[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.account') }
        ],
        phone: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^1[3-9]\d{9}$/, trigger: 'blur', message: t('validation.phone') }
        ],
        nickname: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            { pattern: /^[\p{L}\p{M}\p{N}_-]+$/u, trigger: 'blur', message: t('validation.alpha_dash') }
        ],
        avatar: [
			{ type: 'url', trigger: 'change', message: t('validation.upload') },
			{ type: 'string', min: 1, max: 120, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 120 }) }
        ],
        password: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        repeatPassword: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            {
                required: computed((): boolean => { return saveForm.data.password ? true : false; }),
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password != saveForm.data.repeatPassword) {
                        callback(new Error())
                    }
                    callback()
                }, trigger: 'blur', message: t('validation.repeatPassword')
            }
        ],
        checkPassword: [
            { type: 'string', min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) },
            {
                required: computed((): boolean => { return saveForm.data.account || saveForm.data.phone || saveForm.data.password ? true : false; }), trigger: 'blur', message: t('profile.tip.checkPassword')
            },
            {
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password && saveForm.data.password == saveForm.data.checkPassword) {
                        callback(new Error())
                    }
                    callback()
                }, trigger: 'blur', message: t('validation.newPasswordDiffOldPassword')
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
            param.account || delete param.account
            param.phone || delete param.phone
            param.password ? param.password = md5(param.password) : delete param.password
            delete param.repeatPassword
            param.checkPassword ? param.checkPassword = md5(param.checkPassword) : delete param.checkPassword
            try {
                await request('/login/update', param, true)
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
    }
})
</script>

<template>
    <ElContainer class="common-container">
        <ElMain>
            <ElForm :ref="(el: any) => { saveForm.ref = el }" :model="saveForm.data" :rules="saveForm.rules"
                label-width="auto" :status-icon="true" :scroll-to-error="false">
                <ElFormItem :label="t('profile.name.account')" prop="account">
                    <ElInput v-model="saveForm.data.account" :placeholder="t('profile.name.account')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
                    <label>
                        <ElAlert
                            :title="t('profile.tip.account', { account: adminStore.info.account ? adminStore.info.account : t('common.tip.notSet') })"
                            type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('profile.name.phone')" prop="phone">
                    <ElInput v-model="saveForm.data.phone" :placeholder="t('profile.name.phone')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" style="max-width: 250px;" />
                    <label>
                        <ElAlert
                            :title="t('profile.tip.phone', { phone: adminStore.info.phone ? adminStore.info.phone : t('common.tip.notSet') })"
                            type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('common.name.nickname')" prop="nickname">
                    <ElInput v-model="saveForm.data.nickname" :placeholder="t('common.name.nickname')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.avatar')" prop="avatar">
                    <MyUpload v-model="saveForm.data.avatar" accept="image/*" />
                </ElFormItem>
                <ElFormItem :label="t('common.name.newPassword')" prop="password">
                    <ElInput v-model="saveForm.data.password" :placeholder="t('common.name.newPassword')" minlength="1"
                        maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('common.name.repeatPassword')" prop="repeatPassword">
                    <ElInput v-model="saveForm.data.repeatPassword" :placeholder="t('common.name.repeatPassword')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem :label="t('common.name.oldPassword')" prop="checkPassword">
                    <ElInput v-model="saveForm.data.checkPassword" :placeholder="t('common.name.oldPassword')"
                        minlength="1" maxlength="30" :show-word-limit="true" :clearable="true" :show-password="true"
                        style="max-width: 250px;" />
                    <label>
                        <ElAlert :title="t('profile.tip.checkPassword')" type="info" :show-icon="true"
                            :closable="false" />
                    </label>
                </ElFormItem>
                <ElFormItem>
                    <ElButton type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                        <AutoiconEpCircleCheck />{{ t('common.save') }}
                    </ElButton>
                </ElFormItem>
            </ElForm>
        </ElMain>
    </ElContainer>
</template>