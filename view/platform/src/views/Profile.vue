<script setup lang="tsx">
import md5 from 'js-md5'

const { t } = useI18n()
const adminStore = useAdminStore()

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        nickname: adminStore.info.nickname,
        avatar: adminStore.info.avatar,
    } as { [propName: string]: any },
    rules: {
        nickname: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        avatar: [
            { type: 'url', trigger: 'change', message: t('validation.upload') },
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
        ],
        phone: [
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            { type: 'string', trigger: 'blur', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },
        ],
        email: [
            { type: 'string', trigger: 'blur', max: 60, message: t('validation.max.string', { max: 60 }) },
            { type: 'email', trigger: 'blur', message: t('validation.email') },
        ],
        account: [
            { type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },
        ],
        password: [{ type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) }],
        repeat_password: [
            { required: computed((): boolean => (saveForm.data.password ? true : false)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },
            {
                trigger: 'blur',
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password != saveForm.data.repeat_password) {
                        callback(new Error())
                    }
                    callback()
                },
                message: t('validation.repeat_password'),
            },
        ],
        password_to_check: [
            { required: computed((): boolean => (saveForm.data.phone || saveForm.data.email || saveForm.data.account || saveForm.data.password ? true : false)), message: t('profile.tip.password_to_check') },
            { type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },
            {
                trigger: 'blur',
                validator: (rule: any, value: any, callback: any) => {
                    if (saveForm.data.password && saveForm.data.password == saveForm.data.password_to_check) {
                        callback(new Error())
                    }
                    callback()
                },
                message: t('validation.new_password_diff_old_password'),
            },
        ],
        sms_code_to_bind_phone: [
            { required: computed((): boolean => (saveForm.data.phone ? true : false)), message: t('profile.tip.sms_code_to_bind_phone') },
            { type: 'string', len: 4, message: t('validation.size.string', { size: 4 }) },
        ],
        email_code_to_bind_email: [
            { required: computed((): boolean => (saveForm.data.email ? true : false)), message: t('profile.tip.email_code_to_bind_email') },
            { type: 'string', len: 4, message: t('validation.size.string', { size: 4 }) },
        ],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.phone || delete param.phone
            param.email || delete param.email
            param.account || delete param.account
            param.password ? (param.password = md5(param.password)) : delete param.password
            delete param.repeat_password
            param.password_to_check ? (param.password_to_check = md5(param.password_to_check)) : delete param.password_to_check
            try {
                await request(t('config.VITE_HTTP_API_PREFIX') + '/my/profile/update', param, true)
                //成功则更新用户信息
                for (let k in param) {
                    if (k in adminStore.info) {
                        adminStore.info[k] = param[k]
                    }
                }
            } finally {
                saveForm.loading = false
            }
        })
    },
})

const smsCountdown = reactive({
    isShow: false,
    value: 0,
    finish: () => {
        smsCountdown.isShow = false
        smsCountdown.value = 0
    },
    send: async () => {
        try {
            smsCountdown.isShow = true
            await request(t('config.VITE_HTTP_API_PREFIX') + '/code/send', { scene: 4, to: saveForm.data.phone }, true)
            smsCountdown.value = Date.now() + 5 * 60 * 1000
        } catch (error) {
            smsCountdown.finish()
        }
    },
})

const emailCountdown = reactive({
    isShow: false,
    value: 0,
    finish: () => {
        emailCountdown.isShow = false
        emailCountdown.value = 0
    },
    send: async () => {
        try {
            emailCountdown.isShow = true
            await request(t('config.VITE_HTTP_API_PREFIX') + '/code/send', { scene: 14, to: saveForm.data.email }, true)
            emailCountdown.value = Date.now() + 5 * 60 * 1000
        } catch (error) {
            emailCountdown.finish()
        }
    },
})
</script>

<template>
    <el-container class="common-container">
        <el-main>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="false">
                <el-form-item :label="t('profile.name.nickname')" prop="nickname">
                    <el-input v-model="saveForm.data.nickname" :placeholder="t('profile.name.nickname')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('profile.name.avatar')" prop="avatar">
                    <my-upload v-model="saveForm.data.avatar" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('profile.name.phone')" prop="phone">
                    <el-input v-model="saveForm.data.phone" :placeholder="t('profile.name.phone')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.phone', { phone: adminStore.info.phone ? adminStore.info.phone : t('common.tip.notSet') })" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.email')" prop="email">
                    <el-input v-model="saveForm.data.email" :placeholder="t('profile.name.email')" maxlength="60" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.email', { email: adminStore.info.email ? adminStore.info.email : t('common.tip.notSet') })" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.account')" prop="account">
                    <el-input v-model="saveForm.data.account" :placeholder="t('profile.name.account')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.account', { account: adminStore.info.account ? adminStore.info.account : t('common.tip.notSet') })" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('profile.name.password')" prop="password">
                    <el-input v-model="saveForm.data.password" :placeholder="t('profile.name.password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item v-if="saveForm.data.password" :label="t('profile.name.repeat_password')" prop="repeat_password">
                    <el-input v-model="saveForm.data.repeat_password" :placeholder="t('profile.name.repeat_password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.repeat_password')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item v-if="saveForm.data.phone || saveForm.data.email || saveForm.data.account || saveForm.data.password" :label="t('profile.name.password_to_check')" prop="password_to_check">
                    <el-input v-model="saveForm.data.password_to_check" :placeholder="t('profile.name.password_to_check')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert :title="t('profile.tip.password_to_check')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item v-if="saveForm.data.phone" :label="t('profile.name.sms_code_to_bind_phone')" prop="sms_code_to_bind_phone">
                    <el-input v-model="saveForm.data.sms_code_to_bind_phone" :placeholder="t('profile.name.sms_code_to_bind_phone')" minlength="4" maxlength="4" :show-word-limit="true" :clearable="true" style="max-width: 250px">
                        <template #append>
                            <el-countdown v-if="smsCountdown.isShow && smsCountdown.value > 0" :value="smsCountdown.value" @finish="smsCountdown.finish" format="mm:ss" value-style="color: #909399;" />
                            <el-button v-else :loading="smsCountdown.isShow" @click="smsCountdown.send">{{ t('profile.send_code') }}</el-button>
                        </template>
                    </el-input>
                    <el-alert :title="t('profile.tip.sms_code_to_bind_phone')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item v-if="saveForm.data.email" :label="t('profile.name.email_code_to_bind_email')" prop="email_code_to_bind_email">
                    <el-input v-model="saveForm.data.email_code_to_bind_email" :placeholder="t('profile.name.email_code_to_bind_email')" minlength="4" maxlength="4" :show-word-limit="true" :clearable="true" style="max-width: 250px">
                        <template #append>
                            <el-countdown v-if="emailCountdown.isShow && emailCountdown.value > 0" :value="emailCountdown.value" @finish="emailCountdown.finish" format="mm:ss" value-style="color: #909399;" />
                            <el-button v-else :loading="emailCountdown.isShow" @click="emailCountdown.send">{{ t('profile.send_code') }}</el-button>
                        </template>
                    </el-input>
                    <el-alert :title="t('profile.tip.email_code_to_bind_email')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item>
                    <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading"> <autoicon-ep-circle-check />{{ t('common.save') }}</el-button>
                </el-form-item>
            </el-form>
        </el-main>
    </el-container>
</template>
