<script setup lang="tsx">
import md5 from 'js-md5'
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

let loginNamePrefix = saveCommon.data.id ? '' : useAdminStore().info.org_id + ':'
if (saveCommon.data.phone && saveCommon.data.phone.indexOf(':') !== -1) {
    loginNamePrefix = saveCommon.data.phone.split(':')[0] + `:`
} else if (saveCommon.data.email && saveCommon.data.email.indexOf(':') !== -1) {
    loginNamePrefix = saveCommon.data.email.split(':')[0] + `:`
} else if (saveCommon.data.account && saveCommon.data.account.indexOf(':') !== -1) {
    loginNamePrefix = saveCommon.data.account.split(':')[0] + `:`
}
const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data,
        org_id: saveCommon.data.org_id ? saveCommon.data.org_id : undefined,
        phone: saveCommon.data.phone ? (saveCommon.data.phone.indexOf(':') === -1 ? saveCommon.data.phone : saveCommon.data.phone.split(':')[1]) : undefined,
        email: saveCommon.data.email ? (saveCommon.data.email.indexOf(':') === -1 ? saveCommon.data.email : saveCommon.data.email.split(':')[1]) : undefined,
        account: saveCommon.data.account ? (saveCommon.data.account.indexOf(':') === -1 ? saveCommon.data.account : saveCommon.data.account.split(':')[1]) : undefined,
    } as { [propName: string]: any },
    rules: {
        /* org_id: [
            // { required: true, message: t('validation.required') },
            { type: 'integer', trigger: 'change', min: 1, max: 4294967295, message: t('validation.select') },
        ], */
        nickname: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        avatar: [
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        phone: [
            {
                required: computed((): boolean => (saveForm.data.email || saveForm.data.account ? false : true)),
                message: t('validation.required'),
            },
            { type: 'string', trigger: 'blur', max: 20, message: t('validation.max.string', { max: 20 }) },
            { type: 'string', trigger: 'blur', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },
        ],
        email: [
            {
                required: computed((): boolean => (saveForm.data.phone || saveForm.data.account ? false : true)),
                message: t('validation.required'),
            },
            { type: 'string', trigger: 'blur', max: 60, message: t('validation.max.string', { max: 60 }) },
            { type: 'email', trigger: 'blur', message: t('validation.email') },
        ],
        account: [
            {
                required: computed((): boolean => (saveForm.data.phone || saveForm.data.email ? false : true)),
                message: t('validation.required'),
            },
            { type: 'string', trigger: 'blur', max: 20, message: t('validation.max.string', { max: 20 }) },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },
        ],
        // is_super: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
        password: [
            { required: computed((): boolean => (saveForm.data.id ? false : true)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },
        ],
        role_id_arr: [
            { required: true, message: t('validation.required') },
            { type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'integer', min: 1, max: 4294967295, message: t('validation.select') } }, // 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })
        ],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as { value: any; label: string }[]).map((item) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.org_id === undefined && (param.org_id = 0)
            param.password ? (param.password = md5(param.password)) : delete param.password
            try {
                if (param?.id) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/org/admin/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/org/admin/create', param, true)
                }
                listCommon.ref.getList(true)
                saveCommon.visible = false
            } finally {
                saveForm.loading = false
            }
        })
    },
})

const saveDrawer = reactive({
    ref: null as any,
    size: useSettingStore().saveDrawer.size,
    beforeClose: (done: Function) => {
        if (useSettingStore().saveDrawer.isTipClose) {
            ElMessageBox.confirm('', {
                type: 'info',
                title: t('common.tip.configExit'),
                center: true,
                showClose: false,
            }).then(() => done())
        } else {
            done()
        }
    },
    buttonClose: () => saveDrawer.ref.handleClose(), //saveCommon.visible = false //不会触发beforeClose
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <!-- <el-form-item :label="t('org.admin.name.org_id')" prop="org_id">
                    <my-select v-model="saveForm.data.org_id" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/org/org/list' }" />
                </el-form-item> -->
                <el-form-item :label="t('org.admin.name.nickname')" prop="nickname">
                    <el-input v-model="saveForm.data.nickname" :placeholder="t('org.admin.name.nickname')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('org.admin.name.avatar')" prop="avatar">
                    <my-upload v-model="saveForm.data.avatar" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('org.admin.name.phone')" prop="phone">
                    <el-input v-model="saveForm.data.phone" :placeholder="t('org.admin.name.phone')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px">
                        <template v-if="loginNamePrefix" #prepend>{{ loginNamePrefix }}</template>
                    </el-input>
                    <el-alert :title="t('common.tip.notDuplicate') + (loginNamePrefix ? t('org.admin.tip.login_name_prefix', { login_name_prefix: loginNamePrefix }) : '')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('org.admin.name.email')" prop="email">
                    <el-input v-model="saveForm.data.email" :placeholder="t('org.admin.name.email')" maxlength="60" :show-word-limit="true" :clearable="true" style="max-width: 250px">
                        <template v-if="loginNamePrefix" #prepend>{{ loginNamePrefix }}</template>
                    </el-input>
                    <el-alert :title="t('common.tip.notDuplicate') + (loginNamePrefix ? t('org.admin.tip.login_name_prefix', { login_name_prefix: loginNamePrefix }) : '')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('org.admin.name.account')" prop="account">
                    <el-input v-model="saveForm.data.account" :placeholder="t('org.admin.name.account')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px">
                        <template v-if="loginNamePrefix" #prepend>{{ loginNamePrefix }}</template>
                    </el-input>
                    <el-alert :title="t('common.tip.notDuplicate') + (loginNamePrefix ? t('org.admin.tip.login_name_prefix', { login_name_prefix: loginNamePrefix }) : '')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <!-- <el-form-item :label="t('org.admin.name.is_super')" prop="is_super">
                    <el-switch
                        v-model="saveForm.data.is_super"
                        :active-value="(tm('common.status.whether') as any[])[1].value"
                        :inactive-value="(tm('common.status.whether') as any[])[0].value"
                        :active-text="(tm('common.status.whether') as any[])[1].label"
                        :inactive-text="(tm('common.status.whether') as any[])[0].label"
                        :inline-prompt="true"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item> -->
                <el-form-item :label="t('org.admin.name.password')" prop="password">
                    <el-input v-model="saveForm.data.password" :placeholder="t('org.admin.name.password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert v-if="saveForm.data.id" :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('org.admin.name.role_id_arr')" prop="role_id_arr">
                    <!-- 建议：大表用<my-select>（滚动分页），小表用<my-transfer>（无分页） -->
                    <!-- <my-select v-model="saveForm.data.role_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/role/list' }" :multiple="true" /> -->
                    <my-transfer v-model="saveForm.data.role_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/role/list' }" />
                </el-form-item>
                <el-form-item :label="t('org.admin.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
                        :active-value="(tm('common.status.whether') as any[])[1].value"
                        :inactive-value="(tm('common.status.whether') as any[])[0].value"
                        :active-text="(tm('common.status.whether') as any[])[1].label"
                        :inactive-text="(tm('common.status.whether') as any[])[0].label"
                        :inline-prompt="true"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item>
            </el-form>
        </el-scrollbar>
        <template #footer>
            <el-button @click="saveDrawer.buttonClose">{{ t('common.cancel') }}</el-button>
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">{{ t('common.save') }}</el-button>
        </template>
    </el-drawer>
</template>
