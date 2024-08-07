<script setup lang="tsx">
const router = useRouter()
const route = useRoute()
const { t } = useI18n()

const loginForm = reactive({
    ref: null as any,
    data: {} as { [propName: string]: any },
    rules: {
        org_id: [
            { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'blur', min: 1, max: 4294967295, message: t('validation.between.string', { min: 1, max: 4294967295 }), transform: (value: any) => parseInt(value) },
        ],
        login_name: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', max: 20, message: t('validation.max.string', { max: 20 }) },
        ],
        password: [
            { required: true, message: t('validation.required') },
            { type: 'string', trigger: 'blur', min: 6, max: 30, message: t('validation.between.string', { min: 6, max: 30 }) },
        ],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    loading: false,
    submit: () => {
        loginForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            loginForm.loading = true
            try {
                await useAdminStore().login(loginForm.data.org_id, loginForm.data.login_name, loginForm.data.password)
                router.replace((route.query.redirect ? route.query.redirect : '/') as string)
            } finally {
                loginForm.loading = false
            }
        })
    },
})
</script>

<template>
    <div class="particles"></div>
    <el-tag id="login-container">
        <el-divider>
            <div style="font-size: 25px">{{ t('common.login') }}</div>
        </el-divider>
        <el-form :ref="(el: any) => loginForm.ref = el" :model="loginForm.data" :rules="loginForm.rules" @keyup.enter="loginForm.submit">
            <el-form-item prop="org_id">
                <!-- <el-input-number v-model="loginForm.data.org_id" :placeholder="t('login.name.org_id')" :min="1" :max="4294967295" :precision="0" :controls="false" style="width: 100%;" /> -->
                <el-input v-model="loginForm.data.org_id" :placeholder="t('login.name.org_id')">
                    <template #prefix>
                        <autoicon-ep-school />
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item prop="login_name">
                <el-input v-model="loginForm.data.login_name" :placeholder="t('login.name.login_name')">
                    <template #prefix>
                        <autoicon-ep-user />
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="loginForm.data.password" type="password" :placeholder="t('login.name.password')" :show-password="true">
                    <template #prefix>
                        <autoicon-ep-lock />
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button :loading="loginForm.loading" type="primary" @click="loginForm.submit" style="width: 100%">
                    {{ t('common.login') }}
                </el-button>
            </el-form-item>
        </el-form>
    </el-tag>
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
