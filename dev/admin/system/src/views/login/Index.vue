<script setup lang="ts">
const router = useRouter()
const route = useRoute()
const { t } = useI18n()

const form = reactive({
    login: {
        ref: null as any,
        data: {
            account: '',
            password: ''
        },
        rules: {
            account: [
                { type: 'string', required: true, min: 4, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 4, max: 30 }) }
            ],
            password: [
                { type: 'string', required: true, min: 6, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 30 }) }
            ]
        },
        loading: false,
        submit: () => {
            form.login.ref.validate(async (valid: boolean) => {
                if (!valid) {
                    return false
                }
                form.login.loading = true
                try {
                    await useAdminStore().login(form.login.data.account, form.login.data.password)
                    router.replace(<string>(route.query.redirect ? route.query.redirect : '/'))
                } catch (error) {
                    await errorHandle(<Error>error)
                }
                form.login.loading = false
            })
        }
    }
})
</script>

<template>
    <div class="particles"></div>
    <ElTag id="login-container">
        <ElDivider>
            <div style="font-size: 25px;">{{ t('common.login') }}</div>
        </ElDivider>
        <ElForm :ref="(el:any) => { form.login.ref = el }" :model="form.login.data" :rules="form.login.rules">
            <ElFormItem prop="account">
                <ElInput v-model="form.login.data.account" :placeholder="t('common.account')">
                    <template #prefix>
                        <AutoiconEpUser />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem prop="password">
                <ElInput v-model="form.login.data.password" type="password" :placeholder="t('common.password')"
                    :show-password="true" @keyup.enter="form.login.submit">
                    <template #prefix>
                        <AutoiconEpLock />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem>
                <ElButton :loading="form.login.loading" type="primary" @click="form.login.submit" style="width:100%;">
                    {{ t('common.login') }}
                </ElButton>
            </ElFormItem>
        </ElForm>
    </ElTag>
</template>

<style scoped>
.particles {
    width: 100%;
    height: 100vh;
    overflow: hidden;
    background-image: url('@/assets/login-bg.jpg');
    background-position: center center;
    background-size: cover;
}

#login-container {
    background-color: #fff;
    width: 250px;
    height: 250px;
    margin: auto;
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
}
</style>