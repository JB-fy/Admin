<template>
    <vue-particles class="particles" :particlesNumber="200" />
    <el-tag id="login-container">
        <el-divider>
            <div style="font-size: 25px;">登录</div>
        </el-divider>
        <el-form :ref="(el) => { form.login.ref = el }" :model="form.login.data" :rules="form.login.rules">
            <el-form-item prop="account">
                <el-input v-model="form.login.data.account" placeholder="账号">
                    <template #prefix>
                        <autoicon-ep-user />
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item prop="password">
                <el-input v-model="form.login.data.password" type="password" placeholder="密码" :show-password="true"
                    @keyup.enter="form.login.submit">
                    <template #prefix>
                        <autoicon-ep-lock />
                    </template>
                </el-input>
            </el-form-item>
            <el-form-item>
                <el-button :loading="form.login.loading" type="primary" @click="form.login.submit" style="width:100%;">
                    登录
                </el-button>
            </el-form-item>
        </el-form>
    </el-tag>
</template>

<script>
import VueParticles from 'vue-particles/src/vue-particles/vue-particles.vue'

export default {
    components: {
        VueParticles
    },
    setup: () => {
        const router = useRouter()
        const route = useRoute()
        const store = useStore()
        const state = reactive({
            form: {
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
                        state.form.login.ref.validate(async (valid) => {
                            if (!valid) {
                                return false
                            }
                            state.form.login.loading = true
                            let result = await store.dispatch('user/login', state.form.login.data)
                            state.form.login.loading = false
                            if (result) {
                                router.replace(route.query.redirect ? route.query.redirect : '/')
                            }
                        })
                    }
                }
            }
        })

        return {
            ...toRefs(state)
        }
    }
}
</script>

<style scoped>
.particles {
    width: 100%;
    height: 100vh;
    overflow: hidden;
    background-image: url(~@/assets/login-bg.jpg);
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