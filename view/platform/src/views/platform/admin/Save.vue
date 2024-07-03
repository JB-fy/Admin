<script setup lang="tsx">
import md5 from 'js-md5'
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {
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
        password: [
            { required: computed((): boolean => (saveForm.data.id_arr?.length ? false : true)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },
        ],
        // is_super: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
        role_id_arr: [
            { required: true, message: t('validation.required') },
            { type: 'array', trigger: 'change', message: t('validation.select'), defaultField: { type: 'integer', min: 1, max: 4294967295, message: t('validation.select') } }, // 限制数组数量时用：max: 10, message: t('validation.max.select', { max: 10 })
        ],
        is_stop: [{ type: 'enum', trigger: 'change', enum: (tm('common.status.whether') as any).map((item: any) => item.value), message: t('validation.select') }],
    } as { [propName: string]: { [propName: string]: any } | { [propName: string]: any }[] },
    submit: () => {
        saveForm.ref.validate(async (valid: boolean) => {
            if (!valid) {
                return
            }
            saveForm.loading = true
            const param = removeEmptyOfObj(saveForm.data)
            param.password ? (param.password = md5(param.password)) : delete param.password
            try {
                if (param?.id_arr?.length > 0) {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/platform/admin/create', param, true)
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
                <el-form-item :label="t('platform.admin.name.nickname')" prop="nickname">
                    <el-input v-model="saveForm.data.nickname" :placeholder="t('platform.admin.name.nickname')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('platform.admin.name.avatar')" prop="avatar">
                    <my-upload v-model="saveForm.data.avatar" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('platform.admin.name.phone')" prop="phone">
                    <el-input v-model="saveForm.data.phone" :placeholder="t('platform.admin.name.phone')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('platform.admin.name.email')" prop="email">
                    <el-input v-model="saveForm.data.email" :placeholder="t('platform.admin.name.email')" maxlength="60" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('platform.admin.name.account')" prop="account">
                    <el-input v-model="saveForm.data.account" :placeholder="t('platform.admin.name.account')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('platform.admin.name.password')" prop="password">
                    <el-input v-model="saveForm.data.password" :placeholder="t('platform.admin.name.password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert v-if="saveForm.data.id_arr?.length" :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <!-- <el-form-item :label="t('platform.admin.name.is_super')" prop="is_super">
                    <el-switch
                        v-model="saveForm.data.is_super"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
                        style="--el-switch-on-color: var(--el-color-danger); --el-switch-off-color: var(--el-color-success)"
                    />
                </el-form-item> -->
                <el-form-item :label="t('platform.admin.name.role_id_arr')" prop="role_id_arr">
                    <!-- 建议：大表用<my-select>（滚动分页），小表用<my-transfer>（无分页） -->
                    <!-- <my-select v-model="saveForm.data.role_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/role/list', param: { filter: { scene_code: `platform` } } }" :multiple="true" /> -->
                    <my-transfer v-model="saveForm.data.role_id_arr" :api="{ code: t('config.VITE_HTTP_API_PREFIX') + '/auth/role/list', param: { filter: { scene_code: `platform` } } }" />
                </el-form-item>
                <el-form-item :label="t('platform.admin.name.is_stop')" prop="is_stop">
                    <el-switch
                        v-model="saveForm.data.is_stop"
                        :active-value="1"
                        :inactive-value="0"
                        :inline-prompt="true"
                        :active-text="t('common.yes')"
                        :inactive-text="t('common.no')"
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
