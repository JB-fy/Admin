<script setup lang="ts">
const router = useRouter()
const route = useRoute()
const { t } = useI18n()

const loginForm = reactive({
    ref: null as any,
    data: {
        loginName: '',
        password: ''
    },
    rules: {
        loginName: [{ type: 'string', required: true, max: 30, trigger: 'blur', message: t('validation.max.string', { max: 30 }) }],
        password: [{ type: 'string', required: true, min: 6, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 30 }) }]
    } as any,
    loading: false,
    submit: () => {
        loginForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            loginForm.loading = true
            try {
                await useAdminStore().login(loginForm.data.loginName, loginForm.data.password)
                router.replace((route.query.redirect ? route.query.redirect : '/') as string)
            } catch (error) {}
            loginForm.loading = false
        })
    }
})
</script>

<template>
    <div class="particles"></div>
    <ElTag id="login-container">
        <ElDivider>
            <div style="font-size: 25px">{{ t('common.login') }}</div>
        </ElDivider>
        <ElForm :ref="(el: any) => (loginForm.ref = el)" :model="loginForm.data" :rules="loginForm.rules" @keyup.enter="loginForm.submit">
            <ElFormItem prop="loginName">
                <ElInput v-model="loginForm.data.loginName" :placeholder="t('login.name.loginName')">
                    <template #prefix>
                        <AutoiconEpUser />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem prop="password">
                <ElInput v-model="loginForm.data.password" type="password" :placeholder="t('login.name.password')" :show-password="true">
                    <template #prefix>
                        <AutoiconEpLock />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem>
                <ElButton :loading="loginForm.loading" type="primary" @click="loginForm.submit" style="width: 100%">
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
    background-image: url('@/assets/image/login/login-bg.jpg');
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
