<script setup lang="tsx">
import md5 from 'js-md5'
const { t, tm } = useI18n()

const saveCommon = inject('saveCommon') as { visible: boolean; title: string; data: { [propName: string]: any } }
const listCommon = inject('listCommon') as { ref: any }

const saveForm = reactive({
    ref: null as any,
    loading: false,
    data: {
        gender: 0,
        id_card_gender: 0,
        ...saveCommon.data,
    } as { [propName: string]: any },
    rules: {
        nickname: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        avatar: [
            { type: 'string', trigger: 'blur', max: 200, message: t('validation.max.string', { max: 200 }) },
            { type: 'url', trigger: 'change', message: t('validation.upload') },
        ],
        gender: [{ type: 'enum', trigger: 'change', enum: (tm('users.users.status.gender') as any).map((item: any) => item.value), message: t('validation.select') }],
        birthday: [{ type: 'string', trigger: 'change', message: t('validation.select') }],
        address: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
        phone: [
            { type: 'string', trigger: 'blur', max: 20, message: t('validation.max.string', { max: 20 }) },
            { type: 'string', trigger: 'blur', pattern: /^1[3-9]\d{9}$/, message: t('validation.phone') },
        ],
        email: [
            { type: 'string', trigger: 'blur', max: 60, message: t('validation.max.string', { max: 60 }) },
            { type: 'email', trigger: 'blur', message: t('validation.email') },
        ],
        account: [
            { type: 'string', trigger: 'blur', max: 20, message: t('validation.max.string', { max: 20 }) },
            { type: 'string', trigger: 'blur', pattern: /^[\p{L}][\p{L}\p{N}_]{3,}$/u, message: t('validation.account') },
        ],
        wx_openid: [{ type: 'string', trigger: 'blur', max: 128, message: t('validation.max.string', { max: 128 }) }],
        wx_unionid: [{ type: 'string', trigger: 'blur', max: 64, message: t('validation.max.string', { max: 64 }) }],
        password: [
            { required: computed((): boolean => (saveForm.data.id_arr?.length ? false : true)), message: t('validation.required') },
            { type: 'string', trigger: 'blur', min: 6, max: 20, message: t('validation.between.string', { min: 6, max: 20 }) },
        ],
        id_card_no: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        id_card_name: [{ type: 'string', trigger: 'blur', max: 30, message: t('validation.max.string', { max: 30 }) }],
        id_card_gender: [{ type: 'enum', trigger: 'change', enum: (tm('users.users.status.id_card_gender') as any).map((item: any) => item.value), message: t('validation.select') }],
        id_card_birthday: [{ type: 'string', trigger: 'change', message: t('validation.select') }],
        id_card_address: [{ type: 'string', trigger: 'blur', max: 120, message: t('validation.max.string', { max: 120 }) }],
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
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/users/users/update', param, true)
                } else {
                    await request(t('config.VITE_HTTP_API_PREFIX') + '/users/users/create', param, true)
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
            })
                .then(() => {
                    done()
                })
                .catch(() => {})
        } else {
            done()
        }
    },
    buttonClose: () => {
        //saveCommon.visible = false
        saveDrawer.ref.handleClose() //会触发beforeClose
    },
})
</script>

<template>
    <el-drawer class="save-drawer" :ref="(el: any) => saveDrawer.ref = el" v-model="saveCommon.visible" :title="saveCommon.title" :size="saveDrawer.size" :before-close="saveDrawer.beforeClose">
        <el-scrollbar>
            <el-form :ref="(el: any) => saveForm.ref = el" :model="saveForm.data" :rules="saveForm.rules" label-width="auto" :status-icon="true" :scroll-to-error="true">
                <el-form-item :label="t('users.users.name.nickname')" prop="nickname">
                    <el-input v-model="saveForm.data.nickname" :placeholder="t('users.users.name.nickname')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.avatar')" prop="avatar">
                    <my-upload v-model="saveForm.data.avatar" accept="image/*" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.gender')" prop="gender">
                    <el-radio-group v-model="saveForm.data.gender">
                        <el-radio v-for="(item, index) in (tm('users.users.status.gender') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('users.users.name.birthday')" prop="birthday">
                    <el-date-picker v-model="saveForm.data.birthday" type="date" :placeholder="t('users.users.name.birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.address')" prop="address">
                    <el-input v-model="saveForm.data.address" :placeholder="t('users.users.name.address')" maxlength="120" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.phone')" prop="phone">
                    <el-input v-model="saveForm.data.phone" :placeholder="t('users.users.name.phone')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.email')" prop="email">
                    <el-input v-model="saveForm.data.email" :placeholder="t('users.users.name.email')" maxlength="60" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.account')" prop="account">
                    <el-input v-model="saveForm.data.account" :placeholder="t('users.users.name.account')" maxlength="20" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.wx_openid')" prop="wx_openid">
                    <el-input v-model="saveForm.data.wx_openid" :placeholder="t('users.users.name.wx_openid')" maxlength="128" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.wx_unionid')" prop="wx_unionid">
                    <el-input v-model="saveForm.data.wx_unionid" :placeholder="t('users.users.name.wx_unionid')" maxlength="64" :show-word-limit="true" :clearable="true" style="max-width: 250px" />
                    <el-alert :title="t('common.tip.notDuplicate')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.password')" prop="password">
                    <el-input v-model="saveForm.data.password" :placeholder="t('users.users.name.password')" minlength="6" maxlength="20" :show-word-limit="true" :clearable="true" :show-password="true" style="max-width: 250px" />
                    <el-alert v-if="saveForm.data.id_arr?.length" :title="t('common.tip.notRequired')" type="info" :show-icon="true" :closable="false" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.id_card_no')" prop="id_card_no">
                    <el-input v-model="saveForm.data.id_card_no" :placeholder="t('users.users.name.id_card_no')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.id_card_name')" prop="id_card_name">
                    <el-input v-model="saveForm.data.id_card_name" :placeholder="t('users.users.name.id_card_name')" maxlength="30" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.id_card_gender')" prop="id_card_gender">
                    <el-radio-group v-model="saveForm.data.id_card_gender">
                        <el-radio v-for="(item, index) in (tm('users.users.status.id_card_gender') as any)" :key="index" :value="item.value">
                            {{ item.label }}
                        </el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item :label="t('users.users.name.id_card_birthday')" prop="id_card_birthday">
                    <el-date-picker v-model="saveForm.data.id_card_birthday" type="date" :placeholder="t('users.users.name.id_card_birthday')" format="YYYY-MM-DD" value-format="YYYY-MM-DD" style="width: 160px" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.id_card_address')" prop="id_card_address">
                    <el-input v-model="saveForm.data.id_card_address" :placeholder="t('users.users.name.id_card_address')" maxlength="120" :show-word-limit="true" :clearable="true" />
                </el-form-item>
                <el-form-item :label="t('users.users.name.is_stop')" prop="is_stop">
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
            <el-button type="primary" @click="saveForm.submit" :loading="saveForm.loading">
                {{ t('common.save') }}
            </el-button>
        </template>
    </el-drawer>
</template>
