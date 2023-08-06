<script setup lang="ts">
const router = useRouter()
const route = useRoute()
const { t } = useI18n()

const particleOptions = reactive({
    fpsLimit: 120,
    interactivity: {
        events: {
            onClick: {
                enable: true,
                mode: 'push'
            },
            onHover: {
                enable: true,
                mode: 'repulse'
            },
            resize: true
        },
        modes: {
            bubble: {
                distance: 400,
                duration: 2,
                opacity: 0.8,
                size: 40
            },
            push: {
                quantity: 4
            },
            repulse: {
                distance: 200,
                duration: 0.4
            }
        }
    },
    particles: {
        color: {
            value: '#ffffff'
        },
        links: {
            color: '#ffffff',
            distance: 150,
            enable: true,
            opacity: 0.5,
            width: 1
        },
        move: {
            direction: 'none',
            enable: true,
            outMode: 'bounce',
            random: false,
            speed: 6,
            straight: false
        },
        number: {
            density: {
                enable: true,
                area: 800
            },
            value: 80
        },
        opacity: {
            value: 0.5
        },
        shape: {
            type: 'circle'
        },
        size: {
            random: true,
            value: 5
        }
    },
    detectRetina: true
})

const loginForm = reactive({
    ref: null as any,
    data: {
        account: '',
        password: ''
    },
    rules: {
        account: [
            { type: 'string', required: true, min: 1, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 1, max: 30 }) }
        ],
        password: [
            { type: 'string', required: true, min: 6, max: 30, trigger: 'blur', message: t('validation.between.string', { min: 6, max: 30 }) }
        ]
    } as any,
    loading: false,
    submit: () => {
        loginForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return false
            }
            loginForm.loading = true
            try {
                await useAdminStore().login(loginForm.data.account, loginForm.data.password)
                router.replace((route.query.redirect ? route.query.redirect : '/') as string)
            } catch (error) { }
            loginForm.loading = false
        })
    }
})
</script>

<template>
    <div class="particles"></div>
    <ElTag id="login-container">
        <ElDivider>
            <div style="font-size: 25px;">{{ t('common.login') }}</div>
        </ElDivider>
        <ElForm :ref="(el: any) => { loginForm.ref = el }" :model="loginForm.data" :rules="loginForm.rules"
            @keyup.enter="loginForm.submit">
            <ElFormItem prop="account">
                <ElInput v-model="loginForm.data.account"
                    :placeholder="t('login.name.account') + '/' + t('login.name.phone')">
                    <template #prefix>
                        <AutoiconEpUser />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem prop="password">
                <ElInput v-model="loginForm.data.password" type="password" :placeholder="t('login.name.password')"
                    :show-password="true">
                    <template #prefix>
                        <AutoiconEpLock />
                    </template>
                </ElInput>
            </ElFormItem>
            <ElFormItem>
                <ElButton :loading="loginForm.loading" type="primary" @click="loginForm.submit" style="width:100%;">
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