<script setup lang="ts">
import VueParticles from 'vue-particles/src/vue-particles/vue-particles.vue'
import { useUserStore } from '@/stores/user';

const router = useRouter()
const route = useRoute()
const form = reactive({
    login: {
        ref: null,
        data: {
            account: '',
            password: ''
        },
        rules: {
            account: [
                { type: 'string', required: true, min: 4, max: 30, trigger: 'blur', message: '长度在4到30个字符之间' }
            ],
            password: [
                { type: 'string', required: true, min: 6, max: 30, trigger: 'blur', message: '长度在6到30个字符之间' }
            ]
        },
        loading: false,
        submit: () => {
            form.login.ref.validate(async (valid) => {
                if (!valid) {
                    return false
                }
                form.login.loading = true
                let result = await useUserStore().login(form.login.data.account, form.login.data.password)
                form.login.loading = false
                if (result) {
                    router.replace(route.query.redirect ? route.query.redirect : '/')
                }
            })
        }
    }
})
</script>

<template>
    <div class="particles" :particlesNumber="200" />
    <ElTag id="login-container">
        <ElDivider>
            <div style="font-size: 25px;">登录</div>
        </ElDivider>
        <ElForm :ref="(el) => { form.login.ref = el }" :model="form.login.data" :rules="form.login.rules">
            <ElFormItem prop="account">
                <ElInput v-model="form.login.data.account" placeholder="账号">
                    <template #prefix>
                        <AutoiconEpUser />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem prop="password">
                <ElInput v-model="form.login.data.password" type="password" placeholder="密码" :show-password="true"
                    @keyup.enter="form.login.submit">
                    <template #prefix>
                        <AutoiconEpLock />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem>
                <ElButton :loading="form.login.loading" type="primary" @click="form.login.submit" style="width:100%;">
                    登录
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